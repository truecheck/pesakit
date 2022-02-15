/*
 * MIT License
 *
 * Copyright (c) 2021 PESAKIT - MOBILE MONEY TOOLBOX
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 */

package main

import (
	"github.com/urfave/cli/v2"
	"io"
	"time"
)

type AppDescriptor struct {
	Name        string
	Usage       string
	Description string
	Version     string
	Authors     []*cli.Author
	Email       string
	License     string
	CompiledAt  time.Time
}

func CreateRootCommand(reader io.Reader, writer io.Writer, errWriter io.Writer, descriptor AppDescriptor, commands []*cli.Command, flags []cli.Flag) *cli.App {
	var rootCommand = &cli.App{
		Name:                   descriptor.Name,
		Usage:                  descriptor.Usage,
		Version:                descriptor.Version,
		Description:            descriptor.Description,
		Commands:               commands,
		Flags:                  flags,
		BashComplete:           nil,
		Before:                 nil,
		After:                  nil,
		Action:                 nil,
		CommandNotFound:        nil,
		OnUsageError:           nil,
		Compiled:               descriptor.CompiledAt,
		Authors:                descriptor.Authors,
		Copyright:              descriptor.License,
		Reader:                 reader,
		Writer:                 writer,
		ErrWriter:              errWriter,
		ExitErrHandler:         nil,
		Metadata:               appMetadata,
		ExtraInfo:              nil,
		CustomAppHelpTemplate:  "",
		UseShortOptionHandling: false,
	}

	return rootCommand
}

const (
	// Version is the current version of the cli
	Version  = "0.0.1"
	AppName  = "pesakit"
	AppUsage = "mobile money toolbox for developers,testers and software engineers"
	AppDesc  = `
PESAKIT - Mobile Money Toolbox for Developers, Testers and Software Engineers. It can
be used during the development process for integration testing, debugging and mocking
For the ambitious It can be used with production credentials for production testing to
perform C2B, B2B, B2C, Token generation, Transaction reversal and Account balance enquiry.`
	AppUsageText = "PESAKIT - Mobile Money Toolbox"
)

var author = &cli.Author{
	Name:  "Pius Masengwa Alfred",
	Email: "me.pius1102@gmail.com",
}
var authors = []*cli.Author{
	author,
}
var appMetadata = map[string]interface{}{
	"Created": time.Now().Format("2006-01-02"),
	"Version": Version,
	"Company": "PESAKIT",
	"Twitter": "https://twitter.com/pesabox",
	"Github":  "https://github.com/pesakit",
}

var extraInfo = map[string]string{
	"Snapcraft": "https://snapcraft.io/pesakit",
}

type CommandDescriptor struct {
	Name            string
	Description     string
	Usage           string
	UsageText       string
	DescriptionText string
	Category        string
}

func setFlags(command *cli.Command, flags ...cli.Flag) {
	command.Flags = flags
}
