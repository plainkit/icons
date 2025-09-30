package lucide

import (
	html "github.com/plainkit/html"
)

// Atom creates a Atom Lucide icon.
func Atom(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-atom", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("1")),
		html.SvgPath(html.AD("M20.2 20.2c2.04-2.03.02-7.36-4.5-11.9-4.54-4.52-9.87-6.54-11.9-4.5-2.04 2.03-.02 7.36 4.5 11.9 4.54 4.52 9.87 6.54 11.9 4.5Z")),
		html.SvgPath(html.AD("M15.7 15.7c4.52-4.54 6.54-9.87 4.5-11.9-2.03-2.04-7.36-.02-11.9 4.5-4.52 4.54-6.54 9.87-4.5 11.9 2.03 2.04 7.36.02 11.9-4.5Z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
