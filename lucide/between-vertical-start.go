package lucide

import (
	html "github.com/plainkit/html"
)

// BetweenVerticalStart creates a Between Vertical Start Lucide icon.
func BetweenVerticalStart(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-between-vertical-start", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("7"), html.AHeight("13"), html.AX("3"), html.AY("8"), html.ARx("1"))),
		html.Child(html.SvgPath(html.AD("m15 2-3 3-3-3"))),
		html.Child(html.SvgRect(html.AWidth("7"), html.AHeight("13"), html.AX("14"), html.AY("8"), html.ARx("1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
