package lucide

import (
	html "github.com/plainkit/html"
)

// Sticker creates a Sticker Lucide icon.
func Sticker(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-sticker", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M15.5 3H5a2 2 0 0 0-2 2v14c0 1.1.9 2 2 2h14a2 2 0 0 0 2-2V8.5L15.5 3Z")),
		html.SvgPath(html.AD("M14 3v4a2 2 0 0 0 2 2h4")),
		html.SvgPath(html.AD("M8 13h.01")),
		html.SvgPath(html.AD("M16 13h.01")),
		html.SvgPath(html.AD("M10 16s.8 1 2 1c1.3 0 2-1 2-1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
