package lucide

import (
	html "github.com/plainkit/html"
)

// TrainTrack creates a Train Track Lucide icon.
func TrainTrack(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-train-track", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M2 17 17 2"))),
		html.Child(html.SvgPath(html.AD("m2 14 8 8"))),
		html.Child(html.SvgPath(html.AD("m5 11 8 8"))),
		html.Child(html.SvgPath(html.AD("m8 8 8 8"))),
		html.Child(html.SvgPath(html.AD("m11 5 8 8"))),
		html.Child(html.SvgPath(html.AD("m14 2 8 8"))),
		html.Child(html.SvgPath(html.AD("M7 22 22 7"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
