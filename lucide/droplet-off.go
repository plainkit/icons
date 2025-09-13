package lucide

import x "github.com/bloxui/blox"

// DropletOff creates a Droplet Off Lucide icon.
func DropletOff(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-droplet-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M18.715 13.186C18.29 11.858 17.384 10.607 16 9.5c-2-1.6-3.5-4-4-6.5a10.7 10.7 0 0 1-.884 2.586"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
		x.Child(x.Path(x.D("M8.795 8.797A11 11 0 0 1 8 9.5C6 11.1 5 13 5 15a7 7 0 0 0 13.222 3.208"))),
	)
	return x.Svg(svgArgs...)
}