package lucide

import (
	html "github.com/plainkit/html"
)

// Heading1 creates a Heading 1 Lucide icon.
func Heading1(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-heading-1", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M4 12h8"))),
		html.Child(html.SvgPath(html.AD("M4 18V6"))),
		html.Child(html.SvgPath(html.AD("M12 18V6"))),
		html.Child(html.SvgPath(html.AD("m17 12 3-2v8"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
