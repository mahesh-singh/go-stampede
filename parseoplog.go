package gostampede

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

type Oplog struct {
	Op string                 `json:op`
	Ns string                 `json:ns`
	O  map[string]interface{} `json:o`
}

func LoadOplog(oplog []byte) Oplog {
	//return error and Oplog both
	var oplogObj Oplog

	err := json.Unmarshal(oplog, &oplogObj)

	if err != nil {
		panic(err)
	}

	return oplogObj
}

func (oplog *Oplog) GetInsertStatement() (string, error) {
	//return error and string both

	if oplog.Op == "i" {

		fieldNames := make([]string, 0, len(oplog.O))

		for fieldName, _ := range oplog.O {
			fieldNames = append(fieldNames, fieldName)
		}

		sort.Strings(fieldNames)
		fieldValues := make([]string, 0, len(fieldNames))
		for _, field := range fieldNames {
			fieldValues = append(fieldValues, getStringFormatValue(oplog.O[field]))
		}

		return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", oplog.Ns, strings.Join(fieldNames, ", "), strings.Join(fieldValues, ", ")), nil

	}
	return "", fmt.Errorf("error in generating the insert statement for oplog %v", oplog)
}

func getStringFormatValue(value interface{}) string {
	switch v := value.(type) {
	case float64:
		return fmt.Sprintf("%v", v)
	case bool:
		return fmt.Sprintf("%t", v)
	default:
		return fmt.Sprintf("'%s'", v)
	}

}

//TODO: Check for test coverage
