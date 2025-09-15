package lucide

import x "github.com/plainkit/blox"

// WifiSync creates a Wifi Sync Lucide icon.
func WifiSync(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-wifi-sync", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M11.965 10.105v4L13.5 12.5a5 5 0 0 1 8 1.5"))),
		x.Child(x.Path(x.D("M11.965 14.105h4"))),
		x.Child(x.Path(x.D("M17.965 18.105h4L20.43 19.71a5 5 0 0 1-8-1.5"))),
		x.Child(x.Path(x.D("M2 8.82a15 15 0 0 1 20 0"))),
		x.Child(x.Path(x.D("M21.965 22.105v-4"))),
		x.Child(x.Path(x.D("M5 12.86a10 10 0 0 1 3-2.032"))),
		x.Child(x.Path(x.D("M8.5 16.429h.01"))),
	)
	return x.Svg(svgArgs...)
}
