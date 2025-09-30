package lucide

import (
	html "github.com/plainkit/html"
)

// Milestone creates a Milestone Lucide icon.
func Milestone(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-milestone", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 13v8")),
		html.SvgPath(html.AD("M12 3v3")),
		html.SvgPath(html.AD("M4 6a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h13a2 2 0 0 0 1.152-.365l3.424-2.317a1 1 0 0 0 0-1.635l-3.424-2.318A2 2 0 0 0 17 6z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
