package availability

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNextQuarter(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected string
	}{
		{
			name:     "Q1 date (January) returns Q2 of the same year",
			input:    time.Date(2026, time.January, 10, 0, 0, 0, 0, time.UTC),
			expected: "Q2 2026",
		},
		{
			name:     "Q1 date boundary (March 31st) returns Q2 of the same year",
			input:    time.Date(2026, time.March, 31, 23, 59, 59, 0, time.UTC),
			expected: "Q2 2026",
		},
		{
			name:     "Q2 date (May) returns Q3 of the same year",
			input:    time.Date(2026, time.May, 15, 0, 0, 0, 0, time.UTC),
			expected: "Q3 2026",
		},
		{
			name:     "Q3 date (August) returns Q4 of the same year",
			input:    time.Date(2026, time.August, 1, 0, 0, 0, 0, time.UTC),
			expected: "Q4 2026",
		},
		{
			name:     "Q4 date (October) wraps around to Q1 of the next year",
			input:    time.Date(2026, time.October, 1, 0, 0, 0, 0, time.UTC),
			expected: "Q1 2027",
		},
		{
			name:     "Q4 date boundary (December 31st) wraps around to Q1 of the next year",
			input:    time.Date(2026, time.December, 31, 23, 59, 59, 0, time.UTC),
			expected: "Q1 2027",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NextQuarter(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
