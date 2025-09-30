package lucide

import (
	html "github.com/plainkit/html"
)

// Eye creates a Eye Lucide icon.
func Eye(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-eye", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M2.062 12.348a1 1 0 0 1 0-.696 10.75 10.75 0 0 1 19.876 0 1 1 0 0 1 0 .696 10.75 10.75 0 0 1-19.876 0")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
