package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("He said, \" Iam great\"")
	fmt.Println(`He said, "Iam great"`)
	//compile a regex pattern to match email address
	re:=regexp.MustCompile(`^[a-zA-Z0-9]+[a-zA-Z0-9._%+-]*@[a-zA-Z0-9-]+\.[a-zA-Z]{2,}$`)
	//Test strings
	email1:="use234-.+r@gmail.com"
	email2:="invalid_email"
//regex that accept the only lower case letter only
reg:=regexp.MustCompile(`^[a-z]+$`)
name:="cheta45n"
//regex that accept only number 
// regn:=regexp.MustCompile(`^[0-9]+$`)
//matcfh a string that starts with go 
string1:="going"
rego:=regexp.MustCompile(`^go[a-z]*$`)
fmt.Println("start with go:", rego.MatchString(string1))

//matching a valid username
//4 to 12 characters long Can contain lowercase letters, digits, and underscore (_)
regox:=regexp.MustCompile(`^[a-z0-9_]{4,12}$`)
username:="chetan_1112_"
fmt.Println("valid username:", regox.MatchString(username))
fmt.Println(reg.MatchString(name))

//regex which accept email	//match
	fmt.Println("Email:",re.MatchString(email1))
	
	fmt.Println("Email:",re.MatchString(email2))



	//capturing Groups
	//Compile a regex pattern to capture date coponents;
	re=regexp.MustCompile(`(\d){4}-(\d{2})-(\d{2})`)
	//Test string
	date:="2024-07-30"
//Find all submatches
submatches:=re.FindStringSubmatch(date)
fmt.Println(submatches)
fmt.Println(submatches[0])
fmt.Println(submatches[1])
fmt.Println(submatches[2])
fmt.Println(submatches[3])


//Target String
str:="Hello World"
re=regexp.MustCompile(`[aeiou]`)
result:=re.ReplaceAllString(str,"*")
fmt.Println(result)


//i-case insensitive
//m-multiline model
//s-dot matches
re=regexp.MustCompile(`(?i)go`)
//Test string
text:="Golang is going great"
//match
fmt.Println("match:", re.MatchString(text))


//regular expressing for accepting nepali phone number 977-9868103671||97.......
regex:=regexp.MustCompile(`^977-(98|97)[0-9]{8}$`)

}