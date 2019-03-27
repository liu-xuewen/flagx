package value

import "strconv"

// -- uint Value
type UIntValue uint

func newUIntValue(val uint, p *uint) *UIntValue {
	*p = val
	return (*UIntValue)(p)
}

func (i *UIntValue) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, strconv.IntSize)
	if err != nil {
		err = numError(err)
	}
	*i = UIntValue(v)
	return err
}

func (i *UIntValue) Get() interface{} { return uint(*i) }

func (i *UIntValue) String() string { return strconv.FormatUint(uint64(*i), 10) }
