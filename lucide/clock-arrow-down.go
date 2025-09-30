package lucide

import (
	html "github.com/plainkit/html"
)

// ClockArrowDown creates a Clock Arrow Down Lucide icon.
func ClockArrowDown(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-clock-arrow-down", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 6v6l2 1")),
		html.SvgPath(html.AD("M12.337 21.994a10 10 0 1 1 9.588-8.767")),
		html.SvgPath(html.AD("m14 18 4 4 4-4")),
		html.SvgPath(html.AD("M18 14v8")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
