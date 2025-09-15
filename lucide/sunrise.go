package lucide

import x "github.com/plainkit/html"

// Sunrise creates a Sunrise Lucide icon.
func Sunrise(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-sunrise", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 2v8"))),
		x.Child(x.Path(x.D("m4.93 10.93 1.41 1.41"))),
		x.Child(x.Path(x.D("M2 18h2"))),
		x.Child(x.Path(x.D("M20 18h2"))),
		x.Child(x.Path(x.D("m19.07 10.93-1.41 1.41"))),
		x.Child(x.Path(x.D("M22 22H2"))),
		x.Child(x.Path(x.D("m8 6 4-4 4 4"))),
		x.Child(x.Path(x.D("M16 18a4 4 0 0 0-8 0"))),
	)
	return x.Svg(svgArgs...)
}
