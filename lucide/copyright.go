package lucide

import (
	html "github.com/plainkit/html"
)

// Copyright creates a Copyright Lucide icon.
func Copyright(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-copyright", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
		html.Child(html.SvgPath(html.AD("M14.83 14.83a4 4 0 1 1 0-5.66"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
