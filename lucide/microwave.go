package lucide

import x "github.com/plainkit/blox"

// Microwave creates a Microwave Lucide icon.
func Microwave(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-microwave", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("15"), x.X("2"), x.Y("4"), x.Rx("2"))),
		x.Child(x.Rect(x.RectWidth("8"), x.RectHeight("7"), x.X("6"), x.Y("8"), x.Rx("1"))),
		x.Child(x.Path(x.D("M18 8v7"))),
		x.Child(x.Path(x.D("M6 19v2"))),
		x.Child(x.Path(x.D("M18 19v2"))),
	)
	return x.Svg(svgArgs...)
}
