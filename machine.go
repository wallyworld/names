// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package names

import (
	"regexp"
	"strings"
)

const MachineTagKind = "machine"

const (
	ContainerTypeSnippet = "[a-z]+"
	ContainerSnippet     = "(/" + ContainerTypeSnippet + "/" + NumberSnippet + ")"
	MachineSnippet       = NumberSnippet + ContainerSnippet + "*"
)

var validMachine = regexp.MustCompile("^" + MachineSnippet + "$")

// IsMachine returns whether id is a valid machine id.
var IsMachine = validMachine.MatchString

// IsContainerMachine returns whether id is a valid container machine id.
func IsContainerMachine(id string) bool {
	return validMachine.MatchString(id) && strings.Contains(id, "/")
}

type MachineTag struct {
	id string
}

func (t MachineTag) String() string { return MachineTagKind + "-" + t.id }

// NewMachineTag returns the tag for the machine with the given id.
func NewMachineTag(id string) Tag {
	id = strings.Replace(id, "/", "-", -1)
	return MachineTag{id: id}
}

func machineTagSuffixToId(s string) string {
	return strings.Replace(s, "-", "/", -1)
}
