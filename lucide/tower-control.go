package lucide

import x "github.com/plainkit/blox"

// TowerControl creates a Tower Control Lucide icon.
func TowerControl(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-tower-control", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M18.2 12.27 20 6H4l1.8 6.27a1 1 0 0 0 .95.73h10.5a1 1 0 0 0 .96-.73Z"))),
		x.Child(x.Path(x.D("M8 13v9"))),
		x.Child(x.Path(x.D("M16 22v-9"))),
		x.Child(x.Path(x.D("m9 6 1 7"))),
		x.Child(x.Path(x.D("m15 6-1 7"))),
		x.Child(x.Path(x.D("M12 6V2"))),
		x.Child(x.Path(x.D("M13 2h-2"))),
	)
	return x.Svg(svgArgs...)
}
