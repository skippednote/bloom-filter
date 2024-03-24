package bloom_filter

import (
	"math"

	"github.com/spaolacci/murmur3"
)

type BloomFilter struct {
	bit_vec    []bool
	num_bits   uint
	num_hashes uint
}

func (bf *BloomFilter) Hash(item string, seed uint32) uint64 {
	hash := murmur3.New64WithSeed(seed)
	hash.Write([]byte(string(item)))
	return hash.Sum64() % uint64(bf.num_bits)
}

func (bf *BloomFilter) Insert(item string) {
	for i := uint32(0); i < uint32(bf.num_hashes); i++ {
		index := bf.Hash(item, i) % uint64(bf.num_bits)
		bf.bit_vec[index] = true
	}
}

func (bf *BloomFilter) Contains(item string) bool {
	for i := uint32(0); i < uint32(bf.num_hashes); i++ {
		index := bf.Hash(item, i) % uint64(bf.num_bits)
		if !bf.bit_vec[index] {
			return false
		}
	}
	return true
}

func calculateNumberOfBits(items uint, falsePositiveRate float64) uint {
	n := -1 * float64(items) * math.Log(falsePositiveRate)
	d := math.Pow(math.Log(2), 2)
	return uint(math.Ceil(n / d))
}

func calculateNumberOfHashes(bits uint, items uint) uint {
	return uint(math.Ceil((float64(bits) / float64(items)) * math.Log(2)))
}

func NewBloomFilter(items uint, falsePositiveRate float64) *BloomFilter {
	bits := calculateNumberOfBits(items, falsePositiveRate)
	hashes := calculateNumberOfHashes(bits, items)
	return &BloomFilter{
		bit_vec:    make([]bool, int(bits)),
		num_bits:   bits,
		num_hashes: hashes,
	}
}
