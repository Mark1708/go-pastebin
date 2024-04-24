package gerror

type Base struct {
	Status  int
	Code    string
	Message string
}

type BadArguments struct {
	Base
}
