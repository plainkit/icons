package lucide

import (
	html "github.com/plainkit/html"
)

// AlignHorizontalDistributeEnd creates a Align Horizontal Distribute End Lucide icon.
func AlignHorizontalDistributeEnd(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-align-horizontal-distribute-end", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("6"), html.AHeight("14"), html.AX("4"), html.AY("5"), html.ARx("2")),
		html.SvgRect(html.AWidth("6"), html.AHeight("10"), html.AX("14"), html.AY("7"), html.ARx("2")),
		html.SvgPath(html.AD("M10 2v20")),
		html.SvgPath(html.AD("M20 2v20")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
