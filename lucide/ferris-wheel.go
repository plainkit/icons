package lucide

import x "github.com/plainkit/html"

// FerrisWheel creates a Ferris Wheel Lucide icon.
func FerrisWheel(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-ferris-wheel", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("2"))),
		x.Child(x.Path(x.D("M12 2v4"))),
		x.Child(x.Path(x.D("m6.8 15-3.5 2"))),
		x.Child(x.Path(x.D("m20.7 7-3.5 2"))),
		x.Child(x.Path(x.D("M6.8 9 3.3 7"))),
		x.Child(x.Path(x.D("m20.7 17-3.5-2"))),
		x.Child(x.Path(x.D("m9 22 3-8 3 8"))),
		x.Child(x.Path(x.D("M8 22h8"))),
		x.Child(x.Path(x.D("M18 18.7a9 9 0 1 0-12 0"))),
	)
	return x.Svg(svgArgs...)
}
