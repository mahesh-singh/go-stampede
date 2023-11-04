package gostampede

import (
	"testing"
)

func TestOplog_GetInsertStatement(t *testing.T) {

	tests := []struct {
		name     string
		oplogStr []byte
		want     string
	}{
		{name: "insert operation", oplogStr: []byte(`{
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
			oplog := LoadOplog(tt.oplogStr)
			if got, _ := oplog.GetInsertStatement(); got != tt.want {
				t.Errorf("Oplog.GetInsertStatement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOplog_GetUpdateStatement(t *testing.T) {

	//TODO: Add more test case
	tests := []struct {
		name     string
		oplogStr []byte
		want     string
		wantErr  bool
	}{
		{name: "Update statement with update field", oplogStr: []byte(`{
			"op": "u",
			"ns": "test.student",
			"o": {
			   "$v": 2,
			   "diff": {
				  "u": {
					 "is_graduated": true
				  }
			   }
			},
			 "o2": {
			   "_id": "635b79e231d82a8ab1de863b"
			}
		 }`), want: "UPDATE test.student SET is_graduated = true WHERE _id = '635b79e231d82a8ab1de863b';", wantErr: false},
		{
			name: "Update statement with set null field", oplogStr: []byte(`{
				"op": "u",
				"ns": "test.student",
				"o": {
				   "$v": 2,
				   "diff": {
					  "d": {
						 "roll_no": false
					  }
				   }
				},
				"o2": {
				   "_id": "635b79e231d82a8ab1de863b"
				}
			 }`), want: "UPDATE test.student SET roll_no = NULL WHERE _id = '635b79e231d82a8ab1de863b';", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oplog := LoadOplog(tt.oplogStr)
			got, err := oplog.GetUpdateStatement()
			if (err != nil) != tt.wantErr {
				t.Errorf("Oplog.GetUpdateStatement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Oplog.GetUpdateStatement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOplog_GetDeleteStatement(t *testing.T) {

	tests := []struct {
		name     string
		oplogStr []byte
		want     string
		wantErr  bool
	}{
		{name: "Delete statement", oplogStr: []byte(`{
			"op": "d",
			"ns": "test.student",
			"o": {
			  "_id": "635b79e231d82a8ab1de863b"
			}
		  }`), want: "DELETE FROM test.student WHERE _id = '635b79e231d82a8ab1de863b';", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oplog := LoadOplog(tt.oplogStr)
			got, err := oplog.GetDeleteStatement()
			if (err != nil) != tt.wantErr {
				t.Errorf("Oplog.GetDeleteStatement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Oplog.GetDeleteStatement() = %v, want %v", got, tt.want)
			}
		})
	}
}
