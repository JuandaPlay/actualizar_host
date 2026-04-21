package main

import (
	"fmt"
	hosts "hostseditor/pkg/hosts"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hosts Editor - Admin")
	w.Resize(fyne.NewSize(520, 180))

	status := widget.NewLabel("Estado: listo")

	checkBtn := widget.NewButton("Verificar si la línea existe", func() {
		exists, err := hosts.IsLinePresent()
		if err != nil {
			status.SetText(fmt.Sprintf("Error: %v", err))
			return
		}
		if exists {
			status.SetText("La línea ya existe en hosts.")
		} else {
			status.SetText("La línea no existe en hosts.")
		}
	})

	addBtn := widget.NewButton("Agregar línea", func() {
		err := hosts.AddLineIfMissing()
		if err != nil {
			status.SetText(fmt.Sprintf("Error: %v", err))
			return
		}
		status.SetText("Línea agregada correctamente debajo de localhost (si no existía).")
	})

	content := container.NewVBox(
		widget.NewLabel("Herramienta para modificar el archivo hosts de Windows (requiere administrador)"),
		checkBtn,
		addBtn,
		status,
	)
	w.SetContent(content)
	w.ShowAndRun()
}
