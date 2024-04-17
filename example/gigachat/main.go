package main

import (
	"context"
	"errors"
	"os"

	"github.com/solsw/errorhelper"
	"github.com/solsw/sber/auth"
)

func getIdSecret() (string, string, error) {
	id := os.Getenv("SBER_GIGACHAT_CLIENT_ID")
	if id == "" {
		return "", "", errorhelper.CallerError(errors.New("no Client Id"))
	}
	secret := os.Getenv("SBER_GIGACHAT_CLIENT_SECRET")
	if secret == "" {
		return "", "", errorhelper.CallerError(errors.New("no Client Secret"))
	}
	return id, secret, nil
}

func main() {
	id, secret, err := getIdSecret()
	if err != nil {
		panic(err)
	}
	t, err := auth.NewToken(context.Background(), id, secret)
	if err != nil {
		panic(err)
	}
	var errs []error
	mm, err := models(t.AccessToken)
	errs = append(errs, err)
	errs = append(errs, modelsModel(t, mm.Data[0].Id))
	errs = append(errs, tokensCount(t, mm.Data[0].Id))
	if jerr := errors.Join(errs...); jerr != nil {
		panic(jerr)
	}
}
