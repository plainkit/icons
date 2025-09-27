package lucide

import (
	html "github.com/plainkit/html"
)

// Asterisk creates a Asterisk Lucide icon.
func Asterisk(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-asterisk", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 6v12"))),
		html.Child(html.SvgPath(html.AD("M17.196 9 6.804 15"))),
		html.Child(html.SvgPath(html.AD("m6.804 9 10.392 6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
