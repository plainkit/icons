package lucide

import (
	html "github.com/plainkit/html"
)

// Heading creates a Heading Lucide icon.
func Heading(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-heading", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M6 12h12"))),
		html.Child(html.SvgPath(html.AD("M6 20V4"))),
		html.Child(html.SvgPath(html.AD("M18 20V4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
