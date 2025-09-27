package lucide

import (
	html "github.com/plainkit/html"
)

// CircleAlert creates a Circle Alert Lucide icon.
func CircleAlert(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-alert", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
		html.Child(html.SvgLine(html.AX1("12"), html.AX2("12"), html.AY1("8"), html.AY2("12"))),
		html.Child(html.SvgLine(html.AX1("12"), html.AX2("12.01"), html.AY1("16"), html.AY2("16"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
