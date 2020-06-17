package webAPI

type Error interface {
	error
	HttpCode() int
}
