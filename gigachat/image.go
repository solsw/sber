package gigachat

import (
	"context"
	"encoding/xml"
	"errors"
	"html"
	"net/http"
	"net/url"

	"github.com/solsw/errorhelper"
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
	u, err := url.JoinPath(baseApiUrl, "files", fileId, "content")
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	h := make(http.Header)
	h.Set("Authorization", "Bearer "+accessToken)
	h.Set("Accept", "application/jpg")
	out, err := rest.BodyBody[common.OutError](
		ctx, http.DefaultClient, http.MethodGet, u, h, nil, rest.IsNotStatusOK)
	return out, errorhelper.CallerError(err)
}
