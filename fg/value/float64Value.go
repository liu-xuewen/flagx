package value

import (
	"strconv"
)

// -- float64 Value
type Float64Value float64

func NewFloat64Value(val float64, p *float64) *Float64Value {
	*p = val
	return (*Float64Value)(p)
}

func (f *Float64Value) Set(s string) error {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		err = numError(err)
	}
	*f = Float64Value(v)
	return err
}

func (f *Float64Value) Get() interface{} { return float64(*f) }

func (f *Float64Value) String() string { return strconv.FormatFloat(float64(*f), 'g', -1, 64) }
