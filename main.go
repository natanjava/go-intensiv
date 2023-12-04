package main

import "github.com/natanjava/go-intensiv/internal/entity"

type Car struct {
	Model string
	Color string
}

// metodo
func (c Car) Start() {
	println(c.Model + " has been started")
}

func (c *Car) ChangeColor(color string) {
	c.Color = color // duplica o valor de c.color na memoria - copia do color original
	println("New color is: " + c.Color)
}

func soma(x, y int) int { // funcao
	return x + y
}

func main() {
	/*
		car := Car{ // declarando e atribuindo ao mesmo tempo
			Model: "Ferrrari",
			Color: "Red",
		}
		//car.Model = "Fiat" // alternando o valor do atributo
			    println(car.Model)

				car.Start()
				car.ChangeColor("Blue")
				println(car.Color)
	*/

	order, err := entity.NewOrder("1", 10, 1)
	if err != nil {
		println(err.Error())
	} else {
		println(order.ID)
	}

}
