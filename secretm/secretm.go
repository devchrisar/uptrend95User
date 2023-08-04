package secretm

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/devchrisar/uptrend95User/awsgo"
	"github.com/devchrisar/uptrend95User/models"
)

func GetSecret(s string) (models.SecretRDSJson, error) {
	var secret models.SecretRDSJson
	fmt.Println(" > GetSecret() called" + s)
	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	key, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{SecretId: aws.String(s)})
	if err != nil {
		fmt.Println(err.Error())
		return secret, err
	}
	json.Unmarshal([]byte(*key.SecretString), &secret)
	fmt.Println(" > GetSecret() returning" + s)
	return secret, nil
}
