package gostampede

import (
	"encoding/json"
	"fmt"
)

type Oplog struct {
	Op string //Add json field
	Ns string
	O  json.RawMessage //this should allow to store arbitrary json object
	//Can take map
	/*
		Type fields must start with uppercase char so that during unmarshal, these filed can be access
	*/
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

func (oplog *Oplog) GetInsertStatement() string {
	//return error and string both

	if oplog.Op == "i" {
		var f interface{}
		err := json.Unmarshal(oplog.O, &f)
		if err != nil {
			fmt.Println(err)

		}
		m := f.(map[string]interface{})
		fieldName := ""
		fieldvalue := ""

		for k, v := range m {
			if fieldName == "" {
				fieldName = k
			} else {
				fieldName = fieldName + "," + k
			}

			switch vv := v.(type) {
			case string:
				fieldvalue = fieldvalue + fmt.Sprintf("'%s',", vv)
			case float64:
				fieldvalue = fieldvalue + fmt.Sprintf("%v,", vv)
			}

		}
		if fieldvalue != "" {
			fieldvalue = fieldvalue[0 : len(fieldvalue)-1]
		}

		return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", oplog.Ns, fieldName, fieldvalue)

	}
	return ""
}
