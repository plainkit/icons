package lucide

import (
	html "github.com/plainkit/html"
)

// TicketCheck creates a Ticket Check Lucide icon.
func TicketCheck(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-ticket-check", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M2 9a3 3 0 0 1 0 6v2a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-2a3 3 0 0 1 0-6V7a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2Z"))),
		html.Child(html.SvgPath(html.AD("m9 12 2 2 4-4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
