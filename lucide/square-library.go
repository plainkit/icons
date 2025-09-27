package lucide

import (
	html "github.com/plainkit/html"
)

// SquareLibrary creates a Square Library Lucide icon.
func SquareLibrary(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-library", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M7 7v10"))),
		html.Child(html.SvgPath(html.AD("M11 7v10"))),
		html.Child(html.SvgPath(html.AD("m15 7 2 10"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
