package lucide

import (
	html "github.com/plainkit/html"
)

// Power creates a Power Lucide icon.
func Power(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-power", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 2v10"))),
		html.Child(html.SvgPath(html.AD("M18.4 6.6a9 9 0 1 1-12.77.04"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
