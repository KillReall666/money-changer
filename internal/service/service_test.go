package service

import (
	"gitlab.com/llcmediatel/recruiting/golang-junior-dev/internal/model"
	"reflect"
	"testing"
)

func TestService_Calculate(t *testing.T) {
	type args struct {
		data model.JSONRequest
	}
	tests := []struct {
		name    string
		args    args
		want    [][]float64
		wantErr string
	}{
		{
			name: "normal resp",
			args: args{
				data: model.JSONRequest{
					Amount: 400,
					Banknotes: []float64{
						5000,
						2000,
						1000,
						500,
						200,
						100,
						50,
					},
				},
			},
			want: [][]float64{{200, 200}, {100, 100, 200}, {50, 100, 100, 100}, {50, 50, 100, 200}, {50, 50, 100, 100, 100}, {50, 50, 50, 50, 200}, {50, 50, 50, 50, 100, 100}, {50, 50, 50, 50, 50, 50, 100}, {50, 50, 50, 50, 50, 50, 50, 50}},
		},
		{
			name: "not multiple of amount",
			args: args{
				data: model.JSONRequest{
					Amount: 399,
					Banknotes: []float64{
						5000,
						2000,
						1000,
						500,
						200,
						100,
						50,
					},
				},
			},
			wantErr: "could not calculate exchanges: amount must be a multiple of 50",
		},
		{
			name: "empty banknotes",
			args: args{
				data: model.JSONRequest{
					Amount:    400,
					Banknotes: []float64{},
				},
			},
			wantErr: "could not calculate exchanges: banknotes cannot be empty",
		},
		{
			name: "zero amount",
			args: args{
				data: model.JSONRequest{
					Amount:    0,
					Banknotes: []float64{100, 50, 20},
				},
			},
			wantErr: "could not calculate exchanges: amount must be greater than zero",
		},
		{
			name: "negative amount",
			args: args{
				data: model.JSONRequest{
					Amount:    -100,
					Banknotes: []float64{100, 50, 20},
				},
			},
			wantErr: "could not calculate exchanges: amount must be greater than zero",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{}
			got, err := s.Calculate(tt.args.data)
			if (err != nil && tt.wantErr == "") || (err == nil && tt.wantErr != "") {
				t.Errorf("Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
