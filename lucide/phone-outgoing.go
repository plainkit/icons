package lucide

import x "github.com/plainkit/blox"

// PhoneOutgoing creates a Phone Outgoing Lucide icon.
func PhoneOutgoing(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-phone-outgoing", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m16 8 6-6"))),
		x.Child(x.Path(x.D("M22 8V2h-6"))),
		x.Child(x.Path(x.D("M13.832 16.568a1 1 0 0 0 1.213-.303l.355-.465A2 2 0 0 1 17 15h3a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2A18 18 0 0 1 2 4a2 2 0 0 1 2-2h3a2 2 0 0 1 2 2v3a2 2 0 0 1-.8 1.6l-.468.351a1 1 0 0 0-.292 1.233 14 14 0 0 0 6.392 6.384"))),
	)
	return x.Svg(svgArgs...)
}
