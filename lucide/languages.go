package lucide

import (
	html "github.com/plainkit/html"
)

// Languages creates a Languages Lucide icon.
func Languages(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-languages", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m5 8 6 6"))),
		html.Child(html.SvgPath(html.AD("m4 14 6-6 2-3"))),
		html.Child(html.SvgPath(html.AD("M2 5h12"))),
		html.Child(html.SvgPath(html.AD("M7 2h1"))),
		html.Child(html.SvgPath(html.AD("m22 22-5-10-5 10"))),
		html.Child(html.SvgPath(html.AD("M14 18h6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
