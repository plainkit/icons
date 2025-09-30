package lucide

import (
	html "github.com/plainkit/html"
)

// OctagonX creates a Octagon X Lucide icon.
func OctagonX(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-octagon-x", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m15 9-6 6")),
		html.SvgPath(html.AD("M2.586 16.726A2 2 0 0 1 2 15.312V8.688a2 2 0 0 1 .586-1.414l4.688-4.688A2 2 0 0 1 8.688 2h6.624a2 2 0 0 1 1.414.586l4.688 4.688A2 2 0 0 1 22 8.688v6.624a2 2 0 0 1-.586 1.414l-4.688 4.688a2 2 0 0 1-1.414.586H8.688a2 2 0 0 1-1.414-.586z")),
		html.SvgPath(html.AD("m9 9 6 6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
