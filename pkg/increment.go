package pkg

import "os/exec"

func Increment (version string) {
	exec.Command("git", "tag", "-a", version, "-m", version).Run() // TODO I need to be "git", "push", "origin", "version"
}