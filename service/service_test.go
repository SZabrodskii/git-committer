package service

import (
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type AnekdotServiceTestSuite struct {
	suite.Suite
	service *AnekdotService
}

func (suite *AnekdotServiceTestSuite) SetupSuite() {
	suite.service = NewAnekdotService()
}

func (suite *AnekdotServiceTestSuite) TestGetRandomAnekdot() {
	anekdot, err := suite.service.GetRandomAnekdot()
	suite.Require().NoError(err)

	suite.NotEmpty(anekdot, "Аnekdot should not be empty")

	tempFile, err := os.CreateTemp("", "anekdot_test")
	suite.Require().NoError(err)
	defer os.Remove(tempFile.Name())

	err = suite.service.SaveAnekdotToFile(anekdot, tempFile.Name())
	suite.Require().NoError(err)

	data, err := os.ReadFile(tempFile.Name())
	suite.Require().NoError(err)

	suite.NotEmpty(string(data), "Saved anekdot should not be empty")
}

func (suite *AnekdotServiceTestSuite) TestSaveAnekdotToFile() {
	tempFile, err := os.CreateTemp("", "anekdot_test")
	suite.Require().NoError(err)
	defer os.Remove(tempFile.Name())

	err = suite.service.SaveAnekdotToFile("This is a test anekdot for saving", tempFile.Name())
	suite.Require().NoError(err)

	data, err := os.ReadFile(tempFile.Name())
	suite.Require().NoError(err)

	suite.Equal("This is a test anekdot for saving", string(data))
}

func TestAnekdotServiceTestSuite(t *testing.T) {
	suite.Run(t, new(AnekdotServiceTestSuite))
}
