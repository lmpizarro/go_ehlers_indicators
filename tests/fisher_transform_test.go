package go_ehlers_indicators_test

import (
	"fmt"
	"github.com/lmpizarro/go_timeseries_generator"
	ei "github.com/lmpizarro/go_ehlers_indicators"
	"testing"
)

func TestFisherTransformGraph(t *testing.T) {
	vals := go_timeseries_generator.SineWave(1024)
	fish, err := ei.FisherTransform(vals)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("fish: %v", fish)

	filename := "img/fisher_transform.png"
	err = ei.Plt(fish, filename)
	if err != nil {
		t.Error(err)
	}
}
