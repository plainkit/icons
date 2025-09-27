package lucide

import (
	html "github.com/plainkit/html"
)

// Monitor creates a Monitor Lucide icon.
func Monitor(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-monitor", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("20"), html.AHeight("14"), html.AX("2"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgLine(html.AX1("8"), html.AX2("16"), html.AY1("21"), html.AY2("21"))),
		html.Child(html.SvgLine(html.AX1("12"), html.AX2("12"), html.AY1("17"), html.AY2("21"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
