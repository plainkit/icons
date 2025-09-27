package lucide

import (
	html "github.com/plainkit/html"
)

// SmartphoneCharging creates a Smartphone Charging Lucide icon.
func SmartphoneCharging(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-smartphone-charging", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("14"), html.AHeight("20"), html.AX("5"), html.AY("2"), html.ARx("2"), html.ARy("2"))),
		html.Child(html.SvgPath(html.AD("M12.667 8 10 12h4l-2.667 4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
