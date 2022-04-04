package go_ehlers_indicators
import (
	"math"
	stats "github.com/montanaflynn/stats"
)
// y = (4*Price + 3*Price[1] + 2*Price[2] + Price[3]) / 10
func WMA4(vals []float64) []float64{
	filt := make([]float64, len(vals))

	for i := 3; i < len(vals); i++ {
		filt[i] = (4*vals[i] + 3*vals[i-1] + 2*vals[i-2] + vals[i-3])/10
	}
	return filt
}

func MOM4(vals []float64) []float64{
	filt := make([]float64, len(vals))

	for i := 7; i < len(vals); i++ {

		coef1 := math.Abs(vals[i] - vals[i-4])
		coef2 := math.Abs(vals[i] - vals[i-5])
		coef3 := math.Abs(vals[i] - vals[i-6])
		coef4 := math.Abs(vals[i] - vals[i-7])

		sum_coef := coef1 + coef2 + coef3 + coef4

		filt[i] = (coef1*vals[i] + coef2*vals[i-1] + coef3*vals[i-2] + coef4*vals[i-3])/sum_coef
	}
	return filt
}

func DIST4NCE(vals []float64, length int) []float64{
	filt := make([]float64, len(vals))
	
	coefs := make([]float64, length)

	for i := 2*length; i < len(vals); i++ {

		for count := 0; count < length; count++ {
			coefs[count] = 0
			for lookback := 0; lookback < length; lookback++ {
				coefs[count] += math.Pow(vals[i - count] - vals[i - (count + lookback)],2)
			}
		}

		sum_coef := 0.0
		for count := 0; count < length; count++ {
			sum_coef += coefs[count]
			filt[i] += coefs[count]*vals[i-count]
		}
		filt[i] /= sum_coef 
	}
	return filt
}

func StDevRoll(vals [] float64, length int) []float64 {
	filt := make([]float64, len(vals))

	for i := 0; i < length; i++ {
		filt[i] = 1.0
	}

	for i := length; i < len(vals); i++ {
		filt[i], _ = stats.StandardDeviation(vals[i-length: i])
	}

	return filt
}

func EficiencyRatio(vals [] float64, length int) []float64 {
	filt := make([]float64, len(vals))

	for i := length + 1; i < len(vals); i++ {
		numerator := math.Abs(vals[i] - vals[i-length])
		denominator := 0.0
		for j := 0; j < length ; j++ {
			denominator += math.Abs(vals[i-j] - vals[i-j-1])
		}
		filt[i] = numerator / denominator
	}

	return filt
}


func Kama(vals[] float64, fast_period, slow_period, efficiency_length int) []float64 {
	filt := make([]float64, len(vals))

	eff := EficiencyRatio(vals, efficiency_length)

	fastest := 2.0 / (float64(fast_period) + 1.0)
	slowest := 2.0 / (float64(slow_period) + 1.0)

	for i := 1; i < len(vals); i++ {
		s := math.Pow((eff[i] * (fastest - slowest) + slowest), 2)
		filt[i] = s * vals[i] + (1 - s) * filt[i - 1]
	}
	return filt
}

func KamaDefault(vals[]float64) [] float64 {
	return Kama(vals, 2, 30, 10)
}

func Vidya(vals[] float64, fast_period, slow_period int) []float64 {
	filt := make([]float64, len(vals))

	stdev_fast := StDevRoll(vals, fast_period)
	stdev_slow := StDevRoll(vals, slow_period)

	for i := 1; i < len(vals); i++ {
		k := stdev_fast[i] / stdev_slow[i]
		filt[i] = 0.2 * k * vals[i] + (1 - 0.2 * k) * filt[i-1]
	}

	return filt
}