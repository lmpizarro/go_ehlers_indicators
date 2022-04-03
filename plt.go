package go_ehlers_indicators

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"os"
	"fmt"
	"bufio"
)

type (
	Values struct {
		Xs []float64
		Ys []float64
	}
)

func (v *Values) Len() int {
	return len(v.Xs)
}

func (v *Values) XY(index int) (float64, float64) {
	return v.Xs[index], v.Ys[index]
}

// Plt will create a line plot with given values and filename
func Plt(vals []float64, filename string) error {
	xs := make([]float64, len(vals))
	ys := make([]float64, len(vals))
	for i := 0; i < len(vals); i++ {
		xs[i] = float64(i)
		ys[i] = vals[i]
	}
	values := &Values{
		Xs: xs,
		Ys: ys,
	}

	p := plot.New()
	p.Title.Text = filename

	line, err := plotter.NewLine(values)
	if err != nil {
		return err
	}
	p.Add(line)

	if err := p.Save(297*vg.Millimeter, 210*vg.Millimeter, filename); err != nil {
		return err
	}
	return nil
}

func Wrt(vals, filter_out []float64, filename string) error {

	f, err := os.Create(filename)
    if err != nil {
        return err
    }
    // remember to close the file
    defer f.Close()
	    // create new buffer
    buffer := bufio.NewWriter(f)

	buffer.WriteString("vals,out\n")
	for i, line := range filter_out {
		d := fmt.Sprintf("%f,%f\n", vals[i], line)
        _, err := buffer.WriteString(d)
        if err != nil {
            return err
        }
    }

    // flush buffered data to the file
    if err := buffer.Flush(); err != nil {
         return err
    }

	return nil
}
