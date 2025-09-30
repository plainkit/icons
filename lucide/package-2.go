package lucide

import (
	html "github.com/plainkit/html"
)

// Package2 creates a Package 2 Lucide icon.
func Package2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-package-2", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 3v6")),
		html.SvgPath(html.AD("M16.76 3a2 2 0 0 1 1.8 1.1l2.23 4.479a2 2 0 0 1 .21.891V19a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V9.472a2 2 0 0 1 .211-.894L5.45 4.1A2 2 0 0 1 7.24 3z")),
		html.SvgPath(html.AD("M3.054 9.013h17.893")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
