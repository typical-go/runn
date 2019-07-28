package runn_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/runn"
)

func TestExecute(t *testing.T) {
	err := runn.Execute(
		errors.New("some-error"),
		errors.New("unreachable-error"),
	)

	require.EqualError(t, err, "some-error")

}
