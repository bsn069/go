package path

import (
	"testing"
)

func Test_Look(t *testing.T) {
	strGoExePath, _ := Look("go")
	strGoBinPath := Join(GoRoot(), "bin")
	strGoExePath = Dir(strGoExePath)
	if strGoExePath != strGoBinPath {
		t.Fatal()
	}
}

func Test_Pwd(t *testing.T) {
	strGoPath := GoPath()
	t.Log(strGoPath)
	strCurPath := Join(strGoPath, "src/github.com/bsn069/go/path")
	t.Log(Pwd())
	if Pwd() != strCurPath {
		t.Fatal()
	}
}
