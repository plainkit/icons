package lucide

import (
	html "github.com/plainkit/html"
)

// FileBadge creates a File Badge Lucide icon.
func FileBadge(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-badge", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 22h6a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v3.072")),
		html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4")),
		html.SvgPath(html.AD("m6.69 16.479 1.29 4.88a.5.5 0 0 1-.698.591l-1.843-.849a1 1 0 0 0-.88.001l-1.846.85a.5.5 0 0 1-.693-.593l1.29-4.88")),
		html.SvgCircle(html.ACx("5"), html.ACy("14"), html.AR("3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
