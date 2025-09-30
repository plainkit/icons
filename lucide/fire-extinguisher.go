package lucide

import (
	html "github.com/plainkit/html"
)

// FireExtinguisher creates a Fire Extinguisher Lucide icon.
func FireExtinguisher(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-fire-extinguisher", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M15 6.5V3a1 1 0 0 0-1-1h-2a1 1 0 0 0-1 1v3.5")),
		html.SvgPath(html.AD("M9 18h8")),
		html.SvgPath(html.AD("M18 3h-3")),
		html.SvgPath(html.AD("M11 3a6 6 0 0 0-6 6v11")),
		html.SvgPath(html.AD("M5 13h4")),
		html.SvgPath(html.AD("M17 10a4 4 0 0 0-8 0v10a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2Z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
