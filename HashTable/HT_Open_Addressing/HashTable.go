package ht

type HashTable struct {
	Size  int
	Table []string
}

func CretaeTable(size int) *HashTable {
	return &HashTable{Size: size, Table: make([]string, size)}
}

func (ht *HashTable) Hash(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash += int(key[i])
	}
	return hash % ht.Size
}

// Open Addresssing - Liner Probing
func (ht *HashTable) Insert(key string) {
	index := ht.Hash(key)
	for ht.Table[index] != "" {
		index = (index + 1) % ht.Size //if index is occupied, move to the next index
	}
	ht.Table[index] = key
}

func (ht *HashTable) Search(key string) int {
	index := ht.Hash(key)
	for ht.Table[index] != "" {
		if ht.Table[index] == key {
			println("Key found at index:", index)
			return index
		}
		index = (index + 1) % ht.Size
	}
	// Key not found
	return -1
}

func (ht *HashTable) Delete(key string) {
	index := ht.Hash(key)
	for ht.Table[index] != "" {
		if ht.Table[index] == key {
			ht.Table[index] = ""
			println("Key deleted from index:", index)
			return
		}
		index = (index + 1) % ht.Size
	}
	println("Key not found for deletion")
}

func (ht *HashTable) Display() {
	for i, v := range ht.Table {
		if v != "" {
			println("Index:", i, "Value:", v)
		}
	}
}
