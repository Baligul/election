package formattime

import "time"

func CurrentTime() string {

	// First, we create an instance of a timezone location object
	const layout = "Jan 2, 2006 at 3:04 PM"

	return time.Now().Format(layout)
}
