package entity

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarshal(t *testing.T) {
	expected := Avatar{
		ID:       1,
		Nickname: "kevin",
		Profile:  nil,
	}

	marshal, err := json.Marshal(expected)
	if err != nil {
		t.Fatal(err)
	}

	var got Avatar
	err = json.Unmarshal(marshal, &got)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, got)
}
