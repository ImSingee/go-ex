package init

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}
