package lucide

import (
	html "github.com/plainkit/html"
)

// ExternalLink creates a External Link Lucide icon.
func ExternalLink(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-external-link", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M15 3h6v6")),
		html.SvgPath(html.AD("M10 14 21 3")),
		html.SvgPath(html.AD("M18 13v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
