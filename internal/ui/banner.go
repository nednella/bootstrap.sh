package ui

import (
	"fmt"
	"runtime"
	"time"

	"github.com/charmbracelet/lipgloss"
)

const bannerArt = `     _                 _       _                        _
    | |__   ___   ___ | |_ ___| |_ _ __ __ _ _ __   ___| |__
    | '_ \ / _ \ / _ \| __/ __| __| '__/ _` + "`" + ` | '_ \ / __| '_  \
    | |_) | (_) | (_) | |_\__ \ |_| | | (_| | |_) |\__ \ | | |
    |_.__/ \___/ \___/ \__|___/\__|_|  \__,_| .__(_)___/_| |_|
                                            |_|`

const rule = "════════════════════════════════════════════════════════════════════════════════════════════════════"

var (
	ruleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("8")).MarginTop(1).MarginBottom(1) // dim grey, blank above and below — shared for all rules

	bannerArtStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("14")).Bold(true)              // bright cyan
	bannerCreditStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("8")).MarginLeft(50)           // sits under the art's right edge
	bannerTaglineStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("8")).MarginLeft(4).MarginTop(1)

	farewellBoldStyle = lipgloss.NewStyle().Bold(true)
)

func Banner() {
	platform := "macos-" + runtime.GOARCH
	today := time.Now().Format("2 January 2006")

	fmt.Println(ruleStyle.Render(rule))
	fmt.Println(bannerArtStyle.Render(bannerArt))
	fmt.Println(bannerCreditStyle.Render("by @nednella"))
	fmt.Println(bannerTaglineStyle.Render("Bootstrap a fresh Mac quicker than the time it takes to make a cuppa."))
	fmt.Println(bannerTaglineStyle.Render(fmt.Sprintf("%s (%s)", platform, today)))
	fmt.Println(ruleStyle.Render(rule))
}

func Farewell(elapsed int) {
	fmt.Println(ruleStyle.Render(rule))
	fmt.Println(successStyle.Render("ok") + farewellBoldStyle.Render("Bootstrap complete") + fmt.Sprintf(" in %ds.", elapsed))
	fmt.Println(ruleStyle.Render(rule))
}
