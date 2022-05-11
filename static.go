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

func Exec(file string, args ...string) (string, error) {
	contents, err := Get(file)
	if err != nil {
		return contents, err
	}

	env := []string{
		fmt.Sprintf("SCRIPT_CONTENT=%s", contents),
		fmt.Sprintf("SCRIPT_ARGS=%s", strings.Join(args, " ")),
	}

	cmd := exec.Command("bash")
	cmd.Stdin = strings.NewReader(executorScript)
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), env...)

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}
