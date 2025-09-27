package lucide

import (
	html "github.com/plainkit/html"
)

// Shuffle creates a Shuffle Lucide icon.
func Shuffle(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-shuffle", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m18 14 4 4-4 4"))),
		html.Child(html.SvgPath(html.AD("m18 2 4 4-4 4"))),
		html.Child(html.SvgPath(html.AD("M2 18h1.973a4 4 0 0 0 3.3-1.7l5.454-8.6a4 4 0 0 1 3.3-1.7H22"))),
		html.Child(html.SvgPath(html.AD("M2 6h1.972a4 4 0 0 1 3.6 2.2"))),
		html.Child(html.SvgPath(html.AD("M22 18h-6.041a4 4 0 0 1-3.3-1.8l-.359-.45"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
