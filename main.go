package main

import (
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"html/template"
	"log/slog"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func main() {

	http.HandleFunc("/", indexRender)
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		slog.Error("Error to serve Application.... ", err)
	}

	func() {
		err := open.Run("http://localhost:8000")
		if err != nil {
			fmt.Println("Erro ao abrir o navegador:", err)
		}
	}()

	slog.Info("Application running on port 8000")
	slog.Info("Press Ctrl + C to stop the application")
	slog.Info("Access http://localhost:8000")

}

func indexRender(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{"Mouse", "Mouse Gamer", 100.00, 10},
		{"Teclado", "Teclado Gamer", 200.00, 10},
		{"Monitor", "Monitor Gamer", 300.00, 10},
		{"Mousepad", "Mousepad Gamer", 50.00, 10},
		{"Headset", "Headset Gamer", 150.00, 10},
	}

	temp.ExecuteTemplate(w, "Index", products)

}
