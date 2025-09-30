package lucide

import (
	html "github.com/plainkit/html"
)

// Gamepad creates a Gamepad Lucide icon.
func Gamepad(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-gamepad", args)
	children := []html.SvgArg{
		html.SvgLine(html.AX1("6"), html.AX2("10"), html.AY1("12"), html.AY2("12")),
		html.SvgLine(html.AX1("8"), html.AX2("8"), html.AY1("10"), html.AY2("14")),
		html.SvgLine(html.AX1("15"), html.AX2("15.01"), html.AY1("13"), html.AY2("13")),
		html.SvgLine(html.AX1("18"), html.AX2("18.01"), html.AY1("11"), html.AY2("11")),
		html.SvgRect(html.AWidth("20"), html.AHeight("12"), html.AX("2"), html.AY("6"), html.ARx("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
