package lucide

import (
	html "github.com/plainkit/html"
)

// Gauge creates a Gauge Lucide icon.
func Gauge(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-gauge", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m12 14 4-4"))),
		html.Child(html.SvgPath(html.AD("M3.34 19a10 10 0 1 1 17.32 0"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
