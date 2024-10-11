package git

import (
	"github.com/SZabrodskii/git-committer/config"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"os"
	"os/exec"
	"strings"
	"testing"
)

type RepositoryTestSuite struct {
	suite.Suite
	repo *Repository
}

func (suite *RepositoryTestSuite) SetupSuite() {
	logger, err := zap.NewProduction()
	suite.Require().NoError(err)

	suite.repo = NewRepository(&config.Config{RepoURL: "", RepoName: "."}, logger) // Используем текущую директорию
}

func (suite *RepositoryTestSuite) TestInitExistingRepo() {
	err := suite.repo.Init()
	suite.Require().NoError(err)
	suite.repo.Logger.Info("Repository directory already exists locally")
}

func (suite *RepositoryTestSuite) TestCreateCommit() {
	err := suite.repo.Init()
	suite.Require().NoError(err, "Failed to init repository")

	fileName := "test_file.txt"

	f, err := os.Create(fileName)
	suite.Require().NoError(err)
	defer f.Close()

	defer func() {
		err = os.Remove(fileName)
		suite.Require().NoError(err, "Failed to delete file: %s", fileName)
	}()

	_, err = os.Stat(fileName)
	suite.Require().NoError(err, "File should exist at: %s", fileName)

	commitMessage := "Test commit message"
	err = suite.repo.CreateCommit(fileName, commitMessage)
	suite.Require().NoError(err)

	suite.repo.Logger.Info("Commit created", zap.String("Message", commitMessage))

	cmd := exec.Command("git", "-C", suite.repo.Name, "log", "--pretty=format:%s", "-n", "1")
	output, err := cmd.Output()
	suite.Require().NoError(err)

	history := strings.Split(string(output), "\n")
	suite.Equal(1, len(history))
	suite.Contains(history[0], commitMessage)
}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}
