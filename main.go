// https://aws.amazon.com/blogs/compute/announcing-go-support-for-aws-lambda/

package main

import (
	//"errors"
	"log"
	"strconv"

	//"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type request struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(req request) (string, error) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request: %s\n", req)

	return "response for: id=" + strconv.Itoa(req.ID) + " value=" + req.Value, nil
}

func main() {
	lambda.Start(Handler)
}
