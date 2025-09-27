package lucide

import (
	html "github.com/plainkit/html"
)

// View creates a View Lucide icon.
func View(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-view", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M21 17v2a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-2"))),
		html.Child(html.SvgPath(html.AD("M21 7V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v2"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("1"))),
		html.Child(html.SvgPath(html.AD("M18.944 12.33a1 1 0 0 0 0-.66 7.5 7.5 0 0 0-13.888 0 1 1 0 0 0 0 .66 7.5 7.5 0 0 0 13.888 0"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
