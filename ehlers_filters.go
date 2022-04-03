package go_ehlers_indicators
import "math"
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

