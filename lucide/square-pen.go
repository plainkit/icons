package lucide

import (
	html "github.com/plainkit/html"
)

// SquarePen creates a Square Pen Lucide icon.
func SquarePen(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-pen", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7")),
		html.SvgPath(html.AD("M18.375 2.625a1 1 0 0 1 3 3l-9.013 9.014a2 2 0 0 1-.853.505l-2.873.84a.5.5 0 0 1-.62-.62l.84-2.873a2 2 0 0 1 .506-.852z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
