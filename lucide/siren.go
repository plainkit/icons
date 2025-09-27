package lucide

import (
	html "github.com/plainkit/html"
)

// Siren creates a Siren Lucide icon.
func Siren(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-siren", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M7 18v-6a5 5 0 1 1 10 0v6"))),
		html.Child(html.SvgPath(html.AD("M5 21a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-1a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2z"))),
		html.Child(html.SvgPath(html.AD("M21 12h1"))),
		html.Child(html.SvgPath(html.AD("M18.5 4.5 18 5"))),
		html.Child(html.SvgPath(html.AD("M2 12h1"))),
		html.Child(html.SvgPath(html.AD("M12 2v1"))),
		html.Child(html.SvgPath(html.AD("m4.929 4.929.707.707"))),
		html.Child(html.SvgPath(html.AD("M12 12v6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
