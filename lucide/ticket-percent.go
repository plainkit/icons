package lucide

import (
	html "github.com/plainkit/html"
)

// TicketPercent creates a Ticket Percent Lucide icon.
func TicketPercent(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-ticket-percent", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M2 9a3 3 0 1 1 0 6v2a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-2a3 3 0 1 1 0-6V7a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2Z")),
		html.SvgPath(html.AD("M9 9h.01")),
		html.SvgPath(html.AD("m15 9-6 6")),
		html.SvgPath(html.AD("M15 15h.01")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
