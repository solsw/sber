package gigachat

import (
	"context"
	"net/http"
	"net/url"

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
func Embeddings(ctx context.Context, currToken *common.Token, embeddingsIn *EmbeddingsIn) (*EmbeddingsOut, error) {
	// https://developers.sber.ru/docs/ru/gigachat/api/reference/rest/post-embeddings
	auth, err := common.AuthBearer(ctx, currToken)
	if err != nil {
		return nil, err
	}
	url, err := url.JoinPath(baseApiUrl, "embeddings")
	if err != nil {
		return nil, err
	}
	h := make(http.Header)
	h.Set("Authorization", auth)
	h.Set("Content-Type", "application/json")
	h.Set("Accept", "application/json")
	return rest.JsonJson[EmbeddingsIn, EmbeddingsOut, common.OutError](
		ctx, http.DefaultClient, http.MethodPost, url, h, embeddingsIn, rest.IsNotStatusOK)

}
