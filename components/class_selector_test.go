package components

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateClassSelector_WithMultiFalse(t *testing.T) {
	selector := CreateClassSelector(false)
	assert.Len(t, selector.Objects, 2)
}

func TestCreateClassSelector_WithMultiTrue(t *testing.T) {
	selector := CreateClassSelector(true)
	assert.Len(t, selector.Objects, 2)
}

func TestCreateLevelArray(t *testing.T) {
	arr := createLevelArray()

	assert.Len(t, arr, 20)
}
