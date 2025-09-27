package lucide

import (
	html "github.com/plainkit/html"
)

// NonBinary creates a Non Binary Lucide icon.
func NonBinary(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-non-binary", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 2v10"))),
		html.Child(html.SvgPath(html.AD("m8.5 4 7 4"))),
		html.Child(html.SvgPath(html.AD("m8.5 8 7-4"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("17"), html.AR("5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
