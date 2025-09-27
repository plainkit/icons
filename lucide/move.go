package lucide

import (
	html "github.com/plainkit/html"
)

// Move creates a Move Lucide icon.
func Move(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-move", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 2v20"))),
		html.Child(html.SvgPath(html.AD("m15 19-3 3-3-3"))),
		html.Child(html.SvgPath(html.AD("m19 9 3 3-3 3"))),
		html.Child(html.SvgPath(html.AD("M2 12h20"))),
		html.Child(html.SvgPath(html.AD("m5 9-3 3 3 3"))),
		html.Child(html.SvgPath(html.AD("m9 5 3-3 3 3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
