package stun

import (
	"testing"
)

func TestXORSafe(t *testing.T) {
	dst := make([]byte, 8)
	a := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	b := []byte{8, 7, 7, 6, 6, 3, 4, 1}
	safeXORBytes(dst, a, b)
	safeXORBytes(dst, dst, a)
	for i, v := range dst {
		if b[i] != v {
			t.Error(b[i], "!=", v)
		}
	}
}

func TestXORSafeBSmaller(t *testing.T) {
	dst := make([]byte, 5)
	a := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	b := []byte{8, 7, 7, 6, 6}
	safeXORBytes(dst, a, b)
	safeXORBytes(dst, dst, a)
	for i, v := range dst {
		if b[i] != v {
			t.Error(b[i], "!=", v)
		}
	}
}

func TestXORFast(t *testing.T) {
	if !supportsUnaligned {
		t.Skip("No support for unaligned operations.")
	}
	dst := make([]byte, 8)
	a := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	b := []byte{8, 7, 7, 6, 6, 3, 4, 1}
	xorBytes(dst, a, b)
	xorBytes(dst, dst, a)
	for i, v := range dst {
		if b[i] != v {
			t.Error(b[i], "!=", v)
		}
	}
}

func TestXORFastBSmaller(t *testing.T) {
	if !supportsUnaligned {
		t.Skip("No support for unaligned operations.")
	}
	dst := make([]byte, 5)
	a := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	b := []byte{8, 7, 7, 6, 6}
	xorBytes(dst, a, b)
	xorBytes(dst, dst, a)
	for i, v := range dst {
		if b[i] != v {
			t.Error(b[i], "!=", v)
		}
	}
}

func TestXORFallback(t *testing.T) {
	if !supportsUnaligned {
		t.Skip("No support for unaligned operations.")
	}
	defer func() {
		supportsUnaligned = true
	}()
	supportsUnaligned = false
	dst := make([]byte, 5)
	a := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	b := []byte{8, 7, 7, 6, 6}
	xorBytes(dst, a, b)
	xorBytes(dst, dst, a)
	for i, v := range dst {
		if b[i] != v {
			t.Error(b[i], "!=", v)
		}
	}
}
