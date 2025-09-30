package lucide

import (
	html "github.com/plainkit/html"
)

// TrainFront creates a Train Front Lucide icon.
func TrainFront(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-train-front", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M8 3.1V7a4 4 0 0 0 8 0V3.1")),
		html.SvgPath(html.AD("m9 15-1-1")),
		html.SvgPath(html.AD("m15 15 1-1")),
		html.SvgPath(html.AD("M9 19c-2.8 0-5-2.2-5-5v-4a8 8 0 0 1 16 0v4c0 2.8-2.2 5-5 5Z")),
		html.SvgPath(html.AD("m8 19-2 3")),
		html.SvgPath(html.AD("m16 19 2 3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
