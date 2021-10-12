package main

import (
	"reflect"
	"testing"
)

func Test_parsePath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantOut [][]string
		wantErr bool
	}{
		{"test error", args{"empty.env"}, [][]string(nil), true},
		{"test success", args{"./testdata/example.env"}, [][]string{{"key", "value"}, {"key1", "value"}, {"key2", "value2"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := parsePath(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("parsePath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("parsePath() = %#v, want %#v", gotOut, tt.wantOut)
			}
		})
	}
}

func Test_checkDiff(t *testing.T) {
	type args struct {
		data [][][]string
	}
	tests := []struct {
		name    string
		args    args
		wantOut [][]string
	}{
		{"testData", args{[][][]string{
			{{"key", "value"}, {"key1", "value1"}, {"key2", "value2"}},
			{{"key", "value"}, {"key1", "value"}},
		}}, [][]string{{"key2"}, {"key1=value"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := checkDiff(tt.args.data...); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("checkDiff() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
