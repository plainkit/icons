package lucide

import (
	html "github.com/plainkit/html"
)

// PhilippinePeso creates a Philippine Peso Lucide icon.
func PhilippinePeso(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-philippine-peso", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M20 11H4")),
		html.SvgPath(html.AD("M20 7H4")),
		html.SvgPath(html.AD("M7 21V4a1 1 0 0 1 1-1h4a1 1 0 0 1 0 12H7")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
