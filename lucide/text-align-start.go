package lucide

import (
	html "github.com/plainkit/html"
)

// TextAlignStart creates a Text Align Start Lucide icon.
func TextAlignStart(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-text-align-start", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M21 5H3")),
		html.SvgPath(html.AD("M15 12H3")),
		html.SvgPath(html.AD("M17 19H3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
