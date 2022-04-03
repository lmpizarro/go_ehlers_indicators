package go_ehlers_indicators

import (
	"fmt"
	"github.com/lmpizarro/go_timeseries_generator"
	ei "github.com/lmpizarro/go_ehlers_indicators"
	"testing"
)

func TestReFlex(t *testing.T) {
	vals := go_timeseries_generator.GaussianProcess(1024)
	reflex := ei.ReFlex(vals, 16)
	for i := 0; i < len(reflex); i++ {
		reflex := fmt.Sprintf("reflex[%d]: %f\n", i, reflex)
		fmt.Print(reflex)
	}
}

func TestReFlexGraph(t *testing.T) {
	vals := go_timeseries_generator.GaussianProcess(1024)
	reflex := ei.ReFlex(vals, 16)

	filename := "./img/re_flex.png"
	err := ei.Plt(reflex, filename)
	if err != nil {
		t.Error(err)
	}
}
