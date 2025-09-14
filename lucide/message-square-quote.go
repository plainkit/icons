package lucide

import x "github.com/bloxui/blox"

// MessageSquareQuote creates a Message Square Quote Lucide icon.
func MessageSquareQuote(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-message-square-quote", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14 14a2 2 0 0 0 2-2V8h-2"))),
		x.Child(x.Path(x.D("M22 17a2 2 0 0 1-2 2H6.828a2 2 0 0 0-1.414.586l-2.202 2.202A.71.71 0 0 1 2 21.286V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2z"))),
		x.Child(x.Path(x.D("M8 14a2 2 0 0 0 2-2V8H8"))),
	)
	return x.Svg(svgArgs...)
}
