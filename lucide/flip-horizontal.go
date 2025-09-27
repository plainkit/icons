package lucide

import (
	html "github.com/plainkit/html"
)

// FlipHorizontal creates a Flip Horizontal Lucide icon.
func FlipHorizontal(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-flip-horizontal", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M8 3H5a2 2 0 0 0-2 2v14c0 1.1.9 2 2 2h3"))),
		html.Child(html.SvgPath(html.AD("M16 3h3a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-3"))),
		html.Child(html.SvgPath(html.AD("M12 20v2"))),
		html.Child(html.SvgPath(html.AD("M12 14v2"))),
		html.Child(html.SvgPath(html.AD("M12 8v2"))),
		html.Child(html.SvgPath(html.AD("M12 2v2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
