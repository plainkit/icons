package lucide

import (
	html "github.com/plainkit/html"
)

// SquaresExclude creates a Squares Exclude Lucide icon.
func SquaresExclude(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-squares-exclude", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M16 12v2a2 2 0 0 1-2 2H9a1 1 0 0 0-1 1v3a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V10a2 2 0 0 0-2-2h0")),
		html.SvgPath(html.AD("M4 16a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v3a1 1 0 0 1-1 1h-5a2 2 0 0 0-2 2v2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
