package lucide

import (
	html "github.com/plainkit/html"
)

// Split creates a Split Lucide icon.
func Split(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-split", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M16 3h5v5")),
		html.SvgPath(html.AD("M8 3H3v5")),
		html.SvgPath(html.AD("M12 22v-8.3a4 4 0 0 0-1.172-2.872L3 3")),
		html.SvgPath(html.AD("m15 9 6-6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
