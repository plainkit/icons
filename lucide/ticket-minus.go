package lucide

import x "github.com/bloxui/blox"

// TicketMinus creates a Ticket Minus Lucide icon.
func TicketMinus(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-ticket-minus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 9a3 3 0 0 1 0 6v2a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-2a3 3 0 0 1 0-6V7a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2Z"))),
		x.Child(x.Path(x.D("M9 12h6"))),
	)
	return x.Svg(svgArgs...)
}
