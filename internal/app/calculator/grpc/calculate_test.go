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
						Type:  "calc",
						Op:    lo.ToPtr("+"),
						Var:   "x",
						Left:  lo.ToPtr("1"),
						Right: lo.ToPtr("2"),
					},
					{
						Type: "print",
						Var:  "x",
					},
				},
			},
			wantResp: &desc.CalculateResponse{
				Items: []*desc.CalculateResponse_Item{
					{
						Var:   "x",
						Value: "3",
					},
				},
			},
			wantErr: require.NoError,
		},
		{
			name: "OK. Case 2",
			req: &desc.CalculateRequest{
				Data: []*desc.CalculateRequest_Data{
					{
						Type:  "calc",
						Op:    lo.ToPtr("+"),
						Var:   "x",
						Left:  lo.ToPtr("10"),
						Right: lo.ToPtr("2"),
					},
					{
						Type: "print",
						Var:  "x",
					},
					{
						Type:  "calc",
						Op:    lo.ToPtr("-"),
						Var:   "y",
						Left:  lo.ToPtr("x"),
						Right: lo.ToPtr("3"),
					},
					{
						Type:  "calc",
						Op:    lo.ToPtr("*"),
						Var:   "z",
						Left:  lo.ToPtr("x"),
						Right: lo.ToPtr("y"),
					},
					{
						Type: "print",
						Var:  "w",
					},
					{
						Type:  "calc",
						Op:    lo.ToPtr("*"),
						Var:   "w",
						Left:  lo.ToPtr("z"),
						Right: lo.ToPtr("0"),
					},
				},
			},
			wantResp: &desc.CalculateResponse{
				Items: []*desc.CalculateResponse_Item{
					{
						Var:   "x",
						Value: "12",
					},
					{
						Var:   "w",
						Value: "0",
					},
				},
			},
			wantErr: require.NoError,
		},
		{
			name: "OK. Case 3",
			req: &desc.CalculateRequest{
				Data: []*desc.CalculateRequest_Data{
					{
						Type:  "calc",
						Op:    lo.ToPtr("+"),
						Var:   "x",
						Left:  lo.ToPtr("10"),
						Right: lo.ToPtr("2"),
					},
					{
						Type:  "calc",
						Op:    lo.ToPtr("*"),
						Var:   "y",
						Left:  lo.ToPtr("x"),
						Right: lo.ToPtr("5"),
					},
					{
						Type:  "calc",
						Op:    lo.ToPtr("-"),
						Var:   "q",
						Left:  lo.ToPtr("y"),
						Right: lo.ToPtr("20"),
					},
					{
						Type:  "calc",
						Op:    lo.ToPtr("+"),
						Var:   "unusedA",
						Left:  lo.ToPtr("y"),
						Right: lo.ToPtr("100"),
					},
					{
						Type:  "calc",
						Op:    lo.ToPtr("*"),
						Var:   "unusedB",
						Left:  lo.ToPtr("unusedA"),
						Right: lo.ToPtr("2"),
					},
					{
						Type: "print",
						Var:  "q",
					},
					{
						Type:  "calc",
						Op:    lo.ToPtr("-"),
						Var:   "z",
						Left:  lo.ToPtr("x"),
						Right: lo.ToPtr("15"),
					},
					{
						Type: "print",
						Var:  "z",
					},
					{
						Type:  "calc",
						Op:    lo.ToPtr("+"),
						Var:   "ignoreC",
						Left:  lo.ToPtr("z"),
						Right: lo.ToPtr("y"),
					},
					{
						Type: "print",
						Var:  "x",
					},
				},
			},
			wantResp: &desc.CalculateResponse{
				Items: []*desc.CalculateResponse_Item{
					{
						Var:   "q",
						Value: "40",
					},
					{
						Var:   "z",
						Value: "-3",
					},
					{
						Var:   "x",
						Value: "12",
					},
				},
			},
			wantErr: require.NoError,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			logger := logrus.New()
			s := service.NewService(logger)
			calculator := NewGRPSService(logger, s)
			resp, err := calculator.Calculate(ctx, tt.req)

			tt.wantErr(t, err)
			assert.Equal(t, tt.wantResp, resp)
		})
	}
}
