package dynago

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Dynago struct {
	Svc *dynamodb.DynamoDB
}

func NewClient(provider client.ConfigProvider, cfgs ...*aws.Config) *Dynago {
	dynamo := dynamodb.New(provider, cfgs...)

	return &Dynago{
		Svc: dynamo,
	}
}
