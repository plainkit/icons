package lucide

import (
	html "github.com/plainkit/html"
)

// TextWrap creates a Text Wrap Lucide icon.
func TextWrap(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-text-wrap", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m16 16-3 3 3 3")),
		html.SvgPath(html.AD("M3 12h14.5a1 1 0 0 1 0 7H13")),
		html.SvgPath(html.AD("M3 19h6")),
		html.SvgPath(html.AD("M3 5h18")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
