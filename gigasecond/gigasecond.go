package gigasecond

import (
	"math"
	"time"
)

func AddGigasecond(t time.Time) time.Time {
	gigaSec := time.Duration(math.Pow10(9))
	return t.Add(gigaSec * time.Second)
}
