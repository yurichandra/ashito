package helper

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func FormatDuration(d time.Duration) string {
	if d == 0 {
		return "0s"
	}

	// Handle negative durations
	negative := d < 0
	if negative {
		d = -d
	}

	// Convert to different units
	hours := math.Floor(d.Hours())
	d -= time.Duration(hours) * time.Hour
	minutes := math.Floor(d.Minutes())
	d -= time.Duration(minutes) * time.Minute
	seconds := math.Floor(d.Seconds())
	d -= time.Duration(seconds) * time.Second
	milliseconds := math.Floor(float64(d.Milliseconds()))
	d -= time.Duration(milliseconds) * time.Millisecond
	microseconds := math.Floor(float64(d.Microseconds()))
	d -= time.Duration(microseconds) * time.Microsecond
	nanoseconds := d.Nanoseconds()

	// Build the parts array
	var parts []string
	if hours > 0 {
		parts = append(parts, fmt.Sprintf("%.0fh", hours))
	}
	if minutes > 0 {
		parts = append(parts, fmt.Sprintf("%.0fm", minutes))
	}
	if seconds > 0 {
		parts = append(parts, fmt.Sprintf("%.0fs", seconds))
	}
	if milliseconds > 0 && len(parts) < 2 {
		parts = append(parts, fmt.Sprintf("%.0fms", milliseconds))
	}
	if microseconds > 0 && len(parts) < 2 {
		parts = append(parts, fmt.Sprintf("%.0fµs", microseconds))
	}
	if nanoseconds > 0 && len(parts) < 2 {
		parts = append(parts, fmt.Sprintf("%dns", nanoseconds))
	}

	// Handle zero case
	if len(parts) == 0 {
		return "0s"
	}

	// Join parts and handle negative sign
	result := strings.Join(parts[:min(2, len(parts))], " ")
	if negative {
		return "-" + result
	}
	return result
}
