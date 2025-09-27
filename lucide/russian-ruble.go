package lucide

import (
	html "github.com/plainkit/html"
)

// RussianRuble creates a Russian Ruble Lucide icon.
func RussianRuble(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-russian-ruble", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M6 11h8a4 4 0 0 0 0-8H9v18"))),
		html.Child(html.SvgPath(html.AD("M6 15h8"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
