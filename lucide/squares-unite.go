package lucide

import (
	html "github.com/plainkit/html"
)

// SquaresUnite creates a Squares Unite Lucide icon.
func SquaresUnite(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-squares-unite", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4 16a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v3a1 1 0 0 0 1 1h3a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H10a2 2 0 0 1-2-2v-3a1 1 0 0 0-1-1z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
