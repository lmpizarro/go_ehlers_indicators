package go_ehlers_indicators_test

import (
	"github.com/lmpizarro/go_timeseries_generator"
	ei "github.com/lmpizarro/go_ehlers_indicators"
	"testing"
)

func TestFAMAGraph(t *testing.T) {
	vals := go_timeseries_generator.GaussianProcess(1024)

	fama := ei.FAMADefault(vals)
	filename := "img/fama.png"
	err := ei.Plt(fama, filename)
	if err != nil {
		t.Error(err)
	}
}
