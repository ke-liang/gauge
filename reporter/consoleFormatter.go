// Copyright 2015 ThoughtWorks, Inc.

// This file is part of Gauge.

// Gauge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// Gauge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Gauge.  If not, see <http://www.gnu.org/licenses/>.

package reporter

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/getgauge/gauge/config"
	"github.com/getgauge/gauge/util"
)

const (
	scenarioIndentation = 2
	stepIndentation     = 4
	errorIndentation    = 2
	successSymbol       = "✔"
	failureSymbol       = "✘"
	successChar         = "P"
	failureChar         = "F"
)

func formatScenario(scenarioHeading string) string {
	return fmt.Sprintf("## %s", scenarioHeading)
}

func formatSpec(specHeading string) string {
	return fmt.Sprintf("# %s", specHeading)
}

func indent(text string, indentation int) string {
	return spaces(indentation) + strings.Replace(text, newline, newline+spaces(indentation), -1)
}

func spaces(numOfSpaces int) string {
	return strings.Repeat(" ", numOfSpaces)
}

func getFailureSymbol() string {
	if util.IsWindows() {
		return spaces(1) + failureChar
	}
	return spaces(1) + failureSymbol
}

func getSuccessSymbol() string {
	if util.IsWindows() {
		return spaces(1) + successChar
	}
	return spaces(1) + successSymbol
}

func prepErrorMessage(msg string) string {
	return fmt.Sprintf("Error Message: %s", msg)
}

func prepStepMsg(msg string) string {
	return fmt.Sprintf("\nFailed Step: %s", msg)
}

func prepSpecInfo(fileName string, lineNo int) string {
	return fmt.Sprintf("Specification: %s:%v", getRelativePath(fileName), lineNo)
}

func prepStacktrace(stacktrace string) string {
	return fmt.Sprintf("Stacktrace: \n%s", stacktrace)
}

func formatErrorMessage(msg string, indentation int) string {
	return indent(msg, indentation+errorIndentation) + newline
}

func formatSpecInfo(specInfo string, indentation int) string {
	return indent(specInfo, indentation+errorIndentation) + newline
}

func getRelativePath(path string) string {
	return strings.TrimPrefix(path, config.ProjectRoot+string(filepath.Separator))
}

func formatStacktrace(stacktrace string, indentation int) string {
	return indent(stacktrace, indentation+errorIndentation) + newline
}

func formatStepText(text string, indentation int) string {
	return indent(text, indentation+errorIndentation) + newline
}
