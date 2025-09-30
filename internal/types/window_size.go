package types

import (
	ierr "github.com/flexprice/flexprice/internal/errors"
	"github.com/samber/lo"
)

// WindowSize defines the time window for aggregating usage data.
//
// Supported window sizes and their behavior:
// - MINUTE: 1-minute windows (00:00:00 to 00:00:59)
// - 15MIN: 15-minute windows (00:00:00 to 00:14:59, 00:15:00 to 00:29:59, etc.)
// - 30MIN: 30-minute windows (00:00:00 to 00:29:59, 00:30:00 to 00:59:59, etc.)
// - HOUR: 1-hour windows (00:00:00 to 00:59:59)
// - 3HOUR: 3-hour windows (00:00:00 to 02:59:59, 03:00:00 to 05:59:59, etc.)
// - 6HOUR: 6-hour windows (00:00:00 to 05:59:59, 06:00:00 to 11:59:59, etc.)
// - 12HOUR: 12-hour windows (00:00:00 to 11:59:59, 12:00:00 to 23:59:59)
// - DAY: 1-day windows (00:00:00 to 23:59:59 of the same day)
// - WEEK: 1-week windows (Monday 00:00:00 to Sunday 23:59:59)
// - MONTH: 1-month windows (1st 00:00:00 to last day 23:59:59 of the same month)
//
// Special behavior for MONTH window size:
// - When used with BillingAnchor: Creates custom monthly periods (e.g., 5th to 5th of each month)
// - When used without BillingAnchor: Uses standard calendar months (1st to 1st of each month)
// - All other window sizes ignore BillingAnchor and use standard calendar-based windows
type WindowSize string

// Note: keep values up to date in the meter package
const (
	WindowSizeMinute WindowSize = "MINUTE"
	WindowSize15Min  WindowSize = "15MIN"
	WindowSize30Min  WindowSize = "30MIN"
	WindowSizeHour   WindowSize = "HOUR"
	WindowSize3Hour  WindowSize = "3HOUR"
	WindowSize6Hour  WindowSize = "6HOUR"
	WindowSize12Hour WindowSize = "12HOUR"
	WindowSizeDay    WindowSize = "DAY"
	WindowSizeWeek   WindowSize = "WEEK"
	WindowSizeMonth  WindowSize = "MONTH"
)

func (w WindowSize) Validate() error {
	if w == "" {
		return nil
	}

	if !lo.Contains([]WindowSize{
		WindowSizeMinute,
		WindowSize15Min,
		WindowSize30Min,
		WindowSizeHour,
		WindowSize3Hour,
		WindowSize6Hour,
		WindowSize12Hour,
		WindowSizeDay,
		WindowSizeWeek,
		WindowSizeMonth,
	}, w) {
		return ierr.NewError("invalid window size").
			WithHint("Invalid window size").
			WithReportableDetails(
				map[string]any{
					"window_size": w,
				},
			).
			Mark(ierr.ErrValidation)
	}

	return nil
}
