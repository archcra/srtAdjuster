package helper

import (
	"fmt"
	"math"
	"strconv"
)

func Milliseconds2str(timestamp int) string {
	// return string should be format as: 00:00:10,000
	rMilliseconds := int(math.Mod(float64(timestamp), 1000))
	seconds := (timestamp - rMilliseconds) / 1000
	rSeconds := int(math.Mod(float64(seconds), 60))
	minutes := (seconds - rSeconds) / 60
	rMinutes := int(math.Mod(float64(minutes), 60))
	hours := int(math.Floor(float64(minutes-rMinutes))) / 60

	return fmt.Sprintf("%02d:%02d:%02d,%03d", hours, rMinutes, rSeconds, rMilliseconds)
}

func Str2milliseconds(timestamp string) int {
	// timestamp format as: 00:00:10,000

	hours, _ := strconv.Atoi(timestamp[0:2])
	minutes, _ := strconv.Atoi(timestamp[3:5])
	seconds, _ := strconv.Atoi(timestamp[6:8])
	milliseconds, _ := strconv.Atoi(timestamp[9:12])

	return (hours*3600+minutes*60+seconds)*1000 + milliseconds
}
