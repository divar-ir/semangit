package gitRepoManager

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"os/exec"
	"semangit/src/utils"
	"testing"
)

const CurrentBranch = "master"
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
	s.runGitCommand("commit", "--allow-empty", "-m", "Initial commit")
	s.runGitCommand("checkout", "-b", AnotherBranch)
	s.runGitCommand("commit", "--allow-empty", "-m", "Test commit")
	s.runGitCommand("checkout", CurrentBranch)
	s.repoManager = NewGitRepoManger(s.repoDir)
}

func (s *TestSuite) runGitCommand(args ...string) {
	newArgs := append([]string{"-C", s.repoDir}, args...)
	cmd := exec.Command("git", newArgs...)
	utils.PanicError(cmd.Run())
}

func (s *TestSuite) TearDownTest() {
	utils.PanicError(os.RemoveAll(s.repoDir))
}

func (s *TestSuite) TestCanCheckoutToAnotherBranch() {
	s.repoManager.Checkout(AnotherBranch)
	assert.Fail(s.T(), "TODO: Write this test!")
}
