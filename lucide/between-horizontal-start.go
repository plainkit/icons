package lucide

import (
	html "github.com/plainkit/html"
)

// BetweenHorizontalStart creates a Between Horizontal Start Lucide icon.
func BetweenHorizontalStart(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-between-horizontal-start", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("13"), html.AHeight("7"), html.AX("8"), html.AY("3"), html.ARx("1"))),
		html.Child(html.SvgPath(html.AD("m2 9 3 3-3 3"))),
		html.Child(html.SvgRect(html.AWidth("13"), html.AHeight("7"), html.AX("8"), html.AY("14"), html.ARx("1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
