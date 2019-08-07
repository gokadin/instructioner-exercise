package parser

import (
    "exercise/app/exercise/model"
    "fmt"
    "github.com/labstack/gommon/log"
    "regexp"
    "strconv"
)

func Parse(exercise *model.Exercise) error {
    variables := resolveVariables(exercise.Variables)
    
    err := parseQuestion(exercise, variables)
    if err != nil {
        return err
    }

    err = parseHints(exercise, variables)
    if err != nil {
        return err
    }

    err = parseAnswerFields(exercise, variables)
    if err != nil {
        return err
    }
    
    return nil
}

func parseQuestion(exercise *model.Exercise, variables map[string]interface{}) error {
    exercise.Question = replaceVariables(exercise.Question, variables)
    exercise.Question = replaceCompute(exercise.Question, variables)
    return nil
}

func parseHints(exercise *model.Exercise, variables map[string]interface{}) error {
    for _, hint := range exercise.Hints {
        err := parseHint(hint, variables)
        if err != nil {
            return err
        }
    }
    
    return nil
}

func parseHint(hint *model.Hint, variables map[string]interface{}) error {
    hint.Content = replaceVariables(hint.Content, variables)
    hint.Content = replaceCompute(hint.Content, variables)
    return nil
}

func parseAnswerFields(exercise *model.Exercise, variables map[string]interface{}) error {
    for _, answerField := range exercise.Answer.Choices {
        err := parseAnswerField(answerField, variables)
        if err != nil {
            return err
        }
    }

    return nil
}

func parseAnswerField(answerField *model.AnswerField, variables map[string]interface{}) error {
    answerField.Content = replaceVariables(answerField.Content, variables)
    answerField.Content = replaceCompute(answerField.Content, variables)
    return nil
}

func resolveVariables(variableDefinitions []*model.Variable) map[string]interface{} {
    variables := make(map[string]interface{}, len(variableDefinitions))
    
    for _, variableDefinition := range variableDefinitions {
        value := resolveVariable(variableDefinition)
        variables[variableDefinition.Name] = value
    }
    
    return variables
}

func resolveVariable(variableDefinition *model.Variable) int {
    value, _ := strconv.Atoi(variableDefinition.Default)
    return value
}

func replaceVariables(str string, variables map[string]interface{}) string {
    re, err := regexp.Compile(`@var\((\w+)\)`)
    if err != nil {
        log.Error("Could not replace variable", err)
    }
    
    str = re.ReplaceAllStringFunc(str, func(s string) string {
       variableName := re.ReplaceAllString(s, "$1")
       return fmt.Sprintf("%d", variables[variableName].(int))
    })
    
    return str
}

func replaceCompute(str string, variables map[string]interface{}) string {
    re, err := regexp.Compile(`@compute\((.+)\)`)
    if err != nil {
        log.Error("Could not replace compute", err)
    }

    str = re.ReplaceAllStringFunc(str, func(s string) string {
        computeStr := re.ReplaceAllString(s, "$1")
        computed, err := compute(computeStr, variables)
	    if err != nil {
	        log.Error("Could not compute expression", err)
	    }
	    return fmt.Sprintf("%d", int(computed.(float64)))
    })

    return str
}
