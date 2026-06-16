package main

import (
	"fmt"
	"net/url"
)

func main() {
	//[scheme://][userinfo@]host[:port][/parth][?query][#fragment]

	rawURL := "https://example.com:8080/path?query=param#fragment"
	parsedURL, err := url.Parse(rawURL)
	if err!=nil{
		fmt.Println("Error parsing url:", err)
		return 

	}
	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Schema:", parsedURL.Host)
	fmt.Println("port:", parsedURL.Port())
	fmt.Println("Path", parsedURL.Path)
	fmt.Println("Raw query:", parsedURL.RawQuery)
	fmt.Println("fragment:", parsedURL.Fragment)
	

	rawURL1:="https://example.com/path?name=John&age=30"
	parseURL1, err:=url.Parse(rawURL1)
		if err!=nil{
		fmt.Println("Error parsing url:", err)
		return 
		}
		queryParams:=parseURL1.Query()
		fmt.Println(queryParams)
		fmt.Println("Name:", queryParams.Get("name"))
		fmt.Println("Age:", queryParams.Get("age"))


		//Building Url

		baseUrl:=&url.URL{
			Scheme: "https",
			Host: "example.com",
			Path: "/path",

		}
		query:=baseUrl.Query()
		query.Set("name", "john")
		baseUrl.RawQuery=query.Encode()
		fmt.Println("Built Url:", baseUrl.String())


	value:=url.Values{}
	//Add key-value paris to the value objects 
	value.Add("name", "jane")
	value.Add("age", "30")
	value.Add("city", "london")
	value.Add("country", "UK")
	//encode
	encodedQuery:= value.Encode()
	fmt.Println(encodedQuery)

	//Build a URL
	baseURL:="https://example.com/search"
	fullURL:=baseURL + "?" +encodedQuery
	fmt.Println(fullURL)


}