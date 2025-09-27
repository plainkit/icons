package lucide

import (
	html "github.com/plainkit/html"
)

// BeanOff creates a Bean Off Lucide icon.
func BeanOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bean-off", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M9 9c-.64.64-1.521.954-2.402 1.165A6 6 0 0 0 8 22a13.96 13.96 0 0 0 9.9-4.1"))),
		html.Child(html.SvgPath(html.AD("M10.75 5.093A6 6 0 0 1 22 8c0 2.411-.61 4.68-1.683 6.66"))),
		html.Child(html.SvgPath(html.AD("M5.341 10.62a4 4 0 0 0 6.487 1.208M10.62 5.341a4.015 4.015 0 0 1 2.039 2.04"))),
		html.Child(html.SvgLine(html.AX1("2"), html.AX2("22"), html.AY1("2"), html.AY2("22"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
