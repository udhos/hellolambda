// https://aws.amazon.com/blogs/compute/announcing-go-support-for-aws-lambda/

package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

type Response struct {
	Result string `json:"result"`
}

var NonPositiveID = fmt.Errorf("refusing to handle non-positive ID")

func Handler(req Request) (Response, error) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request: id=%d value=%s", req.ID, req.Value)

	resp := Response{}

	if req.ID < 1 {
		log.Print(NonPositiveID)
		return resp, NonPositiveID
	}

	resp.Result = "response for: id=" + strconv.Itoa(req.ID) + " value=" + req.Value

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
