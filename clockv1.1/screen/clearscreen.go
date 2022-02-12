// There are quite many examples of this code on the internet.
// I found the example on golang libary search results here is the link
// https://pkg.go.dev/github.com/GeistInDerSH/clearscreen

package screen

import (
	"os"
	"os/exec"
	"runtime"
)

var clearScreen map[string]func()

func init() {
	clearScreen = make(map[string]func())

	clearScreen["darwin"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	clearScreen["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	clearScreen["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

}

func ClearScreen() {
	function, exists := clearScreen[runtime.GOOS]
	if exists {
		function()
	}
}

func IsSupportedOS() bool {
	_, exists := clearScreen[runtime.GOOS]
	return exists
}
