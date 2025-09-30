package lucide

import (
	html "github.com/plainkit/html"
)

// Compass creates a Compass Lucide icon.
func Compass(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-compass", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m16.24 7.76-1.804 5.411a2 2 0 0 1-1.265 1.265L7.76 16.24l1.804-5.411a2 2 0 0 1 1.265-1.265z")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
