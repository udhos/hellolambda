// https://aws.amazon.com/blogs/compute/announcing-go-support-for-aws-lambda/

package main_test

import (
	"testing"

	//"github.com/aws-samples/lambda-go-samples"

	main "github.com/udhos/hellolambda"

	//"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	tests := []struct {
		request main.Request
		expect  main.Response
		err     error
	}{
		{
			request: main.Request{ID: 123, Value: "test123"},
			expect:  main.Response{Result: "response for: id=123 value=test123"},
			err:     nil,
		},
		{
			request: main.Request{},
			expect:  main.Response{},
			err:     main.NonPositiveID,
		},
	}

	for _, test := range tests {
		response, err := main.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response)
	}

}
