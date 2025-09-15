package lucide

import x "github.com/plainkit/html"

// Umbrella creates a Umbrella Lucide icon.
func Umbrella(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-umbrella", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 13v7a2 2 0 0 0 4 0"))),
		x.Child(x.Path(x.D("M12 2v2"))),
		x.Child(x.Path(x.D("M20.992 13a1 1 0 0 0 .97-1.274 10.284 10.284 0 0 0-19.923 0A1 1 0 0 0 3 13z"))),
	)
	return x.Svg(svgArgs...)
}
