package lucide

import (
	html "github.com/plainkit/html"
)

// Layers2 creates a Layers 2 Lucide icon.
func Layers2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-layers-2", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M13 13.74a2 2 0 0 1-2 0L2.5 8.87a1 1 0 0 1 0-1.74L11 2.26a2 2 0 0 1 2 0l8.5 4.87a1 1 0 0 1 0 1.74z"))),
		html.Child(html.SvgPath(html.AD("m20 14.285 1.5.845a1 1 0 0 1 0 1.74L13 21.74a2 2 0 0 1-2 0l-8.5-4.87a1 1 0 0 1 0-1.74l1.5-.845"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
