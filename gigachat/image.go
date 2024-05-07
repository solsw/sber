package gigachat

import (
	"context"
	"encoding/xml"
	"errors"
	"html"
	"net/http"
	"net/url"

	"github.com/solsw/errorhelper"
	"github.com/solsw/httphelper"
	"github.com/solsw/httphelper/rest"
	"github.com/solsw/sber/common"
)

// https://developers.sber.ru/docs/ru/gigachat/api/images-generation

type img struct {
	XMLName xml.Name `xml:"img"`
	FileId  string   `xml:"src,attr"`
}

// FileIdFromMessageContent выделяет FileId из строки Message.Content.
func FileIdFromMessageContent(content string) (string, error) {
	// https://developers.sber.ru/docs/ru/gigachat/api/reference/rest/get-file-id#zapros
	u := html.UnescapeString(content)
	var i img
	if err := xml.Unmarshal([]byte(u), &i); err != nil {
		return "", err
	}
	return i.FileId, nil
}

// FilesFileId возвращает файл изображения в бинарном представлении в формате JPG.
func FilesFileId(ctx context.Context, accessToken string, fileId string) ([]byte, error) {
	// https://developers.sber.ru/docs/ru/gigachat/api/reference/rest/get-file-id
	if accessToken == "" {
		return nil, errorhelper.CallerError(errors.New("no accessToken"))
	}
	u, _ := url.JoinPath(baseApiUrl, "files", fileId, "content")
	rq, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	rq.Header.Set("Authorization", "Bearer "+accessToken)
	rq.Header.Set("Accept", "application/jpg")
	jpg, err := rest.ReqBody[common.OutError](http.DefaultClient, rq, httphelper.IsNotStatusOK)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return jpg, nil
}
