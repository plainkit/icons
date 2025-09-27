package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowUp creates a Arrow Up Lucide icon.
func ArrowUp(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-up", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m5 12 7-7 7 7"))),
		html.Child(html.SvgPath(html.AD("M12 19V5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
