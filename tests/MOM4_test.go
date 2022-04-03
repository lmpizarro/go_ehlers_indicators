package go_ehlers_indicators_test

import (
	"github.com/lmpizarro/go_timeseries_generator"
	ei "github.com/lmpizarro/go_ehlers_indicators"
	"testing"
)

func TestMOM4(t *testing.T) {
	vals := go_timeseries_generator.GaussianProcess(1024)
	smooth := ei.MOM4(vals)

	filename := "img/wma4.png"
	err := ei.Plt(smooth, filename)
	if err != nil {
		t.Error(err)
	}
	ei.Wrt(vals, smooth, "data/mom4.csv")
}
