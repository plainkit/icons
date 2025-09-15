package lucide

import x "github.com/plainkit/blox"

// FlaskRound creates a Flask Round Lucide icon.
func FlaskRound(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-flask-round", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 2v6.292a7 7 0 1 0 4 0V2"))),
		x.Child(x.Path(x.D("M5 15h14"))),
		x.Child(x.Path(x.D("M8.5 2h7"))),
	)
	return x.Svg(svgArgs...)
}
