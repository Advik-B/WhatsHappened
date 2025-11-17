package parsing

import (
	"fmt"
	"strings"
	"time"
)

func normalizeTimestamp(date, clock, ampm string) string {
	// WhatsApp may use U+202F or U+00A0 — normalize both
	clock = strings.ReplaceAll(clock, "\u202F", " ")
	clock = strings.ReplaceAll(clock, "\u00A0", " ")

	// Ensure exactly one normal space before am/pm
	return fmt.Sprintf("%s, %s %s", date, clock, strings.ToUpper(ampm))
}

func parseTimeFixed(date, clock, ampm string) (time.Time, error) {
	normalized := normalizeTimestamp(date, clock, ampm)
	layout := "02/01/06, 03:04 PM"
	return time.ParseInLocation(layout, normalized, time.Local)
}
