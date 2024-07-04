package routers

import (
	"bytes"
	"context"
	"encoding/base64"
	"github.com/aws/aws-lambda-go/events"
	"github.com/ngonzalezo/twitterGo/models"
	"io"
	"mime"
	"mime/multipart"
	"strings"
)

type readSeeker struct {
	io.Reader
}

func (rs *readSeeker) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}

func UploadImage(ctx context.Context, uploadType string, request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400
	IDUsuario := claim.ID.Hex()

	var fileName string
	var usuario models.Usuario

	switch uploadType {
	case "A":
		fileName = "avatars/" + IDUsuario + ".jpg"
		usuario.Avatar = fileName
	case "B":
		fileName = "banners/" + IDUsuario + ".jpg"
		usuario.Avatar = fileName
	}

	mediaType, params, err := mime.ParseMediaType(request.Headers["Content-Type"])
	if err != nil {
		r.Status = 500
		r.Message = err.Error()
		return r
	}

	if strings.HasPrefix(mediaType, "multipart/") {
		body, err := base64.StdEncoding.DecodeString(request.Body)
		if err != nil {
			r.Status = 500
			r.Message = err.Error()
			return r
		}

		mr := multipart.NewReader(bytes.NewReader(body), params["boundary"])
		p, err := mr.NextPart()
		if err != nil && err != io.EOF {
			r.Status = 400
			return r
		}

	} else {
		r.Message = "Debe enviar una imagen con el 'Content-Type' de tipo 'multipart/' en el Header"
		r.Status = 400
		return r
	}

}
