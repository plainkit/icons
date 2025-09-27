package lucide

import (
	html "github.com/plainkit/html"
)

// Footprints creates a Footprints Lucide icon.
func Footprints(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-footprints", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M4 16v-2.38C4 11.5 2.97 10.5 3 8c.03-2.72 1.49-6 4.5-6C9.37 2 10 3.8 10 5.5c0 3.11-2 5.66-2 8.68V16a2 2 0 1 1-4 0Z"))),
		html.Child(html.SvgPath(html.AD("M20 20v-2.38c0-2.12 1.03-3.12 1-5.62-.03-2.72-1.49-6-4.5-6C14.63 6 14 7.8 14 9.5c0 3.11 2 5.66 2 8.68V20a2 2 0 1 0 4 0Z"))),
		html.Child(html.SvgPath(html.AD("M16 17h4"))),
		html.Child(html.SvgPath(html.AD("M4 13h4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
