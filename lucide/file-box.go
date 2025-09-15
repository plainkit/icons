package lucide

import x "github.com/plainkit/blox"

// FileBox creates a File Box Lucide icon.
func FileBox(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-file-box", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14.5 22H18a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v4"))),
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Path(x.D("M3 13.1a2 2 0 0 0-1 1.76v3.24a2 2 0 0 0 .97 1.78L6 21.7a2 2 0 0 0 2.03.01L11 19.9a2 2 0 0 0 1-1.76V14.9a2 2 0 0 0-.97-1.78L8 11.3a2 2 0 0 0-2.03-.01Z"))),
		x.Child(x.Path(x.D("M7 17v5"))),
		x.Child(x.Path(x.D("M11.7 14.2 7 17l-4.7-2.8"))),
	)
	return x.Svg(svgArgs...)
}
