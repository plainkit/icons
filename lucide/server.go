package lucide

import (
	html "github.com/plainkit/html"
)

// Server creates a Server Lucide icon.
func Server(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-server", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("20"), html.AHeight("8"), html.AX("2"), html.AY("2"), html.ARx("2"), html.ARy("2")),
		html.SvgRect(html.AWidth("20"), html.AHeight("8"), html.AX("2"), html.AY("14"), html.ARx("2"), html.ARy("2")),
		html.SvgLine(html.AX1("6"), html.AX2("6.01"), html.AY1("6"), html.AY2("6")),
		html.SvgLine(html.AX1("6"), html.AX2("6.01"), html.AY1("18"), html.AY2("18")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
