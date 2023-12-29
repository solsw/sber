package gigachat

import (
	"context"
	"net/http"
	"net/url"

	"github.com/solsw/httphelper"
	"github.com/solsw/timehelper"
	"github.com/solsw/sber/common"
)

// Message.
type Message struct {
	// https://developers.sber.ru/docs/ru/gigachat/api/reference#post-chat-completions
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Input chat/completions object.
type ChatCompletionsIn struct {
	// https://developers.sber.ru/docs/ru/gigachat/api/reference#post-chat-completions
	Model             string    `json:"model"`
	Messages          []Message `json:"messages"`
	Temperature       float64   `json:"temperature,omitempty"`
	TopP              float64   `json:"top_p,omitempty"`
	N                 int64     `json:"n,omitempty"`
	Stream            bool      `json:"stream,omitempty"`
	MaxTokens         int64     `json:"max_tokens,omitempty"`
	RepetitionPenalty float64   `json:"repetition_penalty,omitempty"`
	UpdateInterval    float64   `json:"update_interval,omitempty"`
}

// Choice.
type Choice struct {
	Message      Message `json:"message"`
	Index        int32   `json:"index"`
	FinishReason string  `json:"finish_reason"`
}

// Usage.
type Usage struct {
	PromptTokens     int32 `json:"prompt_tokens"`
	CompletionTokens int32 `json:"completion_tokens"`
	TotalTokens      int32 `json:"total_tokens"`
}

// Output chat/completions object.
type ChatCompletionsOut struct {
	// https://developers.sber.ru/docs/ru/gigachat/api/reference#post-chat-completions
	Choices []Choice           `json:"choices"`
	Created timehelper.UnixSec `json:"created"`
	Model   string             `json:"model"`
	Usage   Usage              `json:"usage"`
	Object  string             `json:"object"`
}

// ChatCompletions returns the model's response taking into account the provided messages.
func ChatCompletions(ctx context.Context, currToken *common.Token, chatCompletionsIn *ChatCompletionsIn) (*ChatCompletionsOut, error) {
	// https://developers.sber.ru/docs/ru/gigachat/api/reference#post-chat-completions
	auth, err := common.AuthBearer(ctx, currToken)
	if err != nil {
		return nil, err
	}
	url, _ := url.JoinPath(baseApiUrl, "chat/completions")
	h := make(http.Header)
	h.Set("Authorization", auth)
	h.Set("Content-Type", "application/json")
	return httphelper.RestInOut[ChatCompletionsIn, ChatCompletionsOut, common.ErrorOut](
		context.Background(), http.DefaultClient, http.MethodPost, url, h, chatCompletionsIn)
}
