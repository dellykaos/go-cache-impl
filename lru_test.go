package cache_test

import (
	"fmt"
	"testing"

	"github.com/dellykaos/go-cache-impl"
	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	assert := assert.New(t)

	t.Run("with capacity", func(t *testing.T) {
		c := cache.NewLRUCache(3)
		c.Put("a", "satu")
		c.Put("b", "dua")
		assert.Equal("satu", c.Get("a"))
		c.Put("c", "tiga")
		assert.Equal("dua", c.Get("b"))
		assert.Equal("tiga", c.Get("c"))
		c.Put("d", "empat")
		assert.Equal("", c.Get("a"))
		assert.Equal("dua", c.Get("b"))
		assert.Equal("tiga", c.Get("c"))
		assert.Equal("empat", c.Get("d"))
		c.Put("d", "lima")
		assert.Equal("lima", c.Get("d"))
		c.Print()
	})

	t.Run("without capacity", func(t *testing.T) {
		c := cache.NewLRUCache(0)
		c.Put("a", "satu")
		assert.Equal("satu", c.Get("a"))
		for i := 0; i <= 1000; i++ {
			c.Put(fmt.Sprintf("%d", i), fmt.Sprintf("cache-%d", i))
		}
		assert.Equal("satu", c.Get("a"))
		assert.Equal("cache-1000", c.Get("1000"))
	})
}
