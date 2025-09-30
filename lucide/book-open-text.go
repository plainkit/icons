package lucide

import (
	html "github.com/plainkit/html"
)

// BookOpenText creates a Book Open Text Lucide icon.
func BookOpenText(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-book-open-text", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 7v14")),
		html.SvgPath(html.AD("M16 12h2")),
		html.SvgPath(html.AD("M16 8h2")),
		html.SvgPath(html.AD("M3 18a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h5a4 4 0 0 1 4 4 4 4 0 0 1 4-4h5a1 1 0 0 1 1 1v13a1 1 0 0 1-1 1h-6a3 3 0 0 0-3 3 3 3 0 0 0-3-3z")),
		html.SvgPath(html.AD("M6 12h2")),
		html.SvgPath(html.AD("M6 8h2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
