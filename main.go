// https://aws.amazon.com/blogs/compute/announcing-go-support-for-aws-lambda/

package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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

	listBuckets()

	resp := Response{}

	if req.ID < 1 {
		// refuse non-positive ID
		log.Print(NonPositiveID)
		return resp, NonPositiveID
	}

	resp.Result = "response for: id=" + strconv.Itoa(req.ID) + " value=" + req.Value

	return resp, nil
}

func listBuckets() {
	svc := s3.New(session.New(&aws.Config{Region: aws.String("us-east-1")}))
	buckets, err := svc.ListBuckets(nil)
	log.Printf("listBuckets: %q error=%v", buckets, err)
}

func main() {
	lambda.Start(Handler)
}
