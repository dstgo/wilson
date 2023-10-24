package code

const (
	RequestOk  = 2000
	LoginOk    = 2001
	RegisterOk = 2002
	LogoutOK   = 2003

	EmailSendOk = 2006
)

const (
	BadRequest     = 4000
	LoginFailed    = 4001
	RegisterFailed = 4002
	LogoutFailed   = 4003

	EmailSendFailed = 4006

	ResourceNotFound = 4040

	Forbidden    = 4030
	UnAuthorized = 4010
)

const (
	InternalServerError = 5000 + iota
	UnknownError
	DatabaseError
	NetworkError
	ProgramError
	FilesystemError
)
