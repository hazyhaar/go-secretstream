/* CLI driver: libsodium crypto_secretstream_xchacha20poly1305
 * Commands:
 *   push-stream <key.bin> <plain.bin> <wire.bin>
 *   pull-stream <key.bin> <wire.bin> <plain.out>
 *   push-script <key.bin> <script.txt> <wire.bin>
 *   pull-script <key.bin> <wire.bin> <script.txt> <plain.out>
 * Exit 1 on crypto failure.
 */
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stddef.h>
#include <stdint.h>
#include <sodium.h>

_Static_assert(sizeof(crypto_secretstream_xchacha20poly1305_state) == 52,
               "libsodium 1.0.18 state size");
_Static_assert(offsetof(crypto_secretstream_xchacha20poly1305_state, nonce) == 32,
               "libsodium 1.0.18 nonce offset");

#define CHUNK 8192

static int read_all(const char *path, unsigned char **buf, size_t *len) {
    FILE *f = fopen(path, "rb");
    if (!f) return -1;
    fseek(f, 0, SEEK_END);
    long n = ftell(f);
    fseek(f, 0, SEEK_SET);
    if (n < 0) { fclose(f); return -1; }
    *buf = malloc((size_t)n + 1);
    if (!*buf) { fclose(f); return -1; }
    if (n > 0 && fread(*buf, 1, (size_t)n, f) != (size_t)n) {
        free(*buf); fclose(f); return -1;
    }
    *len = (size_t)n;
    fclose(f);
    return 0;
}

static int write_all(const char *path, const unsigned char *buf, size_t len) {
    FILE *f = fopen(path, "wb");
    if (!f) return -1;
    if (len && fwrite(buf, 1, len, f) != len) { fclose(f); return -1; }
    fclose(f);
    return 0;
}

static int cmd_push(const char *keyp, const char *plainp, const char *wirep) {
    unsigned char *key = NULL, *plain = NULL;
    size_t klen = 0, plen = 0;
    if (read_all(keyp, &key, &klen) || klen != crypto_secretstream_xchacha20poly1305_KEYBYTES) {
        fprintf(stderr, "bad key\n"); return 1;
    }
    if (read_all(plainp, &plain, &plen)) { fprintf(stderr, "bad plain\n"); return 1; }

    crypto_secretstream_xchacha20poly1305_state st;
    unsigned char header[crypto_secretstream_xchacha20poly1305_HEADERBYTES];
    if (crypto_secretstream_xchacha20poly1305_init_push(&st, header, key) != 0) {
        fprintf(stderr, "init_push failed\n"); return 1;
    }

    size_t cap = crypto_secretstream_xchacha20poly1305_HEADERBYTES +
                 plen + (plen / CHUNK + 2) * crypto_secretstream_xchacha20poly1305_ABYTES + 64;
    unsigned char *out = malloc(cap);
    if (!out) return 1;
    size_t o = 0;
    memcpy(out + o, header, sizeof header); o += sizeof header;

    size_t off = 0;
    if (plen == 0) {
        unsigned char c[crypto_secretstream_xchacha20poly1305_ABYTES];
        unsigned long long clen = 0;
        if (crypto_secretstream_xchacha20poly1305_push(&st, c, &clen, NULL, 0, NULL, 0,
                crypto_secretstream_xchacha20poly1305_TAG_FINAL) != 0) {
            fprintf(stderr, "push empty final failed\n"); return 1;
        }
        memcpy(out + o, c, (size_t)clen); o += (size_t)clen;
    } else {
        while (off < plen) {
            size_t end = off + CHUNK;
            if (end > plen) end = plen;
            size_t mlen = end - off;
            int last = (end >= plen);
            unsigned char tag = last ? crypto_secretstream_xchacha20poly1305_TAG_FINAL
                                     : crypto_secretstream_xchacha20poly1305_TAG_MESSAGE;
            unsigned char *c = out + o;
            unsigned long long clen = 0;
            if (crypto_secretstream_xchacha20poly1305_push(&st, c, &clen, plain + off, mlen, NULL, 0, tag) != 0) {
                fprintf(stderr, "push failed\n"); return 1;
            }
            o += (size_t)clen;
            off = end;
        }
    }
    int rc = write_all(wirep, out, o);
    free(out); free(key); free(plain);
    return rc ? 1 : 0;
}

