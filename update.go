package dynago

import "github.com/aws/aws-sdk-go/service/dynamodb"

type Update struct {
	Table      string
	Key        Map
	Patch      Map
	Expression string
	Returns    string

	From *dynamodb.UpdateItemInput
}

func (c *Dynago) Update(u *Update) (*dynamodb.UpdateItemOutput, error) {
	input := dynamodb.UpdateItemInput{
		TableName:                 &u.Table,
		Key:                       u.Key,
		ExpressionAttributeValues: u.Patch,
		UpdateExpression:          &u.Expression,
		ReturnValues:              &u.Returns,
	}

	if u.From != nil {
		input = *u.From
	}

	return c.Svc.UpdateItem(&input)
}
