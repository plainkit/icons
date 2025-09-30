package lucide

import (
	html "github.com/plainkit/html"
)

// Warehouse creates a Warehouse Lucide icon.
func Warehouse(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-warehouse", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M18 21V10a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1v11")),
		html.SvgPath(html.AD("M22 19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V8a2 2 0 0 1 1.132-1.803l7.95-3.974a2 2 0 0 1 1.837 0l7.948 3.974A2 2 0 0 1 22 8z")),
		html.SvgPath(html.AD("M6 13h12")),
		html.SvgPath(html.AD("M6 17h12")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
