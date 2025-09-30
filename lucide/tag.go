package lucide

import (
	html "github.com/plainkit/html"
)

// Tag creates a Tag Lucide icon.
func Tag(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-tag", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12.586 2.586A2 2 0 0 0 11.172 2H4a2 2 0 0 0-2 2v7.172a2 2 0 0 0 .586 1.414l8.704 8.704a2.426 2.426 0 0 0 3.42 0l6.58-6.58a2.426 2.426 0 0 0 0-3.42z")),
		html.SvgCircle(html.ACx("7.5"), html.ACy("7.5"), html.AR(".5"), html.AFill("currentColor")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
