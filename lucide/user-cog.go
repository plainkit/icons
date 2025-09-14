package lucide

import x "github.com/bloxui/blox"

// UserCog creates a User Cog Lucide icon.
func UserCog(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-user-cog", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 15H6a4 4 0 0 0-4 4v2"))),
		x.Child(x.Path(x.D("m14.305 16.53.923-.382"))),
		x.Child(x.Path(x.D("m15.228 13.852-.923-.383"))),
		x.Child(x.Path(x.D("m16.852 12.228-.383-.923"))),
		x.Child(x.Path(x.D("m16.852 17.772-.383.924"))),
		x.Child(x.Path(x.D("m19.148 12.228.383-.923"))),
		x.Child(x.Path(x.D("m19.53 18.696-.382-.924"))),
		x.Child(x.Path(x.D("m20.772 13.852.924-.383"))),
		x.Child(x.Path(x.D("m20.772 16.148.924.383"))),
		x.Child(x.Circle(x.Cx("18"), x.Cy("15"), x.R("3"))),
		x.Child(x.Circle(x.Cx("9"), x.Cy("7"), x.R("4"))),
	)
	return x.Svg(svgArgs...)
}
