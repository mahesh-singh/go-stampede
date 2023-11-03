package gostampede

import (
	"testing"
)

func TestOplog_GetInsertStatement(t *testing.T) {

	tests := []struct {
		name  string
		oplog []byte
		want  string
	}{
		{name: "insert operation", oplog: []byte(`{
			"op": "i",
			"ns": "test.student",
			"o": {
			  "_id": "635b79e231d82a8ab1de863b",
			  "name": "Selena Miller",
			  "roll_no": 51,
			  "is_graduated": false,
    		  "date_of_birth": "2000-01-30"
			}
		  }`), want: "INSERT INTO test.student (_id, date_of_birth, is_graduated, name, roll_no) VALUES ('635b79e231d82a8ab1de863b', '2000-01-30', false, 'Selena Miller', 51);"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oplog := LoadOplog(tt.oplog)
			if got, _ := oplog.GetInsertStatement(); got != tt.want {
				t.Errorf("Oplog.GetInsertStatement() = %v, want %v", got, tt.want)
			}
		})
	}
}
