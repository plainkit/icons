package lucide

import x "github.com/plainkit/html"

// Satellite creates a Satellite Lucide icon.
func Satellite(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-satellite", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m13.5 6.5-3.148-3.148a1.205 1.205 0 0 0-1.704 0L6.352 5.648a1.205 1.205 0 0 0 0 1.704L9.5 10.5"))),
		x.Child(x.Path(x.D("M16.5 7.5 19 5"))),
		x.Child(x.Path(x.D("m17.5 10.5 3.148 3.148a1.205 1.205 0 0 1 0 1.704l-2.296 2.296a1.205 1.205 0 0 1-1.704 0L13.5 14.5"))),
		x.Child(x.Path(x.D("M9 21a6 6 0 0 0-6-6"))),
		x.Child(x.Path(x.D("M9.352 10.648a1.205 1.205 0 0 0 0 1.704l2.296 2.296a1.205 1.205 0 0 0 1.704 0l4.296-4.296a1.205 1.205 0 0 0 0-1.704l-2.296-2.296a1.205 1.205 0 0 0-1.704 0z"))),
	)
	return x.Svg(svgArgs...)
}
