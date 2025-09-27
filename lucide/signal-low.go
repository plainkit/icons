package lucide

import (
	html "github.com/plainkit/html"
)

// SignalLow creates a Signal Low Lucide icon.
func SignalLow(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-signal-low", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M2 20h.01"))),
		html.Child(html.SvgPath(html.AD("M7 20v-4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
