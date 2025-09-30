package lucide

import (
	html "github.com/plainkit/html"
)

// AudioLines creates a Audio Lines Lucide icon.
func AudioLines(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-audio-lines", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M2 10v3")),
		html.SvgPath(html.AD("M6 6v11")),
		html.SvgPath(html.AD("M10 3v18")),
		html.SvgPath(html.AD("M14 8v7")),
		html.SvgPath(html.AD("M18 5v13")),
		html.SvgPath(html.AD("M22 10v3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
