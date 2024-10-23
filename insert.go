package dynago

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Insert struct {
	Table string
	Item  any

	From *dynamodb.PutItemInput
}

func (c *Dynago) Insert(i Insert) (*dynamodb.PutItemOutput, error) {
	document, err := dynamodbattribute.MarshalMap(i.Item)

	if err != nil {
		return nil, err
	}

	input := dynamodb.PutItemInput{
		TableName: &i.Table,
		Item:      document,
	}

	if i.From != nil {
		input = *i.From
	}

	return c.Svc.PutItem(&input)
}
