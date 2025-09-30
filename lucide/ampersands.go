package lucide

import (
	html "github.com/plainkit/html"
)

// Ampersands creates a Ampersands Lucide icon.
func Ampersands(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-ampersands", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 17c-5-3-7-7-7-9a2 2 0 0 1 4 0c0 2.5-5 2.5-5 6 0 1.7 1.3 3 3 3 2.8 0 5-2.2 5-5")),
		html.SvgPath(html.AD("M22 17c-5-3-7-7-7-9a2 2 0 0 1 4 0c0 2.5-5 2.5-5 6 0 1.7 1.3 3 3 3 2.8 0 5-2.2 5-5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
