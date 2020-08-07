package main

import (
	"log"
	"os"

	"github.com/gotk3/gotk3/gtk"
	app "github.com/yusufpapurcu/FileSearcher/app"
)

const WindowName = "window"
const ChName = "chbox"
const ListboxName = "listbox"
const ButtonName = "button"
const EntryName = "entry"

// UIMain is a main glade file location
var UIMain = os.Getenv("GOPATH") + "/src/github.com/yusufpapurcu/FileSearcher/glade/mainWindow.glade"

func main() {

	gtk.Init(&os.Args)

	bldr, err := getBuilder(UIMain)
	if err != nil {
		panic(err)
	}

	window, err := getWindow(bldr)
	if err != nil {
		panic(err)
	}

	window.SetTitle("File Searcher")
	window.SetDefaultSize(365, 490)
	_, err = window.Connect("destroy", func() {
		gtk.MainQuit()
	})
	if err != nil {
		panic(err)
	}
	window.ShowAll()

	button, err := getButton(bldr)
	if err != nil {
		panic(err)
	}

	entr, err := getEntry(bldr)
	if err != nil {
		panic(err)
	}
	chb, err := getChb(bldr)
	if err != nil {
		panic(err)
	}
	_, err = button.Connect("clicked", func() {
		text, err := entr.GetText()
		if err != nil {
			panic(err)
		}
		if chb.GetActive() {
			err = loadlist(bldr, app.KeywordSearch(".", text))
			if err != nil {
				log.Fatal(err)
			}
			return
		}
		err = loadlist(bldr, app.FileSearch(".", text))
		if err != nil {
			log.Fatal(err)
		}
	})
	if err != nil {
		panic(err)
	}

	gtk.Main()
}

func getBuilder(filename string) (*gtk.Builder, error) {

	b, err := gtk.BuilderNew()
	if err != nil {
		return nil, err
	}

	if filename != "" {
		err = b.AddFromFile(filename)
		if err != nil {
			return nil, err
		}
	}

	return b, nil
}

func getWindow(b *gtk.Builder) (*gtk.Window, error) {

	obj, err := b.GetObject(WindowName)
	if err != nil {
		return nil, err
	}

	window, ok := obj.(*gtk.Window)
	if !ok {
		return nil, err
	}

	return window, nil
}

func getChb(b *gtk.Builder) (*gtk.CheckButton, error) {

	obj, err := b.GetObject(ChName)
	if err != nil {
		return nil, err
	}

	window, ok := obj.(*gtk.CheckButton)
	if !ok {
		return nil, err
	}

	return window, nil
}

func getButton(b *gtk.Builder) (*gtk.Button, error) {

	obj, err := b.GetObject(ButtonName)
	if err != nil {
		return nil, err
	}

	button, ok := obj.(*gtk.Button)
	if !ok {
		return nil, err
	}

	return button, nil
}

func getEntry(b *gtk.Builder) (*gtk.Entry, error) {

	obj, err := b.GetObject(EntryName)
	if err != nil {
		return nil, err
	}

	button, ok := obj.(*gtk.Entry)
	if !ok {
		return nil, err
	}

	return button, nil
}

func getListbox(b *gtk.Builder) (*gtk.ListBox, error) {

	obj, err := b.GetObject(ListboxName)
	if err != nil {
		return nil, err
	}

	lb, ok := obj.(*gtk.ListBox)
	if !ok {
		return nil, err
	}

	return lb, nil
}

func loadlist(b *gtk.Builder, data []string) error {

	lb, err := getListbox(b)
	if err != nil {
		return err
	}

	for index, element := range data {

		lbl, err := gtk.LabelNew(element)
		if err != nil {
			return err
		}

		lbl.SetXAlign(0)
		lbl.SetMarginStart(5)

		row, err := gtk.ListBoxRowNew()
		if err != nil {
			return err
		}

		row.Add(lbl)

		lb.Insert(row, index)
	}

	lb.ShowAll()

	return nil
}
