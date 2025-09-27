package lucide

import (
	html "github.com/plainkit/html"
)

// BetweenHorizontalEnd creates a Between Horizontal End Lucide icon.
func BetweenHorizontalEnd(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-between-horizontal-end", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("13"), html.AHeight("7"), html.AX("3"), html.AY("3"), html.ARx("1"))),
		html.Child(html.SvgPath(html.AD("m22 15-3-3 3-3"))),
		html.Child(html.SvgRect(html.AWidth("13"), html.AHeight("7"), html.AX("3"), html.AY("14"), html.ARx("1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
