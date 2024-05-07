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

// Токен доступа, возвращаемый по запросу к /api/v2/oauth.
type Token struct {
	// https://developers.sber.ru/docs/ru/gigachat/api/authorization#format-otveta25
	AccessToken string              `json:"access_token"`
	ExpiresAt   timehelper.UnixMsec `json:"expires_at"`

	// Client Id
	id string
	// Client Secret
	secret string
}

// NewToken возвращает новый [Token].
func NewToken(ctx context.Context, id, secret string) (*Token, error) {
	// https://developers.sber.ru/docs/ru/gigachat/api/authorization#shag-2-poluchenie-tokena-dostupa-v-obmen-na-avtorizatsionnye-dannye
	// https://developers.sber.ru/docs/ru/gigachat/individuals-quickstart#shag-2-poluchenie-tokena-dostupa
	// https://developers.sber.ru/docs/ru/gigachat/legal-quickstart#shag-4-poluchenie-tokena-dostupa
	// https://developers.sber.ru/docs/ru/gigachat/api/reference/rest/post-token
	u := "https://ngw.devices.sberbank.ru:9443/api/v2/oauth"
	v := neturl.Values{}
	v.Set("scope", "GIGACHAT_API_PERS")
	// v.Set("scope", "GIGACHAT_API_CORP")
	body := strings.NewReader(v.Encode())
	rq, err := http.NewRequestWithContext(ctx, http.MethodPost, u, body)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	rq.Header.Set("Authorization", httphelper.AuthBasic(id, secret))
	rq.Header.Set("RqUID", uuid.NewString())
	rq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	t, err := rest.ReqJson[Token, common.OutError](http.DefaultClient, rq, httphelper.IsNotStatusOK)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	t.id, t.secret = id, secret
	return t, nil
}

// VetToken обновляет [Token], если истёк его срок действия.
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
