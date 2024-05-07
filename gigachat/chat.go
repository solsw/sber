package gigachat

import (
	"context"
	"errors"
	"net/http"
	"net/url"

	"github.com/solsw/errorhelper"
	"github.com/solsw/httphelper"
	"github.com/solsw/httphelper/rest"
	"github.com/solsw/timehelper"
	"github.com/solsw/sber/common"
)

// Message.
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// FewShotExample.
type FewShotExample struct {
	Request string `json:"request"`
	Parames any    `json:"params"`
}

// Function.
type Function struct {
	Name             string           `json:"name"`
	Description      string           `json:"description,omitempty"`
	Parameters       any              `json:"parameters"`
	FewShotExamples  []FewShotExample `json:"few_shot_examples,omitempty"`
	ReturnParameters any              `json:"return_parameters,omitempty"`
}

// Input chat/completions object.
type ChatCompletionsIn struct {
	// https://developers.sber.ru/docs/ru/gigachat/api/reference/rest/post-chat#zapros

	// https://developers.sber.ru/docs/ru/gigachat/api/selecting-a-model
	Model string `json:"model"`
	// https://developers.sber.ru/docs/ru/gigachat/api/keeping-context
	// https://developers.sber.ru/docs/ru/gigachat/api/images-generation
	Messages []Message `json:"messages"`
	// https://developers.sber.ru/docs/ru/gigachat/api/function-calling
	FunctionCall any        `json:"function_call"`
	Functions    []Function `json:"functions"`
	Temperature  float64    `json:"temperature,omitempty"`
	TopP         float64    `json:"top_p,omitempty"`
	N            int64      `json:"n,omitempty"`
	// https://developers.sber.ru/docs/ru/gigachat/api/response-token-streaming
	Stream            bool    `json:"stream,omitempty"`
	MaxTokens         int64   `json:"max_tokens,omitempty"`
	RepetitionPenalty float64 `json:"repetition_penalty,omitempty"`
	UpdateInterval    float64 `json:"update_interval,omitempty"`
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
	// https://developers.sber.ru/docs/ru/gigachat/api/reference/rest/post-chat#responses
	Choices []Choice           `json:"choices"`
	Created timehelper.UnixSec `json:"created"`
	Model   string             `json:"model"`
	Usage   Usage              `json:"usage"`
	Object  string             `json:"object"`
}

// ChatCompletions возвращает ответ модели с учётом переданных сообщений.
func ChatCompletions(ctx context.Context, accessToken string, chatCompletionsIn *ChatCompletionsIn) (*ChatCompletionsOut, error) {
	// https://developers.sber.ru/docs/ru/gigachat/api/reference/rest/post-chat
	if accessToken == "" {
		return nil, errorhelper.CallerError(errors.New("no accessToken"))
	}
	u, _ := url.JoinPath(baseApiUrl, "chat/completions")
	body, err := httphelper.JsonBody(chatCompletionsIn)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	rq, err := http.NewRequestWithContext(ctx, http.MethodPost, u, body)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	rq.Header.Set("Authorization", "Bearer "+accessToken)
	rq.Header.Set("Content-Type", "application/json")
	out, err := rest.ReqJson[ChatCompletionsOut, common.OutError](http.DefaultClient, rq, httphelper.IsNotStatusOK)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return out, nil
}
