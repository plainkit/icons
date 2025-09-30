package lucide

import (
	html "github.com/plainkit/html"
)

// Blocks creates a Blocks Lucide icon.
func Blocks(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-blocks", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 22V7a1 1 0 0 0-1-1H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-5a1 1 0 0 0-1-1H2")),
		html.SvgRect(html.AWidth("8"), html.AHeight("8"), html.AX("14"), html.AY("2"), html.ARx("1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
