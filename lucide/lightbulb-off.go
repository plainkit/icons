package lucide

import (
	html "github.com/plainkit/html"
)

// LightbulbOff creates a Lightbulb Off Lucide icon.
func LightbulbOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-lightbulb-off", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M16.8 11.2c.8-.9 1.2-2 1.2-3.2a6 6 0 0 0-9.3-5"))),
		html.Child(html.SvgPath(html.AD("m2 2 20 20"))),
		html.Child(html.SvgPath(html.AD("M6.3 6.3a4.67 4.67 0 0 0 1.2 5.2c.7.7 1.3 1.5 1.5 2.5"))),
		html.Child(html.SvgPath(html.AD("M9 18h6"))),
		html.Child(html.SvgPath(html.AD("M10 22h4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
