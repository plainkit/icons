package lucide

import (
	html "github.com/plainkit/html"
)

// MegaphoneOff creates a Megaphone Off Lucide icon.
func MegaphoneOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-megaphone-off", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M11.636 6A13 13 0 0 0 19.4 3.2 1 1 0 0 1 21 4v11.344")),
		html.SvgPath(html.AD("M14.378 14.357A13 13 0 0 0 11 14H5a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h1")),
		html.SvgPath(html.AD("m2 2 20 20")),
		html.SvgPath(html.AD("M6 14a12 12 0 0 0 2.4 7.2 2 2 0 0 0 3.2-2.4A8 8 0 0 1 10 14")),
		html.SvgPath(html.AD("M8 8v6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
