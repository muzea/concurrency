package concurrency

import (
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounter1(t *testing.T) {
	var count = int32(0)
	var total = 100
	Run(func(index int) {
		atomic.AddInt32(&count, 1)
	}, 100, total)
	assert.Equal(t, int(count), total)
}

func TestCounter2(t *testing.T) {
	var count = int32(0)
	var total = 1000

	Run(func(index int) {
		atomic.AddInt32(&count, 1)
	}, 100, total)
	assert.Equal(t, int(count), total)
}

func BenchmarkConcurrency1Loop10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Run(func(index int) {}, 1, 10)
	}
}

func BenchmarkNormalLoop10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var cb = func(index int) {}
		for i := 0; i < 10; i++ {
			cb(i)
		}
	}
}

func BenchmarkConcurrency1000Loop10000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Run(func(index int) {}, 1000, 10000)
	}
}
