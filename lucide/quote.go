package lucide

import (
	html "github.com/plainkit/html"
)

// Quote creates a Quote Lucide icon.
func Quote(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-quote", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M16 3a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2 1 1 0 0 1 1 1v1a2 2 0 0 1-2 2 1 1 0 0 0-1 1v2a1 1 0 0 0 1 1 6 6 0 0 0 6-6V5a2 2 0 0 0-2-2z"))),
		html.Child(html.SvgPath(html.AD("M5 3a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2 1 1 0 0 1 1 1v1a2 2 0 0 1-2 2 1 1 0 0 0-1 1v2a1 1 0 0 0 1 1 6 6 0 0 0 6-6V5a2 2 0 0 0-2-2z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
