package go_ehlers_indicators_test

import (
	"fmt"
	"github.com/lmpizarro/go_timeseries_generator"
	ei "github.com/lmpizarro/go_ehlers_indicators"
	"testing"
)

func TestGaussianFilter(t *testing.T) {
	vals := go_timeseries_generator.GaussianProcess(1024)

	for poles := 1; poles < 5; poles++ {
		filt := ei.GaussianFilter(vals, 16, poles)
		filename := fmt.Sprintf("img/gaussian_filter_p%d.png", poles)
		err := ei.Plt(filt, filename)
		if err != nil {
			t.Error(err)
		}
	}
}
