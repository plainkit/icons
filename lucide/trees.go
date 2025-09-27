package lucide

import (
	html "github.com/plainkit/html"
)

// Trees creates a Trees Lucide icon.
func Trees(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-trees", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10 10v.2A3 3 0 0 1 8.9 16H5a3 3 0 0 1-1-5.8V10a3 3 0 0 1 6 0Z"))),
		html.Child(html.SvgPath(html.AD("M7 16v6"))),
		html.Child(html.SvgPath(html.AD("M13 19v3"))),
		html.Child(html.SvgPath(html.AD("M12 19h8.3a1 1 0 0 0 .7-1.7L18 14h.3a1 1 0 0 0 .7-1.7L16 9h.2a1 1 0 0 0 .8-1.7L13 3l-1.4 1.5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
