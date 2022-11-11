package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringExample(t *testing.T) {
	StringExample()

	val, err := client.Get(ctx, "name").Result()
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "tingxuan", val)
}