static int cmd_pull(const char *keyp, const char *wirep, const char *plainp) {
    unsigned char *key = NULL, *wire = NULL;
    size_t klen = 0, wlen = 0;
    if (read_all(keyp, &key, &klen) || klen != crypto_secretstream_xchacha20poly1305_KEYBYTES) {
        fprintf(stderr, "bad key\n"); return 1;
    }
    if (read_all(wirep, &wire, &wlen) ||
        wlen < crypto_secretstream_xchacha20poly1305_HEADERBYTES) {
        fprintf(stderr, "bad wire\n"); return 1;
    }
    crypto_secretstream_xchacha20poly1305_state st;
    if (crypto_secretstream_xchacha20poly1305_init_pull(&st, wire, key) != 0) {
        fprintf(stderr, "init_pull failed\n"); return 1;
    }
    size_t off = crypto_secretstream_xchacha20poly1305_HEADERBYTES;
    size_t pcap = wlen;
    unsigned char *plain = malloc(pcap);
    size_t po = 0;
    if (!plain) return 1;

    while (off < wlen) {
        /* wal-g: full chunks are CHUNK+ABYTES; last may be shorter */
        size_t remain = wlen - off;
        size_t clen = CHUNK + crypto_secretstream_xchacha20poly1305_ABYTES;
        if (clen > remain) clen = remain;
        if (clen < crypto_secretstream_xchacha20poly1305_ABYTES) {
            fprintf(stderr, "truncated chunk\n"); return 1;
        }
        unsigned char m[CHUNK];
        unsigned long long mlen = 0;
        unsigned char tag = 0;
        if (crypto_secretstream_xchacha20poly1305_pull(&st, m, &mlen, &tag, wire + off, clen, NULL, 0) != 0) {
            fprintf(stderr, "pull MAC fail off=%zu clen=%zu\n", off, clen);
            return 1;
        }
        if (po + (size_t)mlen > pcap) {
            pcap *= 2;
            plain = realloc(plain, pcap);
        }
        memcpy(plain + po, m, (size_t)mlen);
        po += (size_t)mlen;
        off += clen;
        if (tag == crypto_secretstream_xchacha20poly1305_TAG_FINAL) break;
    }
    int rc = write_all(plainp, plain, po);
    free(plain); free(key); free(wire);
    return rc ? 1 : 0;
}

static int parse_tag(const char *s, unsigned char *tag) {
    if (strcmp(s, "MESSAGE") == 0) {
        *tag = crypto_secretstream_xchacha20poly1305_TAG_MESSAGE;
        return 0;
    }
    if (strcmp(s, "PUSH") == 0) {
        *tag = crypto_secretstream_xchacha20poly1305_TAG_PUSH;
        return 0;
    }
    if (strcmp(s, "REKEY") == 0) {
        *tag = crypto_secretstream_xchacha20poly1305_TAG_REKEY;
        return 0;
    }
    if (strcmp(s, "FINAL") == 0) {
        *tag = crypto_secretstream_xchacha20poly1305_TAG_FINAL;
        return 0;
    }
    return -1;
}

static void fill_det_msg(unsigned char *m, size_t len) {
    size_t i;
    for (i = 0; i < len; i++) {
        m[i] = (unsigned char)((i * 7 + len) & 0xff);
    }
}

static int parse_hex4(const char *hex, unsigned char out[4]) {
    int i;
    if (hex == NULL || strlen(hex) != 8) return -1;
    for (i = 0; i < 4; i++) {
        unsigned int b = 0;
        if (sscanf(hex + 2 * i, "%2x", &b) != 1) return -1;
        out[i] = (unsigned char)b;
    }
    return 0;
}

static int force_counter(crypto_secretstream_xchacha20poly1305_state *st, const char *hex) {
    unsigned char bytes[4];
    unsigned char *nonce;
    if (parse_hex4(hex, bytes) != 0) return -1;
    nonce = (unsigned char *)st + offsetof(crypto_secretstream_xchacha20poly1305_state, nonce);
    memcpy(nonce, bytes, 4);
    return 0;
}

