package utl

import "time"

// Milliseconds returns the millisecond integer portion of time.Time.
func Milliseconds(t time.Time) int {
	return t.Nanosecond() / int(time.Millisecond)
}
