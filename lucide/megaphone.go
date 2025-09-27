package lucide

import (
	html "github.com/plainkit/html"
)

// Megaphone creates a Megaphone Lucide icon.
func Megaphone(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-megaphone", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M11 6a13 13 0 0 0 8.4-2.8A1 1 0 0 1 21 4v12a1 1 0 0 1-1.6.8A13 13 0 0 0 11 14H5a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2z"))),
		html.Child(html.SvgPath(html.AD("M6 14a12 12 0 0 0 2.4 7.2 2 2 0 0 0 3.2-2.4A8 8 0 0 1 10 14"))),
		html.Child(html.SvgPath(html.AD("M8 6v8"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
