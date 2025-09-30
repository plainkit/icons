package lucide

import (
	html "github.com/plainkit/html"
)

// StickyNote creates a Sticky Note Lucide icon.
func StickyNote(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-sticky-note", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M16 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V8Z")),
		html.SvgPath(html.AD("M15 3v4a2 2 0 0 0 2 2h4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
