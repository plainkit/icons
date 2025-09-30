package lucide

import (
	html "github.com/plainkit/html"
)

// TriangleAlert creates a Triangle Alert Lucide icon.
func TriangleAlert(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-triangle-alert", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m21.73 18-8-14a2 2 0 0 0-3.48 0l-8 14A2 2 0 0 0 4 21h16a2 2 0 0 0 1.73-3")),
		html.SvgPath(html.AD("M12 9v4")),
		html.SvgPath(html.AD("M12 17h.01")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
