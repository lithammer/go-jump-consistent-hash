package jump

import (
	"hash"
	"hash/crc64"
	"hash/fnv"
	"io"
)

// Hash takes a 64 bit key and the number of buckets. It outputs a bucket
// number in the range [0, buckets).
func Hash(key uint64, buckets int32) int32 {
	var b, j int64

	if buckets <= 0 {
		buckets = 1
	}

	for j < int64(buckets) {
		b = j
		key = key*2862933555777941757 + 1
		j = int64(float64(b+1) * (float64(int64(1)<<31) / float64((key>>33)+1)))
	}

	return int32(b)
}

// Hasher is a subset of hash.Hash64 in the standard library.
type Hasher interface {
	// Write (via the embedded io.Writer interface) adds more data to the
	// running hash.
	// It never returns an error.
	io.Writer

	// Reset resets the Hasher to its initial state.
	Reset()

	// Return the result of the added bytes (via io.Writer).
	Sum64() uint64
}

// HashString takes string as key instead of an int and uses a Hasher to
// generate a key compatible with Hash().
func HashString(key string, buckets int32, h Hasher) int32 {
	h.Reset()
	_, err := io.WriteString(h, key)
	if err != nil {
		panic(err)
	}
	return Hash(h.Sum64(), buckets)
}

// Create some Hashers available in the standard library for use with
// HashString().
var (
	// CRC64 uses the 64-bit Cyclic Redundancy Check (CRC-64) with ECMA
	// polynomial.
	CRC64 hash.Hash64 = crc64.New(crc64.MakeTable(crc64.ECMA))
	// FNV1 uses the non-cryptographic hash function FNV-1
	FNV1 hash.Hash64 = fnv.New64()
	// FNV1a uses the non-cryptographic hash function FNV-1a
	FNV1a hash.Hash64 = fnv.New64a()
)
