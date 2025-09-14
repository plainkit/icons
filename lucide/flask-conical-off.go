package lucide

import x "github.com/bloxui/blox"

// FlaskConicalOff creates a Flask Conical Off Lucide icon.
func FlaskConicalOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-flask-conical-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 2v2.343"))),
		x.Child(x.Path(x.D("M14 2v6.343"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
		x.Child(x.Path(x.D("M20 20a2 2 0 0 1-2 2H6a2 2 0 0 1-1.755-2.96l5.227-9.563"))),
		x.Child(x.Path(x.D("M6.453 15H15"))),
		x.Child(x.Path(x.D("M8.5 2h7"))),
	)
	return x.Svg(svgArgs...)
}
