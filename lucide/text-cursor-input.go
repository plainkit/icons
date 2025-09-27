package lucide

import (
	html "github.com/plainkit/html"
)

// TextCursorInput creates a Text Cursor Input Lucide icon.
func TextCursorInput(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-text-cursor-input", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 20h-1a2 2 0 0 1-2-2 2 2 0 0 1-2 2H6"))),
		html.Child(html.SvgPath(html.AD("M13 8h7a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-7"))),
		html.Child(html.SvgPath(html.AD("M5 16H4a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2h1"))),
		html.Child(html.SvgPath(html.AD("M6 4h1a2 2 0 0 1 2 2 2 2 0 0 1 2-2h1"))),
		html.Child(html.SvgPath(html.AD("M9 6v12"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
