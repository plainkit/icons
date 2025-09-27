package lucide

import (
	html "github.com/plainkit/html"
)

// Minus creates a Minus Lucide icon.
func Minus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-minus", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M5 12h14"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
