package lucide

import x "github.com/bloxui/blox"

// CalendarCog creates a Calendar Cog Lucide icon.
func CalendarCog(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-calendar-cog", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m15.228 16.852-.923-.383"))),
		x.Child(x.Path(x.D("m15.228 19.148-.923.383"))),
		x.Child(x.Path(x.D("M16 2v4"))),
		x.Child(x.Path(x.D("m16.47 14.305.382.923"))),
		x.Child(x.Path(x.D("m16.852 20.772-.383.924"))),
		x.Child(x.Path(x.D("m19.148 15.228.383-.923"))),
		x.Child(x.Path(x.D("m19.53 21.696-.382-.924"))),
		x.Child(x.Path(x.D("m20.772 16.852.924-.383"))),
		x.Child(x.Path(x.D("m20.772 19.148.924.383"))),
		x.Child(x.Path(x.D("M21 10.592V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h6"))),
		x.Child(x.Path(x.D("M3 10h18"))),
		x.Child(x.Path(x.D("M8 2v4"))),
		x.Child(x.Circle(x.Cx("18"), x.Cy("18"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
