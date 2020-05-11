package utils

import "testing"

func TestGenNonce(t *testing.T) {
	for i := 0; i < 10; i++ {
		n := GenNonce(15, 88)
		if n < 15 || n > 88 {
			t.Errorf("gen a wrong random number: %d, want 15 to 88", n)
		}
		t.Logf("gen a new random number: %d", n)
	}
}

func TestRandomBytes(t *testing.T) {
	for i := 0; i < 10; i++ {
		b := RandomBytes(64)
		if len(b) != 64 {
			t.Errorf("gen a worng random bytes: %s, want lenth 64, got %d", string(b), len(b))
		}
		t.Logf("gen a new random string: %s", string(b))
	}
}
