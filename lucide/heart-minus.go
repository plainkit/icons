package lucide

import (
	html "github.com/plainkit/html"
)

// HeartMinus creates a Heart Minus Lucide icon.
func HeartMinus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-heart-minus", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m14.876 18.99-1.368 1.323a2 2 0 0 1-3 .019L5 15c-1.5-1.5-3-3.2-3-5.5a5.5 5.5 0 0 1 9.591-3.676.56.56 0 0 0 .818 0A5.49 5.49 0 0 1 22 9.5a5.2 5.2 0 0 1-.244 1.572"))),
		html.Child(html.SvgPath(html.AD("M15 15h6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
