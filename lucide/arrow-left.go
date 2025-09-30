package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowLeft creates a Arrow Left Lucide icon.
func ArrowLeft(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-left", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m12 19-7-7 7-7")),
		html.SvgPath(html.AD("M19 12H5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
