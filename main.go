package main

import (
	"log"

	"github.com/lxn/walk"
)

func main() {
	if flag, err := walk.Clipboard().ContainsText(); err != nil {
		log.Fatalln(err)
		showMsgBox("Error. Failed to access to clipboard.", err)
	} else if flag {
		if text, err := walk.Clipboard().Text(); err != nil {
			showMsgBox("Error. Failed to access to clipboard.", err)
		} else {
			text := convAll(text)
			walk.Clipboard().SetText(text)
		}
	} else {
		showMsgBox("Error. Failed to access to clipboard", nil)
	}
}

func showMsgBox(mesg string, err error) {
	if err != nil {
		log.Println(mesg, err)
	}

	walk.MsgBox(nil, "ClipBoard", mesg, walk.MsgBoxOK)
}
