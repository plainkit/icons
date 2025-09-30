package lucide

import (
	html "github.com/plainkit/html"
)

// Delete creates a Delete Lucide icon.
func Delete(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-delete", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 5a2 2 0 0 0-1.344.519l-6.328 5.74a1 1 0 0 0 0 1.481l6.328 5.741A2 2 0 0 0 10 19h10a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2z")),
		html.SvgPath(html.AD("m12 9 6 6")),
		html.SvgPath(html.AD("m18 9-6 6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
