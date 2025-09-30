package lucide

import (
	html "github.com/plainkit/html"
)

// Slash creates a Slash Lucide icon.
func Slash(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-slash", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M22 2 2 22")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
