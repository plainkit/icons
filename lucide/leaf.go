package lucide

import (
	html "github.com/plainkit/html"
)

// Leaf creates a Leaf Lucide icon.
func Leaf(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-leaf", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M11 20A7 7 0 0 1 9.8 6.1C15.5 5 17 4.48 19 2c1 2 2 4.18 2 8 0 5.5-4.78 10-10 10Z"))),
		html.Child(html.SvgPath(html.AD("M2 21c0-3 1.85-5.36 5.08-6C9.5 14.52 12 13 13 12"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
