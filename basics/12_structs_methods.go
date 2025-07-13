// adapted from: https://withcodeexample.com/golang-tutorial-for-beginners/
package basics

import "fmt"

// Defining a 'Human' struct
type Human struct {
	FirstName string
	LastName  string
	Age       int
}

// Adding a method to the 'Human' struct
func (p Human) FullName() string {
	return p.FirstName + " " + p.LastName
}

// Defining a struct type
type Rectangle struct {
	Width  float64
	Height float64
}

// Method associated with the Rectangle type
// Use a value receiver if the method does not modify the original value.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Method with a pointer receiver
// Use a pointer receiver if the method needs to modify the original value.
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// Defining a struct type
type Point struct {
	X float64
	Y float64
}

// Method with a pointer receiver and a pointer parameter
func (p *Point) MoveTo(newX, newY float64) {
	p.X = newX
	p.Y = newY
}

// Defining a struct type
type Circle struct {
	Radius float64
}

// Method with named return values
func (c Circle) DiameterAndCircumference() (diameter, circumference float64) {
	diameter = 2 * c.Radius
	circumference = 2 * 3.14 * c.Radius
	return
}

// Defining a struct named 'Square'
type Square struct {
	Width float64
}

// Defining an interface named 'Shape'
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Implementing the 'Shape' interface for 'Square'
func (r Square) Area() float64 {
	return r.Width * r.Width
}

func (r Square) Perimeter() float64 {
	return 4 * r.Width
}

// Defining an interface named 'Animal'
type Animal interface {
	Speak() string
}

// Implementing the 'Animal' interface for 'Dog'
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

// Implementing the 'Animal' interface for 'Cat'
type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

func StructsMethods() {
	// Creating an instance of the 'Human' struct
	human := Human{
		FirstName: "Charlie",
		LastName:  "Brown",
		Age:       28,
	}
	// Calling the 'FullName' method
	fullName := human.FullName()
	fmt.Println("Full Name:", fullName)

	// Creating an instance of the Rectangle struct
	rectangle := Rectangle{Width: 5, Height: 10}
	// Calling the method on the struct instance
	area := rectangle.Area()
	fmt.Println("Area:", area)

	// Calling the method with a pointer receiver
	rectangle.Scale(2)
	scaledArea := rectangle.Area()
	fmt.Println("Scaled Area:", scaledArea)

	// Creating an instance of the Point struct
	point := Point{X: 3, Y: 7}
	// Calling the method with a pointer receiver and a pointer parameter
	point.MoveTo(10, 20)
	fmt.Printf("New X: %.2f, New Y: %.2f\n", point.X, point.Y)

	// Creating an instance of the Circle struct
	circle := Circle{Radius: 5}
	// Calling the method with named return values
	diameter, circumference := circle.DiameterAndCircumference()
	fmt.Printf("Diameter: %.2f, Circumference: %.2f\n", diameter, circumference)

	// Creating an instance of the 'Square' struct
	square := Square{Width: 5}
	// Using the 'Shape' interface
	var shape Shape
	shape = square
	// Calling methods defined by the 'Shape' interface
	squareArea := shape.Area()
	perimeter := shape.Perimeter()
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", squareArea, perimeter)

	// Creating instances of 'Dog' and 'Cat'
	dog := Dog{}
	cat := Cat{}
	// Using the 'Animal' interface with different implementations
	var animal Animal
	animal = dog
	fmt.Println("Dog says:", animal.Speak())
	animal = cat
	fmt.Println("Cat says:", animal.Speak())
}
