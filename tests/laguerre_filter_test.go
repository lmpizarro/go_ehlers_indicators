package go_ehlers_indicators_test

import (
	"github.com/lmpizarro/go_timeseries_generator"
	ei "github.com/lmpizarro/go_ehlers_indicators"
	"testing"
)

func TestLaguerreFilter(t *testing.T) {
	vals := go_timeseries_generator.GaussianProcess(1024)
	l := ei.LaguerreFilterDefault(vals)

	filename := "img/laguerre_filter.png"
	err := ei.Plt(l, filename)
	if err != nil {
		t.Error(err)
	}
}
