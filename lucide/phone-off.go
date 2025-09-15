package lucide

import x "github.com/plainkit/html"

// PhoneOff creates a Phone Off Lucide icon.
func PhoneOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-phone-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10.1 13.9a14 14 0 0 0 3.732 2.668 1 1 0 0 0 1.213-.303l.355-.465A2 2 0 0 1 17 15h3a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2 18 18 0 0 1-12.728-5.272"))),
		x.Child(x.Path(x.D("M22 2 2 22"))),
		x.Child(x.Path(x.D("M4.76 13.582A18 18 0 0 1 2 4a2 2 0 0 1 2-2h3a2 2 0 0 1 2 2v3a2 2 0 0 1-.8 1.6l-.468.351a1 1 0 0 0-.292 1.233 14 14 0 0 0 .244.473"))),
	)
	return x.Svg(svgArgs...)
}
