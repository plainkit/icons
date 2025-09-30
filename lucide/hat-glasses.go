package lucide

import (
	html "github.com/plainkit/html"
)

// HatGlasses creates a Hat Glasses Lucide icon.
func HatGlasses(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-hat-glasses", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M14 18a2 2 0 0 0-4 0")),
		html.SvgPath(html.AD("m19 11-2.11-6.657a2 2 0 0 0-2.752-1.148l-1.276.61A2 2 0 0 1 12 4H8.5a2 2 0 0 0-1.925 1.456L5 11")),
		html.SvgPath(html.AD("M2 11h20")),
		html.SvgCircle(html.ACx("17"), html.ACy("18"), html.AR("3")),
		html.SvgCircle(html.ACx("7"), html.ACy("18"), html.AR("3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
