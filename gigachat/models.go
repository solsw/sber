package gigachat

import (
	"context"
	"net/http"
	"net/url"

	"github.com/solsw/generichelper"
	"github.com/solsw/httphelper"
	"github.com/solsw/sber/common"
)

// Output model object.
type Model struct {
	// https://developers.sber.ru/docs/ru/gigachat/api/reference#get-models
	Id      string `json:"id"`
	Object  string `json:"object"`
	OwnedBy string `json:"owned_by"`
}

// Output models object.
type ModelsOut struct {
	// https://developers.sber.ru/docs/ru/gigachat/api/reference#get-models-model
	Object string  `json:"object"`
	Data   []Model `json:"data"`
}

// Models returns available models.
func Models(ctx context.Context, currToken *common.Token) (*ModelsOut, error) {
	// https://developers.sber.ru/docs/ru/gigachat/api/reference#get-models
	auth, err := common.AuthBearer(ctx, currToken)
	if err != nil {
		return nil, err
	}
	url, err := url.JoinPath(baseApiUrl, "models")
	if err != nil {
		return nil, err
	}
	h := make(http.Header)
	h.Set("Authorization", auth)
	return httphelper.RestInOut[generichelper.NoType, ModelsOut, common.ErrorOut](
		context.Background(), http.DefaultClient, http.MethodGet, url, h, nil)
}

// ModelsModel returns description of the model.
func ModelsModel(ctx context.Context, currToken *common.Token, model string) (*Model, error) {
	// https://developers.sber.ru/docs/ru/gigachat/api/reference#get-models-model
	auth, err := common.AuthBearer(ctx, currToken)
	if err != nil {
		return nil, err
	}
	url, err := url.JoinPath(baseApiUrl, "models", model)
	if err != nil {
		return nil, err
	}
	h := make(http.Header)
	h.Set("Authorization", auth)
	return httphelper.RestInOut[generichelper.NoType, Model, common.ErrorOut](
		context.Background(), http.DefaultClient, http.MethodGet, url, h, nil)
}
