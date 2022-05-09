package printer

import (
	"errors"
	"fmt"
	"os"
	. "github.com/dave/jennifer/jen"
)

func CheckIfFileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File does not exist!")
		return false
	} else {
		fmt.Println(err)
		return false
	}
}

type Names struct {
	Names []string `json:"names"`
}

type Cities struct {
	Cities []string `json:"cities"`
}

// @input: Array with strings, that have to be printed
func GenerateGoFile(strings []string, output_path string) {
	//fmt.Println(output_path)
	f := NewFile("main")
	f.Func().Id("main").Params().Block(
		//Qual("fmt", "Println").Call(Lit(output_path)), // dont know why double slashes gets generated
		Id("strings").Op(":=").Index().String().ValuesFunc(func(g *Group) {
			for i := 0; i < len(strings); i++ {
				g.Lit(strings[i])
			}
		}),
		For(
			Id("i").Op(":=").Lit(0),
			Id("i").Op("<").Lit(len(strings)),
			Id("i").Op("++"),	
		).Block(
			Qual("fmt", "Println").Call(Id("strings").Index(Id("i"))),	
		),
	)
	//fmt.Printf("type of f: %T\n", f)
	f.Save(output_path + "main.go")
	
}

