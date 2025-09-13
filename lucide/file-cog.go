package lucide

import x "github.com/bloxui/blox"

// FileCog creates a File Cog Lucide icon.
func FileCog(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-file-cog", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Path(x.D("m2.305 15.53.923-.382"))),
		x.Child(x.Path(x.D("m3.228 12.852-.924-.383"))),
		x.Child(x.Path(x.D("M4.677 21.5a2 2 0 0 0 1.313.5H18a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v2.5"))),
		x.Child(x.Path(x.D("m4.852 11.228-.383-.923"))),
		x.Child(x.Path(x.D("m4.852 16.772-.383.924"))),
		x.Child(x.Path(x.D("m7.148 11.228.383-.923"))),
		x.Child(x.Path(x.D("m7.53 17.696-.382-.924"))),
		x.Child(x.Path(x.D("m8.772 12.852.923-.383"))),
		x.Child(x.Path(x.D("m8.772 15.148.923.383"))),
		x.Child(x.Circle(x.Cx("6"), x.Cy("14"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
