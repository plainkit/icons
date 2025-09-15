package lucide

import x "github.com/plainkit/blox"

// ImagePlay creates a Image Play Lucide icon.
func ImagePlay(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-image-play", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 15.003a1 1 0 0 1 1.517-.859l4.997 2.997a1 1 0 0 1 0 1.718l-4.997 2.997a1 1 0 0 1-1.517-.86z"))),
		x.Child(x.Path(x.D("M21 12.17V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h6"))),
		x.Child(x.Path(x.D("m6 21 5-5"))),
		x.Child(x.Circle(x.Cx("9"), x.Cy("9"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
