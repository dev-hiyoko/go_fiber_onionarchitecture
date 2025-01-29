package util

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/oklog/ulid/v2"
)

type ULID string

var defaultEntropySource *ulid.MonotonicEntropy

func init() {
	defaultEntropySource = ulid.Monotonic(rand.Reader, 0)
}

func NewULID() ULID {
	return ULID(fmt.Sprint(ulid.MustNew(ulid.Timestamp(time.Now()), defaultEntropySource)))
}

func (u ULID) String() string {
	return string(u)
}
