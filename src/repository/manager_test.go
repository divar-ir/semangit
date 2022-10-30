package repository

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"os/exec"
	"semangit/src/utils"
	"strings"
	"testing"
)

const MasterBranch = "master"
const AnotherBranch = "test-branch"

type TestSuite struct {
	suite.Suite
	repoDir     string
	repoManager gitRepoManager
}

func TestGitRepoManager(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) SetupTest() {
	s.repoDir = utils.GetResultOrPanicError(os.MkdirTemp("", "sample-git-repo"))
	s.runGitCommand("init")
	s.runGitCommand("config", "user.email", "test@test.com")
	s.runGitCommand("config", "user.name", "Test")
	s.runGitCommand("commit", "--allow-empty", "-m", "Initial commit")
	s.runGitCommand("checkout", "-b", AnotherBranch)
	s.runGitCommand("commit", "--allow-empty", "-m", "Test commit")
	s.runGitCommand("checkout", MasterBranch)
	s.repoManager = NewGitRepoManger(s.repoDir)
}

func (s *TestSuite) runGitCommand(args ...string) string {
	newArgs := append([]string{"-C", s.repoDir}, args...)
	cmd := exec.Command("git", newArgs...)
	output := utils.GetResultOrPanicError(cmd.Output())
	return string(output)
}

func (s *TestSuite) TearDownTest() {
	utils.PanicError(os.RemoveAll(s.repoDir))
}

func (s *TestSuite) assertBranch(expectedBranch string) {
	currentBranch := s.runGitCommand("rev-parse", "--abbrev-ref", "HEAD")
	currentBranch = strings.Trim(currentBranch, "\n")
	assert.Equal(s.T(), expectedBranch, currentBranch)
}

func (s *TestSuite) TestCanCheckoutToAnotherBranch() {
	s.assertBranch(MasterBranch)
	s.repoManager.Checkout(AnotherBranch)
	s.assertBranch(AnotherBranch)
}
