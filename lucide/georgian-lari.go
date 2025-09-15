package lucide

import x "github.com/plainkit/blox"

// GeorgianLari creates a Georgian Lari Lucide icon.
func GeorgianLari(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-georgian-lari", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M11.5 21a7.5 7.5 0 1 1 7.35-9"))),
		x.Child(x.Path(x.D("M13 12V3"))),
		x.Child(x.Path(x.D("M4 21h16"))),
		x.Child(x.Path(x.D("M9 12V3"))),
	)
	return x.Svg(svgArgs...)
}
