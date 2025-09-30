package lucide

import (
	html "github.com/plainkit/html"
)

// Activity creates a Activity Lucide icon.
func Activity(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-activity", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M22 12h-2.48a2 2 0 0 0-1.93 1.46l-2.35 8.36a.25.25 0 0 1-.48 0L9.24 2.18a.25.25 0 0 0-.48 0l-2.35 8.36A2 2 0 0 1 4.49 12H2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
