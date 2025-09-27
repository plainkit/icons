package lucide

import (
	html "github.com/plainkit/html"
)

// PillBottle creates a Pill Bottle Lucide icon.
func PillBottle(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-pill-bottle", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M18 11h-4a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h4"))),
		html.Child(html.SvgPath(html.AD("M6 7v13a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V7"))),
		html.Child(html.SvgRect(html.AWidth("16"), html.AHeight("5"), html.AX("4"), html.AY("2"), html.ARx("1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
