package lucide

import (
	html "github.com/plainkit/html"
)

// Subscript creates a Subscript Lucide icon.
func Subscript(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-subscript", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m4 5 8 8"))),
		html.Child(html.SvgPath(html.AD("m12 5-8 8"))),
		html.Child(html.SvgPath(html.AD("M20 19h-4c0-1.5.44-2 1.5-2.5S20 15.33 20 14c0-.47-.17-.93-.48-1.29a2.11 2.11 0 0 0-2.62-.44c-.42.24-.74.62-.9 1.07"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
