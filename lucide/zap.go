package lucide

import (
	html "github.com/plainkit/html"
)

// Zap creates a Zap Lucide icon.
func Zap(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-zap", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M4 14a1 1 0 0 1-.78-1.63l9.9-10.2a.5.5 0 0 1 .86.46l-1.92 6.02A1 1 0 0 0 13 10h7a1 1 0 0 1 .78 1.63l-9.9 10.2a.5.5 0 0 1-.86-.46l1.92-6.02A1 1 0 0 0 11 14z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
