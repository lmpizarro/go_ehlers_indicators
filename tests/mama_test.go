package go_ehlers_indicators

import (
	"github.com/lmpizarro/go_timeseries_generator"
	ei "github.com/lmpizarro/go_ehlers_indicators"
	"testing"
)

func TestMAMAGraph(t *testing.T) {
	vals := go_timeseries_generator.GaussianProcess(1024)

	mama := ei.MAMADefault(vals)
	filename := "img/mama.png"
	err := ei.Plt(mama, filename)
	if err != nil {
		t.Error(err)
	}
}
