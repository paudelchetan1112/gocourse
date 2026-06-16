package main

import (
	// "fmt"
	"html/template"
	"os"
)

func main(){
	// tmpl:=template.New("Example")
	// tmpl, err:=template.New("Example").Parse("Welcome, {{.name}}! How are you doing")
tmpl:=template.Must(template.New("Example").Parse("Welcome, {{.name}}! How are you doing? \n"))
	// if err!=nil{
	// 	panic(err)
	// }
	//Define data for the welcome messaage template
	data:=map[string]interface{} {
		"name":"john",

	}
	err=tmpl.Execute(os.Stdout, data)
	if err!=nil{
		panic(err)

	}





}