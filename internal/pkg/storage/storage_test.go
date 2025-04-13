package storage

import (
	"calculator/internal/pkg/model"
	"context"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func Test_AddResult(t *testing.T) {
	tests := []struct {
		name      string
		ctx       context.Context
		variable  string
		expectErr bool
	}{
		{
			name:      "valid context",
			ctx:       context.Background(),
			variable:  "x",
			expectErr: false,
		},
		{
			name: "canceled context",
			ctx: func() context.Context {
				ctx, cancel := context.WithCancel(context.Background())
				cancel()
				return ctx
			}(),
			variable:  "y",
			expectErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := NewStorage()
			err := s.AddResult(tc.ctx, tc.variable)

			if tc.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_GetResults(t *testing.T) {
	tests := []struct {
		name         string
		setup        func(s *Storage)
		ctx          context.Context
		expectErr    bool
		expectedLen  int
		expectedVars []string
	}{
		{
			name: "successfully returns results",
			ctx:  context.Background(),
			setup: func(s *Storage) {
				_ = s.AddResult(context.Background(), "x")
				_ = s.Action(context.Background(), model.Addition, "x", "2", "3")
			},
			expectErr:    false,
			expectedLen:  1,
			expectedVars: []string{"x"},
		},
		{
			name: "unknown variable in result",
			ctx:  context.Background(),
			setup: func(s *Storage) {
				_ = s.AddResult(context.Background(), "y")
			},
			expectErr:   true,
			expectedLen: 0,
		},
		{
			name: "context cancelled",
			ctx: func() context.Context {
				ctx, cancel := context.WithCancel(context.Background())
				cancel()
				return ctx
			}(),
			setup:       func(s *Storage) {},
			expectErr:   true,
			expectedLen: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := NewStorage()
			tc.setup(s)

			results, err := s.GetResults(tc.ctx)

			if tc.expectErr {
				assert.Error(t, err)
				assert.Nil(t, results)
			} else {
				assert.NoError(t, err)
				assert.Len(t, results, tc.expectedLen)
				for i, v := range tc.expectedVars {
					assert.Equal(t, v, results[i].Var)
				}
			}
		})
	}
}

func TestStorage_Action(t *testing.T) {
	tests := []struct {
		name      string
		ctx       context.Context
		setupVars map[string]int64
		op        model.Operation
		left      string
		right     string
		variable  string
		expected  int64
		expectErr bool
	}{
		{
			name:      "addition of constants",
			ctx:       context.Background(),
			op:        model.Addition,
			left:      "2",
			right:     "3",
			variable:  "res",
			expected:  5,
			expectErr: false,
		},
		{
			name:      "subtraction using variables",
			ctx:       context.Background(),
			setupVars: map[string]int64{"a": 10, "b": 3},
			op:        model.Subtraction,
			left:      "a",
			right:     "b",
			variable:  "res",
			expected:  7,
			expectErr: false,
		},
		{
			name:      "multiplication mixed",
			ctx:       context.Background(),
			setupVars: map[string]int64{"x": 4},
			op:        model.Multiplication,
			left:      "x",
			right:     "5",
			variable:  "z",
			expected:  20,
			expectErr: false,
		},
		{
			name:      "division by zero",
			ctx:       context.Background(),
			op:        model.Division,
			left:      "10",
			right:     "0",
			variable:  "crash",
			expectErr: true,
		},
		{
			name:      "unknown operation",
			ctx:       context.Background(),
			op:        model.UnknownOperation,
			left:      "1",
			right:     "2",
			variable:  "fail",
			expectErr: true,
		},
		{
			name:      "undefined variable",
			ctx:       context.Background(),
			op:        model.Addition,
			left:      "not_set",
			right:     "1",
			variable:  "bad",
			expectErr: true,
		},
		{
			name: "context canceled",
			ctx: func() context.Context {
				ctx, cancel := context.WithCancel(context.Background())
				cancel()
				return ctx
			}(),
			op:        model.Addition,
			left:      "1",
			right:     "2",
			variable:  "zzz",
			expectErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := NewStorage()

			// manually preload variables if needed
			for k, v := range tc.setupVars {
				s.Action(context.Background(), model.Addition, k, strconv.FormatInt(v, 10), "0")
			}

			err := s.Action(tc.ctx, tc.op, tc.variable, tc.left, tc.right)

			if tc.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, s.variables[tc.variable])
			}
		})
	}
}
