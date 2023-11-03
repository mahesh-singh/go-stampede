package gostampede

import (
	"strings"
	"testing"
)

func TestGetInsertStatement(t *testing.T) {
	jsonStr := []byte(`{
		"op": "i",
		"ns": "test.student",
		"o": {
		  "_id": "635b79e231d82a8ab1de863b",
		  "name": "Selena Miller",
		  "roll_no": 51
		}
	  }`)
	/*
		  oplog := Oplog{
			op:  "i",
			ns:  "test.student",
			obj: []byte(`{"_id": "635b79e231d82a8ab1de863b", "name": "Selena Miller", "roll_no": 51}`),
		}*/

	oplog := LoadOplog(jsonStr)

	ins := oplog.GetInsertStatement()
	expected := "INSERT INTO test.student (_id,name,roll_no) VALUES ('635b79e231d82a8ab1de863b','Selena Miller',51)"
	if !strings.EqualFold(ins, expected) {
		t.Errorf("expected: %s, output: %s", expected, ins)
	}
}

/*
INSERT INTO test.student (_id,name,roll_no) VALUES ('635b79e231d82a8ab1de863b','Selena Miller',51),
INSERT INTO test.student (roll_no,_id,name) VALUES (51,'635b79e231d82a8ab1de863b','Selena Miller')
*/
/*
func TestOplog_GetInsertStatement(t *testing.T) {
	type fields struct {
		Op string
		Ns string
		O  json.RawMessage
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oplog := &Oplog{
				Op: tt.fields.Op,
				Ns: tt.fields.Ns,
				O:  tt.fields.O,
			}
			if got := oplog.GetInsertStatement(); got != tt.want {
				t.Errorf("Oplog.GetInsertStatement() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/
