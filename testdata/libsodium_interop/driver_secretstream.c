/* CLI driver: libsodium crypto_secretstream_xchacha20poly1305
 * Commands:
 *   push-stream <key.bin> <plain.bin> <wire.bin>
 *   pull-stream <key.bin> <wire.bin> <plain.out>
 * Exit 1 on crypto failure.
 */
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sodium.h>

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

int main(int argc, char **argv) {
    if (sodium_init() < 0) { fprintf(stderr, "sodium_init\n"); return 1; }
    if (argc < 2) {
        fprintf(stderr, "usage: %s push-stream|pull-stream ...\n", argv[0]);
        return 2;
    }
    if (strcmp(argv[1], "push-stream") == 0 && argc == 5)
        return cmd_push(argv[2], argv[3], argv[4]);
    if (strcmp(argv[1], "pull-stream") == 0 && argc == 5)
        return cmd_pull(argv[2], argv[3], argv[4]);
    fprintf(stderr, "bad args\n");
    return 2;
}
