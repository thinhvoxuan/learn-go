package gigasecond

import "time"

// Constant declaration.
const testVersion = 4

// AddGigasecond return time when someone live 10^9 second
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(1e9 * 1e9))
}
