package lucide

import (
	html "github.com/plainkit/html"
)

// GitGraph creates a Git Graph Lucide icon.
func GitGraph(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-git-graph", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("5"), html.ACy("6"), html.AR("3")),
		html.SvgPath(html.AD("M5 9v6")),
		html.SvgCircle(html.ACx("5"), html.ACy("18"), html.AR("3")),
		html.SvgPath(html.AD("M12 3v18")),
		html.SvgCircle(html.ACx("19"), html.ACy("6"), html.AR("3")),
		html.SvgPath(html.AD("M16 15.7A9 9 0 0 0 19 9")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
