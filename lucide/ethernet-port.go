package lucide

import (
	html "github.com/plainkit/html"
)

// EthernetPort creates a Ethernet Port Lucide icon.
func EthernetPort(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-ethernet-port", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m15 20 3-3h2a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h2l3 3z")),
		html.SvgPath(html.AD("M6 8v1")),
		html.SvgPath(html.AD("M10 8v1")),
		html.SvgPath(html.AD("M14 8v1")),
		html.SvgPath(html.AD("M18 8v1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
