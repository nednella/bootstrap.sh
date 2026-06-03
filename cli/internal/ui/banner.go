package ui

import (
	"fmt"
	"runtime"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/nednella/bootstrap.sh/internal"
)

const bannerArt = `     _                 _       _                        _
    | |__   ___   ___ | |_ ___| |_ _ __ __ _ _ __   ___| |__
    | '_ \ / _ \ / _ \| __/ __| __| '__/ _` + "`" + ` | '_ \ / __| '_  \
    | |_) | (_) | (_) | |_\__ \ |_| | | (_| | |_) |\__ \ | | |
    |_.__/ \___/ \___/ \__|___/\__|_|  \__,_| .__(_)___/_| |_|
                                            |_|`

var (
	bannerArtStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("14")).Bold(true) // bright cyan
	bannerCreditStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Align(lipgloss.Right)
	bannerMetaStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Align(lipgloss.Center)
)

func Banner() {
	platform := "macos-" + runtime.GOARCH
	today := time.Now().Format("2 January 2006")
	meta := internal.Version + " · " + platform + " · " + today
	width := lipgloss.Width(bannerArt)

	fmt.Println()
	fmt.Println(bannerArtStyle.Render(bannerArt))
	fmt.Println(bannerCreditStyle.Width(width).Render("by @nednella"))
	fmt.Println()
	fmt.Println(bannerMetaStyle.Width(width).Render(meta))
}
