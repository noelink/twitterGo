package secretmanager

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	awsgo "github.com/ngonzalezo/twitterGo"
	"github.com/ngonzalezo/twitterGo/models"
)

func GetSecret(secretName string) (models.Secret, error) {
	var datosSecret models.Secret
	fmt.Println("Pido secreto")

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		fmt.Println(err.Error())
		return datosSecret, err
	}

	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println(" > Lectura de secret OK " + secretName)
	return datosSecret, nil
}
