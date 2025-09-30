package lucide

import (
	html "github.com/plainkit/html"
)

// Dam creates a Dam Lucide icon.
func Dam(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-dam", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M11 11.31c1.17.56 1.54 1.69 3.5 1.69 2.5 0 2.5-2 5-2 1.3 0 1.9.5 2.5 1")),
		html.SvgPath(html.AD("M11.75 18c.35.5 1.45 1 2.75 1 2.5 0 2.5-2 5-2 1.3 0 1.9.5 2.5 1")),
		html.SvgPath(html.AD("M2 10h4")),
		html.SvgPath(html.AD("M2 14h4")),
		html.SvgPath(html.AD("M2 18h4")),
		html.SvgPath(html.AD("M2 6h4")),
		html.SvgPath(html.AD("M7 3a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1L10 4a1 1 0 0 0-1-1z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
