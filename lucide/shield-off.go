package lucide

import (
	html "github.com/plainkit/html"
)

// ShieldOff creates a Shield Off Lucide icon.
func ShieldOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-shield-off", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m2 2 20 20"))),
		html.Child(html.SvgPath(html.AD("M5 5a1 1 0 0 0-1 1v7c0 5 3.5 7.5 7.67 8.94a1 1 0 0 0 .67.01c2.35-.82 4.48-1.97 5.9-3.71"))),
		html.Child(html.SvgPath(html.AD("M9.309 3.652A12.252 12.252 0 0 0 11.24 2.28a1.17 1.17 0 0 1 1.52 0C14.51 3.81 17 5 19 5a1 1 0 0 1 1 1v7a9.784 9.784 0 0 1-.08 1.264"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
