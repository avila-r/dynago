package dynago

import "github.com/aws/aws-sdk-go/service/dynamodb"

type ListTableParams struct {
	StartFrom string `min:"3" type:"string"`
	Limit     int64  `min:"1" type:"integer"`

	From *dynamodb.ListTablesInput
}

func (c *Dynago) ListTables(params ...ListTableParams) ([]string, error) {
	p := ListTableParams{}
	if len(params) > 0 {
		p = params[0]
	}

	input := dynamodb.ListTablesInput{
		ExclusiveStartTableName: &p.StartFrom,
		Limit:                   &p.Limit,
	}

	if p.From != nil {
		input = *p.From
	}

	tables := []string{}

	for {
		output, err := c.Svc.ListTables(&input)

		if err != nil {
			return nil, err
		}

		for _, t := range output.TableNames {
			tables = append(tables, *t)
		}

		input.ExclusiveStartTableName = output.LastEvaluatedTableName
		if output.LastEvaluatedTableName == nil {
			break
		}
	}

	return tables, nil
}
