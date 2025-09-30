package lucide

import (
	html "github.com/plainkit/html"
)

// TrainFrontTunnel creates a Train Front Tunnel Lucide icon.
func TrainFrontTunnel(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-train-front-tunnel", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M2 22V12a10 10 0 1 1 20 0v10")),
		html.SvgPath(html.AD("M15 6.8v1.4a3 2.8 0 1 1-6 0V6.8")),
		html.SvgPath(html.AD("M10 15h.01")),
		html.SvgPath(html.AD("M14 15h.01")),
		html.SvgPath(html.AD("M10 19a4 4 0 0 1-4-4v-3a6 6 0 1 1 12 0v3a4 4 0 0 1-4 4Z")),
		html.SvgPath(html.AD("m9 19-2 3")),
		html.SvgPath(html.AD("m15 19 2 3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
