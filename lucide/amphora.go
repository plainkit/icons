package lucide

import (
	html "github.com/plainkit/html"
)

// Amphora creates a Amphora Lucide icon.
func Amphora(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-amphora", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 2v5.632c0 .424-.272.795-.653.982A6 6 0 0 0 6 14c.006 4 3 7 5 8")),
		html.SvgPath(html.AD("M10 5H8a2 2 0 0 0 0 4h.68")),
		html.SvgPath(html.AD("M14 2v5.632c0 .424.272.795.652.982A6 6 0 0 1 18 14c0 4-3 7-5 8")),
		html.SvgPath(html.AD("M14 5h2a2 2 0 0 1 0 4h-.68")),
		html.SvgPath(html.AD("M18 22H6")),
		html.SvgPath(html.AD("M9 2h6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
