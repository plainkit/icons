package lucide

import (
	html "github.com/plainkit/html"
)

// FlipHorizontal2 creates a Flip Horizontal 2 Lucide icon.
func FlipHorizontal2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-flip-horizontal-2", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m3 7 5 5-5 5V7"))),
		html.Child(html.SvgPath(html.AD("m21 7-5 5 5 5V7"))),
		html.Child(html.SvgPath(html.AD("M12 20v2"))),
		html.Child(html.SvgPath(html.AD("M12 14v2"))),
		html.Child(html.SvgPath(html.AD("M12 8v2"))),
		html.Child(html.SvgPath(html.AD("M12 2v2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
