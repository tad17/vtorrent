package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gtk"
	"os"
)

const VERSION = "0.01"

func main() {
	fmt.Printf("vtorrent v%s\n", VERSION)

	gtk.Init(&os.Args)

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	setupMainWindow(window)

	gtk.Main()
}

func setupMainWindow(window *gtk.Window) {
	window.SetTitle("Видео торрент")
	window.Connect("destroy", gtk.MainQuit)

	vbox := getVBox()
	window.Add(vbox)

	window.ShowAll()
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetSizeRequest(600, 400)
}

func getVBox() *gtk.VBox {
	vbox := gtk.NewVBox(true, 0)
	vbox.SetBorderWidth(5)

	button := gtk.NewButtonWithLabel("Hello World!")
	vbox.Add(button)

	button2 := gtk.NewButtonWithLabel("Hello World 2!")
	vbox.Add(button2)
	return vbox
}
