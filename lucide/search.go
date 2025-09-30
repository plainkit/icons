package lucide

import (
	html "github.com/plainkit/html"
)

// Search creates a Search Lucide icon.
func Search(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-search", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m21 21-4.34-4.34")),
		html.SvgCircle(html.ACx("11"), html.ACy("11"), html.AR("8")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
