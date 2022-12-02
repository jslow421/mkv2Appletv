package main

import (
	"github.com/alecthomas/kingpin"
	"os"
)

var (
	app   = kingpin.New("mkv2Appletv", "Convert as efficiently as possible media to AppleTV mp4 format.")
	debug = app.Flag("debug", "Enable debug mode.").Short('d').Bool()
	try   = app.Flag(
		"try", "When set to true only the first 10 seconds of conversion will be done").Short('t').Bool()
	input  = app.Flag("input", "Location of input File").Short('i').Required().ExistingFileOrDir()
	output = app.Flag("output", "Location of output mp4 File (not required, if not set output file name will"+
		" be input filename +.mp4)").Short('o').String()

	show    = app.Command("show", "Using ffprobe show relavant information about a input file")
	suggest = app.Command("suggest", "Show what the suggested output of the transformation would look like.")
	convert = app.Command("convert", "Take input and run ffmpeg to generate an optimal mp4 file")
	check   = app.Command("check", "log information about ffprobe and ffmpeg")
	folder  = app.Command("folder", "Convert entire folder to optimal mp4 file")
)

func main() {
	kingpin.Version("0.0.8")
	app.UsageTemplate(kingpin.SeparateOptionalFlagsUsageTemplate)

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {

	case show.FullCommand():
		showFFprobeInfo(*input)

	case suggest.FullCommand():
		suggestConvSettings(*input)

	case convert.FullCommand():
		convertSource(*input, *output)

	case folder.FullCommand():
		//handleFolderConversion(*input)
		handleFolderConversion(*input, *output)
	case check.FullCommand():
		checkVersion()
	}
}
