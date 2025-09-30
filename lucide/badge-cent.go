package lucide

import (
	html "github.com/plainkit/html"
)

// BadgeCent creates a Badge Cent Lucide icon.
func BadgeCent(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-badge-cent", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z")),
		html.SvgPath(html.AD("M12 7v10")),
		html.SvgPath(html.AD("M15.4 10a4 4 0 1 0 0 4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
