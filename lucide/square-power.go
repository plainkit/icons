package lucide

import (
	html "github.com/plainkit/html"
)

// SquarePower creates a Square Power Lucide icon.
func SquarePower(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-power", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 7v4"))),
		html.Child(html.SvgPath(html.AD("M7.998 9.003a5 5 0 1 0 8-.005"))),
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
