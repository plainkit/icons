package lucide

import (
	html "github.com/plainkit/html"
)

// Egg creates a Egg Lucide icon.
func Egg(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-egg", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 2C8 2 4 8 4 14a8 8 0 0 0 16 0c0-6-4-12-8-12")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
