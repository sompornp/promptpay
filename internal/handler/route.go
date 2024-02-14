package handler

import "somporn/promptpay/internal"

type Route struct {
	// simple cache implementation; in a real-world scenario, this would be a Redis or Memcached instance
	Cache           *internal.Cache
	TimeoutInSecond int64
}
