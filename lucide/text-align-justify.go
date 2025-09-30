package lucide

import (
	html "github.com/plainkit/html"
)

// TextAlignJustify creates a Text Align Justify Lucide icon.
func TextAlignJustify(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-text-align-justify", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M3 5h18")),
		html.SvgPath(html.AD("M3 12h18")),
		html.SvgPath(html.AD("M3 19h18")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
