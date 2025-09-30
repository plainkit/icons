package lucide

import (
	html "github.com/plainkit/html"
)

// HousePlug creates a House Plug Lucide icon.
func HousePlug(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-house-plug", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 12V8.964")),
		html.SvgPath(html.AD("M14 12V8.964")),
		html.SvgPath(html.AD("M15 12a1 1 0 0 1 1 1v2a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2v-2a1 1 0 0 1 1-1z")),
		html.SvgPath(html.AD("M8.5 21H5a2 2 0 0 1-2-2v-9a2 2 0 0 1 .709-1.528l7-6a2 2 0 0 1 2.582 0l7 6A2 2 0 0 1 21 10v9a2 2 0 0 1-2 2h-5a2 2 0 0 1-2-2v-2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
