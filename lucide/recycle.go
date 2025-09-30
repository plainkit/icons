package lucide

import (
	html "github.com/plainkit/html"
)

// Recycle creates a Recycle Lucide icon.
func Recycle(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-recycle", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M7 19H4.815a1.83 1.83 0 0 1-1.57-.881 1.785 1.785 0 0 1-.004-1.784L7.196 9.5")),
		html.SvgPath(html.AD("M11 19h8.203a1.83 1.83 0 0 0 1.556-.89 1.784 1.784 0 0 0 0-1.775l-1.226-2.12")),
		html.SvgPath(html.AD("m14 16-3 3 3 3")),
		html.SvgPath(html.AD("M8.293 13.596 7.196 9.5 3.1 10.598")),
		html.SvgPath(html.AD("m9.344 5.811 1.093-1.892A1.83 1.83 0 0 1 11.985 3a1.784 1.784 0 0 1 1.546.888l3.943 6.843")),
		html.SvgPath(html.AD("m13.378 9.633 4.096 1.098 1.097-4.096")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
