package lucide

import (
	html "github.com/plainkit/html"
)

// Slice creates a Slice Lucide icon.
func Slice(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-slice", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M11 16.586V19a1 1 0 0 1-1 1H2L18.37 3.63a1 1 0 1 1 3 3l-9.663 9.663a1 1 0 0 1-1.414 0L8 14"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
