package lucide

import (
	html "github.com/plainkit/html"
)

// Youtube creates a Youtube Lucide icon.
func Youtube(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-youtube", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M2.5 17a24.12 24.12 0 0 1 0-10 2 2 0 0 1 1.4-1.4 49.56 49.56 0 0 1 16.2 0A2 2 0 0 1 21.5 7a24.12 24.12 0 0 1 0 10 2 2 0 0 1-1.4 1.4 49.55 49.55 0 0 1-16.2 0A2 2 0 0 1 2.5 17"))),
		html.Child(html.SvgPath(html.AD("m10 15 5-3-5-3z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
