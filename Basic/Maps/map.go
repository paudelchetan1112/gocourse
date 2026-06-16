package main

import (
	"fmt"
	"maps"
)

func main() {
	// var mapVariable map[keyType] valueType

	//mapVariable=make(map[keyType])
	//using a Map literal
	// mapVariable=map[keyType]valueType{
	// 	key1:value1,
	// 	key2:value2,
	// }
	myMap := make(map[string]int)
	fmt.Println(myMap)
	myMap["Key1"]=9
	fmt.Println(myMap)
	myMap["key2"]=18
	fmt.Println(myMap)
	myMap["key2"]=21
	fmt.Println(myMap)
	// delete(myMap, "key1")
	fmt.Println(myMap)
	// clear(myMap)
	fmt.Println(myMap)
	_, unkownvalue:=myMap["key1"]

fmt.Println(unkownvalue)
myMap2:=map[string]int{"a":1, "b":2}
if maps.Equal(myMap, myMap2){
	fmt.Println("my map and myMap 2 are equal")
}else{
	fmt.Println("not equl")
}
for k, v:=range myMap2{
	fmt.Println(k, v)
}
_,value:=myMap2["a"]
fmt.Println(value)

var myMap4 map[string]string
if myMap4==nil{
	fmt.Println("The map is initialized to nil value")

}else{
	fmt.Println("The map is not initialized to nil value. ")


}
val:=myMap4["key"]
fmt.Println(val)
fmt.Println(len(myMap4))

myMap4=make(map[string]string)

fmt.Println(myMap)


}