package clcolor

import (
	"fmt"
	"runtime"
)

const (
	// TextBlack ...
	TextBlack = iota + 30

	// TextRed ...
	TextRed
	// TextGreen ...
	TextGreen

	// TextYellow ...
	TextYellow

	// TextBlue ...
	TextBlue

	// TextMagenta ...
	TextMagenta

	// TextCyan ...
	TextCyan

	// TextWhite ...
	TextWhite
)

// Black ...
func Black(str string) string {
	return textColor(TextBlack, str)
}

// Red ...
func Red(str string) string {
	return textColor(TextRed, str)
}

// Green ...
func Green(str string) string {
	return textColor(TextGreen, str)
}

// Yellow ...
func Yellow(str string) string {
	return textColor(TextYellow, str)
}

// Blue ...
func Blue(str string) string {
	return textColor(TextBlue, str)
}

// Magenta ...
func Magenta(str string) string {
	return textColor(TextMagenta, str)
}

// Cyan ...
func Cyan(str string) string {
	return textColor(TextCyan, str)
}

// White ...
func White(str string) string {
	return textColor(TextWhite, str)
}

func textColor(color int, str string) string {
	if IsWindows() {
		return str
	}

	switch color {
	case TextBlack:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextBlack, str)
	case TextRed:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextRed, str)
	case TextGreen:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextGreen, str)
	case TextYellow:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextYellow, str)
	case TextBlue:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextBlue, str)
	case TextMagenta:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextMagenta, str)
	case TextCyan:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextCyan, str)
	case TextWhite:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextWhite, str)
	default:
		return str
	}
}

// IsWindows ...
func IsWindows() bool {
	return runtime.GOOS == "windows"
}
