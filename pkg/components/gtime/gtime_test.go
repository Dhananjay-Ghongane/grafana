package gtime

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestParseInterval(t *testing.T) {
	now := time.Now()

	tcs := []struct {
		interval string
		duration time.Duration
		err      string
	}{
		{interval: "1d", duration: now.Sub(now.AddDate(0, 0, -1))},
		{interval: "1w", duration: now.Sub(now.AddDate(0, 0, -7))},
		{interval: "2w", duration: now.Sub(now.AddDate(0, 0, -14))},
		{interval: "1M", duration: now.Sub(now.AddDate(0, -1, 0))},
		{interval: "1y", duration: now.Sub(now.AddDate(-1, 0, 0))},
		{interval: "5y", duration: now.Sub(now.AddDate(-5, 0, 0))},
		{interval: "invalid-duration", err: "time: invalid duration invalid-duration"},
	}

	for i, tc := range tcs {
		tc := tc
		t.Run(fmt.Sprintf("testcase %d", i), func(t *testing.T) {
			res, err := ParseInterval(tc.interval)
			if tc.err == "" {
				require.NoError(t, err, "interval %q", tc.interval)
				require.Equal(t, tc.duration, res, "interval %q", tc.interval)
			} else {
				require.EqualError(t, err, tc.err, "interval %q", tc.interval)
			}
		})
	}
}
