package bloom_filter

import "testing"

func TestBloomFilter(t *testing.T) {
	existing := []string{"Bassam", "Aliya", "Baheej", "Maryam"}
	missing := []string{"Afsa", "Aiza"}

	bf := NewBloomFilter(100, 3)
	for _, person := range existing {
		bf.Insert(person)
	}

	for _, person := range existing {
		if !bf.Contains(person) {
			t.Fatalf("Was expected %s to be in the Bloom Filter", person)
		}
	}

	for _, person := range missing {
		if bf.Contains(person) {
			t.Fatalf("Was not expected %s to be in the Bloom Filter", person)
		}
	}
}
