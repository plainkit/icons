package lucide

import (
	html "github.com/plainkit/html"
)

// Regex creates a Regex Lucide icon.
func Regex(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-regex", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M17 3v10"))),
		html.Child(html.SvgPath(html.AD("m12.67 5.5 8.66 5"))),
		html.Child(html.SvgPath(html.AD("m12.67 10.5 8.66-5"))),
		html.Child(html.SvgPath(html.AD("M9 17a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2v-2z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
