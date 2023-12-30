package common

import (
	"context"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/solsw/httphelper"
	"github.com/solsw/timehelper"
)

// https://developers.sber.ru/docs/ru/gigachat/api/authorization
// https://developers.sber.ru/docs/ru/gigachat/api/integration-individuals
// https://developers.sber.ru/docs/ru/gigachat/api/integration-legal-entities

// Access token.
type Token struct {
	// https://developers.sber.ru/docs/ru/gigachat/api/authorization#format-otveta25
	AccessToken string              `json:"access_token"`
	ExpiresAt   timehelper.UnixMsec `json:"expires_at"`
	// Client Id
	id string
	// Client Secret
	secret string
}

// authBasic returns Basic authorization value.
func authBasic(ctx context.Context, id, secret string) (string, error) {
	// https://developers.sber.ru/docs/ru/gigachat/api/authorization#shag-1-podklyuchenie-giga-chat-api-i-poluchenie-avtorizatsionnyh-dannyh
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(id+":"+secret)), nil
}

// GetToken returns access [Token].
func GetToken(ctx context.Context, id, secret string) (*Token, error) {
	// https://developers.sber.ru/docs/ru/gigachat/api/authorization#shag-2-poluchenie-tokena-dostupa-v-obmen-na-avtorizatsionnye-dannye
	auth, err := authBasic(ctx, id, secret)
	if err != nil {
		return nil, err
	}
	url := "https://ngw.devices.sberbank.ru:9443/api/v2/oauth"
	h := make(http.Header)
	h.Set("Authorization", auth)
	h.Set("RqUID", uuid.NewString())
	h.Set("Content-Type", "application/x-www-form-urlencoded")
	in := "scope=GIGACHAT_API_PERS"
	// in := "scope=GIGACHAT_API_CORP"
	t, err := httphelper.RestInOut[string, Token, ErrorOut](context.Background(), http.DefaultClient, http.MethodPost, url, h, &in)
	if err != nil {
		return nil, err
	}
	t.id = id
	t.secret = secret
	return t, nil
}

func adjustToken(ctx context.Context, token *Token) error {
	if token != nil && time.Until(time.Time(token.ExpiresAt)) > 1*time.Minute {
		return nil
	}
	t, err := GetToken(ctx, token.id, token.secret)
	if err != nil {
		return err
	}
	*token = *t
	return nil
}

// AuthBearer returns Bearer authorization value.
func AuthBearer(ctx context.Context, token *Token) (string, error) {
	if err := adjustToken(ctx, token); err != nil {
		return "", err
	}
	return "Bearer " + token.AccessToken, nil
}
