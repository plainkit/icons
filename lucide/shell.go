package lucide

import (
	html "github.com/plainkit/html"
)

// Shell creates a Shell Lucide icon.
func Shell(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-shell", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M14 11a2 2 0 1 1-4 0 4 4 0 0 1 8 0 6 6 0 0 1-12 0 8 8 0 0 1 16 0 10 10 0 1 1-20 0 11.93 11.93 0 0 1 2.42-7.22 2 2 0 1 1 3.16 2.44"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
