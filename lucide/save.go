package lucide

import (
	html "github.com/plainkit/html"
)

// Save creates a Save Lucide icon.
func Save(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-save", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M15.2 3a2 2 0 0 1 1.4.6l3.8 3.8a2 2 0 0 1 .6 1.4V19a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2z"))),
		html.Child(html.SvgPath(html.AD("M17 21v-7a1 1 0 0 0-1-1H8a1 1 0 0 0-1 1v7"))),
		html.Child(html.SvgPath(html.AD("M7 3v4a1 1 0 0 0 1 1h7"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
