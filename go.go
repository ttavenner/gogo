package main

import (
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilog"
)

func main() {
	var a, _ = astilectron.New(astilectron.Options{AppName: "GoGo"})

	defer a.Close()

	a.Start()

	var w, _ = a.NewWindow("http://127.0.0.1:3000", &astilectron.WindowOptions{
		Center: astilectron.PtrBool(true),
		Height: astilectron.PtrInt(600),
		Width:  astilectron.PtrInt(600),
	})

	w.Create()

	a.On(astilectron.EventNameAppCrash, func(e astilectron.Event) (deleteListener bool) {
		astilog.Error("App has crashed")
		return
	})

	w.On(astilectron.EventNameWindowEventResize, func(e astilectron.Event) (deleteListener bool) {
		astilog.Info("Window resized")
		return
	})
}
