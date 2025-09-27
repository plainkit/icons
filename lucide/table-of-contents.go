package lucide

import (
	html "github.com/plainkit/html"
)

// TableOfContents creates a Table Of Contents Lucide icon.
func TableOfContents(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-table-of-contents", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M16 5H3"))),
		html.Child(html.SvgPath(html.AD("M16 12H3"))),
		html.Child(html.SvgPath(html.AD("M16 19H3"))),
		html.Child(html.SvgPath(html.AD("M21 5h.01"))),
		html.Child(html.SvgPath(html.AD("M21 12h.01"))),
		html.Child(html.SvgPath(html.AD("M21 19h.01"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
