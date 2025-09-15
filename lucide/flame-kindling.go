package lucide

import x "github.com/plainkit/blox"

// FlameKindling creates a Flame Kindling Lucide icon.
func FlameKindling(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-flame-kindling", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 2c1 3 2.5 3.5 3.5 4.5A5 5 0 0 1 17 10a5 5 0 1 1-10 0c0-.3 0-.6.1-.9a2 2 0 1 0 3.3-2C8 4.5 11 2 12 2Z"))),
		x.Child(x.Path(x.D("m5 22 14-4"))),
		x.Child(x.Path(x.D("m5 18 14 4"))),
	)
	return x.Svg(svgArgs...)
}
