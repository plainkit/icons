package lucide

import (
	html "github.com/plainkit/html"
)

// AudioWaveform creates a Audio Waveform Lucide icon.
func AudioWaveform(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-audio-waveform", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M2 13a2 2 0 0 0 2-2V7a2 2 0 0 1 4 0v13a2 2 0 0 0 4 0V4a2 2 0 0 1 4 0v13a2 2 0 0 0 4 0v-4a2 2 0 0 1 2-2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
