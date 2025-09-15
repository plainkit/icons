package lucide

import x "github.com/plainkit/html"

// TicketPercent creates a Ticket Percent Lucide icon.
func TicketPercent(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-ticket-percent", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 9a3 3 0 1 1 0 6v2a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-2a3 3 0 1 1 0-6V7a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2Z"))),
		x.Child(x.Path(x.D("M9 9h.01"))),
		x.Child(x.Path(x.D("m15 9-6 6"))),
		x.Child(x.Path(x.D("M15 15h.01"))),
	)
	return x.Svg(svgArgs...)
}
