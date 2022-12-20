package generator

import (
	"errors"
	"io/fs"
	"os"
	"os/exec"

	"github.com/dave/jennifer/jen"
)

func Generate(projectName string) error {
	if projectName == "" {
		return errors.New("project name must be identified, please use -p flag")
	}
	if err := genGoMod(projectName); err != nil {
		return err
	}
	return genMainFile()
}

func genGoMod(projectName string) error {
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

func genMainFile() error {
	mainFile := jen.NewFilePath("main")
	mainFile.Func().Id("main").Params().Block(
		jen.Err().Op(":=").Qual("github.com/joho/godotenv", "Load").Call(),
		jen.If(jen.Err().Op("!=").Nil()).Block(
			jen.Panic(jen.Err()),
		),

		jen.Qual("github.com/nkien0204/lets-go/cmd", "Execute").Call(),
	)
	return mainFile.Save("main.go")
}
