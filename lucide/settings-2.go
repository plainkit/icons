package lucide

import x "github.com/bloxui/blox"

// Settings2 creates a Settings 2 Lucide icon.
func Settings2(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-settings-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14 17H5"))),
		x.Child(x.Path(x.D("M19 7h-9"))),
		x.Child(x.Circle(x.Cx("17"), x.Cy("17"), x.R("3"))),
		x.Child(x.Circle(x.Cx("7"), x.Cy("7"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
