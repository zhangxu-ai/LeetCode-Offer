package main

import (
	"errors"
	"sync"
)

const (
	NullKey  = -32768
	HashSize = 100
)

type HashTable struct {
	data  [HashSize]interface{}
	count int
	mx    sync.Mutex
}

func NewHashTable() *HashTable {
	da := new([HashSize]interface{})
	for i := 0; i < HashSize; i++ {
		da[i] = NullKey
	}
	ht := &HashTable{
		data:  *da,
		count: 0,
	}
	return ht
}

func (ht *HashTable) Hash(key int) int {
	return key % HashSize
}

func (ht *HashTable) Insert(key int) error {
	if ht.count == HashSize {
		return errors.New("no free space")
	}
	addr := ht.Hash(key)
	//如果不为空，则说明冲突了
	for ht.data[addr] != NullKey {
		//开放定址法
		addr = (addr + 1) % HashSize
	}
	//直到有空位后插入关键字
	ht.mx.Lock()
	ht.data[addr] = key
	ht.count++
	ht.mx.Unlock()
	return nil
}
func (ht *HashTable) Search(key int) (int, error) {
	addr := ht.Hash(key)
	a := addr
	//不相等，则发生了冲突
	for ht.data[addr] != key {
		//求下一个地址
		addr = (addr + 1) % HashSize
		//如果下一个地址为空，说明这个key从未插入过
		//或者有到头部了，说明找了一圈都没找到
		if ht.data[addr] == NullKey || addr == a {
			return -1, errors.New("no key")
		}
	}
	return addr, nil
}
