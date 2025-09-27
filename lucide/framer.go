package lucide

import (
	html "github.com/plainkit/html"
)

// Framer creates a Framer Lucide icon.
func Framer(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-framer", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M5 16V9h14V2H5l14 14h-7m-7 0 7 7v-7m-7 0h7"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
