package lucide

import (
	html "github.com/plainkit/html"
)

// ClockArrowUp creates a Clock Arrow Up Lucide icon.
func ClockArrowUp(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-clock-arrow-up", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 6v6l1.56.78"))),
		html.Child(html.SvgPath(html.AD("M13.227 21.925a10 10 0 1 1 8.767-9.588"))),
		html.Child(html.SvgPath(html.AD("m14 18 4-4 4 4"))),
		html.Child(html.SvgPath(html.AD("M18 22v-8"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
