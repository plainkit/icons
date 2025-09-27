package lucide

import (
	html "github.com/plainkit/html"
)

// AlignVerticalDistributeEnd creates a Align Vertical Distribute End Lucide icon.
func AlignVerticalDistributeEnd(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-align-vertical-distribute-end", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("14"), html.AHeight("6"), html.AX("5"), html.AY("14"), html.ARx("2"))),
		html.Child(html.SvgRect(html.AWidth("10"), html.AHeight("6"), html.AX("7"), html.AY("4"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M2 20h20"))),
		html.Child(html.SvgPath(html.AD("M2 10h20"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
