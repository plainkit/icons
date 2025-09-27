package lucide

import (
	html "github.com/plainkit/html"
)

// CircleDotDashed creates a Circle Dot Dashed Lucide icon.
func CircleDotDashed(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-dot-dashed", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10.1 2.18a9.93 9.93 0 0 1 3.8 0"))),
		html.Child(html.SvgPath(html.AD("M17.6 3.71a9.95 9.95 0 0 1 2.69 2.7"))),
		html.Child(html.SvgPath(html.AD("M21.82 10.1a9.93 9.93 0 0 1 0 3.8"))),
		html.Child(html.SvgPath(html.AD("M20.29 17.6a9.95 9.95 0 0 1-2.7 2.69"))),
		html.Child(html.SvgPath(html.AD("M13.9 21.82a9.94 9.94 0 0 1-3.8 0"))),
		html.Child(html.SvgPath(html.AD("M6.4 20.29a9.95 9.95 0 0 1-2.69-2.7"))),
		html.Child(html.SvgPath(html.AD("M2.18 13.9a9.93 9.93 0 0 1 0-3.8"))),
		html.Child(html.SvgPath(html.AD("M3.71 6.4a9.95 9.95 0 0 1 2.7-2.69"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
