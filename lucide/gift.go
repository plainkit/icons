package lucide

import x "github.com/plainkit/html"

// Gift creates a Gift Lucide icon.
func Gift(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-gift", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("4"), x.X("3"), x.Y("8"), x.Rx("1"))),
		x.Child(x.Path(x.D("M12 8v13"))),
		x.Child(x.Path(x.D("M19 12v7a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-7"))),
		x.Child(x.Path(x.D("M7.5 8a2.5 2.5 0 0 1 0-5A4.8 8 0 0 1 12 8a4.8 8 0 0 1 4.5-5 2.5 2.5 0 0 1 0 5"))),
	)
	return x.Svg(svgArgs...)
}
