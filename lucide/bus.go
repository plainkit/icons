package lucide

import x "github.com/bloxui/blox"

// Bus creates a Bus Lucide icon.
func Bus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-bus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M8 6v6"))),
		x.Child(x.Path(x.D("M15 6v6"))),
		x.Child(x.Path(x.D("M2 12h19.6"))),
		x.Child(x.Path(x.D("M18 18h3s.5-1.7.8-2.8c.1-.4.2-.8.2-1.2 0-.4-.1-.8-.2-1.2l-1.4-5C20.1 6.8 19.1 6 18 6H4a2 2 0 0 0-2 2v10h3"))),
		x.Child(x.Circle(x.Cx("7"), x.Cy("18"), x.R("2"))),
		x.Child(x.Path(x.D("M9 18h5"))),
		x.Child(x.Circle(x.Cx("16"), x.Cy("18"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
