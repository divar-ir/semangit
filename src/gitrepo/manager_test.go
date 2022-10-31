package gitrepo

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"os/exec"
	"path"
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
	return s.runCommand("git", newArgs...)
}

func (s *TestSuite) runCommand(name string, args ...string) string {
	cmd := exec.Command(name, args...)
	output := utils.GetResultOrPanicError(cmd.CombinedOutput())
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

func (s *TestSuite) TestCanListChangedFiles() {
	fromRev := s.runGitCommand("log", "-1", "--pretty=%H")
	fromRev = strings.TrimSpace(fromRev)
	s.NotEmpty(fromRev)

	s.runCommand("touch", path.Join(s.repoDir, "newfile.txt"))
	s.runGitCommand("add", ".")
	s.runGitCommand("commit", "-m", "Add new file")

	toRev := s.runGitCommand("log", "-1", "--pretty=%H")
	toRev = strings.TrimSpace(toRev)
	s.NotEmpty(toRev)
	changedFiles := s.repoManager.ListChangedFiles(fromRev, toRev)
	s.Equal([]string{"newfile.txt"}, changedFiles)
}

func (s *TestSuite) TestChangedFilesIncludeBothOldAndNewFilenamesOnRename() {
	s.runCommand("touch", path.Join(s.repoDir, "file1.txt"))
	s.runGitCommand("add", ".")
	s.runGitCommand("commit", "-m", "Add file1")

	fromRev := s.runGitCommand("log", "-1", "--pretty=%H")
	fromRev = strings.TrimSpace(fromRev)
	s.NotEmpty(fromRev)

	s.runCommand("mv", path.Join(s.repoDir, "file1.txt"), path.Join(s.repoDir, "file2.txt"))
	s.runGitCommand("add", ".")
	s.runGitCommand("commit", "-m", "Rename file1 to file2")

	toRev := s.runGitCommand("log", "-1", "--pretty=%H")
	toRev = strings.TrimSpace(toRev)
	s.NotEmpty(toRev)
	changedFiles := s.repoManager.ListChangedFiles(fromRev, toRev)
	changedFilesStr := strings.Join(changedFiles, " ")
	s.True(strings.Contains(changedFilesStr, "file1.txt"))
	s.True(strings.Contains(changedFilesStr, "file2.txt"))
}
