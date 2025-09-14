package lucide

import x "github.com/bloxui/blox"

// Bot creates a Bot Lucide icon.
func Bot(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-bot", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 8V4H8"))),
		x.Child(x.Rect(x.RectWidth("16"), x.RectHeight("12"), x.X("4"), x.Y("8"), x.Rx("2"))),
		x.Child(x.Path(x.D("M2 14h2"))),
		x.Child(x.Path(x.D("M20 14h2"))),
		x.Child(x.Path(x.D("M15 13v2"))),
		x.Child(x.Path(x.D("M9 13v2"))),
	)
	return x.Svg(svgArgs...)
}
