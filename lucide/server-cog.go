package lucide

import x "github.com/bloxui/blox"

// ServerCog creates a Server Cog Lucide icon.
func ServerCog(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-server-cog", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m10.852 14.772-.383.923"))),
		x.Child(x.Path(x.D("M13.148 14.772a3 3 0 1 0-2.296-5.544l-.383-.923"))),
		x.Child(x.Path(x.D("m13.148 9.228.383-.923"))),
		x.Child(x.Path(x.D("m13.53 15.696-.382-.924a3 3 0 1 1-2.296-5.544"))),
		x.Child(x.Path(x.D("m14.772 10.852.923-.383"))),
		x.Child(x.Path(x.D("m14.772 13.148.923.383"))),
		x.Child(x.Path(x.D("M4.5 10H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-.5"))),
		x.Child(x.Path(x.D("M4.5 14H4a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2h-.5"))),
		x.Child(x.Path(x.D("M6 18h.01"))),
		x.Child(x.Path(x.D("M6 6h.01"))),
		x.Child(x.Path(x.D("m9.228 10.852-.923-.383"))),
		x.Child(x.Path(x.D("m9.228 13.148-.923.383"))),
	)
	return x.Svg(svgArgs...)
}
