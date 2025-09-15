package lucide

import x "github.com/plainkit/blox"

// BugOff creates a Bug Off Lucide icon.
func BugOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-bug-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 20v-8"))),
		x.Child(x.Path(x.D("M14.12 3.88 16 2"))),
		x.Child(x.Path(x.D("M15 7.13V6a3 3 0 0 0-5.14-2.1L8 2"))),
		x.Child(x.Path(x.D("M18 12.34V11a4 4 0 0 0-4-4h-1.3"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
		x.Child(x.Path(x.D("M21 5a4 4 0 0 1-3.55 3.97"))),
		x.Child(x.Path(x.D("M22 13h-3.34"))),
		x.Child(x.Path(x.D("M3 21a4 4 0 0 1 3.81-4"))),
		x.Child(x.Path(x.D("M6 13H2"))),
		x.Child(x.Path(x.D("M7.7 7.7A4 4 0 0 0 6 11v3a6 6 0 0 0 11.13 3.13"))),
	)
	return x.Svg(svgArgs...)
}
