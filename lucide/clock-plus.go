package lucide

import (
	html "github.com/plainkit/html"
)

// ClockPlus creates a Clock Plus Lucide icon.
func ClockPlus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-clock-plus", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 6v6l3.644 1.822")),
		html.SvgPath(html.AD("M16 19h6")),
		html.SvgPath(html.AD("M19 16v6")),
		html.SvgPath(html.AD("M21.92 13.267a10 10 0 1 0-8.653 8.653")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
