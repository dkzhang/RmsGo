package webapiError

type errStruct struct {
	errorType int
	errorInfo string
	msg       string
}

type Err interface {
	Error() string
	Type() int
	Msg() string
}

func (e errStruct) Error() string {
	return e.errorInfo
}

func (e errStruct) Type() int {
	return e.errorType
}

func (e errStruct) Msg() string {
	return e.msg
}

func WaErr(et int, ei string, m string) errStruct {
	return errStruct{
		errorType: et,
		errorInfo: ei,
		msg:       m,
	}
}

const (
	TypeBadRequest        = 400
	TypeDatabaseError     = 404
	TypeAuthorityError    = 403
	TypeNotYetImplemented = 500
)
