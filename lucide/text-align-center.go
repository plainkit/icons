package lucide

import (
	html "github.com/plainkit/html"
)

// TextAlignCenter creates a Text Align Center Lucide icon.
func TextAlignCenter(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-text-align-center", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M21 5H3")),
		html.SvgPath(html.AD("M17 12H7")),
		html.SvgPath(html.AD("M19 19H5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
