package lucide

import x "github.com/plainkit/blox"

// Magnet creates a Magnet Lucide icon.
func Magnet(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-magnet", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m12 15 4 4"))),
		x.Child(x.Path(x.D("M2.352 10.648a1.205 1.205 0 0 0 0 1.704l2.296 2.296a1.205 1.205 0 0 0 1.704 0l6.029-6.029a1 1 0 1 1 3 3l-6.029 6.029a1.205 1.205 0 0 0 0 1.704l2.296 2.296a1.205 1.205 0 0 0 1.704 0l6.365-6.367A1 1 0 0 0 8.716 4.282z"))),
		x.Child(x.Path(x.D("m5 8 4 4"))),
	)
	return x.Svg(svgArgs...)
}
