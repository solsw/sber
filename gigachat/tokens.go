package gigachat

import (
	"context"
	"errors"
	"net/http"
	"net/url"

	"github.com/solsw/errorhelper"
	"github.com/solsw/httphelper"
	"github.com/solsw/httphelper/rest"
	"github.com/solsw/sber/common"
)

// Input tokens/count object.
type TokensCountIn struct {
	// https://developers.sber.ru/docs/ru/gigachat/api/reference/rest/post-tokens-count#zapros
	Model string   `json:"model"`
	Input []string `json:"input"`
}

// Output tokens/count object.
type TokensCountOut struct {
	// https://developers.sber.ru/docs/ru/gigachat/api/reference/rest/post-tokens-count#responses
	Object     string `json:"object"`
	Tokens     int    `json:"tokens"`
	Characters int    `json:"characters"`
}

// TokensCount возвращает объект с информацией о количестве токенов,
// посчитанных заданной моделью в строках, переданных в массиве input.
func TokensCount(ctx context.Context, accessToken string, tokensCountIn *TokensCountIn) (*[]TokensCountOut, error) {
	// https://developers.sber.ru/docs/ru/gigachat/api/reference/rest/post-tokens-count
	if accessToken == "" {
		return nil, errorhelper.CallerError(errors.New("no accessToken"))
	}
	u, _ := url.JoinPath(baseApiUrl, "tokens/count")
	body, err := httphelper.JsonBody(tokensCountIn)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	rq, err := http.NewRequestWithContext(ctx, http.MethodPost, u, body)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	rq.Header.Set("Authorization", "Bearer "+accessToken)
	rq.Header.Set("Content-Type", "application/json")
	out, err := rest.ReqJson[[]TokensCountOut, common.OutError](http.DefaultClient, rq, httphelper.IsNotStatusOK)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return out, nil
}
