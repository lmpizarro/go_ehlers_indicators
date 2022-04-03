package go_ehlers_indicators

import (
	"github.com/lmpizarro/go_timeseries_generator"
	ei "github.com/lmpizarro/go_ehlers_indicators"
	"testing"
)

func TestLaguerreRSI(t *testing.T) {
	vals := go_timeseries_generator.GaussianProcess(512)
	lrsi := ei.LaguerreRSIDefault(vals)

	filename := "img/laguerre_rsi.png"
	err := ei.Plt(lrsi, filename)
	if err != nil {
		t.Error(err)
	}
}