static int cmd_push_script(const char *keyp, const char *scriptp, const char *wirep) {
    unsigned char *key = NULL;
    size_t klen = 0;
    FILE *sf;
    char line[256];
    int lineno = 0;
    crypto_secretstream_xchacha20poly1305_state st;
    unsigned char header[crypto_secretstream_xchacha20poly1305_HEADERBYTES];
    unsigned char *out;
    size_t cap, o;

    if (read_all(keyp, &key, &klen) || klen != crypto_secretstream_xchacha20poly1305_KEYBYTES) {
        fprintf(stderr, "bad key\n"); return 1;
    }
    if (crypto_secretstream_xchacha20poly1305_init_push(&st, header, key) != 0) {
        fprintf(stderr, "init_push failed\n"); free(key); return 1;
    }
    cap = crypto_secretstream_xchacha20poly1305_HEADERBYTES + 64;
    out = malloc(cap);
    if (!out) { free(key); return 1; }
    o = 0;
    memcpy(out + o, header, sizeof header); o += sizeof header;

    sf = fopen(scriptp, "r");
    if (!sf) { fprintf(stderr, "bad script\n"); free(out); free(key); return 1; }
    while (fgets(line, sizeof line, sf)) {
        char cmd[32], a1[32], a2[32];
        int n;
        lineno++;
        if (line[0] == '#' || line[0] == '\n' || line[0] == '\r') continue;
        n = sscanf(line, "%31s %31s %31s", cmd, a1, a2);
        if (n < 1) continue;
        if (strcmp(cmd, "force-counter") == 0) {
            if (n < 2 || force_counter(&st, a1) != 0) {
                fprintf(stderr, "script line %d: bad force-counter\n", lineno);
                fclose(sf); free(out); free(key); return 1;
            }
            continue;
        }
        if (strcmp(cmd, "rekey") == 0) {
            crypto_secretstream_xchacha20poly1305_rekey(&st);
            continue;
        }
        if (strcmp(cmd, "msg") == 0) {
            unsigned char tag = 0;
            unsigned long mlen = 0;
            unsigned char *m = NULL;
            unsigned char *c;
            unsigned long long clen = 0;
            if (n < 3 || parse_tag(a1, &tag) != 0 || sscanf(a2, "%lu", &mlen) != 1) {
                fprintf(stderr, "script line %d: bad msg\n", lineno);
                fclose(sf); free(out); free(key); return 1;
            }
            if (mlen > 0) {
                m = malloc(mlen);
                if (!m) { fclose(sf); free(out); free(key); return 1; }
                fill_det_msg(m, (size_t)mlen);
            }
            if (o + mlen + crypto_secretstream_xchacha20poly1305_ABYTES + 16 > cap) {
                cap = (o + (size_t)mlen + crypto_secretstream_xchacha20poly1305_ABYTES + 16) * 2;
                out = realloc(out, cap);
                if (!out) { free(m); fclose(sf); free(key); return 1; }
            }
            c = out + o;
            if (crypto_secretstream_xchacha20poly1305_push(&st, c, &clen, m, mlen, NULL, 0, tag) != 0) {
                fprintf(stderr, "push failed line %d\n", lineno);
                free(m); fclose(sf); free(out); free(key); return 1;
            }
            o += (size_t)clen;
            free(m);
            continue;
        }
        fprintf(stderr, "script line %d: unknown cmd\n", lineno);
        fclose(sf); free(out); free(key); return 1;
    }
    fclose(sf);
    {
        int rc = write_all(wirep, out, o);
        free(out); free(key);
        return rc ? 1 : 0;
    }
}

