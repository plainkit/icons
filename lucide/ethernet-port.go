package lucide

import x "github.com/bloxui/blox"

// EthernetPort creates a Ethernet Port Lucide icon.
func EthernetPort(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-ethernet-port", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m15 20 3-3h2a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h2l3 3z"))),
		x.Child(x.Path(x.D("M6 8v1"))),
		x.Child(x.Path(x.D("M10 8v1"))),
		x.Child(x.Path(x.D("M14 8v1"))),
		x.Child(x.Path(x.D("M18 8v1"))),
	)
	return x.Svg(svgArgs...)
}