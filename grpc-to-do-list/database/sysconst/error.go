package sysconst

import "errors"

var (
	ErrorTokenExpired             error = errors.New("token is expired")
	ErrorTokenNotValidYet         error = errors.New("token not active yet")
	ErrorTokenMalformed           error = errors.New("that's not even a token")
	ErrorTokenInvalid             error = errors.New("couldn't handle this token")
	ErrorPlayerInvalid            error = errors.New("can not find player")
	ErrorTeamInvalid              error = errors.New("can not find team")
	ErrorUserNameDuplicateInvalid error = errors.New("username duplicate")
	ErrorInvalidParameter         error = errors.New("invalid parameter")
	ErrDataOutOfSync              error = errors.New("the data is out of sync")
)

var (
	ErrNonAffected  error = errors.New("nothing affected")
	ErrNotFound           = errors.New("data not found")
	ErrEnumNotFound       = errors.New("enum not found")
)

func GetResultErrorMessage(rs error) string {
	var result = ""
	switch rs {
	case ErrorInvalidParameter:
		result = "錯誤參數"
	}
	return result
}
