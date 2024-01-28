package gigachat

import (
	"context"
	"net/http"
	"net/url"

	"github.com/solsw/httphelper/rest"
	"github.com/solsw/sber/common"
)

// Output model object.
type Model struct {
	// https://developers.sber.ru/docs/ru/gigachat/api/reference/rest/get-models#responses
	Id      string `json:"id"`
	Object  string `json:"object"`
	OwnedBy string `json:"owned_by"`
}

// Output models object.
type ModelsOut struct {
	// https://developers.sber.ru/docs/ru/gigachat/api/reference/rest/get-models#responses
	Object string  `json:"object"`
	Data   []Model `json:"data"`
}

// Models возвращает массив объектов с данными доступных моделей.
func Models(ctx context.Context, currToken *common.Token) (*ModelsOut, error) {
	// https://developers.sber.ru/docs/ru/gigachat/api/reference/rest/get-models
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
	return rest.BodyJson[ModelsOut, common.OutError](ctx, http.DefaultClient, http.MethodGet, url, h, nil, nil)
}

// ModelsModel возвращает объект с описанием указанной модели.
func ModelsModel(ctx context.Context, currToken *common.Token, model string) (*Model, error) {
	// https://developers.sber.ru/docs/ru/gigachat/api/reference/rest/get-model
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
	return rest.BodyJson[Model, common.OutError](ctx, http.DefaultClient, http.MethodGet, url, h, nil, nil)
}
