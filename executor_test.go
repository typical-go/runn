package runn_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/runn"
)

type RunnerImplementationWithError struct{}

func (i RunnerImplementationWithError) Run() error { return errors.New("some-error") }

type RunnerImplementationNoError struct{}

func (i RunnerImplementationNoError) Run() error { return nil }

func TestExecutor_All(t *testing.T) {
	testcases := []struct {
		stopWhenError bool
		stmts         []interface{}
		err           error
	}{
		{
			false,
			[]interface{}{
				RunnerImplementationWithError{},
				errors.New("error1"),
				errors.New("error2"),
			},
			errors.New("some-error; error1; error2"),
		},
		{
			true,
			[]interface{}{
				errors.New("error1"),
				errors.New("unreachable-error"),
			},
			errors.New("error1"),
		},
		{
			false,
			[]interface{}{},
			nil,
		},
		{
			true,
			[]interface{}{},
			nil,
		},
		{
			false,
			[]interface{}{
				RunnerImplementationNoError{},
				RunnerImplementationWithError{},
			},
			errors.New("some-error"),
		},
	}

	for _, tt := range testcases {
		err := runn.Executor{
			StopWhenError: tt.stopWhenError,
		}.Execute(tt.stmts...)

		if tt.err == nil {
			require.NoError(t, err)
		} else {
			require.EqualError(t, err, tt.err.Error())
		}
	}
}
