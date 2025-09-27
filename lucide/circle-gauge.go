package lucide

import (
	html "github.com/plainkit/html"
)

// CircleGauge creates a Circle Gauge Lucide icon.
func CircleGauge(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-gauge", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M15.6 2.7a10 10 0 1 0 5.7 5.7"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("2"))),
		html.Child(html.SvgPath(html.AD("M13.4 10.6 19 5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
