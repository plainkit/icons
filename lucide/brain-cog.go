package lucide

import x "github.com/plainkit/blox"

// BrainCog creates a Brain Cog Lucide icon.
func BrainCog(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-brain-cog", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m10.852 14.772-.383.923"))),
		x.Child(x.Path(x.D("m10.852 9.228-.383-.923"))),
		x.Child(x.Path(x.D("m13.148 14.772.382.924"))),
		x.Child(x.Path(x.D("m13.531 8.305-.383.923"))),
		x.Child(x.Path(x.D("m14.772 10.852.923-.383"))),
		x.Child(x.Path(x.D("m14.772 13.148.923.383"))),
		x.Child(x.Path(x.D("M17.598 6.5A3 3 0 1 0 12 5a3 3 0 0 0-5.63-1.446 3 3 0 0 0-.368 1.571 4 4 0 0 0-2.525 5.771"))),
		x.Child(x.Path(x.D("M17.998 5.125a4 4 0 0 1 2.525 5.771"))),
		x.Child(x.Path(x.D("M19.505 10.294a4 4 0 0 1-1.5 7.706"))),
		x.Child(x.Path(x.D("M4.032 17.483A4 4 0 0 0 11.464 20c.18-.311.892-.311 1.072 0a4 4 0 0 0 7.432-2.516"))),
		x.Child(x.Path(x.D("M4.5 10.291A4 4 0 0 0 6 18"))),
		x.Child(x.Path(x.D("M6.002 5.125a3 3 0 0 0 .4 1.375"))),
		x.Child(x.Path(x.D("m9.228 10.852-.923-.383"))),
		x.Child(x.Path(x.D("m9.228 13.148-.923.383"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
