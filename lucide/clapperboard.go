package lucide

import (
	html "github.com/plainkit/html"
)

// Clapperboard creates a Clapperboard Lucide icon.
func Clapperboard(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-clapperboard", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M20.2 6 3 11l-.9-2.4c-.3-1.1.3-2.2 1.3-2.5l13.5-4c1.1-.3 2.2.3 2.5 1.3Z")),
		html.SvgPath(html.AD("m6.2 5.3 3.1 3.9")),
		html.SvgPath(html.AD("m12.4 3.4 3.1 4")),
		html.SvgPath(html.AD("M3 11h18v8a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2Z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
