package go_ehlers_indicators_test

import (
	"github.com/lmpizarro/go_timeseries_generator"
	ei "github.com/lmpizarro/go_ehlers_indicators"
	"testing"
)

func TestStDevRoll(t *testing.T) {
	vals := go_timeseries_generator.GaussianProcess(1024)
	smooth := ei.StDevRoll(vals, 20)

	filename := "img/stdev_roll.png"
	err := ei.Plt(smooth, filename)
	if err != nil {
		t.Error(err)
	}
	ei.Wrt(vals, smooth, "data/stdev_roll.csv")
}
