package dynago

import "github.com/aws/aws-sdk-go/service/dynamodb"

type Delete struct {
	Table string
	Key   Map

	From *dynamodb.DeleteItemInput
}

func (c *Dynago) Delete(d Delete) (*dynamodb.DeleteItemOutput, error) {
	input := dynamodb.DeleteItemInput{
		TableName: &d.Table,
		Key:       d.Key,
	}

	if d.From != nil {
		input = *d.From
	}

	return c.Svc.DeleteItem(&input)
}
