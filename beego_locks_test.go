package beego_locks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGlobalBeegoLockFactory_InitiallyNil(t *testing.T) {
	assert.Nil(t, GlobalBeegoLockFactory)
}
