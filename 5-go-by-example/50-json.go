package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// demonstrate encoding and decoding of custom types below
type response1 struct {
	Page   int
	Fruits []string
}

// only exported fields will be encoded/decoded in JSON
// Fields must start with capital letters to be exported
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

// Go offers built-in support for JSON encoding and decoding, including to and from built-in and custom data types
func main() {
	// encoding basic data types to JSON strings
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// some for slices and maps, which encode to JSON arrays and objects
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// the JSON package can automatically encode your custom data types
	// it will only include exported fields in the encoded output and will by default use those names as the JSON keys
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// use tags on struct field declarations to customize the encoded JSON key names
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// generic decoding JSON data into Go values
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// provide a variable where the JSON package can put the decoded data
	var dat map[string]interface{}

	// actual decoding, and a check for associated errors
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// in order to use the values in the decoded map, convert them to their appropriate type
	num := dat["num"].(float64)
	fmt.Println(num)

	// accessing nested data requires a series of conversions
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// decode JSON into custom data types
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// stream JSON encodings directly to os.Writers
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
