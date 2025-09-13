package lucide

import x "github.com/bloxui/blox"

// Ticket creates a Ticket Lucide icon.
func Ticket(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-ticket", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 9a3 3 0 0 1 0 6v2a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-2a3 3 0 0 1 0-6V7a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2Z"))),
		x.Child(x.Path(x.D("M13 5v2"))),
		x.Child(x.Path(x.D("M13 17v2"))),
		x.Child(x.Path(x.D("M13 11v2"))),
	)
	return x.Svg(svgArgs...)
}
