package lucide

import x "github.com/plainkit/blox"

// Bug creates a Bug Lucide icon.
func Bug(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-bug", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 20v-9"))),
		x.Child(x.Path(x.D("M14 7a4 4 0 0 1 4 4v3a6 6 0 0 1-12 0v-3a4 4 0 0 1 4-4z"))),
		x.Child(x.Path(x.D("M14.12 3.88 16 2"))),
		x.Child(x.Path(x.D("M21 21a4 4 0 0 0-3.81-4"))),
		x.Child(x.Path(x.D("M21 5a4 4 0 0 1-3.55 3.97"))),
		x.Child(x.Path(x.D("M22 13h-4"))),
		x.Child(x.Path(x.D("M3 21a4 4 0 0 1 3.81-4"))),
		x.Child(x.Path(x.D("M3 5a4 4 0 0 0 3.55 3.97"))),
		x.Child(x.Path(x.D("M6 13H2"))),
		x.Child(x.Path(x.D("m8 2 1.88 1.88"))),
		x.Child(x.Path(x.D("M9 7.13V6a3 3 0 1 1 6 0v1.13"))),
	)
	return x.Svg(svgArgs...)
}
