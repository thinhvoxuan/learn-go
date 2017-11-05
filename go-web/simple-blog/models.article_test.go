package main

import (
	"reflect"
	"testing"
)

func Test_getAllArticles(t *testing.T) {
	tests := []struct {
		name string
		want []article
	}{
		{"testcase 1", articleList},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAllArticles(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllArticles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getArticleByID(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    *article
		wantErr bool
	}{
		{"testcase 1", args{1}, &articleList[0], false},
		{"testcase 2", args{2}, &articleList[1], false},
		{"testcase 3", args{3}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getArticleByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("getArticleByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getArticleByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
