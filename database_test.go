package main

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestGetOwnerId(t *testing.T) {
	type args struct {
		ownerName string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Success test",
			args: args{
				ownerName: "Bulbasaur",
			},
			want: 2,
		},
		{
			name: "Success test - No Matching Case",
			args: args{
				ownerName: "VolToRb",
			},
			want: 94,
		},
		{
			name: "Fail test - No Matching Case",
			args: args{
				ownerName: "Som3_0wner",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOwnerId(tt.args.ownerName); got != tt.want {
				t.Errorf("GetOwnerId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAdminsFromDatabase(t *testing.T) {
	type args struct {
		parameters []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "General test",
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetAdminsFromDatabase(tt.args.parameters...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAdminsFromDatabase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestQueryDatabase(t *testing.T) {
	type args struct {
		query             string
		administratorName string
		ownerId           int
		criticalityId     int
		updateParams      []int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Query Select",
			args: args{
				query:             "SELECT * FROM admins WHERE adminame = ? AND ownerId = ? AND criticalityId = ?",
				administratorName: "Ragnar",
				ownerId:           68,
				criticalityId:     3,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := QueryDatabase(tt.args.query, tt.args.administratorName, tt.args.ownerId, tt.args.criticalityId, tt.args.updateParams...); (err != nil) != tt.wantErr {
				t.Errorf("QueryDatabase() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
