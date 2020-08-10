package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Body string `json:"body"`
}

func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	resBody := Response{Body: "Hello, world! from " + req.HTTPMethod}
	res, _ := json.Marshal(resBody)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(res),
		Headers: map[string]string{
			"Access-Control-Allow-Headers": "Content-Type, Authorization",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "*",
		},
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
