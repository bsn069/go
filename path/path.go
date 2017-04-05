package path

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/bsn069/go/utils"
)

func Look(strFileName string) (string, error) {
	return exec.LookPath(strFileName)
}

func Pwd() string {
	strDir, _ := os.Getwd()
	return strDir
}

func Format(strPath string) string {
	return filepath.FromSlash(strPath)
}

func Join(strName ...string) string {
	return filepath.Join(strName...)
}

func GoPath() string {
	return utils.Env("GOPATH")
}

func GoRoot() string {
	return utils.Env("GOROOT")
}

func Dir(strPath string) string {
	return filepath.Dir(strPath)
}