static int cmd_pull_script(const char *keyp, const char *wirep, const char *scriptp, const char *plainp) {
    unsigned char *key = NULL, *wire = NULL;
    size_t klen = 0, wlen = 0;
    FILE *sf;
    char line[256];
    int lineno = 0;
    crypto_secretstream_xchacha20poly1305_state st;
    size_t off;
    unsigned char *plain;
    size_t pcap, po;

    if (read_all(keyp, &key, &klen) || klen != crypto_secretstream_xchacha20poly1305_KEYBYTES) {
        fprintf(stderr, "bad key\n"); return 1;
    }
    if (read_all(wirep, &wire, &wlen) ||
        wlen < crypto_secretstream_xchacha20poly1305_HEADERBYTES) {
        fprintf(stderr, "bad wire\n"); free(key); return 1;
    }
    if (crypto_secretstream_xchacha20poly1305_init_pull(&st, wire, key) != 0) {
        fprintf(stderr, "init_pull failed\n"); free(key); free(wire); return 1;
    }
    off = crypto_secretstream_xchacha20poly1305_HEADERBYTES;
    pcap = wlen;
    plain = malloc(pcap);
    if (!plain) { free(key); free(wire); return 1; }
    po = 0;

    sf = fopen(scriptp, "r");
    if (!sf) { fprintf(stderr, "bad script\n"); free(plain); free(key); free(wire); return 1; }
    while (fgets(line, sizeof line, sf)) {
        char cmd[32], a1[32], a2[32];
        int n;
        lineno++;
        if (line[0] == '#' || line[0] == '\n' || line[0] == '\r') continue;
        n = sscanf(line, "%31s %31s %31s", cmd, a1, a2);
        if (n < 1) continue;
        if (strcmp(cmd, "force-counter") == 0) {
            if (n < 2 || force_counter(&st, a1) != 0) {
                fprintf(stderr, "script line %d: bad force-counter\n", lineno);
                fclose(sf); free(plain); free(key); free(wire); return 1;
            }
            continue;
        }
        if (strcmp(cmd, "rekey") == 0) {
            crypto_secretstream_xchacha20poly1305_rekey(&st);
            continue;
        }
        if (strcmp(cmd, "msg") == 0) {
            unsigned char parsed_tag = 0;
            unsigned long mlen_expect = 0;
            size_t clen;
            unsigned char *m;
            unsigned long long mlen = 0;
            unsigned char tag = 0;
            if (n < 3 || parse_tag(a1, &parsed_tag) != 0 || sscanf(a2, "%lu", &mlen_expect) != 1) {
                fprintf(stderr, "script line %d: bad msg\n", lineno);
                fclose(sf); free(plain); free(key); free(wire); return 1;
            }
            (void)parsed_tag;
            clen = (size_t)mlen_expect + crypto_secretstream_xchacha20poly1305_ABYTES;
            if (off + clen > wlen) {
                fprintf(stderr, "truncated chunk line %d\n", lineno);
                fclose(sf); free(plain); free(key); free(wire); return 1;
            }
            m = malloc(mlen_expect == 0 ? 1 : (size_t)mlen_expect);
            if (!m) { fclose(sf); free(plain); free(key); free(wire); return 1; }
            if (crypto_secretstream_xchacha20poly1305_pull(&st, m, &mlen, &tag, wire + off, clen, NULL, 0) != 0) {
                fprintf(stdout, "MAC fail\n");
                free(m); fclose(sf); free(plain); free(key); free(wire); return 1;
            }
            if (po + (size_t)mlen > pcap) {
                pcap = (po + (size_t)mlen) * 2;
                plain = realloc(plain, pcap);
                if (!plain) { free(m); fclose(sf); free(key); free(wire); return 1; }
            }
            memcpy(plain + po, m, (size_t)mlen);
            po += (size_t)mlen;
            off += clen;
            printf("tag=%u len=%llu ok\n", (unsigned)tag, mlen);
            free(m);
            continue;
        }
        fprintf(stderr, "script line %d: unknown cmd\n", lineno);
        fclose(sf); free(plain); free(key); free(wire); return 1;
    }
    fclose(sf);
    {
        int rc = write_all(plainp, plain, po);
        free(plain); free(key); free(wire);
        return rc ? 1 : 0;
    }
}

int main(int argc, char **argv) {
    if (sodium_init() < 0) { fprintf(stderr, "sodium_init\n"); return 1; }
    if (argc < 2) {
        fprintf(stderr, "usage: %s push-stream|pull-stream|push-script|pull-script ...\n", argv[0]);
        return 2;
    }
    if (strcmp(argv[1], "push-stream") == 0 && argc == 5)
        return cmd_push(argv[2], argv[3], argv[4]);
    if (strcmp(argv[1], "pull-stream") == 0 && argc == 5)
        return cmd_pull(argv[2], argv[3], argv[4]);
    if (strcmp(argv[1], "push-script") == 0 && argc == 5)
        return cmd_push_script(argv[2], argv[3], argv[4]);
    if (strcmp(argv[1], "pull-script") == 0 && argc == 6)
        return cmd_pull_script(argv[2], argv[3], argv[4], argv[5]);
    fprintf(stderr, "bad args\n");
    return 2;
}
