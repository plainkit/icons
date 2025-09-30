package lucide

import (
	html "github.com/plainkit/html"
)

// Armchair creates a Armchair Lucide icon.
func Armchair(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-armchair", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M19 9V6a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v3")),
		html.SvgPath(html.AD("M3 16a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-5a2 2 0 0 0-4 0v1.5a.5.5 0 0 1-.5.5h-9a.5.5 0 0 1-.5-.5V11a2 2 0 0 0-4 0z")),
		html.SvgPath(html.AD("M5 18v2")),
		html.SvgPath(html.AD("M19 18v2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
