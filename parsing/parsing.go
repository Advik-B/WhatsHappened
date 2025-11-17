package parsing

import (
	"bufio"
	"regexp"
	"strings"

	"github.com/Advik-B/WhatsHappened/models"
)

const timeBreak = string(rune(0x202F))

var reUser = regexp.MustCompile(`(?i)^(\d{2}/\d{2}/\d{2}),\s+(\d{1,2}:\d{2})(?:` + regexp.QuoteMeta(timeBreak) + `)?\s*(am|pm)\s*-\s*([^:]+?):\s*(.*)$`)
var reSystem = regexp.MustCompile(`(?i)^(\d{2}/\d{2}/\d{2}),\s+(\d{1,2}:\d{2})(?:` + regexp.QuoteMeta(timeBreak) + `)?\s*(am|pm)\s*-\s*(.*)$`)

func ParseChat(multilineChat string) ([]models.ParsedMessage, error) {
	// Case-insensitive regexes — (?i) at the start
	//parseTime := func(date, clock, ampm string) (time.Time, error) {
	//	return time.ParseInLocation(
	//		timeFormat,
	//		date+", "+clock+timeBreak+strings.ToLower(ampm),
	//		time.Local,
	//	)
	//}

	scanner := bufio.NewScanner(strings.NewReader(multilineChat))
	var msgs []models.ParsedMessage
	var cur *models.ParsedMessage

	for scanner.Scan() {
		line := strings.TrimRight(scanner.Text(), "\r\n")

		if m := reUser.FindStringSubmatch(line); m != nil {
			if cur != nil {
				msgs = append(msgs, *cur)
			}
			cur = &models.ParsedMessage{}
			t, err := parseTimeFixed(m[1], m[2], m[3])
			if err != nil {
				return nil, err
			}
			cur.Time = t
			cur.Sender = strings.TrimSpace(m[4])
			cur.Content = m[5]
			continue
		}

		if m := reSystem.FindStringSubmatch(line); m != nil {
			if reUser.MatchString(line) {
				// system regex must NOT override user regex
				continue
			}
			if cur != nil {
				msgs = append(msgs, *cur)
			}
			cur = &models.ParsedMessage{}
			t, err := parseTimeFixed(m[1], m[2], m[3])
			if err != nil {
				return nil, err
			}
			cur.Time = t
			cur.Sender = "" // system message
			cur.Content = m[4]
			continue
		}

		// continuation line
		if cur != nil {
			cur.Content += "\n" + line
		}
	}

	if cur != nil {
		msgs = append(msgs, *cur)
	}
	return msgs, scanner.Err()
}
