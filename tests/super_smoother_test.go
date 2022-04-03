package go_ehlers_indicators_test

import (
	"github.com/lmpizarro/go_timeseries_generator"
	ei "github.com/lmpizarro/go_ehlers_indicators"
	"testing"
)

func TestSuperSmoother(t *testing.T) {
	vals := go_timeseries_generator.GaussianProcess(1024)
	smooth := ei.SuperSmoother(vals, 16)

	filename := "img/super_smoother.png"
	err := ei.Plt(smooth, filename)
	if err != nil {
		t.Error(err)
	}
}
