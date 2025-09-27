package lucide

import (
	html "github.com/plainkit/html"
)

// BookOpen creates a Book Open Lucide icon.
func BookOpen(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-book-open", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 7v14"))),
		html.Child(html.SvgPath(html.AD("M3 18a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h5a4 4 0 0 1 4 4 4 4 0 0 1 4-4h5a1 1 0 0 1 1 1v13a1 1 0 0 1-1 1h-6a3 3 0 0 0-3 3 3 3 0 0 0-3-3z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
