package lucide

import (
	html "github.com/plainkit/html"
)

// Shrub creates a Shrub Lucide icon.
func Shrub(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-shrub", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 22v-5.172a2 2 0 0 0-.586-1.414L9.5 13.5")),
		html.SvgPath(html.AD("M14.5 14.5 12 17")),
		html.SvgPath(html.AD("M17 8.8A6 6 0 0 1 13.8 20H10A6.5 6.5 0 0 1 7 8a5 5 0 0 1 10 0z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
