package lucide

import (
	html "github.com/plainkit/html"
)

// BookDashed creates a Book Dashed Lucide icon.
func BookDashed(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-book-dashed", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 17h1.5"))),
		html.Child(html.SvgPath(html.AD("M12 22h1.5"))),
		html.Child(html.SvgPath(html.AD("M12 2h1.5"))),
		html.Child(html.SvgPath(html.AD("M17.5 22H19a1 1 0 0 0 1-1"))),
		html.Child(html.SvgPath(html.AD("M17.5 2H19a1 1 0 0 1 1 1v1.5"))),
		html.Child(html.SvgPath(html.AD("M20 14v3h-2.5"))),
		html.Child(html.SvgPath(html.AD("M20 8.5V10"))),
		html.Child(html.SvgPath(html.AD("M4 10V8.5"))),
		html.Child(html.SvgPath(html.AD("M4 19.5V14"))),
		html.Child(html.SvgPath(html.AD("M4 4.5A2.5 2.5 0 0 1 6.5 2H8"))),
		html.Child(html.SvgPath(html.AD("M8 22H6.5a1 1 0 0 1 0-5H8"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
