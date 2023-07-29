package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Parse the Slack request body into a JSON object
	var body struct {
		Text string `json:"text"`
	}
	err := json.Unmarshal([]byte(request.Body), &body)
	if err != nil {
		log.Printf("Error parsing request body: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 400}, nil
	}

	// Check if the request is coming from Slack
	if request.Headers["X-Slack-Signature"] == "" {
		log.Printf("Invalid request: missing Slack signature")
		return events.APIGatewayProxyResponse{StatusCode: 400}, nil
	}

	// Validate the Slack request signature
	// You'll need to replace the <SLACK_SIGNING_SECRET> placeholder with your app's signing secret
	// You can find your app's signing secret in the "App Credentials" section of your Slack app's "Basic Information" page
	if !validateSlackRequest(request.Body, request.Headers["X-Slack-Signature"], request.Headers["X-Slack-Request-Timestamp"], "<SLACK_SIGNING_SECRET>") {
		log.Printf("Invalid request: invalid Slack signature")
		return events.APIGatewayProxyResponse{StatusCode: 400}, nil
	}

	// Send a response back to Slack
	response := map[string]string{
		"text": fmt.Sprintf("You said: %s", body.Text),
	}
	responseJson, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshaling response: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}
	return events.APIGatewayProxyResponse{
		Body:       string(responseJson),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}

func validateSlackRequest(body string, signature string, timestamp string, signingSecret string) bool {
	// Implement Slack request signature validation logic here
	// You can find sample code for validating Slack request signatures in Go at https://github.com/slackapi/golang-slack-sdk/tree/main/examples/http-handler/verify_signature

	return true // Change this to return true if validation succeeds, or false otherwise
}
