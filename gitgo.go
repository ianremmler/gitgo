package gitgo

import (
	"os/exec"
	"strings"
)

type GitGo struct{}

func New() *GitGo {
	return &GitGo{}
}

func (g *GitGo) run(op string, args []string) (string, error) {
	c := exec.Command("git", append([]string{op}, args...)...)
	out, err := c.CombinedOutput()
	return string(out), err
}

func (g *GitGo) Run(op string, args ...string) (string, error) {
	return g.run(op, args)
}

func (g *GitGo) Init(args ...string) (string, error) {
	return g.run("init", args)
}

func (g *GitGo) Add(args ...string) (string, error) {
	return g.run("add", args)
}

func (g *GitGo) Commit(msg string, args ...string) (string, error) {
	return g.run("commit", append([]string{"-m", msg}, args...))
}

func (g *GitGo) Reset(args ...string) (string, error) {
	return g.run("reset", args)
}

func (g *GitGo) Checkout(args ...string) (string, error) {
	return g.run("checkout", args)
}

func (g *GitGo) NewBranch(name string) (string, error) {
	return g.Run("branch", name)
}

func (g *GitGo) CheckoutNewBranch(name string) (string, error) {
	return g.Checkout("-b", name)
}

func (g *GitGo) CurBranch() (string, error) {
	out, err := g.Run("rev-parse", "--abbrev-ref", "HEAD")
	return strings.TrimSpace(out), err
}

func (g *GitGo) Branches(pattern string) ([]string, error) {
	out, err := g.Run("for-each-ref", "--format=%(refname:short)", "refs/heads/"+pattern)
	if err != nil {
		return nil, err
	}
	return strings.Split(out, "\n"), nil
}

func (g *GitGo) FileContents(branch, filename string) (string, error) {
	return g.Run("show", branch+":"+filename)
}

func (g *GitGo) Blame(branch, filename string) (string, error) {
	if branch == "" {
		return g.Run("blame", filename)
	}
	return g.Run("blame", branch, filename)
}
