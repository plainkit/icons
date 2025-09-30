package lucide

import (
	html "github.com/plainkit/html"
)

// Umbrella creates a Umbrella Lucide icon.
func Umbrella(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-umbrella", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 13v7a2 2 0 0 0 4 0")),
		html.SvgPath(html.AD("M12 2v2")),
		html.SvgPath(html.AD("M20.992 13a1 1 0 0 0 .97-1.274 10.284 10.284 0 0 0-19.923 0A1 1 0 0 0 3 13z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
