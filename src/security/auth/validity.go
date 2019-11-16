package auth


import (
	"time"
)

func (backend *JWTAuthenticationStructure) GetTokenRemainingValidity(timestamp interface{}) int {
	if validity, ok := timestamp.(float64); ok {
		tm := time.Unix(int64(validity), 0)
		remainer := tm.Sub(time.Now())
		if remainer > 0 {
			return int(remainer.Seconds() + _expireOffset)
		}
	}
	return _expireOffset
}