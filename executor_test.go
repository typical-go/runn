package runn_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/runn"
)

func TestExecutor_All(t *testing.T) {
	executor := runn.Executor{
		StopWhenError: false,
	}

	t.Run("GIVEN multiple error", func(t *testing.T) {
		err := executor.Execute(
			errors.New("error1"),
			errors.New("error2"),
		)
		require.EqualError(t, err, "error1; error2;")
	})

	t.Run("GIVEN no error/statement", func(t *testing.T) {
		err := executor.Execute()
		require.NoError(t, err)
	})

}

func TestExecutor_StopWhenError(t *testing.T) {
	executor := runn.Executor{
		StopWhenError: true,
	}

	t.Run("GIVEN multiple error", func(t *testing.T) {
		err := executor.Execute(
			errors.New("error1"),
			errors.New("unreachable-error"),
		)
		require.EqualError(t, err, "error1")
	})

	t.Run("GIVEN no error/statement", func(t *testing.T) {
		err := executor.Execute()
		require.NoError(t, err)
	})

}
