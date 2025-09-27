package lucide

import (
	html "github.com/plainkit/html"
)

// HeadphoneOff creates a Headphone Off Lucide icon.
func HeadphoneOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-headphone-off", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M21 14h-1.343"))),
		html.Child(html.SvgPath(html.AD("M9.128 3.47A9 9 0 0 1 21 12v3.343"))),
		html.Child(html.SvgPath(html.AD("m2 2 20 20"))),
		html.Child(html.SvgPath(html.AD("M20.414 20.414A2 2 0 0 1 19 21h-1a2 2 0 0 1-2-2v-3"))),
		html.Child(html.SvgPath(html.AD("M3 14h3a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-7a9 9 0 0 1 2.636-6.364"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
