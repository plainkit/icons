package lucide

import (
	html "github.com/plainkit/html"
)

// TriangleDashed creates a Triangle Dashed Lucide icon.
func TriangleDashed(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-triangle-dashed", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10.17 4.193a2 2 0 0 1 3.666.013")),
		html.SvgPath(html.AD("M14 21h2")),
		html.SvgPath(html.AD("m15.874 7.743 1 1.732")),
		html.SvgPath(html.AD("m18.849 12.952 1 1.732")),
		html.SvgPath(html.AD("M21.824 18.18a2 2 0 0 1-1.835 2.824")),
		html.SvgPath(html.AD("M4.024 21a2 2 0 0 1-1.839-2.839")),
		html.SvgPath(html.AD("m5.136 12.952-1 1.732")),
		html.SvgPath(html.AD("M8 21h2")),
		html.SvgPath(html.AD("m8.102 7.743-1 1.732")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
