package youdaoclient

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddSKUs(t *testing.T) {
	youdao := Translation("tendency")
	assert.NotNil(t, youdao)
	assert.EqualValues(t, youdao.ErrorCode, 0)
}
