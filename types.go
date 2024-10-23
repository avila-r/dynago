package dynago

import "github.com/aws/aws-sdk-go/service/dynamodb"

type (
	Map map[string]*dynamodb.AttributeValue
)
