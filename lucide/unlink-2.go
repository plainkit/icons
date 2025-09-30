package lucide

import (
	html "github.com/plainkit/html"
)

// Unlink2 creates a Unlink 2 Lucide icon.
func Unlink2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-unlink-2", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M15 7h2a5 5 0 0 1 0 10h-2m-6 0H7A5 5 0 0 1 7 7h2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
