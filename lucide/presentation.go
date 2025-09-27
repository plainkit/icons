package lucide

import (
	html "github.com/plainkit/html"
)

// Presentation creates a Presentation Lucide icon.
func Presentation(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-presentation", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M2 3h20"))),
		html.Child(html.SvgPath(html.AD("M21 3v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V3"))),
		html.Child(html.SvgPath(html.AD("m7 21 5-5 5 5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
