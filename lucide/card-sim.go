package lucide

import (
	html "github.com/plainkit/html"
)

// CardSim creates a Card Sim Lucide icon.
func CardSim(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-card-sim", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 14v4")),
		html.SvgPath(html.AD("M14.172 2a2 2 0 0 1 1.414.586l3.828 3.828A2 2 0 0 1 20 7.828V20a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2z")),
		html.SvgPath(html.AD("M8 14h8")),
		html.SvgRect(html.AWidth("8"), html.AHeight("8"), html.AX("8"), html.AY("10"), html.ARx("1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
