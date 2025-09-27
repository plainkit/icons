package lucide

import (
	html "github.com/plainkit/html"
)

// SkipForward creates a Skip Forward Lucide icon.
func SkipForward(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-skip-forward", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M21 4v16"))),
		html.Child(html.SvgPath(html.AD("M6.029 4.285A2 2 0 0 0 3 6v12a2 2 0 0 0 3.029 1.715l9.997-5.998a2 2 0 0 0 .003-3.432z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
