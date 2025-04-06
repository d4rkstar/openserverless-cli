// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package tools

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func validateTool() error {
	flagSet := flag.NewFlagSet("validate", flag.ContinueOnError)

	flagSet.Usage = printValidateHelp

	helpFlag := flagSet.Bool("h", false, "Print this help message.")
	envFlag := flagSet.Bool("e", false, "Retrieve value from the environment variable with the given name.")
	mailFlag := flagSet.Bool("m", false, "Check if the value is a valid email address.")
	numberFlag := flagSet.Bool("n", false, "Check if the value is a number.")
	regexFlag := flagSet.String("r", "", "Check if the value matches the given regular expression.")

	err := flagSet.Parse(os.Args[1:])
	if err != nil {
		return err
	}

	if *helpFlag {
		flagSet.Usage()
		return nil
	}

	if flagSet.NArg() != 1 && flagSet.NArg() != 2 {
		flagSet.Usage()
		return errors.New("invalid number of arguments")
	}

	arg := flagSet.Arg(0)
	customErr := flagSet.Arg(1)
	if customErr == "" {
		customErr = "validation failed"
	}
	value := arg

	if *envFlag {
		value = os.Getenv(arg)
		if value == "" {
			return fmt.Errorf("variable '%s' not set", arg)
		}
	}

	if *mailFlag {
		if !isValidEmail(value) {
			return fmt.Errorf("%s", customErr)
		}
		return nil
	}

	if *numberFlag {
		if !isValidNumber(value) {
			return fmt.Errorf("%s", customErr)
		}
		return nil
	}

	if *regexFlag != "" {
		valid, err := isValidByRegex(value, *regexFlag)
		if err != nil {
			return err
		}

		if !valid {
			return fmt.Errorf("%s", customErr)
		}
	}

	return nil
}

func isValidNumber(number string) bool {
	_, err := strconv.ParseFloat(number, 64)
	return err == nil
}

func isValidEmail(email string) bool {
	// Regular expression pattern for email validation
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Create a regular expression object
	regExp := regexp.MustCompile(pattern)

	// Use the regular expression to match the email string
	return regExp.MatchString(email)
}

func isValidByRegex(value string, regex string) (bool, error) {
	// Create a regular expression object
	regExp, err := regexp.Compile(regex)
	if err != nil {
		return false, err
	}

	// Use the regular expression to match the email string
	return regExp.MatchString(value), nil
}

func printValidateHelp() {
	fmt.Println(MarkdownHelp("validate"))
}
