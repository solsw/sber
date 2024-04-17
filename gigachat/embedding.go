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

// Input embeddings object.
type EmbeddingsIn struct {
	// https://developers.sber.ru/docs/ru/gigachat/api/reference/rest/post-embeddings#zapros
	Model string   `json:"model"`
	Input []string `json:"input"`
}

// Embedding
type Embedding struct {
	Object    string    `json:"object"`
	Embedding []float64 `json:"embedding"`
	Index     int64     `json:"index"`
}

// Output embeddings object.
type EmbeddingsOut struct {
	// https://developers.sber.ru/docs/ru/gigachat/api/reference/rest/post-embeddings#responses
	Object string      `json:"object"`
	Data   []Embedding `json:"data"`
	Model  string      `json:"model"`
}

// Embeddings возвращает векторные представления соответствующих текстовых запросов.
func Embeddings(ctx context.Context, accessToken string, embeddingsIn *EmbeddingsIn) (*EmbeddingsOut, error) {
	// https://developers.sber.ru/docs/ru/gigachat/api/reference/rest/post-embeddings
	// https://developers.sber.ru/docs/ru/gigachat/api/embeddings
	if accessToken == "" {
		return nil, errorhelper.CallerError(errors.New("no accessToken"))
	}
	u, err := url.JoinPath(baseApiUrl, "embeddings")
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	h := make(http.Header)
	h.Set("Authorization", "Bearer "+accessToken)
	h.Set("Content-Type", "application/json")
	h.Set("Accept", "application/json")
	out, err := rest.JsonJson[EmbeddingsIn, EmbeddingsOut, common.OutError](
		ctx, http.DefaultClient, http.MethodPost, u, h, embeddingsIn, rest.IsNotStatusOK)
	return out, errorhelper.CallerError(err)
}
