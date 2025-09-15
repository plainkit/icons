package lucide

import x "github.com/plainkit/blox"

// Droplet creates a Droplet Lucide icon.
func Droplet(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-droplet", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 22a7 7 0 0 0 7-7c0-2-1-3.9-3-5.5s-3.5-4-4-6.5c-.5 2.5-2 4.9-4 6.5C6 11.1 5 13 5 15a7 7 0 0 0 7 7z"))),
	)
	return x.Svg(svgArgs...)
}
