package lucide

import (
	html "github.com/plainkit/html"
)

// ThermometerSnowflake creates a Thermometer Snowflake Lucide icon.
func ThermometerSnowflake(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-thermometer-snowflake", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m10 20-1.25-2.5L6 18")),
		html.SvgPath(html.AD("M10 4 8.75 6.5 6 6")),
		html.SvgPath(html.AD("M10.585 15H10")),
		html.SvgPath(html.AD("M2 12h6.5L10 9")),
		html.SvgPath(html.AD("M20 14.54a4 4 0 1 1-4 0V4a2 2 0 0 1 4 0z")),
		html.SvgPath(html.AD("m4 10 1.5 2L4 14")),
		html.SvgPath(html.AD("m7 21 3-6-1.5-3")),
		html.SvgPath(html.AD("m7 3 3 6h2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
