package webpagetest

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingResultWithPlrAsNumber(t *testing.T) {
	var response, err = ioutil.ReadFile("./testdata/TestResultPlrAsNumber.json")
	assert.Nil(t, err)
	_, err = parseResultResponse(response)
	assert.Nil(t, err)
}

func TestParsingResultWithPlrAsString(t *testing.T) {
	var response, err = ioutil.ReadFile("./testdata/TestResultPlrAsString.json")
	assert.Nil(t, err)
	_, err = parseResultResponse(response)
	assert.Nil(t, err)
}
