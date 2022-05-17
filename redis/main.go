package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/renatobrittoaraujo/learn/redis/redis_helper"
)

const (
	redisError = "redis: nil"
)

type Combinatorics struct {
	rdb redis_helper.Redis
}

var cache_hits uint64 = 0

func (c *Combinatorics) Combinations(n, r uint64) uint64 {
	keyname := fmt.Sprintf("%d %d", n, r)
	sval, err := c.rdb.Read(keyname)
	if err == nil {
		val, err := strconv.ParseUint(sval, 10, 64)
		if err == nil {
			cache_hits++
			return val
		}
	}

	if n <= 1 {
		return 1
	}
	if r <= 1 {
		return n
	}
	var val uint64
	if rand.Int()&1 == 1 {
		val = (c.Combinations(n-1, r) + c.Combinations(n, r-1)) % (1e9 + 7)
	} else {
		val = (c.Combinations(n, r-1) + c.Combinations(n-1, r)) % (1e9 + 7)
	}
	c.rdb.Write(keyname, fmt.Sprintf("%d", val))
	return val
}

func main() {
	ctx := context.Background()
	redis := redis_helper.NewRedis(ctx)
	c := &Combinatorics{
		redis,
	}

	for i := 0; i < 1000; i++ {
		x := uint64(rand.Intn(10000))
		y := uint64(rand.Intn(10000))
		t := c.Combinations(x, y)
		fmt.Println("RESULT FOR", x, y, "=", t)
		fmt.Println("HITS:", cache_hits)
	}

}
