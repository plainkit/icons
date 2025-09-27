package lucide

import (
	html "github.com/plainkit/html"
)

// Heading2 creates a Heading 2 Lucide icon.
func Heading2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-heading-2", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M4 12h8"))),
		html.Child(html.SvgPath(html.AD("M4 18V6"))),
		html.Child(html.SvgPath(html.AD("M12 18V6"))),
		html.Child(html.SvgPath(html.AD("M21 18h-4c0-4 4-3 4-6 0-1.5-2-2.5-4-1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
