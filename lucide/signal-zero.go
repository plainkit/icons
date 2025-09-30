package lucide

import (
	html "github.com/plainkit/html"
)

// SignalZero creates a Signal Zero Lucide icon.
func SignalZero(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-signal-zero", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M2 20h.01")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
