package lucide

import x "github.com/bloxui/blox"

// WifiCog creates a Wifi Cog Lucide icon.
func WifiCog(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-wifi-cog", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m14.305 19.53.923-.382"))),
		x.Child(x.Path(x.D("m15.228 16.852-.923-.383"))),
		x.Child(x.Path(x.D("m16.852 15.228-.383-.923"))),
		x.Child(x.Path(x.D("m16.852 20.772-.383.924"))),
		x.Child(x.Path(x.D("m19.148 15.228.383-.923"))),
		x.Child(x.Path(x.D("m19.53 21.696-.382-.924"))),
		x.Child(x.Path(x.D("M2 7.82a15 15 0 0 1 20 0"))),
		x.Child(x.Path(x.D("m20.772 16.852.924-.383"))),
		x.Child(x.Path(x.D("m20.772 19.148.924.383"))),
		x.Child(x.Path(x.D("M5 11.858a10 10 0 0 1 11.5-1.785"))),
		x.Child(x.Path(x.D("M8.5 15.429a5 5 0 0 1 2.413-1.31"))),
		x.Child(x.Circle(x.Cx("18"), x.Cy("18"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
