package lucide

import (
	html "github.com/plainkit/html"
)

// Scaling creates a Scaling Lucide icon.
func Scaling(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-scaling", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"))),
		html.Child(html.SvgPath(html.AD("M14 15H9v-5"))),
		html.Child(html.SvgPath(html.AD("M16 3h5v5"))),
		html.Child(html.SvgPath(html.AD("M21 3 9 15"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
