package lucide

import (
	html "github.com/plainkit/html"
)

// SunMedium creates a Sun Medium Lucide icon.
func SunMedium(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-sun-medium", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("4")),
		html.SvgPath(html.AD("M12 3v1")),
		html.SvgPath(html.AD("M12 20v1")),
		html.SvgPath(html.AD("M3 12h1")),
		html.SvgPath(html.AD("M20 12h1")),
		html.SvgPath(html.AD("m18.364 5.636-.707.707")),
		html.SvgPath(html.AD("m6.343 17.657-.707.707")),
		html.SvgPath(html.AD("m5.636 5.636.707.707")),
		html.SvgPath(html.AD("m17.657 17.657.707.707")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
