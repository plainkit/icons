package lucide

import x "github.com/plainkit/blox"

// ThermometerSnowflake creates a Thermometer Snowflake Lucide icon.
func ThermometerSnowflake(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-thermometer-snowflake", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m10 20-1.25-2.5L6 18"))),
		x.Child(x.Path(x.D("M10 4 8.75 6.5 6 6"))),
		x.Child(x.Path(x.D("M10.585 15H10"))),
		x.Child(x.Path(x.D("M2 12h6.5L10 9"))),
		x.Child(x.Path(x.D("M20 14.54a4 4 0 1 1-4 0V4a2 2 0 0 1 4 0z"))),
		x.Child(x.Path(x.D("m4 10 1.5 2L4 14"))),
		x.Child(x.Path(x.D("m7 21 3-6-1.5-3"))),
		x.Child(x.Path(x.D("m7 3 3 6h2"))),
	)
	return x.Svg(svgArgs...)
}
