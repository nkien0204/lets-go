package generator

import (
	"os/exec"

	"github.com/dave/jennifer/jen"
)

func Generate(projectName string) {
	genGoMod(projectName)
	genMainFile()
}

func genGoMod(projectName string) {
	cmd := exec.Command("go", "mod", "init", projectName)
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func genMainFile() {
	mainFile := jen.NewFilePath("main")
	mainFile.Func().Id("main").Params().Block(
		jen.Err().Op(":=").Qual("github.com/joho/godotenv", "Load").Call(),
		jen.If(jen.Err().Op("!=").Nil()).Block(
			jen.Panic(jen.Err()),
		),

		jen.Qual("github.com/nkien0204/lets-go/cmd", "Execute").Call(),
	)
	mainFile.Save("main.go")
}
