package value

import "strconv"

// -- int Value
type IntValue int

func newIntValue(val int, p *int) *IntValue {
	*p = val
	return (*IntValue)(p)
}

func (i *IntValue) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, strconv.IntSize)
	if err != nil {
		err = numError(err)
	}
	*i = IntValue(v)
	return err
}

func (i *IntValue) Get() interface{} { return int(*i) }

func (i *IntValue) String() string { return strconv.Itoa(int(*i)) }
