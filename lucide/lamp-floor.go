package lucide

import (
	html "github.com/plainkit/html"
)

// LampFloor creates a Lamp Floor Lucide icon.
func LampFloor(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-lamp-floor", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 10v12")),
		html.SvgPath(html.AD("M17.929 7.629A1 1 0 0 1 17 9H7a1 1 0 0 1-.928-1.371l2-5A1 1 0 0 1 9 2h6a1 1 0 0 1 .928.629z")),
		html.SvgPath(html.AD("M9 22h6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
