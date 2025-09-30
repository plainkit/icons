package lucide

import (
	html "github.com/plainkit/html"
)

// Receipt creates a Receipt Lucide icon.
func Receipt(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-receipt", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4 2v20l2-1 2 1 2-1 2 1 2-1 2 1 2-1 2 1V2l-2 1-2-1-2 1-2-1-2 1-2-1-2 1Z")),
		html.SvgPath(html.AD("M16 8h-6a2 2 0 1 0 0 4h4a2 2 0 1 1 0 4H8")),
		html.SvgPath(html.AD("M12 17.5v-11")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
