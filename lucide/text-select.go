package lucide

import x "github.com/bloxui/blox"

// TextSelect creates a Text Select Lucide icon.
func TextSelect(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-text-select", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14 21h1"))),
		x.Child(x.Path(x.D("M14 3h1"))),
		x.Child(x.Path(x.D("M19 3a2 2 0 0 1 2 2"))),
		x.Child(x.Path(x.D("M21 14v1"))),
		x.Child(x.Path(x.D("M21 19a2 2 0 0 1-2 2"))),
		x.Child(x.Path(x.D("M21 9v1"))),
		x.Child(x.Path(x.D("M3 14v1"))),
		x.Child(x.Path(x.D("M3 9v1"))),
		x.Child(x.Path(x.D("M5 21a2 2 0 0 1-2-2"))),
		x.Child(x.Path(x.D("M5 3a2 2 0 0 0-2 2"))),
		x.Child(x.Path(x.D("M7 12h10"))),
		x.Child(x.Path(x.D("M7 16h6"))),
		x.Child(x.Path(x.D("M7 8h8"))),
		x.Child(x.Path(x.D("M9 21h1"))),
		x.Child(x.Path(x.D("M9 3h1"))),
	)
	return x.Svg(svgArgs...)
}
