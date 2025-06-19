package htchaining

type HashTable struct {
	size    int
	buckets []*Bucket
}

type Bucket struct { // Each bucket is a linked list
	Head *BucketNode
}

type BucketNode struct { // BucketNode represents a node in the linked list
	key  string
	next *BucketNode
}

func CreateTable(size int) *HashTable {
	table := HashTable{size: size, buckets: make([]*Bucket, size)}
	for index, _ := range table.buckets {
		table.buckets[index] = &Bucket{Head: nil} // Every index contains a Bucket/linked list
	}
	return &table
}

// HashTable methods
func (ht *HashTable) Hash(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash += int(key[i])
	}
	return hash % ht.size
}

func (ht *HashTable) Insert(key string) {
	index := ht.Hash(key)
	Bucket := ht.buckets[index]
	Bucket.Insert(key)
}

func (ht *HashTable) Delele(key string) {
	index := ht.Hash(key)
	Bucket := ht.buckets[index]
	Bucket.Delete(key)
}

func (ht *HashTable) Search(key string) bool {
	index := ht.Hash(key)
	Bucket := ht.buckets[index]
	currentNode := Bucket.Head

	for currentNode != nil {
		if currentNode.key == key {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (ht *HashTable) Display() {
	for index, bucket := range ht.buckets {
		print("Bucket", index, ":")
		currentNode := bucket.Head
		if currentNode == nil {
			println("  Empty")
			continue
		}

		for currentNode != nil {
			print(" ->", currentNode.key)
			currentNode = currentNode.next
		}
		println()
	}
}

// Bucket methods
func (bkt *Bucket) Insert(key string) {
	bucketnode := &BucketNode{key: key, next: nil}
	if bkt.Head == nil {
		bkt.Head = bucketnode
		return
	}

	bucketnode.next = bkt.Head
	bkt.Head = bucketnode
}

func (bkt *Bucket) Delete(key string) {

	if bkt.Head == nil {
		println("Not Found")
		return
	}

	if bkt.Head.key == key {
		bkt.Head = bkt.Head.next
		return
	}

	currentNode := bkt.Head
	var prev *BucketNode

	for currentNode != nil {
		if currentNode.key == key {
			prev.next = currentNode.next
			return
		}
		prev = currentNode
		currentNode = currentNode.next
	}
	println("Not Found")
}
