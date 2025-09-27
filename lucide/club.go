package lucide

import (
	html "github.com/plainkit/html"
)

// Club creates a Club Lucide icon.
func Club(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-club", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M17.28 9.05a5.5 5.5 0 1 0-10.56 0A5.5 5.5 0 1 0 12 17.66a5.5 5.5 0 1 0 5.28-8.6Z"))),
		html.Child(html.SvgPath(html.AD("M12 17.66L12 22"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
