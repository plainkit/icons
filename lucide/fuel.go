package lucide

import (
	html "github.com/plainkit/html"
)

// Fuel creates a Fuel Lucide icon.
func Fuel(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-fuel", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M14 13h2a2 2 0 0 1 2 2v2a2 2 0 0 0 4 0v-6.998a2 2 0 0 0-.59-1.42L18 5")),
		html.SvgPath(html.AD("M14 21V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v16")),
		html.SvgPath(html.AD("M2 21h13")),
		html.SvgPath(html.AD("M3 9h11")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
