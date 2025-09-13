package lucide

import x "github.com/bloxui/blox"

// HdmiPort creates a Hdmi Port Lucide icon.
func HdmiPort(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-hdmi-port", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M22 9a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h1l2 2h12l2-2h1a1 1 0 0 0 1-1Z"))),
		x.Child(x.Path(x.D("M7.5 12h9"))),
	)
	return x.Svg(svgArgs...)
}
