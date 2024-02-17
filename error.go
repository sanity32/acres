package acres

type ErrorCode interface {
	String() string
	Int() int
}
