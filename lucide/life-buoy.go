package lucide

import x "github.com/plainkit/blox"

// LifeBuoy creates a Life Buoy Lucide icon.
func LifeBuoy(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-life-buoy", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Path(x.D("m4.93 4.93 4.24 4.24"))),
		x.Child(x.Path(x.D("m14.83 9.17 4.24-4.24"))),
		x.Child(x.Path(x.D("m14.83 14.83 4.24 4.24"))),
		x.Child(x.Path(x.D("m9.17 14.83-4.24 4.24"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("4"))),
	)
	return x.Svg(svgArgs...)
}
