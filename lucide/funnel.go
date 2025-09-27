package lucide

import (
	html "github.com/plainkit/html"
)

// Funnel creates a Funnel Lucide icon.
func Funnel(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-funnel", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10 20a1 1 0 0 0 .553.895l2 1A1 1 0 0 0 14 21v-7a2 2 0 0 1 .517-1.341L21.74 4.67A1 1 0 0 0 21 3H3a1 1 0 0 0-.742 1.67l7.225 7.989A2 2 0 0 1 10 14z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
