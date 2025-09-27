package lucide

import (
	html "github.com/plainkit/html"
)

// GalleryHorizontalEnd creates a Gallery Horizontal End Lucide icon.
func GalleryHorizontalEnd(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-gallery-horizontal-end", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M2 7v10"))),
		html.Child(html.SvgPath(html.AD("M6 5v14"))),
		html.Child(html.SvgRect(html.AWidth("12"), html.AHeight("18"), html.AX("10"), html.AY("3"), html.ARx("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
