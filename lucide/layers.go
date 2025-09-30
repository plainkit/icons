package lucide

import (
	html "github.com/plainkit/html"
)

// Layers creates a Layers Lucide icon.
func Layers(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-layers", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12.83 2.18a2 2 0 0 0-1.66 0L2.6 6.08a1 1 0 0 0 0 1.83l8.58 3.91a2 2 0 0 0 1.66 0l8.58-3.9a1 1 0 0 0 0-1.83z")),
		html.SvgPath(html.AD("M2 12a1 1 0 0 0 .58.91l8.6 3.91a2 2 0 0 0 1.65 0l8.58-3.9A1 1 0 0 0 22 12")),
		html.SvgPath(html.AD("M2 17a1 1 0 0 0 .58.91l8.6 3.91a2 2 0 0 0 1.65 0l8.58-3.9A1 1 0 0 0 22 17")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
