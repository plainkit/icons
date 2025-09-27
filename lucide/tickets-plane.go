package lucide

import (
	html "github.com/plainkit/html"
)

// TicketsPlane creates a Tickets Plane Lucide icon.
func TicketsPlane(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-tickets-plane", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10.5 17h1.227a2 2 0 0 0 1.345-.52L18 12"))),
		html.Child(html.SvgPath(html.AD("m12 13.5 3.75.5"))),
		html.Child(html.SvgPath(html.AD("m4.5 8 10.58-5.06a1 1 0 0 1 1.342.488L18.5 8"))),
		html.Child(html.SvgPath(html.AD("M6 10V8"))),
		html.Child(html.SvgPath(html.AD("M6 14v1"))),
		html.Child(html.SvgPath(html.AD("M6 19v2"))),
		html.Child(html.SvgRect(html.AWidth("20"), html.AHeight("13"), html.AX("2"), html.AY("8"), html.ARx("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
