package secret

var signs = []string{"wink", "double blink", "close your eyes", "jump"}

// Handshake returns the sequence of events for the given code
func Handshake(code uint) []string {
	reversed := false
	if code&16 == 16 {
		reversed = true
	}

	var events []string
	for i := 0; i < 4; i++ {
		if code>>uint(i)&1 == 1 {
			events = add(events, signs[i], reversed)
		}
	}

	return events
}

func add(events []string, event string, reversed bool) []string {
	if reversed {
		return append([]string{event}, events...)
	}
	return append(events, event)
}
