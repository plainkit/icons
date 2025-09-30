package lucide

import (
	html "github.com/plainkit/html"
)

// PowerOff creates a Power Off Lucide icon.
func PowerOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-power-off", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M18.36 6.64A9 9 0 0 1 20.77 15")),
		html.SvgPath(html.AD("M6.16 6.16a9 9 0 1 0 12.68 12.68")),
		html.SvgPath(html.AD("M12 2v4")),
		html.SvgPath(html.AD("m2 2 20 20")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
