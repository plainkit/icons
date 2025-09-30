package lucide

import (
	html "github.com/plainkit/html"
)

// SwissFranc creates a Swiss Franc Lucide icon.
func SwissFranc(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-swiss-franc", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 21V3h8")),
		html.SvgPath(html.AD("M6 16h9")),
		html.SvgPath(html.AD("M10 9.5h7")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
