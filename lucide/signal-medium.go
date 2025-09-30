package lucide

import (
	html "github.com/plainkit/html"
)

// SignalMedium creates a Signal Medium Lucide icon.
func SignalMedium(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-signal-medium", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M2 20h.01")),
		html.SvgPath(html.AD("M7 20v-4")),
		html.SvgPath(html.AD("M12 20v-8")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
