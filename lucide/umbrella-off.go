package lucide

import (
	html "github.com/plainkit/html"
)

// UmbrellaOff creates a Umbrella Off Lucide icon.
func UmbrellaOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-umbrella-off", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 13v7a2 2 0 0 0 4 0")),
		html.SvgPath(html.AD("M12 2v2")),
		html.SvgPath(html.AD("M18.656 13h2.336a1 1 0 0 0 .97-1.274 10.284 10.284 0 0 0-12.07-7.51")),
		html.SvgPath(html.AD("m2 2 20 20")),
		html.SvgPath(html.AD("M5.961 5.957a10.28 10.28 0 0 0-3.922 5.769A1 1 0 0 0 3 13h10")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
