package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func main() {
	lambda.Start(router)
}

func router(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	log.Printf("Received request: Path=%s, Method=%s\n", request.RequestContext.HTTP.Path, request.RequestContext.HTTP.Method)

	switch request.RequestContext.HTTP.Path {
	case "/login":
		if request.RequestContext.HTTP.Method == http.MethodPost {
			return handleLogin(request)
		}
	default:
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 404,
			Body:       "Path not found",
		}, nil
	}
	return events.APIGatewayV2HTTPResponse{
		StatusCode: 405,
		Body:       "Method not allowed",
	}, nil
}

func handleLogin(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	type RequestBody struct {
		Email string `json:"email"`
	}

	var body RequestBody
	if err := json.Unmarshal([]byte(request.Body), &body); err != nil {
		log.Println("Failed to parse request body:", err)
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 400,
			Body:       "Invalid request body",
		}, nil
	}

	log.Printf("Received email: %s\n", body.Email)

	// Initialize AWS SDK session
	sess := session.Must(session.NewSession())
	cognito := cognitoidentityprovider.New(sess, aws.NewConfig().WithRegion("eu-north-1"))

	// Trigger email OTP using Cognito Admin API
	_, err := cognito.AdminCreateUser(&cognitoidentityprovider.AdminCreateUserInput{
		UserPoolId: aws.String("your_user_pool_id"), // Replace with your user pool ID
		Username:   aws.String(body.Email),
	})
	if err != nil {
		log.Println("Failed to send OTP:", err)
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 500,
			Body:       "Failed to send OTP",
		}, nil
	}

	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body:       "OTP sent successfully!",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}
