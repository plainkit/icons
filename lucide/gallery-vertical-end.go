package lucide

import (
	html "github.com/plainkit/html"
)

// GalleryVerticalEnd creates a Gallery Vertical End Lucide icon.
func GalleryVerticalEnd(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-gallery-vertical-end", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M7 2h10"))),
		html.Child(html.SvgPath(html.AD("M5 6h14"))),
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("12"), html.AX("3"), html.AY("10"), html.ARx("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
