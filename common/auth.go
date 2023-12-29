package common

import (
	"context"
	"encoding/base64"
	"errors"
	"net/http"
	"os"
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
}

func getIdSecret() (string, string, error) {
	id := os.Getenv("SBER_GIGACHAT_CLIENT_ID")
	if id == "" {
		return "", "", errors.New("no Client Id")
	}
	secret := os.Getenv("SBER_GIGACHAT_CLIENT_SECRET")
	if secret == "" {
		return "", "", errors.New("no Client Secret")
	}
	return id, secret, nil
}

// AuthBasic returns Basic authorization value.
func AuthBasic(ctx context.Context) (string, error) {
	// https://developers.sber.ru/docs/ru/gigachat/api/authorization#shag-1-podklyuchenie-giga-chat-api-i-poluchenie-avtorizatsionnyh-dannyh
	id, secret, err := getIdSecret()
	if err != nil {
		return "", err
	}
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(id+":"+secret)), nil
}

// GetToken returns access [Token].
func GetToken(ctx context.Context, currToken *Token) (*Token, error) {
	// https://developers.sber.ru/docs/ru/gigachat/api/authorization#shag-2-poluchenie-tokena-dostupa-v-obmen-na-avtorizatsionnye-dannye
	if currToken != nil && time.Until(time.Time(currToken.ExpiresAt)) > 1*time.Minute {
		return currToken, nil
	}
	auth, err := AuthBasic(ctx)
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
	if currToken != nil {
		currToken = t
	}
	return t, nil
}

// AuthBearer returns Bearer authorization value.
func AuthBearer(ctx context.Context, currToken *Token) (string, error) {
	t, err := GetToken(ctx, currToken)
	if err != nil {
		return "", err
	}
	return "Bearer " + t.AccessToken, nil
}
