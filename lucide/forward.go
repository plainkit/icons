package lucide

import (
	html "github.com/plainkit/html"
)

// Forward creates a Forward Lucide icon.
func Forward(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-forward", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m15 17 5-5-5-5")),
		html.SvgPath(html.AD("M4 18v-2a4 4 0 0 1 4-4h12")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
