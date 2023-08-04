package uptrend95User

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/devchrisar/uptrend95User/awsgo"
	"github.com/devchrisar/uptrend95User/db"
	"github.com/devchrisar/uptrend95User/models"
	"os"
)

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.AwsGoInit()
	if !ValidateParams() {
		fmt.Println("Missing required environment variables")
		err := errors.New("missing required environment variables")
		return event, err
	}
	var values models.Signup
	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			values.UserEmail = att
			fmt.Println(" > email: " + att)
		case "sub":
			values.UserUUID = att
			fmt.Println(" > sub: " + att)
		}
	}
	err := db.GetSecret()
	if err != nil {
		fmt.Println(" > error getting secret" + err.Error())
		return event, err
	}
	err = db.SignUp(values)
	return event, err
}

func ValidateParams() bool {
	var valid bool
	_, valid = os.LookupEnv("SecretName")
	return valid
}
