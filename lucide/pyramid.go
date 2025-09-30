package lucide

import (
	html "github.com/plainkit/html"
)

// Pyramid creates a Pyramid Lucide icon.
func Pyramid(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-pyramid", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M2.5 16.88a1 1 0 0 1-.32-1.43l9-13.02a1 1 0 0 1 1.64 0l9 13.01a1 1 0 0 1-.32 1.44l-8.51 4.86a2 2 0 0 1-1.98 0Z")),
		html.SvgPath(html.AD("M12 2v20")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
