package ui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
)

var (
	headerStyle      = lipgloss.NewStyle().MarginTop(1).MarginBottom(1)
	headerArrowStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("14")).Bold(true) // bright cyan
	headerTitleStyle = lipgloss.NewStyle().Bold(true)
	infoStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("12")).MarginLeft(2).Width(8) // bright blue
	successStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("10")).MarginLeft(2).Width(8) // bright green
	warnStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("11")).MarginLeft(2).Width(8) // bright yellow
	errorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("9")).MarginLeft(2).Width(8) // bright red
	dryStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("8")).MarginLeft(2).Width(8) // grey
	dryBodyStyle = lipgloss.NewStyle().Faint(true)
)

func Die(msg string) {
	Error(msg)
	os.Exit(1)
}

func Dry(msg string) {
	fmt.Println(dryStyle.Render("dry") + dryBodyStyle.Render(msg))
}

func Error(msg string) {
	fmt.Fprintln(os.Stderr, errorStyle.Render("error")+msg)
}

func Header(title string) {
	line := headerArrowStyle.Render("==> ") + headerTitleStyle.Render(title)
	fmt.Println(headerStyle.Render(line))
}

func Info(msg string) {
	fmt.Println(infoStyle.Render("info") + msg)
}

func Success(msg string) {
	fmt.Println(successStyle.Render("ok") + msg)
}

func Warn(msg string) {
	fmt.Println(warnStyle.Render("warn") + msg)
}
