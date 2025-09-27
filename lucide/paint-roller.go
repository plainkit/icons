package lucide

import (
	html "github.com/plainkit/html"
)

// PaintRoller creates a Paint Roller Lucide icon.
func PaintRoller(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-paint-roller", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("16"), html.AHeight("6"), html.AX("2"), html.AY("2"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M10 16v-2a2 2 0 0 1 2-2h8a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2h-2"))),
		html.Child(html.SvgRect(html.AWidth("4"), html.AHeight("6"), html.AX("8"), html.AY("16"), html.ARx("1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
