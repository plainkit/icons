package lucide

import (
	html "github.com/plainkit/html"
)

// MarsStroke creates a Mars Stroke Lucide icon.
func MarsStroke(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-mars-stroke", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m14 6 4 4")),
		html.SvgPath(html.AD("M17 3h4v4")),
		html.SvgPath(html.AD("m21 3-7.75 7.75")),
		html.SvgCircle(html.ACx("9"), html.ACy("15"), html.AR("6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
