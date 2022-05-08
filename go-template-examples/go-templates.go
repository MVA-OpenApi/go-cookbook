package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Age  int    `json:"Age"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// use these for the test template
type testObject struct {
	Name []string
	Type []string
	ID   int
}

func (a testObject) PrintFunction() string {
	return "PrintFunction"
}

func (a testObject) PrintName() string {
	return fmt.Sprintf("%v and %v", a.Name[0], a.Name[1])
}

func main() {
	var users Users
	file, err := os.Create("sample.go")
	check(err)
	defer file.Close()
	fmt.Println(file)
	jsonFile, err := os.Open("example.json")
	check(err)
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &users)
	for i := 0; i < len(users.Users); i++ {
		fmt.Fprintf(file, "type %s struct {\n", users.Users[i].Name)
		fmt.Fprintf(file, "\tName string `json:\"name\"`\n")
		fmt.Fprintf(file, "\tType string `json:\"type\"`\n")
		fmt.Fprintf(file, "\tAge int `json:\"age\"`\n")
		fmt.Fprintf(file, "}\n")
	}

	// Create a Template
	a := testObject{[]string{"x", "y"}, []string{"int", "int"}, 2}

	template, err := template.ParseFiles("template.go.tmpl", "split.go.tmpl")
	if err != nil {
		fmt.Println(err.Error())
	}
	template.Execute(os.Stdout, a)
}
