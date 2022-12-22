package off

import (
	"errors"
	"io/fs"
	"os"
	"os/exec"

	"github.com/dave/jennifer/jen"
)

type OfflineGenerator struct {
	ProjectName string
}

func (off *OfflineGenerator) Generate() error {
	if off.ProjectName == "" {
		return errors.New("project name must be identified, please use -p flag")
	}
	if err := off.genGoMod(off.ProjectName); err != nil {
		return err
	}
	return off.genMainFile()
}

func (off *OfflineGenerator) genGoMod(projectName string) error {
	if _, err := os.Stat(projectName); !errors.Is(err, fs.ErrNotExist) {
		return fs.ErrExist
	}
	if err := os.Mkdir(projectName, 0755); err != nil {
		return err
	}
	if err := os.Chdir(projectName); err != nil {
		return err
	}
	cmd := exec.Command("go", "mod", "init", projectName)
	return cmd.Run()
}

func (off *OfflineGenerator) genMainFile() error {
	mainFile := jen.NewFilePath("main")
	mainFile.Func().Id("main").Params().Block(
		jen.Qual("github.com/nkien0204/lets-go/cmd", "Execute").Call(),
	)
	return mainFile.Save("main.go")
}
