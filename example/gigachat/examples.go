package main

import (
	"context"
	"encoding/json"
	"fmt"

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

func modelsModel(t *auth.Token, mn string) error {
	if err := auth.VetToken(context.Background(), t); err != nil {
		return errorhelper.CallerError(err)
	}
	m, err := gigachat.ModelsModel(context.Background(), t.AccessToken, mn)
	if err != nil {
		return errorhelper.CallerError(err)
	}
	bb, _ := json.MarshalIndent(m, "", "  ")
	fmt.Println(string(bb))
	return nil
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
