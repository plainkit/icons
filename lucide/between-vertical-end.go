package lucide

import (
	html "github.com/plainkit/html"
)

// BetweenVerticalEnd creates a Between Vertical End Lucide icon.
func BetweenVerticalEnd(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-between-vertical-end", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("7"), html.AHeight("13"), html.AX("3"), html.AY("3"), html.ARx("1"))),
		html.Child(html.SvgPath(html.AD("m9 22 3-3 3 3"))),
		html.Child(html.SvgRect(html.AWidth("7"), html.AHeight("13"), html.AX("14"), html.AY("3"), html.ARx("1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
