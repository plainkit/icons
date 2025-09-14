package lucide

import x "github.com/bloxui/blox"

// Smartphone creates a Smartphone Lucide icon.
func Smartphone(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-smartphone", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("14"), x.RectHeight("20"), x.X("5"), x.Y("2"), x.Rx("2"), x.Ry("2"))),
		x.Child(x.Path(x.D("M12 18h.01"))),
	)
	return x.Svg(svgArgs...)
}
