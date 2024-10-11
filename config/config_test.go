package config_test

import (
	"bytes"
	"github.com/SZabrodskii/git-committer/config"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ConfigTestSuite struct {
	suite.Suite
	logger *zap.Logger
}

func (suite *ConfigTestSuite) SetupTest() {
	encoderCfg := zap.NewDevelopmentEncoderConfig()
	encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), zapcore.AddSync(bytes.NewBuffer([]byte{})), zapcore.DebugLevel)
	suite.logger = zap.New(core)
}

func (suite *ConfigTestSuite) TestLoadConfig_Success() {
	viper.SetConfigFile("./config.json")
	cfg, err := config.NewConfig()
	suite.Require().NoError(err)
	suite.Require().NotNil(cfg)
	suite.Equal(3, cfg.MinCommits)
	suite.Equal(7, cfg.MaxCommits)
	suite.Equal("https://github.com/SZabrodskii/git-committer.git", cfg.RepoURL)
}

func (suite *ConfigTestSuite) TestLoadConfig_FromEmbedded() {
	viper.SetConfigFile("./testdata/nonexistent.json")

	cfg, err := config.NewConfig()
	suite.Require().NoError(err)
	suite.Require().NotNil(cfg)
	suite.Equal(3, cfg.MinCommits)
	suite.Equal(7, cfg.MaxCommits)
	suite.Equal("https://github.com/SZabrodskii/git-committer.git", cfg.RepoURL)
}

func (suite *ConfigTestSuite) TestLoadConfig_InvalidFormat() {
	viper.SetConfigFile("./invalid_config.json")

	err := viper.ReadInConfig()
	suite.Error(err, "Expected an error due to invalid JSON format")
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
