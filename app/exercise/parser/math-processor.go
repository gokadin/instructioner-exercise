package parser

import (
	"github.com/knetic/govaluate"
)

func compute(str string, variables map[string]interface{}) (interface{}, error) {
	expression, err := govaluate.NewEvaluableExpression(str)
	if err != nil {
		return nil, err
	}
	
	evaluated, err := expression.Evaluate(variables)
	if err != nil {
		return nil, err
	}
	
    return evaluated, nil
}