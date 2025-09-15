package lucide

import x "github.com/plainkit/blox"

// Dice1 creates a Dice 1 Lucide icon.
func Dice1(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-dice-1", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"), x.Ry("2"))),
		x.Child(x.Path(x.D("M12 12h.01"))),
	)
	return x.Svg(svgArgs...)
}
