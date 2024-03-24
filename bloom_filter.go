package bloom_filter

import "github.com/spaolacci/murmur3"

type BloomFilter struct {
	bit_vec    []bool
	num_bits   uint
	num_hashes uint
}

func (bf *BloomFilter) Hash(item string, seed int) uint64 {
	hash := murmur3.New64WithSeed(uint32(seed))
	hash.Write([]byte(string(item)))
	return hash.Sum64() % uint64(bf.num_bits)
}

func (bf *BloomFilter) Insert(item string) {
	for i := 0; i < int(bf.num_hashes); i++ {
		index := bf.Hash(item, i) % uint64(bf.num_bits)
		bf.bit_vec[index] = true
	}
}

func (bf *BloomFilter) Contains(item string) bool {
	for i := 0; i < int(bf.num_hashes); i++ {
		index := bf.Hash(item, i) % uint64(bf.num_bits)
		if !bf.bit_vec[index] {
			return false
		}
	}
	return true
}

func NewBloomFilter(num_bits, num_hashes uint) BloomFilter {
	return BloomFilter{
		bit_vec:    make([]bool, num_bits),
		num_bits:   num_bits,
		num_hashes: num_hashes,
	}
}
