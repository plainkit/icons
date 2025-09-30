package lucide

import (
	html "github.com/plainkit/html"
)

// Terminal creates a Terminal Lucide icon.
func Terminal(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-terminal", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 19h8")),
		html.SvgPath(html.AD("m4 17 6-6-6-6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
