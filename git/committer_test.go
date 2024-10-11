package git

import (
	"github.com/SZabrodskii/git-committer/config"
	"github.com/SZabrodskii/git-committer/service"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"testing"
)

type CommitterTestSuite struct {
	suite.Suite
	committer  *GitCommitter
	repo       *Repository
	anekdotSvc *service.AnekdotService
	config     *config.Config
}

func (suite *CommitterTestSuite) SetupSuite() {
	suite.repo = &Repository{
		Logger: zap.NewExample(),
	}
	suite.anekdotSvc = &service.AnekdotService{}

	suite.config = &config.Config{
		MinCommits:        1,
		MaxCommits:        3,
		Days:              3,
		IncludeWeekends:   true,
		WeekendMinCommits: 1,
		WeekendMaxCommits: 2,
		RepoURL:           "https://github.com/SZabrodskii/git-committer.git",
		CommitTemplate:    "Test commit",
	}

	suite.committer = NewGitCommitter(suite.config, suite.repo, suite.anekdotSvc)

	if err := os.MkdirAll("example", 0755); err != nil {
		suite.T().Fatalf("Failed to create example directory: %v", err)
	}
}

func (suite *CommitterTestSuite) TearDownSuite() {
	if err := os.RemoveAll("example"); err != nil {
		suite.T().Fatalf("Failed to remove example directory: %v", err)
	}
}

func (suite *CommitterTestSuite) TestGenerateCommits() {
	suite.config.Days = 30 // Тестируем 30 дней
	suite.config.MinCommits = 2
	suite.config.MaxCommits = 3

	// Выполняем генерацию коммитов
	err := suite.committer.Commit()
	suite.NoError(err, "expected no error during commit generation")

	// Проверяем созданные файлы
	files, err := os.ReadDir("example")
	suite.Require().NoError(err, "expected no error reading example directory")

	// Ожидаемое количество файлов в зависимости от настроек
	expectedCommits := (suite.config.MaxCommits * suite.config.Days) // максимальное количество коммитов
	suite.Require().LessOrEqual(len(files), expectedCommits, "expected number of files to be less than or equal to expected commits")
	suite.Require().GreaterOrEqual(len(files), suite.config.Days*suite.config.MinCommits, "expected at least min commits for each day")

	// Проверяем формат имен файлов
	for _, file := range files {
		suite.Require().Regexp(`anekdot_\d{4}-\d{2}-\d{2}_\d`, file.Name(), "expected file name to match pattern")
	}

	// После проверки удаляем созданные файлы
	for _, file := range files {
		err := os.Remove(filepath.Join("example", file.Name()))
		suite.NoError(err, "expected no error deleting test file")
	}
}

func TestGitCommitterTestSuite(t *testing.T) {
	suite.Run(t, new(CommitterTestSuite))
}

func TestCreateFile(t *testing.T) {
	rootDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current working directory: %v", err)
	}

	exampleDir := filepath.Join(rootDir, "..", "example")
	filePath := filepath.Join(exampleDir, "testfile.txt")

	if _, err := os.Stat(exampleDir); os.IsNotExist(err) {
		err = os.Mkdir(exampleDir, 0755)
		if err != nil {
			t.Fatalf("failed to create example directory: %v", err)
		}
	}

	file, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("failed to create file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString("This is a test.")
	if err != nil {
		t.Fatalf("failed to write to file: %v", err)
	}

	t.Log("File created successfully")

	defer func() {
		err := os.Remove(filePath)
		if err != nil {
			t.Logf("failed to remove file: %v", err)
		} else {
			t.Log("File removed successfully")
		}
	}()
}
