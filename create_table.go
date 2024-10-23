package dynago

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type (
	Attribute struct {
		AttributeName string
		AttributeType string
	}

	Key struct {
		AttributeName string
		KeyType       string
	}

	ProvisionedTP struct {
		Read  int64
		Write int64
	}
)

type CreateTableParams struct {
	Attributes    []Attribute
	Keys          []Key
	ProvisionedTP ProvisionedTP

	From *dynamodb.CreateTableInput
}

var (
	DefaultInput = &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: Pointer("_id"),
				AttributeType: Pointer("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: Pointer("_id"),
				KeyType:       Pointer("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  Pointer[int64](10),
			WriteCapacityUnits: Pointer[int64](10),
		},
	}
)

func (c *Dynago) NewTable(name string, params ...CreateTableParams) error {
	input := DefaultInput

	var p *CreateTableParams
	if len(params) > 0 {
		p = &params[0]
	}

	if p != nil {
		input = FromParams(*p)
	}

	if p.From != nil {
		input = p.From
	}

	input.TableName = &name

	_, err := c.Svc.CreateTable(input)

	return err
}

func FromParams(params CreateTableParams) *dynamodb.CreateTableInput {
	attributes := []*dynamodb.AttributeDefinition{}
	for _, atr := range params.Attributes {
		attributes = append(attributes, &dynamodb.AttributeDefinition{
			AttributeName: &atr.AttributeName,
			AttributeType: &atr.AttributeType,
		})
	}

	keys := []*dynamodb.KeySchemaElement{}
	for _, key := range params.Keys {
		keys = append(keys, &dynamodb.KeySchemaElement{
			AttributeName: &key.AttributeName,
			KeyType:       &key.KeyType,
		})
	}

	provisioned_tp := &dynamodb.ProvisionedThroughput{
		WriteCapacityUnits: &params.ProvisionedTP.Write,
		ReadCapacityUnits:  &params.ProvisionedTP.Read,
	}

	new := dynamodb.CreateTableInput{
		AttributeDefinitions:  attributes,
		KeySchema:             keys,
		ProvisionedThroughput: provisioned_tp,
	}

	return &new
}
