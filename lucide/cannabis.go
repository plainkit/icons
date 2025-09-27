package lucide

import (
	html "github.com/plainkit/html"
)

// Cannabis creates a Cannabis Lucide icon.
func Cannabis(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-cannabis", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 22v-4"))),
		html.Child(html.SvgPath(html.AD("M7 12c-1.5 0-4.5 1.5-5 3 3.5 1.5 6 1 6 1-1.5 1.5-2 3.5-2 5 2.5 0 4.5-1.5 6-3 1.5 1.5 3.5 3 6 3 0-1.5-.5-3.5-2-5 0 0 2.5.5 6-1-.5-1.5-3.5-3-5-3 1.5-1 4-4 4-6-2.5 0-5.5 1.5-7 3 0-2.5-.5-5-2-7-1.5 2-2 4.5-2 7-1.5-1.5-4.5-3-7-3 0 2 2.5 5 4 6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
