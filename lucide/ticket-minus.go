package lucide

import (
	html "github.com/plainkit/html"
)

// TicketMinus creates a Ticket Minus Lucide icon.
func TicketMinus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-ticket-minus", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M2 9a3 3 0 0 1 0 6v2a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-2a3 3 0 0 1 0-6V7a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2Z")),
		html.SvgPath(html.AD("M9 12h6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
