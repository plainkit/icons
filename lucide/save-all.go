package lucide

import (
	html "github.com/plainkit/html"
)

// SaveAll creates a Save All Lucide icon.
func SaveAll(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-save-all", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10 2v3a1 1 0 0 0 1 1h5"))),
		html.Child(html.SvgPath(html.AD("M18 18v-6a1 1 0 0 0-1-1h-6a1 1 0 0 0-1 1v6"))),
		html.Child(html.SvgPath(html.AD("M18 22H4a2 2 0 0 1-2-2V6"))),
		html.Child(html.SvgPath(html.AD("M8 18a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9.172a2 2 0 0 1 1.414.586l2.828 2.828A2 2 0 0 1 22 6.828V16a2 2 0 0 1-2.01 2z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
