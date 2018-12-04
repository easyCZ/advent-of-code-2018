package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestGuard_MostSleptMinute(t *testing.T) {
	scenarios := []struct {
		id        int64
		intervals []string
		expected  int
	}{
		{
			id: 10,
			intervals: []string{
				"1518-11-01 00:05,1518-11-01 00:25",
				"1518-11-01 00:30,1518-11-01 00:55",
				"1518-11-03 00:24,1518-11-03 00:29",
			},
			expected: 24,
		},
		{
			id: 99,
			intervals: []string{
				"1518-11-02 00:40,1518-11-02 00:50",
				"1518-11-04 00:36,1518-11-04 00:46",
				"1518-11-05 00:45,1518-11-05 00:55",
			},
			expected: 45,
		},
	}

	for _, scenario := range scenarios {
		intervals := make([]*Interval, len(scenario.intervals))
		for i, rawInterval := range scenario.intervals {
			tokens := strings.Split(rawInterval, ",")
			from, err := time.Parse(dateLayout, tokens[0])
			assert.NoError(t, err)

			to, err := time.Parse(dateLayout, tokens[1])
			assert.NoError(t, err)
			intervals[i] = &Interval{
				from: from,
				to:   to,
			}
		}
		guard := &Guard{
			id:     scenario.id,
			sleeps: intervals,
		}
		_, minute := guard.mostSleptMinute()
		assert.Equal(t, scenario.expected, minute)
	}
}
