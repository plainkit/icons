package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowDownToDot creates a Arrow Down To Dot Lucide icon.
func ArrowDownToDot(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-down-to-dot", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 2v14")),
		html.SvgPath(html.AD("m19 9-7 7-7-7")),
		html.SvgCircle(html.ACx("12"), html.ACy("21"), html.AR("1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
