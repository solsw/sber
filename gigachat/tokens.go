package gigachat

import (
	"context"
	"net/http"
	"net/url"

	"github.com/solsw/httphelper"
	"github.com/solsw/sber/common"
)

// Input tokens/count object.
type TokensCountIn struct {
	// https://developers.sber.ru/docs/ru/gigachat/api/reference#post-tokens-count
	Model string   `json:"model"`
	Input []string `json:"input"`
}

// Output tokens/count object.
type TokensCountOut struct {
	// https://developers.sber.ru/docs/ru/gigachat/api/reference#post-tokens-count
	Object     string `json:"object"`
	Tokens     int    `json:"tokens"`
	Characters int    `json:"characters"`
}

// TokensCount returns quantity of tokens calculated by the model for every string in input array.
func TokensCount(ctx context.Context, currToken *common.Token, tokensCountIn *TokensCountIn) (*[]TokensCountOut, error) {
	// https://developers.sber.ru/docs/ru/gigachat/api/reference#post-tokens-count
	auth, err := common.AuthBearer(ctx, currToken)
	if err != nil {
		return nil, err
	}
	url, _ := url.JoinPath(baseApiUrl, "tokens/count")
	h := make(http.Header)
	h.Set("Authorization", auth)
	h.Set("Content-Type", "application/json")
	return httphelper.RestInOut[TokensCountIn, []TokensCountOut, common.ErrorOut](
		context.Background(), http.DefaultClient, http.MethodPost, url, h, tokensCountIn)
}
