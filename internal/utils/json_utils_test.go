package utils

import (
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
)

type JSONUtilsTestSuite struct {
	suite.Suite
}

type homeAddress struct {
	Address string
	Num     int
}

type info struct {
	Name string
	Age  int
	Home homeAddress
	ID   int
}

const desiredValue = `
{
	"Name": "foo",
	"Age": 50,
	"Home": {
		"Address": "foo,bar",
		"Num": 10
	},
	"ID": 1
}
`

func TestJSONUtils(t *testing.T) {
	suite.Run(t, new(JSONUtilsTestSuite))
}

func (s *JSONUtilsTestSuite) TestInterfaceToString() {
	myInfo := info{
		Name: "foo",
		Age:  50,
		Home: homeAddress{
			Address: "foo,bar",
			Num:     10,
		},
		ID: 1,
	}

	myInfoString := InterfaceToString(myInfo)
	s.Equal(strings.TrimSpace(myInfoString), strings.TrimSpace(desiredValue))
}
