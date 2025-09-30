package lucide

import (
	html "github.com/plainkit/html"
)

// VenetianMask creates a Venetian Mask Lucide icon.
func VenetianMask(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-venetian-mask", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M18 11c-1.5 0-2.5.5-3 2")),
		html.SvgPath(html.AD("M4 6a2 2 0 0 0-2 2v4a5 5 0 0 0 5 5 8 8 0 0 1 5 2 8 8 0 0 1 5-2 5 5 0 0 0 5-5V8a2 2 0 0 0-2-2h-3a8 8 0 0 0-5 2 8 8 0 0 0-5-2z")),
		html.SvgPath(html.AD("M6 11c1.5 0 2.5.5 3 2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
