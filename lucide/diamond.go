package lucide

import (
	html "github.com/plainkit/html"
)

// Diamond creates a Diamond Lucide icon.
func Diamond(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-diamond", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M2.7 10.3a2.41 2.41 0 0 0 0 3.41l7.59 7.59a2.41 2.41 0 0 0 3.41 0l7.59-7.59a2.41 2.41 0 0 0 0-3.41l-7.59-7.59a2.41 2.41 0 0 0-3.41 0Z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
