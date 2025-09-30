package lucide

import (
	html "github.com/plainkit/html"
)

// GalleryThumbnails creates a Gallery Thumbnails Lucide icon.
func GalleryThumbnails(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-gallery-thumbnails", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("14"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("M4 21h1")),
		html.SvgPath(html.AD("M9 21h1")),
		html.SvgPath(html.AD("M14 21h1")),
		html.SvgPath(html.AD("M19 21h1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
