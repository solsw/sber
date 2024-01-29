package gigachat

import (
	"context"
	"encoding/xml"
	"html"
	"net/http"
	"net/url"

	"github.com/solsw/httphelper/rest"
	"github.com/solsw/sber/common"
)

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

// FilesFileId возвращает файл изображения в бинарном представлении, в формате JPG.
func FilesFileId(ctx context.Context, currToken *common.Token, fileId string) ([]byte, error) {
	// https://developers.sber.ru/docs/ru/gigachat/api/reference/rest/get-file-id
	auth, err := common.AuthBearer(ctx, currToken)
	if err != nil {
		return nil, err
	}
	url, err := url.JoinPath(baseApiUrl, "files", fileId, "content")
	if err != nil {
		return nil, err
	}
	h := make(http.Header)
	h.Set("Authorization", auth)
	h.Set("Accept", "application/jpg")
	return rest.BodyBody[common.OutError](
		ctx, http.DefaultClient, http.MethodGet, url, h, nil, rest.IsNotStatusOK)
}
