package lucide

import x "github.com/bloxui/blox"

// Ambulance creates a Ambulance Lucide icon.
func Ambulance(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-ambulance", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 10H6"))),
		x.Child(x.Path(x.D("M14 18V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v11a1 1 0 0 0 1 1h2"))),
		x.Child(x.Path(x.D("M19 18h2a1 1 0 0 0 1-1v-3.28a1 1 0 0 0-.684-.948l-1.923-.641a1 1 0 0 1-.578-.502l-1.539-3.076A1 1 0 0 0 16.382 8H14"))),
		x.Child(x.Path(x.D("M8 8v4"))),
		x.Child(x.Path(x.D("M9 18h6"))),
		x.Child(x.Circle(x.Cx("17"), x.Cy("18"), x.R("2"))),
		x.Child(x.Circle(x.Cx("7"), x.Cy("18"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
