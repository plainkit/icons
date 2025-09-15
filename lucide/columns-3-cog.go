package lucide

import x "github.com/plainkit/blox"

// Columns3Cog creates a Columns 3 Cog Lucide icon.
func Columns3Cog(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-columns-3-cog", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10.5 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v5.5"))),
		x.Child(x.Path(x.D("m14.3 19.6 1-.4"))),
		x.Child(x.Path(x.D("M15 3v7.5"))),
		x.Child(x.Path(x.D("m15.2 16.9-.9-.3"))),
		x.Child(x.Path(x.D("m16.6 21.7.3-.9"))),
		x.Child(x.Path(x.D("m16.8 15.3-.4-1"))),
		x.Child(x.Path(x.D("m19.1 15.2.3-.9"))),
		x.Child(x.Path(x.D("m19.6 21.7-.4-1"))),
		x.Child(x.Path(x.D("m20.7 16.8 1-.4"))),
		x.Child(x.Path(x.D("m21.7 19.4-.9-.3"))),
		x.Child(x.Path(x.D("M9 3v18"))),
		x.Child(x.Circle(x.Cx("18"), x.Cy("18"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
