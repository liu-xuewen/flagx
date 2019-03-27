package value

import "strconv"

// -- int64 Value
type Int64Value int64

func newInt64Value(val int64, p *int64) *Int64Value {
	*p = val
	return (*Int64Value)(p)
}

func (i *Int64Value) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		err = numError(err)
	}
	*i = Int64Value(v)
	return err
}

func (i *Int64Value) Get() interface{} { return int64(*i) }

func (i *Int64Value) String() string { return strconv.FormatInt(int64(*i), 10) }
