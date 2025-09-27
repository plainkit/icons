package lucide

import (
	html "github.com/plainkit/html"
)

// Piano creates a Piano Lucide icon.
func Piano(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-piano", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M18.5 8c-1.4 0-2.6-.8-3.2-2A6.87 6.87 0 0 0 2 9v11a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-8.5C22 9.6 20.4 8 18.5 8"))),
		html.Child(html.SvgPath(html.AD("M2 14h20"))),
		html.Child(html.SvgPath(html.AD("M6 14v4"))),
		html.Child(html.SvgPath(html.AD("M10 14v4"))),
		html.Child(html.SvgPath(html.AD("M14 14v4"))),
		html.Child(html.SvgPath(html.AD("M18 14v4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
