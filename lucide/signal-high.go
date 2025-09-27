package lucide

import (
	html "github.com/plainkit/html"
)

// SignalHigh creates a Signal High Lucide icon.
func SignalHigh(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-signal-high", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M2 20h.01"))),
		html.Child(html.SvgPath(html.AD("M7 20v-4"))),
		html.Child(html.SvgPath(html.AD("M12 20v-8"))),
		html.Child(html.SvgPath(html.AD("M17 20V8"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
