package lucide

import x "github.com/bloxui/blox"

// RouteOff creates a Route Off Lucide icon.
func RouteOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-route-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("6"), x.Cy("19"), x.R("3"))),
		x.Child(x.Path(x.D("M9 19h8.5c.4 0 .9-.1 1.3-.2"))),
		x.Child(x.Path(x.D("M5.2 5.2A3.5 3.53 0 0 0 6.5 12H12"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
		x.Child(x.Path(x.D("M21 15.3a3.5 3.5 0 0 0-3.3-3.3"))),
		x.Child(x.Path(x.D("M15 5h-4.3"))),
		x.Child(x.Circle(x.Cx("18"), x.Cy("5"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
