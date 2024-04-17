package auth

import (
	"context"
	"errors"
	"net/http"
	neturl "net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/solsw/errorhelper"
	"github.com/solsw/httphelper"
	"github.com/solsw/httphelper/rest"
	"github.com/solsw/timehelper"
	"github.com/solsw/sber/common"
)

// https://developers.sber.ru/docs/ru/gigachat/api/authorization
// https://developers.sber.ru/docs/ru/gigachat/api/integration-individuals
// https://developers.sber.ru/docs/ru/gigachat/api/integration-legal-entities

// Token as returned from /api/v2/oauth endpoint.
type Token struct {
	// https://developers.sber.ru/docs/ru/gigachat/api/authorization#format-otveta25
	AccessToken string              `json:"access_token"`
	ExpiresAt   timehelper.UnixMsec `json:"expires_at"`

	// Client Id
	id string
	// Client Secret
	secret string
}

// NewToken returns new [Token].
func NewToken(ctx context.Context, id, secret string) (*Token, error) {
	// https://developers.sber.ru/docs/ru/gigachat/api/authorization#shag-2-poluchenie-tokena-dostupa-v-obmen-na-avtorizatsionnye-dannye
	// https://developers.sber.ru/docs/ru/gigachat/individuals-quickstart#shag-2-poluchenie-tokena-dostupa
	// https://developers.sber.ru/docs/ru/gigachat/api/reference/rest/post-token
	h := make(http.Header)
	h.Set("Authorization", httphelper.AuthBasic(id, secret))
	h.Set("RqUID", uuid.NewString())
	h.Set("Content-Type", "application/x-www-form-urlencoded")
	u := "https://ngw.devices.sberbank.ru:9443/api/v2/oauth"
	body := neturl.Values{}
	body.Set("scope", "GIGACHAT_API_PERS")
	// body.Set("scope", "GIGACHAT_API_CORP")
	in := strings.NewReader(body.Encode())
	t, err := rest.BodyJson[Token, common.OutError](
		ctx, http.DefaultClient, http.MethodPost, u, h, in, rest.IsNotStatusOK)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	t.id, t.secret = id, secret
	return t, nil
}

// VetToken updates 't', if needed.
func VetToken(ctx context.Context, t *Token) error {
	if t == nil {
		return errorhelper.CallerError(errors.New("nil token"))
	}
	if time.Until(time.Time(t.ExpiresAt)) > 1*time.Minute {
		return nil
	}
	// access token expired
	newt, err := NewToken(ctx, t.id, t.secret)
	if err != nil {
		return errorhelper.CallerError(err)
	}
	*t = *newt
	return nil
}
