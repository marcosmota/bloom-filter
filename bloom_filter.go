package bloom_filter

import (
	"github.com/spaolacci/murmur3"
	"math"
)

type BloomFilter struct {
	LenghtElement         int
	Probability           float64
	NumberOfBits          int
	NumberOfHashFunctions int
	bitArray              []bool
}

func NewBloomFilter(lenghtElement int, probability float64) BloomFilter {
	numberOfBits := _calculateNumberOfBits(probability, lenghtElement)
	numberOfHashFunctions := _calculateNumberOfHashFunctions(numberOfBits, lenghtElement)
	bitArray := make([]bool, numberOfBits)
	return BloomFilter{
		LenghtElement:         lenghtElement,
		Probability:           probability,
		NumberOfBits:          numberOfBits,
		NumberOfHashFunctions: numberOfHashFunctions,
		bitArray:              bitArray,
	}
}

func _calculateNumberOfBits(probability float64, lengthElements int) int {
	return int(math.Ceil(-(float64(lengthElements) * math.Log(probability)) / math.Pow(math.Ln2, 2)))
}

func _calculateNumberOfHashFunctions(numberOfBits int, lengthElements int) int {
	return int(math.Ceil(float64(numberOfBits/lengthElements) * math.Ln2))
}

func (b BloomFilter) Lookup(key string) bool {
	for i := 0; i < b.NumberOfHashFunctions; i++ {
		hash := murmur3.Sum32WithSeed([]byte(key), uint32(i))
		index := hash % uint32(b.NumberOfBits)
		if b.bitArray[index] == false {
			return false
		}
	}
	return true
}

func (b BloomFilter) Put(key string) {
	for i := 0; i < b.NumberOfHashFunctions; i++ {
		hash := murmur3.Sum32WithSeed([]byte(key), uint32(i))
		index := hash % uint32(b.NumberOfBits)
		b.bitArray[index] = true
	}
}
