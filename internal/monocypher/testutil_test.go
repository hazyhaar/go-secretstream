package monocypher_test

func keyNonce() (key, nonce []byte) {
	key = make([]byte, 32)
	nonce = make([]byte, 24)
	for i := range key {
		key[i] = byte(i + 1)
	}
	for i := range nonce {
		nonce[i] = byte(i + 10)
	}
	return key, nonce
}

func mkPT(n int, pattern string) []byte {
	pt := make([]byte, n)
	switch pattern {
	case "zero":
		// leave zeros
	case "i%251":
		for i := range pt {
			pt[i] = byte(i % 251)
		}
	case "(i*17+3)%251":
		for i := range pt {
			pt[i] = byte((i*17 + 3) % 251)
		}
	default:
		for i := range pt {
			pt[i] = byte((i*17 + 3) % 251)
		}
	}
	return pt
}
