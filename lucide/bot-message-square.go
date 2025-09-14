package lucide

import x "github.com/bloxui/blox"

// BotMessageSquare creates a Bot Message Square Lucide icon.
func BotMessageSquare(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-bot-message-square", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 6V2H8"))),
		x.Child(x.Path(x.D("M15 11v2"))),
		x.Child(x.Path(x.D("M2 12h2"))),
		x.Child(x.Path(x.D("M20 12h2"))),
		x.Child(x.Path(x.D("M20 16a2 2 0 0 1-2 2H8.828a2 2 0 0 0-1.414.586l-2.202 2.202A.71.71 0 0 1 4 20.286V8a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2z"))),
		x.Child(x.Path(x.D("M9 11v2"))),
	)
	return x.Svg(svgArgs...)
}
