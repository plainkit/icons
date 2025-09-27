package lucide

import (
	html "github.com/plainkit/html"
)

// CigaretteOff creates a Cigarette Off Lucide icon.
func CigaretteOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-cigarette-off", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 12H3a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h13"))),
		html.Child(html.SvgPath(html.AD("M18 8c0-2.5-2-2.5-2-5"))),
		html.Child(html.SvgPath(html.AD("m2 2 20 20"))),
		html.Child(html.SvgPath(html.AD("M21 12a1 1 0 0 1 1 1v2a1 1 0 0 1-.5.866"))),
		html.Child(html.SvgPath(html.AD("M22 8c0-2.5-2-2.5-2-5"))),
		html.Child(html.SvgPath(html.AD("M7 12v4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
