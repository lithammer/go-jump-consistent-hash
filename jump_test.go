package jump

import "testing"

func TestHashInBucketRange(t *testing.T) {
	h := Hash(1, 1)
	if h != 0 {
		t.Error("expected bucket to be 0, got", h)
	}

	h = Hash(42, 57)
	if h != 43 {
		t.Error("expected bucket to be 43, got", h)
	}

	h = Hash(0xDEAD10CC, 1)
	if h != 0 {
		t.Error("expected bucket to be 0, got", h)
	}

	h = Hash(0xDEAD10CC, 666)
	if h != 361 {
		t.Error("expected bucket to be 361, got", h)
	}

	h = Hash(256, 1024)
	if h != 520 {
		t.Error("expected bucket to be 520, got", h)
	}

}

func TestNegativeBucket(t *testing.T) {
	h := Hash(0, -10)
	if h != 0 {
		t.Error("expected bucket to be 0, got", h)
	}

	h = Hash(0xDEAD10CC, -666)
	if h != 0 {
		t.Error("expected bucket to be 0, got", h)
	}
}

func ExampleHash() {
	Hash(256, 1024)
	// Output: 520
}
