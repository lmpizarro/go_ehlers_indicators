package go_ehlers_indicators

import (
	"github.com/lmpizarro/go_timeseries_generator"
	ei "github.com/lmpizarro/go_ehlers_indicators"
	"testing"
)

func TestZeroLagGraph(t *testing.T) {
	vals := go_timeseries_generator.GaussianProcess(1024)
	zl := ei.ZeroLagDefault(vals, 16)

	filename := "img/zero_lag.png"
	err := ei.Plt(zl, filename)
	if err != nil {
		t.Error(err)
	}
}

// TestZeroLagGraphStep tests the zero lag filter on a step function by graphing it
func TestZeroLagGraphStep(t *testing.T) {
	vals := go_timeseries_generator.StepFunction(1024, 500, 100)

	filename := "img/zero_lag_step.png"
	err := ei.Plt(vals, filename)
	if err != nil {
		t.Error(err)
	}
}
