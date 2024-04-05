package lib

import (
	"log"

	"github.com/fatih/color"
)

// Define custom logger initialization constant
var MyLogger = log.New(color.Output, "", 0)
