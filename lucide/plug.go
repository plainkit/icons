package lucide

import (
	html "github.com/plainkit/html"
)

// Plug creates a Plug Lucide icon.
func Plug(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-plug", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 22v-5")),
		html.SvgPath(html.AD("M9 8V2")),
		html.SvgPath(html.AD("M15 8V2")),
		html.SvgPath(html.AD("M18 8v5a4 4 0 0 1-4 4h-4a4 4 0 0 1-4-4V8Z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
