package main

import (
	"errors"
	"fmt"
)

const HT_CAPACITY = 5

type KVNode struct {
	key   string
	value any
	next  *KVNode
}

func NewKVNode(key string, value any) *KVNode {
	return &KVNode{key: key, value: value}
}

type HashTable [HT_CAPACITY]*KVNode

func NewHashTable() *HashTable {
	return &HashTable{nil}
}

func Hash(seed string) int {
	return len(seed) % HT_CAPACITY
}

func (ht *HashTable) Put(key string, value any) {
	newNode := NewKVNode(key, value)
	index := Hash(key)

	if ht[index] == nil {
		ht[index] = newNode
	} else {
		curr := ht[index]
		for {
			// if key already exists, just override the value
			if curr.key == key {
				curr.value = value
				return
			}

			if curr.next == nil {
				break
			}

			curr = curr.next
		}
		curr.next = newNode
	}
}

func (ht *HashTable) Get(key string) (error, value any) {
	index := Hash(key)
	target := ht[index]

	for target != nil {
		if target.key == key {
			return nil, target.value
		}
		target = target.next
	}

	return errors.New("key not found"), nil
}

func HashTableDemo() {
	ht := NewHashTable()

	ht.Put("apple", 456)
	ht.Put("apple6", 123)
	ht.Put("apple2", 589)
	_, v := ht.Get("apple2")

	fmt.Println(v)
	ht.Put("apple2", 963)

	fmt.Println()
	fmt.Println(ht)
	fmt.Println(ht[0])
	fmt.Println(ht[0].next)
	fmt.Println(ht[1])
	fmt.Println(ht[1].next)
	fmt.Println(ht[1].next.next)

	_, v = ht.Get("apple2")

	fmt.Println(v)
}
