package atm

import (
	"context"
	"crypto/rsa"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/twitchtv/twirp"
)

func GenerateToken(merchantID string, key *rsa.PrivateKey, exp time.Duration) string {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"uid": merchantID,
		"exp": time.Now().Add(exp).Unix(),
	})

	str, err := token.SignedString(key)
	if err != nil {
		panic(err)
	}

	return str
}

// WithToken store an authorization token in a context.Context
// 如果是调用需要鉴权的 rpc 方法，需要通过这种方式将 token 附带上去
func WithToken(ctx context.Context, token string) context.Context {
	header := make(http.Header)
	header.Set("Authorization", "Bearer "+token)

	ctx, err := twirp.WithHTTPRequestHeaders(ctx, header)
	if err != nil {
		panic(fmt.Errorf("twirp error setting headers: %s", err))
	}

	return ctx
}
