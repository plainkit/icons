package lucide

import x "github.com/plainkit/html"

// Cog creates a Cog Lucide icon.
func Cog(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-cog", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M11 10.27 7 3.34"))),
		x.Child(x.Path(x.D("m11 13.73-4 6.93"))),
		x.Child(x.Path(x.D("M12 22v-2"))),
		x.Child(x.Path(x.D("M12 2v2"))),
		x.Child(x.Path(x.D("M14 12h8"))),
		x.Child(x.Path(x.D("m17 20.66-1-1.73"))),
		x.Child(x.Path(x.D("m17 3.34-1 1.73"))),
		x.Child(x.Path(x.D("M2 12h2"))),
		x.Child(x.Path(x.D("m20.66 17-1.73-1"))),
		x.Child(x.Path(x.D("m20.66 7-1.73 1"))),
		x.Child(x.Path(x.D("m3.34 17 1.73-1"))),
		x.Child(x.Path(x.D("m3.34 7 1.73 1"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("2"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("8"))),
	)
	return x.Svg(svgArgs...)
}
