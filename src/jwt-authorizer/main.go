package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	jwt "github.com/dgrijalva/jwt-go"
)

var signingKey string

func init() {
	signingKey = os.Getenv("SIGNING_KEY") //read from environment
	if signingKey == "" {
		panic("Environment variable not defined: <SIGNING_KEY>")
	}
}

func handler(request events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	authorizationHeader := request.AuthorizationToken
	authorizationValue := strings.Split(authorizationHeader, " ")

	// Expected - Authorization: Bearer <TOKEN_HERE>
	if len(authorizationValue) != 2 || !strings.EqualFold("Bearer", authorizationValue[0]) {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized")
	}

	token, err := jwt.Parse(authorizationValue[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(signingKey), nil
	})

	if err != nil || token.Valid == false {
		log.Printf("==> Error: %s ||  Token valid ? %t\n", err.Error(), token.Valid)
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New(err.Error())
	}

	return generatePolicy("user", "Allow", request.MethodArn), nil

}

func main() {
	lambda.Start(handler)
}

func generatePolicy(principalID, effect, resource string) events.APIGatewayCustomAuthorizerResponse {
	authResponse := events.APIGatewayCustomAuthorizerResponse{PrincipalID: principalID}

	if effect != "" && resource != "" {
		authResponse.PolicyDocument = events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   effect,
					Resource: []string{resource},
				},
			},
		}
	}
	return authResponse
}
