package value

import "time"

// -- time.Duration Value
type DurationValue time.Duration

func (d *DurationValue) Set(s string) error {
	v, err := time.ParseDuration(s)
	if err != nil {
		err = errParse
	}
	*d = DurationValue(v)
	return err
}

func (d *DurationValue) Get() interface{} { return time.Duration(*d) }

func (d *DurationValue) String() string { return (*time.Duration)(d).String() }
