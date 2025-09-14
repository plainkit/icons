package lucide

import x "github.com/bloxui/blox"

// Snowflake creates a Snowflake Lucide icon.
func Snowflake(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-snowflake", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m10 20-1.25-2.5L6 18"))),
		x.Child(x.Path(x.D("M10 4 8.75 6.5 6 6"))),
		x.Child(x.Path(x.D("m14 20 1.25-2.5L18 18"))),
		x.Child(x.Path(x.D("m14 4 1.25 2.5L18 6"))),
		x.Child(x.Path(x.D("m17 21-3-6h-4"))),
		x.Child(x.Path(x.D("m17 3-3 6 1.5 3"))),
		x.Child(x.Path(x.D("M2 12h6.5L10 9"))),
		x.Child(x.Path(x.D("m20 10-1.5 2 1.5 2"))),
		x.Child(x.Path(x.D("M22 12h-6.5L14 15"))),
		x.Child(x.Path(x.D("m4 10 1.5 2L4 14"))),
		x.Child(x.Path(x.D("m7 21 3-6-1.5-3"))),
		x.Child(x.Path(x.D("m7 3 3 6h4"))),
	)
	return x.Svg(svgArgs...)
}
