package lucide

import (
	html "github.com/plainkit/html"
)

// Ban creates a Ban Lucide icon.
func Ban(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-ban", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M4.929 4.929 19.07 19.071"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
