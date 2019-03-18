package gigasecond

// import path for the time package from the standard library
import (
	"math"
	"time"
)

const testVersion = 4

func AddGigasecond(in time.Time) time.Time {
	return in.Add(time.Duration(math.Pow(10.0, 9.0)) * time.Second)
}
