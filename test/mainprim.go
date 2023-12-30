package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/solsw/sber/common"
	"github.com/solsw/sber/gigachat"
)

func getIdSecret() (string, string, error) {
	id := os.Getenv("SBER_GIGACHAT_CLIENT_ID")
	if id == "" {
		return "", "", errors.New("no Client Id")
	}
	secret := os.Getenv("SBER_GIGACHAT_CLIENT_SECRET")
	if secret == "" {
		return "", "", errors.New("no Client Secret")
	}
	return id, secret, nil
}

func mainprim() error {
	id, secret, err := getIdSecret()
	if err != nil {
		return err
	}
	token, err := common.GetToken(context.Background(), id, secret)
	if err != nil {
		return err
	}

	mm, err := gigachat.Models(context.Background(), token)
	if err != nil {
		return err
	}
	bb, _ := json.MarshalIndent(mm, "", "  ")
	fmt.Println(string(bb))

	m, err := gigachat.ModelsModel(context.Background(), token, mm.Data[0].Id)
	if err != nil {
		return err
	}
	bb, _ = json.MarshalIndent(m, "", "  ")
	fmt.Println(string(bb))

	tcIn := gigachat.TokensCountIn{
		Model: mm.Data[0].Id,
		Input: []string{"Почему не стоит выходить из комнаты?", "В чем заключается ошибка?"},
	}
	tcOut, err := gigachat.TokensCount(context.Background(), token, &tcIn)
	if err != nil {
		return err
	}
	bb, _ = json.MarshalIndent(tcOut, "", "  ")
	fmt.Println(string(bb))

	ccIn := gigachat.ChatCompletionsIn{
		Model: "GigaChat-Pro",
		Messages: []gigachat.Message{
			{
				Role:    "user",
				Content: "Расскажи мне сказку про Царя-Колбаску.",
			},
		},
	}
	ccOut, err := gigachat.ChatCompletions(context.Background(), token, &ccIn)
	if err != nil {
		return err
	}
	bb, _ = json.MarshalIndent(ccOut, "", "  ")
	fmt.Println(string(bb))
	return nil
}
