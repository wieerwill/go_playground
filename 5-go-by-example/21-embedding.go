// embedding of structs and interfaces to express a more seamless composition of types

package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// container embeds a base
// embedding looks like a field without a name
type container struct {
	base
	str string
}

func main() {
	// when creating structs with literals, initialize the embedding explicitly
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// access the base’s fields directly
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// full path using the embedded type name
	fmt.Println("also num:", co.base.num)

	// container embeds base
	// methods of base also become methods of a container
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// container implements describer interface because it embeds base
	var d describer = co
	fmt.Println("describer:", d.describe())
}
