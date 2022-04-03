package go_ehlers_indicators_test

import (
	"fmt"
	"github.com/lmpizarro/go_timeseries_generator"
	ei "github.com/lmpizarro/go_ehlers_indicators"
	"testing"
)

func TestTrendFlex(t *testing.T) {
	vals := go_timeseries_generator.GaussianProcess(1024)
	tf := ei.TrendFlex(vals, 16)
	for i := 0; i < len(tf); i++ {
		fmt.Printf("trendflex[%d]: %f\n", i, tf[i])
	}
}

func TestTrendFlexGraph(t *testing.T) {
	vals := go_timeseries_generator.GaussianProcess(1024)
	tf := ei.TrendFlex(vals, 16)

	filename := "./img/trend_flex.png"
	err := ei.Plt(tf, filename)
	if err != nil {
		t.Error(err)
	}
}
