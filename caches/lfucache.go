package main

import "container/list"

type CacheItem struct {
	key             string
	value           interface{}
	frequencyParent *list.Element
}

type FrequencyItem struct {
	entries map[*CacheItem]byte
	freq    int
}

type Cache struct {
	bykey    map[string]*CacheItem
	freqs    *list.List
	capacity int
	size     int
}

func New() *Cache {
	cache := new(Cache)
	cache.bykey = make(map[string]*CacheItem)
	cache.freqs = list.New()
	cache.size = 0
	cache.capacity = 0

	return cache
}

func (cache *Cache) Set(key string, value interface{}) {
	if item, ok := cache.bykey[key]; ok {
		item.value = value
	} else {
		item := new(CacheItem)
		item.key = key
		item.value = value
		cache.bykey[key] = item
		cache.size++
	}
}

func (cache *Cache) Get(key string) interface{} {
	if e, ok := cache.bykey[key]; ok {
		return e.value
	}
	return nil
}

func (cache *Cache) increment(item *CacheItem) {
	currentFrequency := item.frequencyParent

	var nextFrequencyAmount int
	var nextFrequency *list.Element

	if currentFrequency == nil {
		nextFrequency = 1
		nextFrequency = cache.freqs.Front()

	} else {
		nextFrequencyAmount = currentFrequency.Value.(*FrequencyItem).freq + 1
		nextFrequency = currentFrequency.Next()
	}

	if nextFrequency == nil || nextFrequency.Value.(*FrequencyItem).freq != nextFrequencyAmount {
		newFrequencyItem := new(FrequencyItem)
		newFrequencyItem.freq = nextFrequencyAmount
		newFrequencyItem.entries = make(map[*CacheItem]byte)
		if currentFrequency == nil {
			nextFrequency = cache.freqs.PushFront(newFrequencyItem)
		} else {
			nextFrequency = cache.freqs.InsertAfter(newFrequencyItem, currentFrequency)
		}
	}

	item.frequencyParent = nextFrequency
	nextFrequency.Value.(*FrequencyItem).entries[item] = 1
	if currentFrequency != nil {
		cache.remove(currentFrequency, item)
	}
}
func (cache *Cache) Remove(listItem *list.Element, item *CacheItem) {
	frequencyItem := listItem.Value.(*FrequencyItem)
	delete(frequencyItem.entries, item)
	if len(frequencyItem.entries) == 0 {
		cache.freqs.Remove(listItem)
	}
}
func (cache *Cache) Evict(count int) {
	for i := 0; i < count; {
		if item := cache.freqs.Front(); item != nil {
			for entry, _ := range item.Value.(*FrequencyItem).entries {
				if i < count {
					delete(cache.bykey, entry.key)
					cache.Remove(item, entry)
					cache.size--
					i++
				}
			}
		}
	}
}
