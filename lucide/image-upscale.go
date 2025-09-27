package lucide

import (
	html "github.com/plainkit/html"
)

// ImageUpscale creates a Image Upscale Lucide icon.
func ImageUpscale(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-image-upscale", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M16 3h5v5"))),
		html.Child(html.SvgPath(html.AD("M17 21h2a2 2 0 0 0 2-2"))),
		html.Child(html.SvgPath(html.AD("M21 12v3"))),
		html.Child(html.SvgPath(html.AD("m21 3-5 5"))),
		html.Child(html.SvgPath(html.AD("M3 7V5a2 2 0 0 1 2-2"))),
		html.Child(html.SvgPath(html.AD("m5 21 4.144-4.144a1.21 1.21 0 0 1 1.712 0L13 19"))),
		html.Child(html.SvgPath(html.AD("M9 3h3"))),
		html.Child(html.SvgRect(html.AWidth("10"), html.AHeight("10"), html.AX("3"), html.AY("11"), html.ARx("1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
