package auth_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/guru-invest/guru.framework/src/security/auth"
)

func TestGetExpiredToken(t *testing.T) {
	const _expireOffset = 3600

	Remaining := auth.InitJWTAuthenticationBackend("InvalidToken").GetTokenRemainingValidity("")

	if Remaining == _expireOffset {
		return
	}

	fmt.Printf("\033[1;31m%s%d\033[0m", "RemainingValidity must be ", _expireOffset)
	t.Fail()
}

func TestGetTokenRemainingValidity(t *testing.T) {

	const _expireOffset = 3600

	now := time.Now()
	tenSecondsInTheFuture := now.Add(10 * time.Second)

	Remaining := auth.InitJWTAuthenticationBackend("InvalidToken").GetTokenRemainingValidity(float64(tenSecondsInTheFuture.Unix()) + float64(0.000000000))

	if Remaining > _expireOffset {
		return
	}

	fmt.Printf("\033[1;31m%s%d\033[0m", "RemainingValidity must be higher than ", _expireOffset)
	t.Fail()
}
