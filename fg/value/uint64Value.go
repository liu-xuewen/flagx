package value

import "strconv"

// -- uint64 Value
type UInt64Value uint64

func newUInt64Value(val uint64, p *uint64) *UInt64Value {
	*p = val
	return (*UInt64Value)(p)
}

func (i *UInt64Value) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 64)
	if err != nil {
		err = numError(err)
	}
	*i = UInt64Value(v)
	return err
}

func (i *UInt64Value) Get() interface{} { return uint64(*i) }

func (i *UInt64Value) String() string { return strconv.FormatUint(uint64(*i), 10) }
