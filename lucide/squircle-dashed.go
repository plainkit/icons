package lucide

import (
	html "github.com/plainkit/html"
)

// SquircleDashed creates a Squircle Dashed Lucide icon.
func SquircleDashed(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-squircle-dashed", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M13.77 3.043a34 34 0 0 0-3.54 0"))),
		html.Child(html.SvgPath(html.AD("M13.771 20.956a33 33 0 0 1-3.541.001"))),
		html.Child(html.SvgPath(html.AD("M20.18 17.74c-.51 1.15-1.29 1.93-2.439 2.44"))),
		html.Child(html.SvgPath(html.AD("M20.18 6.259c-.51-1.148-1.291-1.929-2.44-2.438"))),
		html.Child(html.SvgPath(html.AD("M20.957 10.23a33 33 0 0 1 0 3.54"))),
		html.Child(html.SvgPath(html.AD("M3.043 10.23a34 34 0 0 0 .001 3.541"))),
		html.Child(html.SvgPath(html.AD("M6.26 20.179c-1.15-.508-1.93-1.29-2.44-2.438"))),
		html.Child(html.SvgPath(html.AD("M6.26 3.82c-1.149.51-1.93 1.291-2.44 2.44"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
