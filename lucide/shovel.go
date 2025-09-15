package lucide

import x "github.com/plainkit/blox"

// Shovel creates a Shovel Lucide icon.
func Shovel(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-shovel", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21.56 4.56a1.5 1.5 0 0 1 0 2.122l-.47.47a3 3 0 0 1-4.212-.03 3 3 0 0 1 0-4.243l.44-.44a1.5 1.5 0 0 1 2.121 0z"))),
		x.Child(x.Path(x.D("M3 22a1 1 0 0 1-1-1v-3.586a1 1 0 0 1 .293-.707l3.355-3.355a1.205 1.205 0 0 1 1.704 0l3.296 3.296a1.205 1.205 0 0 1 0 1.704l-3.355 3.355a1 1 0 0 1-.707.293z"))),
		x.Child(x.Path(x.D("m9 15 7.879-7.878"))),
	)
	return x.Svg(svgArgs...)
}
