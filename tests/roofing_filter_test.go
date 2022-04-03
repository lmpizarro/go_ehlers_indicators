package go_ehlers_indicators_test

import (
	"github.com/lmpizarro/go_timeseries_generator"
	ei "github.com/lmpizarro/go_ehlers_indicators"
	"testing"
)

func TestRoofingFilter(t *testing.T) {
	vals := go_timeseries_generator.GaussianProcess(1024)
	filt := ei.RoofingFilter(vals)

	filename := "img/roofing_filter.png"
	err := ei.Plt(filt, filename)
	if err != nil {
		t.Error(err)
	}
}
