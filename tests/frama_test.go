package go_ehlers_indicators

import (
	"github.com/lmpizarro/go_timeseries_generator"
	"github.com/lmpizarro/go_ehlers_indicators"
	"testing"
)

func TestFRAMAGraph(t *testing.T) {
	candles := go_timeseries_generator.GaussianOHLCVDefault(1024)
	highs := make([]float64, len(candles))
	lows := make([]float64, len(candles))
	// fill in higs and low
	for i := 0; i < len(candles); i++ {
		highs[i] = candles[i].High
		lows[i] = candles[i].Low
	}
	frama, err := go_ehlers_indicators.FRAMA(highs, lows, 16)
	if err != nil {
		t.Error(err)
	}

	filename := "img/frama.png"
	err = go_ehlers_indicators.Plt(frama, filename)
	if err != nil {
		t.Error(err)
	}
}
