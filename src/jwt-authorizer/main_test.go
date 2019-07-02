package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestValidTokenAuthentication(t *testing.T) {
	request := events.APIGatewayCustomAuthorizerRequest{
		AuthorizationToken: "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.drt_po6bHhDOF_FJEHTrK-KD8OGjseJZpHwHIgsnoTM",
		MethodArn:          "testARN",
	}
	response, err := handler(request)
	assert.Nil(t, err)

	assert.NotNil(t, response)
	assert.Equal(t, response.PolicyDocument.Version, "2012-10-17", "Versions do not match")
	assert.Equal(t, response.PolicyDocument.Statement, []events.IAMPolicyStatement{
		{
			Action:   []string{"execute-api:Invoke"},
			Effect:   "Allow",
			Resource: []string{"testARN"},
		},
	}, "Policy Statements do not match")
}

func TestInvalidToken(t *testing.T) {
	request := events.APIGatewayCustomAuthorizerRequest{
		AuthorizationToken: "INVALID TOKEN",
		MethodArn:          "testARN",
	}
	_, err := handler(request)
	assert.NotNil(t, err)
}
