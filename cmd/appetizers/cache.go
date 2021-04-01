package main

import (
	"container/list"
	"fmt"
	"strings"
)

type Element interface{}

type Cache struct {
	Size     int
	Checker  map[interface{}]*list.Element
	InMemory *list.List
}

func (c *Cache) AddInt32(e Element) {
	key := e.(int32)
	if _, present := c.Checker[key]; !present {
		if c.InMemory.Len() == c.Size {
			lru := c.InMemory.Back()
			value := c.InMemory.Remove(lru)
			delete(c.Checker, value)
			// TODO: Notify eviction - fmt.Println("evicted lru:", v)
		}
		c.Checker[key] = c.InMemory.PushFront(key)
		// TODO: Notify add - fmt.Println("evicted lru:", v)
	}
}

func (c *Cache) GetInt32(e Element) bool {
	key := e.(int32)
	if mru, present := c.Checker[key]; present {
		value := c.InMemory.Remove(mru)
		c.Checker[value] = c.InMemory.PushFront(value)
		return true
	}
	return false
}

func (c *Cache) Snapshot() string {
	var sb strings.Builder
	if c.InMemory.Len() == 0 {
		return "cache is empty"
	}
	for node := c.InMemory.Front(); node != nil; node = node.Next() {
		sb.WriteString(fmt.Sprintf("[%v]", node.Value))
	}
	return sb.String()
}
