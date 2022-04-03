package go_ehlers_indicators_test

import (
	"fmt"
	"github.com/lmpizarro/go_timeseries_generator"
	"github.com/lmpizarro/go_ehlers_indicators"
	"testing"
)

func TestCenterOfGravity(t *testing.T) {
	vals := go_timeseries_generator.GaussianProcess(1024)
	cgo := go_ehlers_indicators.CenterOfGravity(vals, 16)
	for i := 0; i < len(cgo); i++ {
		fmt.Printf("cgo[%d]: %f\n", i, cgo[i])
	}
}

func TestCenterOfGravityGraph(t *testing.T) {
	vals := go_timeseries_generator.GaussianProcess(1024)
	cgo := go_ehlers_indicators.CenterOfGravity(vals, 16)
	filename := "./img/center_of_gravity.png"
	err := go_ehlers_indicators.Plt(cgo, filename)
	if err != nil {
		t.Error(err)
	}
}
