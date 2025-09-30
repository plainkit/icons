package lucide

import (
	html "github.com/plainkit/html"
)

// ThermometerSun creates a Thermometer Sun Lucide icon.
func ThermometerSun(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-thermometer-sun", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 9a4 4 0 0 0-2 7.5")),
		html.SvgPath(html.AD("M12 3v2")),
		html.SvgPath(html.AD("m6.6 18.4-1.4 1.4")),
		html.SvgPath(html.AD("M20 4v10.54a4 4 0 1 1-4 0V4a2 2 0 0 1 4 0Z")),
		html.SvgPath(html.AD("M4 13H2")),
		html.SvgPath(html.AD("M6.34 7.34 4.93 5.93")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
