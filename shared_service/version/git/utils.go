package git

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func runCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	return strings.TrimSpace(out.String()), err
}

func Push() {
	branch := fmt.Sprintf("feature-%s", time.Now().Format("2006-01-02_15-04-05"))

	changer, err := runCommand("git", "status", "--porcelain")
	{
		if err != nil {
			panic(err)
		}
		if changer == "" {
			panic("uncommitted changes")
		}
	}

	if _, err := runCommand("git", "checkout", "-b", branch); err != nil {
		panic(err)
	}

	if _, err := runCommand("git", "add", "."); err != nil {
		fmt.Println("Ошибка при git add:", err)
		os.Exit(1)
	}

	if _, err := runCommand("git", "commit", "-m", "Автоматический коммит"); err != nil {
		fmt.Println("Ошибка при git commit:", err)
		panic(err)
	}

	if _, err := runCommand("git", "push", "origin", branch); err != nil {
		fmt.Println("Ошибка при git push:", err)
		panic(err)
	}

	Merge(branch)
}

func Merge(branch string) {
	mainBranch, err := runCommand("git", "rev-parse", "--abbrev-ref", "origin/HEAD")
	{
		if err != nil {
			panic(err)
		}
	}

	if _, err := runCommand("git", "checkout", mainBranch); err != nil {
		fmt.Println("Ошибка при git push:", err)
		panic(err)
	}

	if _, err := runCommand("git", "pull", "origin", mainBranch); err != nil {
		fmt.Println("Ошибка при git push:", err)
		panic(err)
	}

	if _, err := runCommand("git", "merge", "--no-ff", branch, "-m", fmt.Sprintf("Merge %s into %s", branch, mainBranch)); err != nil {
		fmt.Println("Ошибка при git push:", err)
		panic(err)
	}

	if _, err := runCommand("git", "push", "origin", mainBranch); err != nil {
		fmt.Println("Ошибка при git push:", err)
		panic(err)
	}

	if _, err := runCommand("git", "branch", "-d", branch); err != nil {
		fmt.Println("Ошибка при git push:", err)
		panic(err)
	}

	// git push origin --delete $(BRANCH_NAME)
	if _, err := runCommand("git", "push", "origin", "--delete", branch); err != nil {
		fmt.Println("Ошибка при git push:", err)
		panic(err)
	}

	if _, err := runCommand("git", "rev-parse", "--abbrev-ref", "origin/HEAD"); err != nil {
		fmt.Println("Ошибка при git push:", err)
		panic(err)
	}

}
