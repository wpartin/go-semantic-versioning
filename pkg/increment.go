package pkg

import "os/exec"

func Increment (version string) {
	exec.Command("git", "tag", "-a", version, "-m", version).Run()
	exec.Command("git", "push", "origin", version).Run()
}