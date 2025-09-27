package lucide

import (
	html "github.com/plainkit/html"
)

// FileBadge2 creates a File Badge 2 Lucide icon.
func FileBadge2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-badge-2", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m13.69 12.479 1.29 4.88a.5.5 0 0 1-.697.591l-1.844-.849a1 1 0 0 0-.88.001l-1.846.85a.5.5 0 0 1-.693-.593l1.29-4.88"))),
		html.Child(html.SvgPath(html.AD("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7z"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("10"), html.AR("3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
