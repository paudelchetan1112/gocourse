package main

import (
	"fmt"
	// "regexp"
	// "strconv"
	"strings"
	// "unicode/utf8"
)

func main(){
// 	str := "Hello Go!"
// 	fmt.Println(len(str))
// 	str1:="Hello"
// 	str2:="World" 
// 	result:=str1 + " " + str2
// 	fmt.Println(result)

// 	fmt.Printf("%s", str[0])
// 	fmt.Print(str[0:8])

// 	//string converrsion

// num:=18
// str3:=strconv.Itoa(num)
// fmt.Println(len(str3))


// //string splitting [string to array]
// fruits:="apple, orange, banana"
// parts:=strings.Split(fruits, ",")
// fmt.Println(parts)
// fruits1:="apple-banana-orange"
// parts1:=strings.Split(fruits1, "-")
// fmt.Println(parts1)


// //join  [array to string]
// countries:=[]string{"Germany", "France", "Italy"}

// joined:=strings.Join(countries, ", ")
// fmt.Println(joined)

// //strings.Contains
// fmt.Println(strings.Contains(str, "Go"))

// //strings.Replace 
// replaced:=strings.Replace(str, "Go", "Golang",1)
// fmt.Println(replaced)

// //string.Trimspace
// strwspace:= " Hello everyone     how are    you and     thenb    ht little and "
// fmt.Println(strings.TrimSpace(strwspace))
// fmt.Println(strings.ToUpper(strwspace))

// fmt.Println(strings.Count(strwspace, "e"))
// fmt.Println(strings.HasPrefix("hello", "he"))
// fmt.Println(strings.HasPrefix("hello", "lo"))



// str5:="Hello, 123, Go 345 my name is chetan paudel and i 18 year old"
// re:=regexp.MustCompile(`\d+`)
// fmt.Println(re.FindAllString(str5, -1))

// str6:="नेपाली हो नेपाली "
// fmt.Println(utf8.RuneCountInString(str6))

//String Builder
var builder strings.Builder
builder.WriteString("Hello")
builder.WriteString(", ")
builder.WriteString("World!")

//convert builder to a string

result := builder.String()
fmt.Println(result)

//suing Writerune to add a character

builder.WriteRune(' ')
builder.WriteString("How are you ")
result=builder.String()
fmt.Println(result)

//Reset the builder
builder.Reset()
builder.WriteString("Starting fresh!")
result=builder.String()
fmt.Println(result)




}