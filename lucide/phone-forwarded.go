package lucide

import x "github.com/plainkit/html"

// PhoneForwarded creates a Phone Forwarded Lucide icon.
func PhoneForwarded(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-phone-forwarded", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14 6h8"))),
		x.Child(x.Path(x.D("m18 2 4 4-4 4"))),
		x.Child(x.Path(x.D("M13.832 16.568a1 1 0 0 0 1.213-.303l.355-.465A2 2 0 0 1 17 15h3a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2A18 18 0 0 1 2 4a2 2 0 0 1 2-2h3a2 2 0 0 1 2 2v3a2 2 0 0 1-.8 1.6l-.468.351a1 1 0 0 0-.292 1.233 14 14 0 0 0 6.392 6.384"))),
	)
	return x.Svg(svgArgs...)
}
