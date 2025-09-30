package lucide

import (
	html "github.com/plainkit/html"
)

// FastForward creates a Fast Forward Lucide icon.
func FastForward(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-fast-forward", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 6a2 2 0 0 1 3.414-1.414l6 6a2 2 0 0 1 0 2.828l-6 6A2 2 0 0 1 12 18z")),
		html.SvgPath(html.AD("M2 6a2 2 0 0 1 3.414-1.414l6 6a2 2 0 0 1 0 2.828l-6 6A2 2 0 0 1 2 18z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
