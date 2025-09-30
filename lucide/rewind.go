package lucide

import (
	html "github.com/plainkit/html"
)

// Rewind creates a Rewind Lucide icon.
func Rewind(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-rewind", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 6a2 2 0 0 0-3.414-1.414l-6 6a2 2 0 0 0 0 2.828l6 6A2 2 0 0 0 12 18z")),
		html.SvgPath(html.AD("M22 6a2 2 0 0 0-3.414-1.414l-6 6a2 2 0 0 0 0 2.828l6 6A2 2 0 0 0 22 18z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
