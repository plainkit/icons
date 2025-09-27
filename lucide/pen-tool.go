package lucide

import (
	html "github.com/plainkit/html"
)

// PenTool creates a Pen Tool Lucide icon.
func PenTool(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-pen-tool", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M15.707 21.293a1 1 0 0 1-1.414 0l-1.586-1.586a1 1 0 0 1 0-1.414l5.586-5.586a1 1 0 0 1 1.414 0l1.586 1.586a1 1 0 0 1 0 1.414z"))),
		html.Child(html.SvgPath(html.AD("m18 13-1.375-6.874a1 1 0 0 0-.746-.776L3.235 2.028a1 1 0 0 0-1.207 1.207L5.35 15.879a1 1 0 0 0 .776.746L13 18"))),
		html.Child(html.SvgPath(html.AD("m2.3 2.3 7.286 7.286"))),
		html.Child(html.SvgCircle(html.ACx("11"), html.ACy("11"), html.AR("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
