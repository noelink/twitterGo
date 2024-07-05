package routers

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/ngonzalezo/twitterGo/bd"
	"github.com/ngonzalezo/twitterGo/models"
	"io"
	"strings"
)

type readSeeker struct {
	io.Reader
}

func (rs *readSeeker) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}

func UploadImage(ctx context.Context, uploadType string, request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {

	var response models.RespApi
	response.Status = 400

	email := claim.Email
	var fileName string
	var user models.Usuario
	bucket := aws.String(ctx.Value(models.Key("bucketName")).(string))
	requestBody := request.Body

	switch uploadType {
	case "A":
		{
			fileName = "avatars/" + strings.Split(email, ".")[0] + ".jpg"
			user.Avatar = fileName
		}
	case "B":
		{
			fileName = "banners/" + email + ".jpg"
			user.Banner = fileName
		}
	}

	var imageRequest models.ImageRequest

	jsonError := json.Unmarshal([]byte(requestBody), &imageRequest)
	if jsonError != nil {
		response.Status = 500
		response.Message = jsonError.Error()
		return response
	}
	sendImage, err := base64.StdEncoding.DecodeString(imageRequest.Image)
	if err != nil {
		response.Status = 500
		response.Message = "base64->" + err.Error() + imageRequest.Image
		return response
	}

	slideBiteImage := bytes.NewReader(sendImage)
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, slideBiteImage); err != nil {
		response.Status = 500
		response.Message = "Error file->" + err.Error()
		return response
	}
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")})
	if err != nil {
		response.Status = 500
		response.Message = "Error aws->" + err.Error()
		return response
	}
	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: bucket,
		Key:    aws.String(fileName),
		Body:   &readSeeker{buf},
	})
	if err != nil {
		response.Status = 500
		response.Message = "error bucket->" + err.Error()
		return response
	}
	status, err := bd.ModificoRegistro(user, email)
	if err != nil || !status {
		response.Status = 400
		response.Message = "Error al modificar la base de datos"
		return response
	}

	response.Status = 200
	response.Message = "Todo ok"
	return response
}
