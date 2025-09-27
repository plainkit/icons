package lucide

import (
	html "github.com/plainkit/html"
)

// Merge creates a Merge Lucide icon.
func Merge(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-merge", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m8 6 4-4 4 4"))),
		html.Child(html.SvgPath(html.AD("M12 2v10.3a4 4 0 0 1-1.172 2.872L4 22"))),
		html.Child(html.SvgPath(html.AD("m20 22-5-5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
