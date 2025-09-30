package lucide

import (
	html "github.com/plainkit/html"
)

// Skull creates a Skull Lucide icon.
func Skull(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-skull", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m12.5 17-.5-1-.5 1h1z")),
		html.SvgPath(html.AD("M15 22a1 1 0 0 0 1-1v-1a2 2 0 0 0 1.56-3.25 8 8 0 1 0-11.12 0A2 2 0 0 0 8 20v1a1 1 0 0 0 1 1z")),
		html.SvgCircle(html.ACx("15"), html.ACy("12"), html.AR("1")),
		html.SvgCircle(html.ACx("9"), html.ACy("12"), html.AR("1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
