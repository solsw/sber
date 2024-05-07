package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/solsw/errorhelper"
	"github.com/solsw/sber/auth"
	"github.com/solsw/sber/gigachat"
)

func models(accessToken string) (*gigachat.ModelsOut, error) {
	mm, err := gigachat.Models(context.Background(), accessToken)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	bb, _ := json.MarshalIndent(mm, "", "  ")
	fmt.Println(string(bb))
	return mm, nil
}

func tokensCount(t *auth.Token, mn string) error {
	if err := auth.VetToken(context.Background(), t); err != nil {
		return errorhelper.CallerError(err)
	}
	tcIn := gigachat.TokensCountIn{
		Model: mn,
		Input: []string{"Почему не стоит выходить из комнаты?", "В чём заключается ошибка?"},
	}
	tcOut, err := gigachat.TokensCount(context.Background(), t.AccessToken, &tcIn)
	if err != nil {
		return errorhelper.CallerError(err)
	}
	bb, _ := json.MarshalIndent(tcOut, "", "  ")
	fmt.Println(string(bb))
	return nil
}

func chatCompletions(t *auth.Token, mn string) error {
	if err := auth.VetToken(context.Background(), t); err != nil {
		return errorhelper.CallerError(err)
	}
	ccIn := gigachat.ChatCompletionsIn{
		Model: mn,
		Messages: []gigachat.Message{
			{
				Role:    "user",
				Content: "Расскажи мне сказку про Царя-Колбаску.",
			},
		},
	}
	ccOut, err := gigachat.ChatCompletions(context.Background(), t.AccessToken, &ccIn)
	if err != nil {
		return errorhelper.CallerError(err)
	}
	bb, _ := json.MarshalIndent(ccOut, "", "  ")
	fmt.Println(string(bb))
	return nil
}

func image(t *auth.Token, mn string) (*gigachat.ChatCompletionsOut, error) {
	if err := auth.VetToken(context.Background(), t); err != nil {
		return nil, errorhelper.CallerError(err)
	}
	ccIn := gigachat.ChatCompletionsIn{
		Model: mn,
		Messages: []gigachat.Message{
			{
				Role:    "user",
				Content: "Нарисуй Царя-Колбаску.",
			},
		},
	}
	ccOut, err := gigachat.ChatCompletions(context.Background(), t.AccessToken, &ccIn)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	bb, _ := json.MarshalIndent(ccOut, "", "  ")
	fmt.Println(string(bb))
	return ccOut, nil
}

func filesFileId(t *auth.Token, content string) error {
	if err := auth.VetToken(context.Background(), t); err != nil {
		return errorhelper.CallerError(err)
	}
	fid, err := gigachat.FileIdFromMessageContent(content)
	if err != nil {
		return errorhelper.CallerError(err)
	}
	fmt.Println(fid)
	bb, err := gigachat.FilesFileId(context.Background(), t.AccessToken, fid)
	if err != nil {
		return errorhelper.CallerError(err)
	}
	if err := os.WriteFile("Царь-Колбаска.jpg", bb, os.ModePerm); err != nil {
		return errorhelper.CallerError(err)
	}
	return nil
}

func embeddings(t *auth.Token) error {
	if err := auth.VetToken(context.Background(), t); err != nil {
		return errorhelper.CallerError(err)
	}
	eIn := gigachat.EmbeddingsIn{
		Model: "Embeddings",
		Input: []string{"Однажды в студёную, зимнюю пору я из лесу вышел, был сильный мороз."},
	}
	eOut, err := gigachat.Embeddings(context.Background(), t.AccessToken, &eIn)
	if err != nil {
		return errorhelper.CallerError(err)
	}
	bb, _ := json.MarshalIndent(eOut, "", "  ")
	fmt.Println(string(bb))
	return nil
}
