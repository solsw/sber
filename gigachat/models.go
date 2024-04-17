package gigachat

import (
	"context"
	"errors"
	"net/http"
	"net/url"

	"github.com/solsw/errorhelper"
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

// Models [возвращает] массив объектов с данными доступных моделей.
//
// [возвращает]: https://developers.sber.ru/docs/ru/gigachat/api/reference/rest/get-models
func Models(ctx context.Context, accessToken string) (*ModelsOut, error) {
	if accessToken == "" {
		return nil, errorhelper.CallerError(errors.New("no accessToken"))
	}
	u, err := url.JoinPath(baseApiUrl, "models")
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	h := make(http.Header)
	h.Set("Authorization", "Bearer "+accessToken)
	out, err := rest.BodyJson[ModelsOut, common.OutError](
		ctx, http.DefaultClient, http.MethodGet, u, h, nil, rest.IsNotStatusOK)
	return out, errorhelper.CallerError(err)
}

// ModelsModel [возвращает] объект с описанием указанной модели.
//
// [возвращает]: https://developers.sber.ru/docs/ru/gigachat/api/reference/rest/get-model
func ModelsModel(ctx context.Context, accessToken string, model string) (*Model, error) {
	if accessToken == "" {
		return nil, errorhelper.CallerError(errors.New("no accessToken"))
	}
	u, err := url.JoinPath(baseApiUrl, "models", model)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	h := make(http.Header)
	h.Set("Authorization", "Bearer "+accessToken)
	out, err := rest.BodyJson[Model, common.OutError](
		ctx, http.DefaultClient, http.MethodGet, u, h, nil, rest.IsNotStatusOK)
	return out, errorhelper.CallerError(err)
}
