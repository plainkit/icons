package lucide

import (
	html "github.com/plainkit/html"
)

// Library creates a Library Lucide icon.
func Library(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-library", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m16 6 4 14")),
		html.SvgPath(html.AD("M12 6v14")),
		html.SvgPath(html.AD("M8 8v12")),
		html.SvgPath(html.AD("M4 4v16")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
