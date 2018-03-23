// https://aws.amazon.com/blogs/compute/announcing-go-support-for-aws-lambda/

package main

import (
	"log"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

func Handler(req Request) (string, error) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request: id=%d value=%s", req.ID, req.Value)

	return "response for: id=" + strconv.Itoa(req.ID) + " value=" + req.Value, nil
}

func main() {
	lambda.Start(Handler)
}
