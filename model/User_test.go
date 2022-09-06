package model

import (
	"testing"
)

func TestScryptPassword(t *testing.T) {
	hashpass := ScryptPassword("123456")

	t.Log(hashpass)
	if hashpass == "" {
		t.Error()
	}
}
