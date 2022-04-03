package go_ehlers_indicators_test

import (
	"github.com/lmpizarro/go_timeseries_generator"
	ei "github.com/lmpizarro/go_ehlers_indicators"
	"testing"
)

func TestDISTANCE(t *testing.T) {
	vals := go_timeseries_generator.GaussianProcess(3000)
	smooth := ei.DIST4NCE(vals, 20)

	filename := "img/wma4.png"
	err := ei.Plt(smooth, filename)
	if err != nil {
		t.Error(err)
	}

	ei.Wrt(vals, smooth, "data/distance.csv")

}
