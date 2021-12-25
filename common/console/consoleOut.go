package console

import (
	"fmt"
)

const format = "%s%s%s"
const colorBlue = "\x1b[34;1m"
const colorCyan = "\x1b[36;1m"
const colorGreen = "\x1b[32;1m"
const colorMagenta = "\x1b[35;1m"
const colorRed = "\x1b[31;1m"
const colorWhite = "\x1b[37;1m"
const colorYellow = "\x1b[33;1m"
const colorNone = "\x1b[0m"

// Blue prints formatted text in blue
func Blue(text string) {
	fmt.Printf(format, colorBlue, text, colorNone)
}

// Cyan prints formatted text in cyan
func Cyan(text string) {
	fmt.Printf(format, colorCyan, text, colorNone)
}

// Green prints formatted text in green
func Green(text string) {
	fmt.Printf(format, colorGreen, text, colorNone)
}

// Magenta prints formatted text in magenta
func Magenta(text string) {
	fmt.Printf(format, colorMagenta, text, colorNone)
}

// Red prints formatted text in red
func Red(text string) {
	fmt.Printf(format, colorRed, text, colorNone)
}

// White prints formatted text in white
func White(text string) {
	fmt.Printf(format, colorWhite, text, colorNone)
}

// Yellow prints formatted text in yellow
func Yellow(text string) {
	fmt.Printf(format, colorYellow, text, colorNone)
}
