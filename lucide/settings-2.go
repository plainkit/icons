package lucide

import (
	html "github.com/plainkit/html"
)

// Settings2 creates a Settings 2 Lucide icon.
func Settings2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-settings-2", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M14 17H5")),
		html.SvgPath(html.AD("M19 7h-9")),
		html.SvgCircle(html.ACx("17"), html.ACy("17"), html.AR("3")),
		html.SvgCircle(html.ACx("7"), html.ACy("7"), html.AR("3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
