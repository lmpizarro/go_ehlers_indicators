package go_ehlers_indicators

import (
	"github.com/lmpizarro/go_timeseries_generator"
	ei "github.com/lmpizarro/go_ehlers_indicators"
	"testing"
)

func TestInverseFisherTransform(t *testing.T) {
	sine := go_timeseries_generator.SineWave(1024)
	ift := ei.InverseFisherTransform(sine)

	for i := 0; i < len(ift); i++ {
		if ift[i] > 1 || ift[i] < -1 {
			t.Error(ift[i] > 1 || ift[i] < -1)
		}
	}
}

func TestInverseFisherTransformGraph(t *testing.T) {
	sine := go_timeseries_generator.SineWave(1024)
	ift := ei.InverseFisherTransform(sine)

	filename := "img/inverse_fisher_transform.png"
	err := ei.Plt(ift, filename)
	if err != nil {
		t.Error(err)
	}
}
