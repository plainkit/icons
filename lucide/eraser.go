package lucide

import (
	html "github.com/plainkit/html"
)

// Eraser creates a Eraser Lucide icon.
func Eraser(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-eraser", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M21 21H8a2 2 0 0 1-1.42-.587l-3.994-3.999a2 2 0 0 1 0-2.828l10-10a2 2 0 0 1 2.829 0l5.999 6a2 2 0 0 1 0 2.828L12.834 21")),
		html.SvgPath(html.AD("m5.082 11.09 8.828 8.828")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
