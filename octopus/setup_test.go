package octopus

import (
	"testing"

	"golang.org/x/time/rate"
)

func TestValidateRateLimits(t *testing.T) {
	tests := []struct {
		name      string
		rate      rate.Limit
		burst     rate.Limit
		wantPanic bool
	}{
		{name: "disabled", rate: -1, burst: -1, wantPanic: false},
		{name: "valid", rate: 5, burst: 8, wantPanic: false},
		{name: "zero rate", rate: 0, burst: 1, wantPanic: true},
		{name: "burst lower than rate", rate: 8, burst: 5, wantPanic: true},
		{name: "negative rate with positive burst", rate: -1, burst: 1, wantPanic: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				recovered := recover()
				if tt.wantPanic && recovered == nil {
					t.Fatal("expected panic")
				}
				if !tt.wantPanic && recovered != nil {
					t.Fatalf("unexpected panic: %v", recovered)
				}
			}()

			validateRateLimits(tt.rate, tt.burst)
		})
	}
}

func TestSetupTimeToQuitPanicsOnInvalidValue(t *testing.T) {
	o := NewWithDefaultOptions()
	o.TimeToQuit = 0

	defer func() {
		if recover() == nil {
			t.Fatal("expected setupTimeToQuit to panic when TimeToQuit <= 0")
		}
	}()
	o.setupTimeToQuit()
}
