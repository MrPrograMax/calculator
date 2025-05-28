package grpc

import (
	"calculator/internal/pkg/service"
	desc "calculator/pkg/api"
	"context"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestImplementation_Calculate(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	tests := []struct {
		name     string
		req      *desc.CalculateRequest
		wantResp *desc.CalculateResponse
		wantErr  require.ErrorAssertionFunc
	}{
		{
			name: "OK. Case 1",
			req: &desc.CalculateRequest{
				Data: []*desc.CalculateRequest_Data{
					{
						Type:       "calc",
						Op:         lo.ToPtr("+"),
						Var:        "x",
						LeftValue:  &desc.CalculateRequest_Data_LeftConst{LeftConst: 1},
						RightValue: &desc.CalculateRequest_Data_RightConst{RightConst: 2},
					},
					{
						Type: "print",
						Var:  "x",
					},
				},
			},
			wantResp: &desc.CalculateResponse{
				Items: []*desc.CalculateResponse_Item{
					{Var: "x", Value: "3"},
				},
			},
			wantErr: require.NoError,
		},
		{
			name: "OK. Case 2",
			req: &desc.CalculateRequest{
				Data: []*desc.CalculateRequest_Data{
					{
						Type:       "calc",
						Op:         lo.ToPtr("+"),
						Var:        "x",
						LeftValue:  &desc.CalculateRequest_Data_LeftConst{LeftConst: 10},
						RightValue: &desc.CalculateRequest_Data_RightConst{RightConst: 2},
					},
					{
						Type: "print",
						Var:  "x",
					},
					{
						Type:       "calc",
						Op:         lo.ToPtr("-"),
						Var:        "y",
						LeftValue:  &desc.CalculateRequest_Data_LeftVar{LeftVar: "x"},
						RightValue: &desc.CalculateRequest_Data_RightConst{RightConst: 3},
					},
					{
						Type:       "calc",
						Op:         lo.ToPtr("*"),
						Var:        "z",
						LeftValue:  &desc.CalculateRequest_Data_LeftVar{LeftVar: "x"},
						RightValue: &desc.CalculateRequest_Data_RightVar{RightVar: "y"},
					},
					{
						Type: "print",
						Var:  "w",
					},
					{
						Type:       "calc",
						Op:         lo.ToPtr("*"),
						Var:        "w",
						LeftValue:  &desc.CalculateRequest_Data_LeftVar{LeftVar: "z"},
						RightValue: &desc.CalculateRequest_Data_RightConst{RightConst: 0},
					},
				},
			},
			wantResp: &desc.CalculateResponse{
				Items: []*desc.CalculateResponse_Item{
					{Var: "x", Value: "12"},
					{Var: "w", Value: "0"},
				},
			},
			wantErr: require.NoError,
		},
		{
			name: "OK. Case 3",
			req: &desc.CalculateRequest{
				Data: []*desc.CalculateRequest_Data{
					{Type: "calc", Op: lo.ToPtr("+"), Var: "x", LeftValue: &desc.CalculateRequest_Data_LeftConst{LeftConst: 10}, RightValue: &desc.CalculateRequest_Data_RightConst{RightConst: 2}},
					{Type: "calc", Op: lo.ToPtr("*"), Var: "y", LeftValue: &desc.CalculateRequest_Data_LeftVar{LeftVar: "x"}, RightValue: &desc.CalculateRequest_Data_RightConst{RightConst: 5}},
					{Type: "calc", Op: lo.ToPtr("-"), Var: "q", LeftValue: &desc.CalculateRequest_Data_LeftVar{LeftVar: "y"}, RightValue: &desc.CalculateRequest_Data_RightConst{RightConst: 20}},
					{Type: "calc", Op: lo.ToPtr("+"), Var: "unusedA", LeftValue: &desc.CalculateRequest_Data_LeftVar{LeftVar: "y"}, RightValue: &desc.CalculateRequest_Data_RightConst{RightConst: 100}},
					{Type: "calc", Op: lo.ToPtr("*"), Var: "unusedB", LeftValue: &desc.CalculateRequest_Data_LeftVar{LeftVar: "unusedA"}, RightValue: &desc.CalculateRequest_Data_RightConst{RightConst: 2}},
					{Type: "print", Var: "q"},
					{Type: "calc", Op: lo.ToPtr("-"), Var: "z", LeftValue: &desc.CalculateRequest_Data_LeftVar{LeftVar: "x"}, RightValue: &desc.CalculateRequest_Data_RightConst{RightConst: 15}},
					{Type: "print", Var: "z"},
					{Type: "calc", Op: lo.ToPtr("+"), Var: "ignoreC", LeftValue: &desc.CalculateRequest_Data_LeftVar{LeftVar: "z"}, RightValue: &desc.CalculateRequest_Data_RightVar{RightVar: "y"}},
					{Type: "print", Var: "x"},
				},
			},
			wantResp: &desc.CalculateResponse{
				Items: []*desc.CalculateResponse_Item{
					{Var: "q", Value: "40"},
					{Var: "z", Value: "-3"},
					{Var: "x", Value: "12"},
				},
			},
			wantErr: require.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := logrus.New()
			svc := service.NewService(logger)
			calculator := NewGRPSService(logger, svc)
			resp, err := calculator.Calculate(ctx, tt.req)

			tt.wantErr(t, err)
			assert.Equal(t, tt.wantResp, resp)
		})
	}
}
