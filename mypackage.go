package mypackage

const (
	_ = iota
	ONE
	TWO
	THREE
)

const (
	W, K = iota, iota + 1
)

func GetValue() string {
	return "Hello get value"
}
