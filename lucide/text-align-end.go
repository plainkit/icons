package lucide

import (
	html "github.com/plainkit/html"
)

// TextAlignEnd creates a Text Align End Lucide icon.
func TextAlignEnd(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-text-align-end", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M21 5H3"))),
		html.Child(html.SvgPath(html.AD("M21 12H9"))),
		html.Child(html.SvgPath(html.AD("M21 19H7"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
