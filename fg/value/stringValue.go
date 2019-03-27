package value

// -- string Value
type StringValue string

func NewStringValue(val string, p *string) *StringValue {
	*p = val //在NewStringValue这里对p进行了赋值
	return (*StringValue)(p)
}

func (s *StringValue) Set(val string) error {
	*s = StringValue(val)
	return nil
}

func (s *StringValue) Get() interface{} { return string(*s) }

func (s *StringValue) String() string { return string(*s) }
