package formattime

import "time"

func CurrentTime() string {

	// First, we create an instance of a timezone location object
	loc, _ := time.LoadLocation("Asia/Kolkata")
	const layout = "Jan 2, 2006 at 3:04pm (MST)"

	t, err := time.ParseInLocation(layout, time.Now().Format(layout),  loc)
	if err != nil {
		return ""
	}

	return t.Format(layout)
}