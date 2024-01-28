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

func modelsModel(t *common.Token, mn string) error {
	m, err := gigachat.ModelsModel(context.Background(), t, mn)
	if err != nil {
		return err
	}
	bb, _ := json.MarshalIndent(m, "", "  ")
	fmt.Println(string(bb))
	return nil
}

func tokensCount(t *common.Token, mn string) error {
	tcIn := gigachat.TokensCountIn{
		Model: mn,
		Input: []string{"Почему не стоит выходить из комнаты?", "В чем заключается ошибка?"},
	}
	tcOut, err := gigachat.TokensCount(context.Background(), t, &tcIn)
	if err != nil {
		return err
	}
	bb, _ := json.MarshalIndent(tcOut, "", "  ")
	fmt.Println(string(bb))
	return nil
}

func chatCompletions(t *common.Token, mn string) error {
	ccIn := gigachat.ChatCompletionsIn{
		Model: mn,
		Messages: []gigachat.Message{
			{
				Role:    "user",
				Content: "Расскажи мне сказку про Царя-Колбаску.",
			},
		},
	}
	ccOut, err := gigachat.ChatCompletions(context.Background(), t, &ccIn)
	if err != nil {
		return err
	}
	bb, _ := json.MarshalIndent(ccOut, "", "  ")
	fmt.Println(string(bb))
	return nil
}

func image(t *common.Token, mn string) error {
	ccIn := gigachat.ChatCompletionsIn{
		Model: mn,
		Messages: []gigachat.Message{
			{
				Role:    "user",
				Content: "Нарисуй Царя-Колбаску.",
			},
		},
	}
	ccOut, err := gigachat.ChatCompletions(context.Background(), t, &ccIn)
	if err != nil {
		return err
	}
	bb, _ := json.MarshalIndent(ccOut, "", "  ")
	fmt.Println(string(bb))
	return nil
}

func filesFileId(t *common.Token) error {
	content := "\u003cimg src=\"8882p0090sa2010q1482g4g539b74mjaan42amr89s02gmta816q4mtt9rdjy1jd810q4n2xan57yq8z1h77c8dct8wky001355jy30f0ge0q9y85me2emapagd7j1jm150qyma3awdqjqam1x77cmj3b902wmrt150japaq0md4xvtstm\" fuse=\"true\"/\u003e"
	fid, err := gigachat.FileIdFromMessageContent(content)
	if err != nil {
		return err
	}
	fmt.Println(fid)
	bb, err := gigachat.FilesFileId(context.Background(), t, fid)
	if err != nil {
		return err
	}
	if err := os.WriteFile("s.jpg", bb, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func embeddings(t *common.Token) error {
	eIn := gigachat.EmbeddingsIn{
		Model: "Embeddings",
		Input: []string{"Однажды в студёную, зимнюю пору."},
	}
	eOut, err := gigachat.Embeddings(context.Background(), t, &eIn)
	if err != nil {
		return err
	}
	bb, _ := json.MarshalIndent(eOut, "", "  ")
	fmt.Println(string(bb))
	return nil
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

	// if err := modelsModel(token, mm.Data[0].Id); err != nil {
	// 	return err
	// }
	// if err := tokensCount(token, mm.Data[0].Id); err != nil {
	// 	return err
	// }
	// if err := chatCompletions(token, "GigaChat-Pro"); err != nil {
	// 	return err
	// }
	// if err := image(token, "GigaChat-Pro"); err != nil {
	// 	return err
	// }
	if err := filesFileId(token); err != nil {
		return err
	}
	// if err := embeddings(token); err != nil {
	// 	return err
	// }

	return nil
}
