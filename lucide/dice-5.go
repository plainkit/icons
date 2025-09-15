package lucide

import x "github.com/plainkit/html"

// Dice5 creates a Dice 5 Lucide icon.
func Dice5(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-dice-5", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"), x.Ry("2"))),
		x.Child(x.Path(x.D("M16 8h.01"))),
		x.Child(x.Path(x.D("M8 8h.01"))),
		x.Child(x.Path(x.D("M8 16h.01"))),
		x.Child(x.Path(x.D("M16 16h.01"))),
		x.Child(x.Path(x.D("M12 12h.01"))),
	)
	return x.Svg(svgArgs...)
}
