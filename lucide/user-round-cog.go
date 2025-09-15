package lucide

import x "github.com/plainkit/html"

// UserRoundCog creates a User Round Cog Lucide icon.
func UserRoundCog(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-user-round-cog", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m14.305 19.53.923-.382"))),
		x.Child(x.Path(x.D("m15.228 16.852-.923-.383"))),
		x.Child(x.Path(x.D("m16.852 15.228-.383-.923"))),
		x.Child(x.Path(x.D("m16.852 20.772-.383.924"))),
		x.Child(x.Path(x.D("m19.148 15.228.383-.923"))),
		x.Child(x.Path(x.D("m19.53 21.696-.382-.924"))),
		x.Child(x.Path(x.D("M2 21a8 8 0 0 1 10.434-7.62"))),
		x.Child(x.Path(x.D("m20.772 16.852.924-.383"))),
		x.Child(x.Path(x.D("m20.772 19.148.924.383"))),
		x.Child(x.Circle(x.Cx("10"), x.Cy("8"), x.R("5"))),
		x.Child(x.Circle(x.Cx("18"), x.Cy("18"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
