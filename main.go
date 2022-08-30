//go:generate fyne package -icon .\icon.png
package main

import (
	"puttycmd/putty"
)

func main() {
	putty.Run()
}
