package static

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//go:embed executor.sh
var executorScript string

//go:embed bin
//go:embed scripts
var memFS embed.FS

//go:embed etc/common.sh
var commonScript string

// Get returns the contents of the specified embedded file
func Get(file string) (string, error) {
	bytes, err := memFS.ReadFile(file)
	if err != nil {
		return "", err
	}

	content := string(bytes)
	if strings.HasPrefix(file, "scripts/") {
		content = strings.Replace(content, "etc/common.sh", commonScript, 1)
	}

	return content, nil
}

// Run executes the specified embedded file and waits for it to complete
func Run(file string, args ...string) error {
	cmd, err := Command(file, args...)
	if err != nil {
		return err
	}

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

// Output executes the specified embedded file and returns its standard output
func Output(file string, args ...string) (string, error) {
	cmd, err := Command(file, args...)
	if err != nil {
		return "", err
	}

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

// Command returns the Cmd struct to execute the embedded file with the given arguments.
func Command(file string, args ...string) (*exec.Cmd, error) {
	contents, err := Get(file)
	if err != nil {
		return nil, err
	}

	env := []string{
		fmt.Sprintf("SCRIPT_CONTENT=%s", contents),
		fmt.Sprintf("SCRIPT_ARGS=%s", strings.Join(args, " ")),
	}

	cmd := exec.Command("bash")
	cmd.Stdin = strings.NewReader(executorScript)
	cmd.Env = append(os.Environ(), env...)
	cmd.Stderr = os.Stderr

	return cmd, nil
}
