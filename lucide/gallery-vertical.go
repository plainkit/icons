package lucide

import (
	html "github.com/plainkit/html"
)

// GalleryVertical creates a Gallery Vertical Lucide icon.
func GalleryVertical(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-gallery-vertical", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M3 2h18")),
		html.SvgRect(html.AWidth("18"), html.AHeight("12"), html.AX("3"), html.AY("6"), html.ARx("2")),
		html.SvgPath(html.AD("M3 22h18")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
