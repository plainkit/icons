package lucide

import (
	html "github.com/plainkit/html"
)

// Turntable creates a Turntable Lucide icon.
func Turntable(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-turntable", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10 12.01h.01"))),
		html.Child(html.SvgPath(html.AD("M18 8v4a8 8 0 0 1-1.07 4"))),
		html.Child(html.SvgCircle(html.ACx("10"), html.ACy("12"), html.AR("4"))),
		html.Child(html.SvgRect(html.AWidth("20"), html.AHeight("16"), html.AX("2"), html.AY("4"), html.ARx("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
