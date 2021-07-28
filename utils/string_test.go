package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_To_Under_Score(t *testing.T) {
	assert.Equal(t, "saya_belajar_golang", ToUnderScore("SayaBelajarGolang"))
}
