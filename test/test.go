package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/solsw/sber/common"
	"github.com/solsw/sber/gigachat"
)

func main() {
	currToken, err := common.GetToken(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	mm, err := gigachat.Models(context.Background(), currToken)
	if err != nil {
		panic(err)
	}
	bb, _ := json.MarshalIndent(mm, "", "  ")
	fmt.Println(string(bb))

	// m, err := gigachat.ModelsModel(context.Background(), currToken, mm.Data[0].Id)
	// if err != nil {
	// 	panic(err)
	// }
	// bb, _ = json.MarshalIndent(m, "", "  ")
	// fmt.Println(string(bb))

	// tcIn := gigachat.TokensCountIn{
	// 	Model: mm.Data[0].Id,
	// 	Input: []string{"Почему не стоит выходить из комнаты?", "В чем заключается ошибка?"},
	// }
	// tco, err := gigachat.TokensCount(context.Background(), currToken, &tcIn)
	// if err != nil {
	// 	panic(err)
	// }
	// bb, _ = json.MarshalIndent(tco, "", "  ")
	// fmt.Println(string(bb))

	ccIn := gigachat.ChatCompletionsIn{
		Model: "GigaChat-Pro",
		Messages: []gigachat.Message{
			{
				Role:    "user",
				Content: "Расскажи мне сказку про Царя-Колбаску.",
			},
		},
	}
	cco, err := gigachat.ChatCompletions(context.Background(), currToken, &ccIn)
	if err != nil {
		panic(err)
	}
	bb, _ = json.MarshalIndent(cco, "", "  ")
	fmt.Println(string(bb))
}
