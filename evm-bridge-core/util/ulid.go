package util

import (
	"math/rand"
	"sync"
	"time"

	"github.com/oklog/ulid/v2"
)

var mu sync.Mutex
var entropy = ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)

func ULID() string {
	mu.Lock()
	defer mu.Unlock()

	return ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
}
