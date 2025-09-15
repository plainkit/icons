package lucide

import x "github.com/plainkit/html"

// TicketCheck creates a Ticket Check Lucide icon.
func TicketCheck(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-ticket-check", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 9a3 3 0 0 1 0 6v2a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-2a3 3 0 0 1 0-6V7a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2Z"))),
		x.Child(x.Path(x.D("m9 12 2 2 4-4"))),
	)
	return x.Svg(svgArgs...)
}
