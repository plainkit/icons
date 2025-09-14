package lucide

import x "github.com/bloxui/blox"

// FlaskConical creates a Flask Conical Lucide icon.
func FlaskConical(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-flask-conical", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14 2v6a2 2 0 0 0 .245.96l5.51 10.08A2 2 0 0 1 18 22H6a2 2 0 0 1-1.755-2.96l5.51-10.08A2 2 0 0 0 10 8V2"))),
		x.Child(x.Path(x.D("M6.453 15h11.094"))),
		x.Child(x.Path(x.D("M8.5 2h7"))),
	)
	return x.Svg(svgArgs...)
}
