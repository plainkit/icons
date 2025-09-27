package lucide

import (
	html "github.com/plainkit/html"
)

// Sofa creates a Sofa Lucide icon.
func Sofa(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-sofa", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M20 9V6a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v3"))),
		html.Child(html.SvgPath(html.AD("M2 16a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-5a2 2 0 0 0-4 0v1.5a.5.5 0 0 1-.5.5h-11a.5.5 0 0 1-.5-.5V11a2 2 0 0 0-4 0z"))),
		html.Child(html.SvgPath(html.AD("M4 18v2"))),
		html.Child(html.SvgPath(html.AD("M20 18v2"))),
		html.Child(html.SvgPath(html.AD("M12 4v9"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
