package lucide

import (
	html "github.com/plainkit/html"
)

// GalleryHorizontal creates a Gallery Horizontal Lucide icon.
func GalleryHorizontal(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-gallery-horizontal", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M2 3v18")),
		html.SvgRect(html.AWidth("12"), html.AHeight("18"), html.AX("6"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("M22 3v18")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
