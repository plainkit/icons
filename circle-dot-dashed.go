package lucide

import x "github.com/bloxui/blox"

// CircleDotDashed creates a Circle Dot Dashed Lucide icon.
func CircleDotDashed(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-circle-dot-dashed", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10.1 2.18a9.93 9.93 0 0 1 3.8 0"))),
		x.Child(x.Path(x.D("M17.6 3.71a9.95 9.95 0 0 1 2.69 2.7"))),
		x.Child(x.Path(x.D("M21.82 10.1a9.93 9.93 0 0 1 0 3.8"))),
		x.Child(x.Path(x.D("M20.29 17.6a9.95 9.95 0 0 1-2.7 2.69"))),
		x.Child(x.Path(x.D("M13.9 21.82a9.94 9.94 0 0 1-3.8 0"))),
		x.Child(x.Path(x.D("M6.4 20.29a9.95 9.95 0 0 1-2.69-2.7"))),
		x.Child(x.Path(x.D("M2.18 13.9a9.93 9.93 0 0 1 0-3.8"))),
		x.Child(x.Path(x.D("M3.71 6.4a9.95 9.95 0 0 1 2.7-2.69"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("1"))),
	)
	return x.Svg(svgArgs...)
}
