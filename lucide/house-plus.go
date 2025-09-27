package lucide

import (
	html "github.com/plainkit/html"
)

// HousePlus creates a House Plus Lucide icon.
func HousePlus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-house-plus", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12.35 21H5a2 2 0 0 1-2-2v-9a2 2 0 0 1 .71-1.53l7-6a2 2 0 0 1 2.58 0l7 6A2 2 0 0 1 21 10v2.35"))),
		html.Child(html.SvgPath(html.AD("M14.8 12.4A1 1 0 0 0 14 12h-4a1 1 0 0 0-1 1v8"))),
		html.Child(html.SvgPath(html.AD("M15 18h6"))),
		html.Child(html.SvgPath(html.AD("M18 15v6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
