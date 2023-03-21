package html

import (
	"embed"
	"io"
	"text/template"
)

type ThemeMode bool

type ConfigParams struct {
	Theme ThemeMode
}

type DashboardParams struct {
	ConfigParams
	Greeting string
}

const (
	ThemeModeDark  ThemeMode = true
	ThemeModeLight ThemeMode = false
)

//go:embed templates/*
var files embed.FS

var (
	dashboard = parse("templates/dashboard.html")
)

func Dashboard(w io.Writer, p DashboardParams) error {
	return dashboard.Execute(w, p)
}

func parse(file string) *template.Template {
	return template.Must(
		template.New("layout.html").ParseFS(files, "templates/layout.html", file))
}
