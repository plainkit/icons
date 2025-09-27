package lucide

import (
	html "github.com/plainkit/html"
)

// CreativeCommons creates a Creative Commons Lucide icon.
func CreativeCommons(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-creative-commons", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
		html.Child(html.SvgPath(html.AD("M10 9.3a2.8 2.8 0 0 0-3.5 1 3.1 3.1 0 0 0 0 3.4 2.7 2.7 0 0 0 3.5 1"))),
		html.Child(html.SvgPath(html.AD("M17 9.3a2.8 2.8 0 0 0-3.5 1 3.1 3.1 0 0 0 0 3.4 2.7 2.7 0 0 0 3.5 1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
