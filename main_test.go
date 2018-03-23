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
		expect  string
		err     error
	}{
		{
			// Test that the handler responds with the correct response
			// when a valid name is provided in the HTTP body
			request: main.Request{ID: 123, Value: "test123"},
			expect:  "response for: id=123 value=test123",
			err:     nil,
		},
		{
			// Test that the handler responds ErrNameNotProvided
			// when no name is provided in the HTTP body
			request: main.Request{},
			expect:  "response for: id=0 value=",
			err:     nil,
		},
	}

	for _, test := range tests {
		response, err := main.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response)
	}

}
