package lucide

import x "github.com/plainkit/html"

// TicketsPlane creates a Tickets Plane Lucide icon.
func TicketsPlane(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-tickets-plane", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10.5 17h1.227a2 2 0 0 0 1.345-.52L18 12"))),
		x.Child(x.Path(x.D("m12 13.5 3.75.5"))),
		x.Child(x.Path(x.D("m4.5 8 10.58-5.06a1 1 0 0 1 1.342.488L18.5 8"))),
		x.Child(x.Path(x.D("M6 10V8"))),
		x.Child(x.Path(x.D("M6 14v1"))),
		x.Child(x.Path(x.D("M6 19v2"))),
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("13"), x.X("2"), x.Y("8"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
