package lucide

import x "github.com/bloxui/blox"

// Sun creates a Sun Lucide icon.
func Sun(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-sun", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("4"))),
		x.Child(x.Path(x.D("M12 2v2"))),
		x.Child(x.Path(x.D("M12 20v2"))),
		x.Child(x.Path(x.D("m4.93 4.93 1.41 1.41"))),
		x.Child(x.Path(x.D("m17.66 17.66 1.41 1.41"))),
		x.Child(x.Path(x.D("M2 12h2"))),
		x.Child(x.Path(x.D("M20 12h2"))),
		x.Child(x.Path(x.D("m6.34 17.66-1.41 1.41"))),
		x.Child(x.Path(x.D("m19.07 4.93-1.41 1.41"))),
	)
	return x.Svg(svgArgs...)
}
