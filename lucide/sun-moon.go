package lucide

import (
	html "github.com/plainkit/html"
)

// SunMoon creates a Sun Moon Lucide icon.
func SunMoon(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-sun-moon", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 2v2"))),
		html.Child(html.SvgPath(html.AD("M14.837 16.385a6 6 0 1 1-7.223-7.222c.624-.147.97.66.715 1.248a4 4 0 0 0 5.26 5.259c.589-.255 1.396.09 1.248.715"))),
		html.Child(html.SvgPath(html.AD("M16 12a4 4 0 0 0-4-4"))),
		html.Child(html.SvgPath(html.AD("m19 5-1.256 1.256"))),
		html.Child(html.SvgPath(html.AD("M20 12h2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
