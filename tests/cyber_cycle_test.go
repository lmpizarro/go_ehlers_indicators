package go_ehlers_indicators_test

import (
	"fmt"
	"github.com/lmpizarro/go_timeseries_generator"
	ei "github.com/lmpizarro/go_ehlers_indicators"
	"testing"
)

func TestCyberCycle(t *testing.T) {
	vals := go_timeseries_generator.GaussianProcess(1024)
	cc := ei.CyberCycle(vals, 16)
	for i := 0; i < len(cc); i++ {
		fmt.Printf("cc[%d]: %f\n", i, cc[i])
	}
}

func TestCyberCycleGraph(t *testing.T) {
	vals := go_timeseries_generator.GaussianProcess(1024)
	cc := ei.CyberCycle(vals, 16)

	filename := "./img/cyber_cycle.png"
	err := ei.Plt(cc, filename)
	if err != nil {
		t.Error(err)
	}
}
