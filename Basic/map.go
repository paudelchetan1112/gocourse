package main

import (
	"fmt"
	"maps"
)

func main() {
	//var mapVariable map[keyType]valueType

	//mapVariable=make(map[keyType]valueType)
	// mapVariable=map[keyType]valueType{
	// 	key1:value1,
	// 	key2:value2,
	// 	key3:value3
	// }

	myMap := make(map[string]int)
	fmt.Println(myMap)
	myMap["key1"]=11
	myMap["code"]=18
	fmt.Println(myMap["key1"])
myMap["code"]=21
fmt.Println(myMap["code"])
// delete(myMap, "key1")
// fmt.Println(myMap)
// clear(myMap)

value, unknownValue:=myMap["key1"]
fmt.Println(value)
_,isvlaue:=myMap["key2"]
fmt.Println(isvlaue)

fmt.Println(unknownValue)

myMap2:=map[string]int{"a":1,"b":2}
myMap3:=map[string]int{"a":1, "b":2}
if maps.Equal(myMap2,myMap3){
	fmt.Println("myMap3 and Mymap are equal")

}
for _, v:=range myMap2{
	fmt.Println(v)

}
_,ok:=myMap["key1"]
if ok{
	fmt.Println("value is exist with key1:",ok)

}
var myMap4 map[string]string
// val:=myMap4["key"]
// fmt.Println("value:",val)
myMap4=make(map[string]string)
myMap4["key"]="value"
fmt.Println(myMap4)
fmt.Println("length of mymap4", len(myMap4))
myMap5:=make(map[string]map[string]string)
myMap5["key"]=myMap4
fmt.Println(myMap5)

}