package dynago

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Query struct {
	Table string
	Key   Map

	From *dynamodb.GetItemInput
}

func (c *Dynago) FindTo(q Query, t *any) error {
	input := dynamodb.GetItemInput{
		TableName: &q.Table,
		Key:       q.Key,
	}

	if q.From != nil {
		input = *q.From
	}

	output, err := c.Svc.GetItem(&input)

	if err != nil {
		return err
	}

	return dynamodbattribute.UnmarshalMap(output.Item, &t)
}

func (c *Dynago) Find(q Query) (*dynamodb.GetItemOutput, error) {
	input := dynamodb.GetItemInput{
		TableName: &q.Table,
		Key:       q.Key,
	}

	if q.From != nil {
		input = *q.From
	}

	return c.Svc.GetItem(&input)
}
