package lucide

import x "github.com/plainkit/blox"

// Images creates a Images Lucide icon.
func Images(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-images", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m22 11-1.296-1.296a2.4 2.4 0 0 0-3.408 0L11 16"))),
		x.Child(x.Path(x.D("M4 8a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2"))),
		x.Child(x.Circle(x.Cx("13"), x.Cy("7"), x.R("1"), x.Fill("currentColor"))),
		x.Child(x.Rect(x.RectWidth("14"), x.RectHeight("14"), x.X("8"), x.Y("2"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
