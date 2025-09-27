package lucide

import (
	html "github.com/plainkit/html"
)

// Pill creates a Pill Lucide icon.
func Pill(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-pill", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m10.5 20.5 10-10a4.95 4.95 0 1 0-7-7l-10 10a4.95 4.95 0 1 0 7 7Z"))),
		html.Child(html.SvgPath(html.AD("m8.5 8.5 7 7"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
