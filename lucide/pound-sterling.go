package lucide

import (
	html "github.com/plainkit/html"
)

// PoundSterling creates a Pound Sterling Lucide icon.
func PoundSterling(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-pound-sterling", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M18 7c0-5.333-8-5.333-8 0"))),
		html.Child(html.SvgPath(html.AD("M10 7v14"))),
		html.Child(html.SvgPath(html.AD("M6 21h12"))),
		html.Child(html.SvgPath(html.AD("M6 13h10"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
