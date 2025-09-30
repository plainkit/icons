package lucide

import (
	html "github.com/plainkit/html"
)

// TicketX creates a Ticket X Lucide icon.
func TicketX(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-ticket-x", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M2 9a3 3 0 0 1 0 6v2a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-2a3 3 0 0 1 0-6V7a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2Z")),
		html.SvgPath(html.AD("m9.5 14.5 5-5")),
		html.SvgPath(html.AD("m9.5 9.5 5 5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
