package go_ehlers_indicators_test

import (
	"github.com/lmpizarro/go_timeseries_generator"
	ei "github.com/lmpizarro/go_ehlers_indicators"
	"testing"
	"fmt"
)

func TestKamaDefault(t *testing.T) {
	vals := go_timeseries_generator.GaussianProcess(1024)
	smooth := ei.KamaDefault(vals)

	filename := "KamaDefault"
	err := ei.Plt(smooth, fmt.Sprintf("img/%s.png",filename))
	if err != nil {
		t.Error(err)
	}


	ei.Wrt(vals, smooth, fmt.Sprintf("data/%s.csv",filename))
}
