package lucide

import (
	html "github.com/plainkit/html"
)

// ClockAlert creates a Clock Alert Lucide icon.
func ClockAlert(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-clock-alert", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 6v6l4 2"))),
		html.Child(html.SvgPath(html.AD("M20 12v5"))),
		html.Child(html.SvgPath(html.AD("M20 21h.01"))),
		html.Child(html.SvgPath(html.AD("M21.25 8.2A10 10 0 1 0 16 21.16"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
