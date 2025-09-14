package lucide

import x "github.com/bloxui/blox"

// Bitcoin creates a Bitcoin Lucide icon.
func Bitcoin(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-bitcoin", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M11.767 19.089c4.924.868 6.14-6.025 1.216-6.894m-1.216 6.894L5.86 18.047m5.908 1.042-.347 1.97m1.563-8.864c4.924.869 6.14-6.025 1.215-6.893m-1.215 6.893-3.94-.694m5.155-6.2L8.29 4.26m5.908 1.042.348-1.97M7.48 20.364l3.126-17.727"))),
	)
	return x.Svg(svgArgs...)
}
