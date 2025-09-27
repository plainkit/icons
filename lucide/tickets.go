package lucide

import (
	html "github.com/plainkit/html"
)

// Tickets creates a Tickets Lucide icon.
func Tickets(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-tickets", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m4.5 8 10.58-5.06a1 1 0 0 1 1.342.488L18.5 8"))),
		html.Child(html.SvgPath(html.AD("M6 10V8"))),
		html.Child(html.SvgPath(html.AD("M6 14v1"))),
		html.Child(html.SvgPath(html.AD("M6 19v2"))),
		html.Child(html.SvgRect(html.AWidth("20"), html.AHeight("13"), html.AX("2"), html.AY("8"), html.ARx("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
