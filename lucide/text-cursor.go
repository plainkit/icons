package lucide

import (
	html "github.com/plainkit/html"
)

// TextCursor creates a Text Cursor Lucide icon.
func TextCursor(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-text-cursor", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M17 22h-1a4 4 0 0 1-4-4V6a4 4 0 0 1 4-4h1")),
		html.SvgPath(html.AD("M7 22h1a4 4 0 0 0 4-4v-1")),
		html.SvgPath(html.AD("M7 2h1a4 4 0 0 1 4 4v1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
