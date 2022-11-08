package gitrepo

import (
	"github.com/stretchr/testify/suite"
	"os"
	"os/exec"
	"path"
	"semangit/internal/utils"
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
	s.repoDir = utils.GetResultOrPanic(os.MkdirTemp("", "sample-git-repo"))
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
	output := utils.GetResultOrPanic(cmd.CombinedOutput())
	return string(output)
}

func (s *TestSuite) TearDownTest() {
	utils.PanicError(os.RemoveAll(s.repoDir))
}

func (s *TestSuite) assertBranch(expectedBranch string) {
	currentBranch := s.runGitCommand("rev-parse", "--abbrev-ref", "HEAD")
	currentBranch = strings.Trim(currentBranch, "\n")
	s.Equal(expectedBranch, currentBranch)
}

func (s *TestSuite) TestCanCheckoutToAnotherBranch() {
	s.assertBranch(MasterBranch)
	s.repoManager.Checkout(AnotherBranch)
	s.assertBranch(AnotherBranch)
}

func (s *TestSuite) TestCanListChangedFiles() {
	srcRev := s.runGitCommand("log", "-1", "--pretty=%H")
	srcRev = strings.TrimSpace(srcRev)
	s.NotEmpty(srcRev)

	s.runCommand("touch", path.Join(s.repoDir, "newfile.txt"))
	s.runGitCommand("add", ".")
	s.runGitCommand("commit", "-m", "Add new file")

	destRev := s.runGitCommand("log", "-1", "--pretty=%H")
	destRev = strings.TrimSpace(destRev)
	s.NotEmpty(destRev)
	changedFiles := s.repoManager.ListChangedFiles(srcRev, destRev)
	s.Equal([]string{"newfile.txt"}, changedFiles)
}

func (s *TestSuite) TestChangedFilesIncludeBothOldAndNewFilenamesOnRename() {
	s.runCommand("touch", path.Join(s.repoDir, "file1.txt"))
	s.runGitCommand("add", ".")
	s.runGitCommand("commit", "-m", "Add file1")

	srcRev := s.runGitCommand("log", "-1", "--pretty=%H")
	srcRev = strings.TrimSpace(srcRev)
	s.NotEmpty(srcRev)

	s.runCommand("mv", path.Join(s.repoDir, "file1.txt"), path.Join(s.repoDir, "file2.txt"))
	s.runGitCommand("add", ".")
	s.runGitCommand("commit", "-m", "Rename file1 to file2")

	destRev := s.runGitCommand("log", "-1", "--pretty=%H")
	destRev = strings.TrimSpace(destRev)
	s.NotEmpty(destRev)
	changedFiles := s.repoManager.ListChangedFiles(srcRev, destRev)
	changedFilesStr := strings.Join(changedFiles, " ")
	s.True(strings.Contains(changedFilesStr, "file1.txt"))
	s.True(strings.Contains(changedFilesStr, "file2.txt"))
}
