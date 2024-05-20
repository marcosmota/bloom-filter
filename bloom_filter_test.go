package bloom_filter

import (
	"testing"
)

func TestNewBloomFilter(t *testing.T) {
	lengthElement := 100
	probability := 0.1
	bf := NewBloomFilter(lengthElement, probability)

	if bf.LenghtElement != lengthElement {
		t.Errorf("Expected LengthElement %d, got %d", lengthElement, bf.LenghtElement)
	}

	if bf.Probability != probability {
		t.Errorf("Expected Probability %f, got %f", probability, bf.Probability)
	}

	expectedNumberOfBits := 480
	if bf.NumberOfBits != expectedNumberOfBits {
		t.Errorf("Expected NumberOfBits %d, got %d", expectedNumberOfBits, bf.NumberOfBits)
	}

	expectedNumberOfHashFunctions := 3
	if bf.NumberOfHashFunctions != expectedNumberOfHashFunctions {
		t.Errorf("Expected NumberOfHashFunctions %d, got %d", expectedNumberOfHashFunctions, bf.NumberOfHashFunctions)
	}

	if len(bf.bitArray) != expectedNumberOfBits {
		t.Errorf("Expected bitArray length %d, got %d", expectedNumberOfBits, len(bf.bitArray))
	}
}

func TestBloomFilter_PutAndGet(t *testing.T) {
	bf := NewBloomFilter(100, 0.01)
	item := "test_item"

	// Test Lookup for an item that hasn't been added
	if bf.Lookup(item) {
		t.Errorf("Expected Lookup to return false for an item that hasn't been added")
	}

	// Add the item and test Lookup again
	bf.Put(item)
	if !bf.Lookup(item) {
		t.Errorf("Expected Lookup to return true for an item that has been added")
	}
}

func TestBloomFilter_FalsePositiveRate(t *testing.T) {
	bf := NewBloomFilter(1000, 0.01)
	insertedItems := []string{"item1", "item2", "item3", "item4", "item5"}

	// Add items to the Bloom filter
	for _, item := range insertedItems {
		bf.Put(item)
	}

	// Check for a non-inserted item
	nonInsertedItem := "non_inserted_item"
	if bf.Lookup(nonInsertedItem) {
		t.Errorf("Expected Lookup to return false for an item that hasn't been added")
	}

	// You can also add more tests here to statistically evaluate the false positive rate if desired
}

func TestBloomFilter_InsertAndCheckMultipleItems(t *testing.T) {
	bf := NewBloomFilter(1000, 0.01)
	items := []string{"apple", "banana", "cherry", "date", "elderberry"}

	// Add items to the Bloom filter
	for _, item := range items {
		bf.Put(item)
	}

	// Check each item
	for _, item := range items {
		if !bf.Lookup(item) {
			t.Errorf("Expected Lookup to return true for an item that has been added: %s", item)
		}
	}

	// Check a non-inserted item
	nonInsertedItem := "fig"
	if bf.Lookup(nonInsertedItem) {
		t.Errorf("Expected Lookup to return false for a non-inserted item: %s", nonInsertedItem)
	}
}
