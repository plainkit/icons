package lucide

import (
	html "github.com/plainkit/html"
)

// SmartphoneNfc creates a Smartphone Nfc Lucide icon.
func SmartphoneNfc(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-smartphone-nfc", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("7"), html.AHeight("12"), html.AX("2"), html.AY("6"), html.ARx("1"))),
		html.Child(html.SvgPath(html.AD("M13 8.32a7.43 7.43 0 0 1 0 7.36"))),
		html.Child(html.SvgPath(html.AD("M16.46 6.21a11.76 11.76 0 0 1 0 11.58"))),
		html.Child(html.SvgPath(html.AD("M19.91 4.1a15.91 15.91 0 0 1 .01 15.8"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
