package lucide

import (
	html "github.com/plainkit/html"
)

// Hamburger creates a Hamburger Lucide icon.
func Hamburger(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-hamburger", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 16H4a2 2 0 1 1 0-4h16a2 2 0 1 1 0 4h-4.25")),
		html.SvgPath(html.AD("M5 12a2 2 0 0 1-2-2 9 7 0 0 1 18 0 2 2 0 0 1-2 2")),
		html.SvgPath(html.AD("M5 16a2 2 0 0 0-2 2 3 3 0 0 0 3 3h12a3 3 0 0 0 3-3 2 2 0 0 0-2-2q0 0 0 0")),
		html.SvgPath(html.AD("m6.67 12 6.13 4.6a2 2 0 0 0 2.8-.4l3.15-4.2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
