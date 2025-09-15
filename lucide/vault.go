package lucide

import x "github.com/plainkit/html"

// Vault creates a Vault Lucide icon.
func Vault(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-vault", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Circle(x.Cx("7.5"), x.Cy("7.5"), x.R(".5"), x.Fill("currentColor"))),
		x.Child(x.Path(x.D("m7.9 7.9 2.7 2.7"))),
		x.Child(x.Circle(x.Cx("16.5"), x.Cy("7.5"), x.R(".5"), x.Fill("currentColor"))),
		x.Child(x.Path(x.D("m13.4 10.6 2.7-2.7"))),
		x.Child(x.Circle(x.Cx("7.5"), x.Cy("16.5"), x.R(".5"), x.Fill("currentColor"))),
		x.Child(x.Path(x.D("m7.9 16.1 2.7-2.7"))),
		x.Child(x.Circle(x.Cx("16.5"), x.Cy("16.5"), x.R(".5"), x.Fill("currentColor"))),
		x.Child(x.Path(x.D("m13.4 13.4 2.7 2.7"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
