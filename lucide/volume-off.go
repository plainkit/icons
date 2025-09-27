package lucide

import (
	html "github.com/plainkit/html"
)

// VolumeOff creates a Volume Off Lucide icon.
func VolumeOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-volume-off", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M16 9a5 5 0 0 1 .95 2.293"))),
		html.Child(html.SvgPath(html.AD("M19.364 5.636a9 9 0 0 1 1.889 9.96"))),
		html.Child(html.SvgPath(html.AD("m2 2 20 20"))),
		html.Child(html.SvgPath(html.AD("m7 7-.587.587A1.4 1.4 0 0 1 5.416 8H3a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h2.416a1.4 1.4 0 0 1 .997.413l3.383 3.384A.705.705 0 0 0 11 19.298V11"))),
		html.Child(html.SvgPath(html.AD("M9.828 4.172A.686.686 0 0 1 11 4.657v.686"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
