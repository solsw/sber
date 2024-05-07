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
	u, _ := url.JoinPath(baseApiUrl, "models")
	rq, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	rq.Header.Set("Authorization", "Bearer "+accessToken)
	out, err := rest.ReqJson[ModelsOut, common.OutError](http.DefaultClient, rq, httphelper.IsNotStatusOK)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return out, nil
}
