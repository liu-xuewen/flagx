package fg

type Value interface {
	String() string
	Set(string) error
}
