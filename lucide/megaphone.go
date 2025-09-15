package lucide

import x "github.com/plainkit/blox"

// Megaphone creates a Megaphone Lucide icon.
func Megaphone(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-megaphone", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M11 6a13 13 0 0 0 8.4-2.8A1 1 0 0 1 21 4v12a1 1 0 0 1-1.6.8A13 13 0 0 0 11 14H5a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2z"))),
		x.Child(x.Path(x.D("M6 14a12 12 0 0 0 2.4 7.2 2 2 0 0 0 3.2-2.4A8 8 0 0 1 10 14"))),
		x.Child(x.Path(x.D("M8 6v8"))),
	)
	return x.Svg(svgArgs...)
}
