package utils_test

import (
	"testing"

	"github.com/avran02/medods/internal/utils"
	"github.com/stretchr/testify/assert"
)

var testStrings = []string{
	"test",
	"test1",
	"test2",
	"123",
	"ajsdfiuhasduivcbasjidbvsa",
	"1234512456245692384",
	"jsadbvkfh1384uryfasuidf",
	"eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdG9rZW5faWQiOiI4MTVjOTdhMi02MmYxLTQ5MzctOTZmNy1hMTBjMDliYjNhZTkiLCJ1c2VyX2lwIjoiMTcyLjIxLjAuMSIsInN1YiI6IiAzMzYwYWVlYy0xMjgzLTRiNjEtYmZhNi0yMWFhM2M2MTgwNGQgIiwiZXhwIjoxNzIzNjQyMjg0LCJqdGkiOiJkMDc1ZDU5Ny0yYzI2LTRmNTAtYjU5Mi1mMTEzMTdjYWUxYzUifQ.9drECLSZKS9-H_il72ctYTvCHOYkjGmqqk4QaJYtbI2vYPfs1ZZPC5sdn9suuAJLe8FwS4Y9b1DhEwKrvoASJQ",
}

func TestHash(t *testing.T) {
	for _, str := range testStrings {
		t.Run(str, func(t *testing.T) {
			hash, err := utils.Hash(str)
			assert.NoError(t, err)
			assert.NotEmpty(t, hash)

			assert.NoError(t, utils.CompareHashAndPassword(str, hash))
		})
	}
}
