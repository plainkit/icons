package lucide

import (
	html "github.com/plainkit/html"
)

// Anchor creates a Anchor Lucide icon.
func Anchor(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-anchor", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 22V8"))),
		html.Child(html.SvgPath(html.AD("M5 12H2a10 10 0 0 0 20 0h-3"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("5"), html.AR("3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
