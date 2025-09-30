package lucide

import (
	html "github.com/plainkit/html"
)

// BookOpenCheck creates a Book Open Check Lucide icon.
func BookOpenCheck(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-book-open-check", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 21V7")),
		html.SvgPath(html.AD("m16 12 2 2 4-4")),
		html.SvgPath(html.AD("M22 6V4a1 1 0 0 0-1-1h-5a4 4 0 0 0-4 4 4 4 0 0 0-4-4H3a1 1 0 0 0-1 1v13a1 1 0 0 0 1 1h6a3 3 0 0 1 3 3 3 3 0 0 1 3-3h6a1 1 0 0 0 1-1v-1.3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
