package gigasecond

import "time"

// AddGigasecond adds one billion seconds to t.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}
