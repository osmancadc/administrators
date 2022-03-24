package main

import (
	"testing"
)

func TestGetCriticalityId(t *testing.T) {
	type args struct {
		criticality string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Error Test",
			args: args{
				criticality: "some",
			},
			want: -1,
		},
		{
			name: "Success Test (Low)",
			args: args{
				criticality: "low",
			},
			want: 1,
		},
		{
			name: "Success Test (Medium)",
			args: args{
				criticality: "medium",
			},
			want: 2,
		},
		{
			name: "Success Test (High)",
			args: args{
				criticality: "high",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCriticalityId(tt.args.criticality); got != tt.want {
				t.Errorf("GetCriticalityId() = %v, want %v", got, tt.want)
			}
		})
	}
}
