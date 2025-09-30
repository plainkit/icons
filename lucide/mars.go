package lucide

import (
	html "github.com/plainkit/html"
)

// Mars creates a Mars Lucide icon.
func Mars(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-mars", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M16 3h5v5")),
		html.SvgPath(html.AD("m21 3-6.75 6.75")),
		html.SvgCircle(html.ACx("10"), html.ACy("14"), html.AR("6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
