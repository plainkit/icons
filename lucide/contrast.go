package lucide

import (
	html "github.com/plainkit/html"
)

// Contrast creates a Contrast Lucide icon.
func Contrast(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-contrast", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10")),
		html.SvgPath(html.AD("M12 18a6 6 0 0 0 0-12v12z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
