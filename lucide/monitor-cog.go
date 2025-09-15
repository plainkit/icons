package lucide

import x "github.com/plainkit/blox"

// MonitorCog creates a Monitor Cog Lucide icon.
func MonitorCog(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-monitor-cog", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 17v4"))),
		x.Child(x.Path(x.D("m14.305 7.53.923-.382"))),
		x.Child(x.Path(x.D("m15.228 4.852-.923-.383"))),
		x.Child(x.Path(x.D("m16.852 3.228-.383-.924"))),
		x.Child(x.Path(x.D("m16.852 8.772-.383.923"))),
		x.Child(x.Path(x.D("m19.148 3.228.383-.924"))),
		x.Child(x.Path(x.D("m19.53 9.696-.382-.924"))),
		x.Child(x.Path(x.D("m20.772 4.852.924-.383"))),
		x.Child(x.Path(x.D("m20.772 7.148.924.383"))),
		x.Child(x.Path(x.D("M22 13v2a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h7"))),
		x.Child(x.Path(x.D("M8 21h8"))),
		x.Child(x.Circle(x.Cx("18"), x.Cy("6"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
