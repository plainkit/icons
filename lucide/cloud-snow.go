package lucide

import x "github.com/plainkit/blox"

// CloudSnow creates a Cloud Snow Lucide icon.
func CloudSnow(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-cloud-snow", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242"))),
		x.Child(x.Path(x.D("M8 15h.01"))),
		x.Child(x.Path(x.D("M8 19h.01"))),
		x.Child(x.Path(x.D("M12 17h.01"))),
		x.Child(x.Path(x.D("M12 21h.01"))),
		x.Child(x.Path(x.D("M16 15h.01"))),
		x.Child(x.Path(x.D("M16 19h.01"))),
	)
	return x.Svg(svgArgs...)
}
