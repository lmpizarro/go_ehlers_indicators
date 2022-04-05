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

func Alfa(length int) float64{
	 return 2.0 / (float64(length) + 1.0)
}

func Kama(vals[] float64, fast_period, slow_period, efficiency_length int) []float64 {
	filt := make([]float64, len(vals))

	eff := EficiencyRatio(vals, efficiency_length)

	fastest := Alfa(fast_period) 
	slowest := Alfa(slow_period) 

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

func Zema(vals[] float64,  length_alfa, length_k, delay int) []float64 {
	filt := make([]float64, len(vals))

	alfa := Alfa(length_alfa)
	k := Alfa(length_k)
	
	for i := delay; i < len(vals); i++ {
		momentum := vals[i] - vals[i-delay]
		filt[i] = alfa * (vals[i] + k * (momentum)) + (1 - alfa) * filt[i-1]
	}

	return filt
}

func SimpleKalman(measurement[] float64) [] float64{
	filt := make([]float64, len(measurement))

	
	uncertain_estimate := measurement[0] // Pn,n-1 
	uncertain_measu := measurement[0]    // Rn
	kg := uncertain_estimate / (uncertain_estimate + uncertain_measu)

	for i := 1; i < len(measurement); i++{
		filt[i] = filt[i-1] + kg * (measurement[i] - filt[i-1])
		// uncertain_estimate = (1 - kg) * uncertain_estimate
		uncertain_measu = 20 * math.Abs(measurement[i] - measurement[i-1])
		kg = uncertain_estimate / (uncertain_estimate + uncertain_measu)
	}
	return filt
}

/*
    State Extrapolation Equations 

	x_n+1 = x_n + Dt * v_n`
	v_n+1 = v_n

	ex_n,n-1 = ex_n-1,n-1 + Dt * ev_n-1,n-1
	ev_n,n-1 = ev_n-1,n-1

	State Update Equation

	ev_n,n = ev_n,n-1 + beta * (z_n - ex_n,n-1) / Dt
	ex_n,n = ex_n,n-1 + alfa * (z_n - ex_n,n-1)

*/