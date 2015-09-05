package gin

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

type Builder interface {
	Build() error
	Binary() string
	Errors() string
}

type builder struct {
	dir          string
	mainFilePath string
	binary       string
	errors       string
	useGodep     bool
}

//2015/09/05 stefan.zan: add mainFilePath
func NewBuilder(dir string, mainFilePath string, bin string, useGodep bool) Builder {
	if len(bin) == 0 {
		bin = "bin"
	}

	// does not work on Windows without the ".exe" extension
	if runtime.GOOS == "windows" {
		if !strings.HasSuffix(bin, ".exe") { // check if it already has the .exe extension
			bin += ".exe"
		}
	}
	//2015/09/05 stefan.zan: add mainFilePath
	return &builder{dir: dir, mainFilePath: mainFilePath, binary: bin, useGodep: useGodep}
}

func (b *builder) Binary() string {
	return b.binary
}

func (b *builder) Errors() string {
	return b.errors
}

func (b *builder) Build() error {
	var command *exec.Cmd
	if b.useGodep {
		command = exec.Command("godep", "go", "build", "-o", b.binary)
	} else {
		//2015/09/05 stefan.zan:  modified: add mainFilePath
		command = exec.Command("go", "build", "-o", b.binary, b.mainFilePath)
	}
	command.Dir = b.dir

	output, err := command.CombinedOutput()

	if command.ProcessState.Success() {
		b.errors = ""
	} else {
		b.errors = string(output)
	}

	if len(b.errors) > 0 {
		return fmt.Errorf(b.errors)
	}

	return err
}
