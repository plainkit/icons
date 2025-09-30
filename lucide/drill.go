package lucide

import (
	html "github.com/plainkit/html"
)

// Drill creates a Drill Lucide icon.
func Drill(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-drill", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 18a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1H5a3 3 0 0 1-3-3 1 1 0 0 1 1-1z")),
		html.SvgPath(html.AD("M13 10H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1l-.81 3.242a1 1 0 0 1-.97.758H8")),
		html.SvgPath(html.AD("M14 4h3a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-3")),
		html.SvgPath(html.AD("M18 6h4")),
		html.SvgPath(html.AD("m5 10-2 8")),
		html.SvgPath(html.AD("m7 18 2-8")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
