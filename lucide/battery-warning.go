package lucide

import (
	html "github.com/plainkit/html"
)

// BatteryWarning creates a Battery Warning Lucide icon.
func BatteryWarning(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-battery-warning", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 17h.01")),
		html.SvgPath(html.AD("M10 7v6")),
		html.SvgPath(html.AD("M14 6h2a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-2")),
		html.SvgPath(html.AD("M22 14v-4")),
		html.SvgPath(html.AD("M6 18H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
