package lucide

import (
	html "github.com/plainkit/html"
)

// Type creates a Type Lucide icon.
func Type(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-type", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 4v16"))),
		html.Child(html.SvgPath(html.AD("M4 7V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v2"))),
		html.Child(html.SvgPath(html.AD("M9 20h6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
