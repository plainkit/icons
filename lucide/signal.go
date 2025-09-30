package lucide

import (
	html "github.com/plainkit/html"
)

// Signal creates a Signal Lucide icon.
func Signal(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-signal", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M2 20h.01")),
		html.SvgPath(html.AD("M7 20v-4")),
		html.SvgPath(html.AD("M12 20v-8")),
		html.SvgPath(html.AD("M17 20V8")),
		html.SvgPath(html.AD("M22 4v16")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
