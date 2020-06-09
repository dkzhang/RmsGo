package userTempDM

type UserTempDM interface {
	IsSmsLock(userID int) bool
	LockSms(userID int) error

	SetPassword(userID int) error
	ValidatePassword(userID int) (bool, error)

	CreateToken(userID int) (token string, err error)
	ValidateToken(token string) (userID int, err error)
	DeleteToken(userID int) (err error)
}
