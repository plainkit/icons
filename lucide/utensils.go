package lucide

import (
	html "github.com/plainkit/html"
)

// Utensils creates a Utensils Lucide icon.
func Utensils(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-utensils", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M3 2v7c0 1.1.9 2 2 2h4a2 2 0 0 0 2-2V2"))),
		html.Child(html.SvgPath(html.AD("M7 2v20"))),
		html.Child(html.SvgPath(html.AD("M21 15V2a5 5 0 0 0-5 5v6c0 1.1.9 2 2 2h3Zm0 0v7"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
