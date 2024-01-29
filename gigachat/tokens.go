package gigachat

import (
	"context"
	"net/http"
	"net/url"

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
func TokensCount(ctx context.Context, currToken *common.Token, tokensCountIn *TokensCountIn) (*[]TokensCountOut, error) {
	// https://developers.sber.ru/docs/ru/gigachat/api/reference/rest/post-tokens-count
	auth, err := common.AuthBearer(ctx, currToken)
	if err != nil {
		return nil, err
	}
	url, _ := url.JoinPath(baseApiUrl, "tokens/count")
	h := make(http.Header)
	h.Set("Authorization", auth)
	h.Set("Content-Type", "application/json")
	return rest.JsonJson[TokensCountIn, []TokensCountOut, common.OutError](
		ctx, http.DefaultClient, http.MethodPost, url, h, tokensCountIn, rest.IsNotStatusOK)
}
