package lucide

import (
	html "github.com/plainkit/html"
)

// TextQuote creates a Text Quote Lucide icon.
func TextQuote(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-text-quote", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M17 5H3"))),
		html.Child(html.SvgPath(html.AD("M21 12H8"))),
		html.Child(html.SvgPath(html.AD("M21 19H8"))),
		html.Child(html.SvgPath(html.AD("M3 12v7"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
