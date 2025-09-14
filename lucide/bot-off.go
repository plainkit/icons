package lucide

import x "github.com/bloxui/blox"

// BotOff creates a Bot Off Lucide icon.
func BotOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-bot-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M13.67 8H18a2 2 0 0 1 2 2v4.33"))),
		x.Child(x.Path(x.D("M2 14h2"))),
		x.Child(x.Path(x.D("M20 14h2"))),
		x.Child(x.Path(x.D("M22 22 2 2"))),
		x.Child(x.Path(x.D("M8 8H6a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h12a2 2 0 0 0 1.414-.586"))),
		x.Child(x.Path(x.D("M9 13v2"))),
		x.Child(x.Path(x.D("M9.67 4H12v2.33"))),
	)
	return x.Svg(svgArgs...)
}
