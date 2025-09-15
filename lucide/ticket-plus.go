package lucide

import x "github.com/plainkit/html"

// TicketPlus creates a Ticket Plus Lucide icon.
func TicketPlus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-ticket-plus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 9a3 3 0 0 1 0 6v2a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-2a3 3 0 0 1 0-6V7a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2Z"))),
		x.Child(x.Path(x.D("M9 12h6"))),
		x.Child(x.Path(x.D("M12 9v6"))),
	)
	return x.Svg(svgArgs...)
}
