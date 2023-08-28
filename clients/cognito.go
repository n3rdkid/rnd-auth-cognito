package clients

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type CognitoClient interface {
	SignUp(email string, password string) (string, error)
	ConfirmSignUp(email string, code string) (string, error)
	SignIn(email string, password string) (*cognito.InitiateAuthOutput, error)
}

type awsCognitoClient struct {
	cognitoClient *cognito.CognitoIdentityProvider
	appClientId   string
}

func NewCognitoClient(cognitoRegion string, cognitoAppClientID string) CognitoClient {
	conf := &aws.Config{
		Region: aws.String(cognitoRegion),
	}
	sess, err := session.NewSession(conf)
	client := cognito.New(sess)
	if err != nil {
		panic(err)
	}
	return &awsCognitoClient{
		cognitoClient: client,
		appClientId:   cognitoAppClientID,
	}
}

func (ctx *awsCognitoClient) SignUp(email string, password string) (string, error) {
	user := &cognito.SignUpInput{
		ClientId: aws.String(ctx.appClientId),
		Username: aws.String(email),
		Password: aws.String(password),
	}
	result, err := ctx.cognitoClient.SignUp(user)
	if err != nil {
		return "", err
	}
	return result.String(), err
}

func (ctx *awsCognitoClient) ConfirmSignUp(email string, code string) (string, error) {
	user := &cognito.ConfirmSignUpInput{
		ClientId:         aws.String(ctx.appClientId),
		Username:         aws.String(email),
		ConfirmationCode: aws.String(code),
	}
	result, err := ctx.cognitoClient.ConfirmSignUp(user)
	if err != nil {
		return "", err
	}
	return result.String(), err
}

func (ctx *awsCognitoClient) SignIn(email string, password string) (*cognito.InitiateAuthOutput, error) {
	auth := &cognito.InitiateAuthInput{
		ClientId: aws.String(ctx.appClientId),
		AuthFlow: aws.String("USER_PASSWORD_AUTH"),
		AuthParameters: aws.StringMap(map[string]string{
			"USERNAME": email,
			"PASSWORD": password,
		}),
	}
	result, err := ctx.cognitoClient.InitiateAuth(auth)
	if err != nil {
		return nil, err
	}
	fmt.Println(*result.AuthenticationResult.IdToken)
	return result, nil
}
