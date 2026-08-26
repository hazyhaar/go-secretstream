/* One-shot layout probe for crypto_secretstream_xchacha20poly1305_state.
 * Expected: sizeof == 52, offsetof(nonce) == 32 (libsodium 1.0.18).
 */
#include <stdio.h>
#include <stddef.h>
#include <sodium.h>

int main(void) {
    printf("sizeof=%zu offsetof_nonce=%zu offsetof_k=%zu\n",
           sizeof(crypto_secretstream_xchacha20poly1305_state),
           offsetof(crypto_secretstream_xchacha20poly1305_state, nonce),
           offsetof(crypto_secretstream_xchacha20poly1305_state, k));
    return 0;
}
