package lucide

import x "github.com/plainkit/blox"

// Mail creates a Mail Lucide icon.
func Mail(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-mail", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m22 7-8.991 5.727a2 2 0 0 1-2.009 0L2 7"))),
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("16"), x.X("2"), x.Y("4"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
