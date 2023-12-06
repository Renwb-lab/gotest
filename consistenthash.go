package main

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
)

type Hash struct {
	hashKeys []uint32                         // hash后的数值
	mapping  map[uint32]string                // hash后数值和原本数值的影射
	hash     func(data []byte) uint32         // hash 方法
	replic   func(key string) (keys []string) // 将一个key映射到多个keys
}

func NewHash(
	hash func(data []byte) uint32,
	replic func(key string) (keys []string),
) *Hash {
	fn := hash
	if fn == nil {
		fn = crc32.ChecksumIEEE
	}
	if replic == nil {
		// 默认不映射多个虚拟节点
		replic = func(key string) (keys []string) {
			keys = append(keys, key)
			return
		}
	}
	return &Hash{
		hashKeys: make([]uint32, 0),
		mapping:  make(map[uint32]string),
		hash:     hash,
		replic:   replic,
	}
}

func (h *Hash) Add(keys ...string) {
	for _, key := range keys {
		for _, newKeys := range h.replic(key) {
			hashKey := h.hash([]byte(newKeys))       // 计算hash后的结果
			h.hashKeys = append(h.hashKeys, hashKey) // 添加到hashKeys中，
			h.mapping[hashKey] = key                 // 完成映射
		}
	}
	sort.Slice(h.hashKeys, func(i, j int) bool {
		return h.hashKeys[i] < h.hashKeys[j]
	})
}

func (h *Hash) Get(key string) string {
	hashKey := h.hash([]byte(key))
	idx := sort.Search(len(h.hashKeys), func(i int) bool {
		return h.hashKeys[i] >= hashKey
	})
	// 如果超出最后数值的话，则调整为0
	if idx == len(h.hashKeys) {
		idx = 0
	}
	// 找到新映射的HashKey
	newHashKey := h.hashKeys[idx]
	return h.mapping[newHashKey]
}

type ReplicaHash struct {
	replicFunc func(key string) (keys []string)
	Hash
}

func main156() {
	fn := func(data []byte) uint32 {
		i, _ := strconv.Atoi(string(data))
		return uint32(i)
	}
	hash := NewHash(fn, nil)
	hash.Add("1", "2", "5", "10")
	keys := []string{"1", "2", "3", "5", "6", "10", "11"}
	for i := range keys {
		fmt.Printf("the mapping result of %s is %s\n", keys[i], hash.Get(keys[i]))
	}

	replicFunc := func(key string) (keys []string) {
		num, _ := strconv.Atoi(string(key))
		multi := 1
		for i := 0; i < 3; i += 1 {
			keys = append(keys, strconv.Itoa(num*multi))
			multi *= 10
		}
		return
	}
	fmt.Println("======================")
	replicaHash := NewHash(fn, replicFunc)
	replicaHash.Add("1", "2", "5", "10")
	for i := range keys {
		fmt.Printf("the mapping result of %s is %s\n", keys[i], replicaHash.Get(keys[i]))
	}
}

// 执行结果：
// the mapping result of 1 is 1
// the mapping result of 2 is 2
// the mapping result of 3 is 5
// the mapping result of 5 is 5
// the mapping result of 6 is 10
// the mapping result of 10 is 10
// the mapping result of 11 is 1 // 区别项
// ======================
// the mapping result of 1 is 1
// the mapping result of 2 is 2
// the mapping result of 3 is 5
// the mapping result of 5 is 5
// the mapping result of 6 is 10
// the mapping result of 10 is 10
// the mapping result of 11 is 2 // 区别项
