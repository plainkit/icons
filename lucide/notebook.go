package lucide

import x "github.com/plainkit/blox"

// Notebook creates a Notebook Lucide icon.
func Notebook(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-notebook", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 6h4"))),
		x.Child(x.Path(x.D("M2 10h4"))),
		x.Child(x.Path(x.D("M2 14h4"))),
		x.Child(x.Path(x.D("M2 18h4"))),
		x.Child(x.Rect(x.RectWidth("16"), x.RectHeight("20"), x.X("4"), x.Y("2"), x.Rx("2"))),
		x.Child(x.Path(x.D("M16 2v20"))),
	)
	return x.Svg(svgArgs...)
}
