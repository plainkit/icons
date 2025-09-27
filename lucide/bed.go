package lucide

import (
	html "github.com/plainkit/html"
)

// Bed creates a Bed Lucide icon.
func Bed(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bed", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M2 4v16"))),
		html.Child(html.SvgPath(html.AD("M2 8h18a2 2 0 0 1 2 2v10"))),
		html.Child(html.SvgPath(html.AD("M2 17h20"))),
		html.Child(html.SvgPath(html.AD("M6 8v9"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
