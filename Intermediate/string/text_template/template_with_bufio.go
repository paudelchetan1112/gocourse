package main

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter any name:")
name, _:=reader.ReadString('\n')

name=strings.TrimSpace(name)
//Define named templates for different type of 
templates:=map[string]string{
	"welcome":"welcome, {{.name}}! we're glad you joined.", 
	"notification":"{{.name}}, you have a new notification: {{.notification}}", 
	"error":"Oops! An error occurred: {{.errorMessage}}", 

}
//Parse and store templates
parsedTemplates:=make(map[string]*template.Template)
for name, tmpl:=range templates{
	parsedTemplates[name]=template.Must(template.New(name).Parse(tmpl))

}
for{
	fmt.Println("\nMenu:")
	fmt.Println("1. Join")
	fmt.Println("2. Get Notification")
	fmt.Println("3. Get Error")
	fmt.Println("4. Exit")
	fmt.Println("Choose an option:")
	choice,_:=reader.ReadString('\n')
	choice=strings.TrimSpace(choice)
	var data map[string]interface{}
	var tmpl *template.Template
	switch choice {
	case "1":
		tmpl=parsedTemplates["welcome"]
		data=map[string]interface{}{"name":name}
	case "2":
		fmt.Println("Enter your notification Message")
		notification, _:=reader.ReadString('\n')
		notification=strings.TrimSpace(notification)
		tmpl=parsedTemplates["notification"]
		data=map[string]interface{}{"name":name, "notification":notification}
	
case "3":
	fmt.Println("Enter your Error Message:")
	errorMessage, _:=reader.ReadString('\n')
	errorMessage=strings.TrimSpace(errorMessage)
	tmpl=parsedTemplates["errorMessage"]
	data=map[string]interface{}{"name":name, "error":errorMessage}
	
case "4":
	fmt.Println("Exiting...")
	return
default:
	fmt.Println("Invalid Chocie please select a valid option.")

}
	//rener and print the template to the console
	err:=tmpl.Execute(os.Stdout, data)
	if err!=nil{
	fmt.Println("Error executing template:", err)
	}
}
}
	



