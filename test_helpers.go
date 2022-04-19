/***** BEGIN LICENSE BLOCK *****
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this file,
# You can obtain one at http://mozilla.org/MPL/2.0/.
#
# The Initial Developer of the Original Code is the Mozilla Foundation.
# Portions created by the Initial Developer are Copyright (C) 2013-2015
# the Initial Developer. All Rights Reserved.
#
# Contributor(s):
#   Rob Miller (rmiller@mozilla.com)
#
# ***** END LICENSE BLOCK *****/

package testsupport

import (
	"github.com/wulog/engine/message"

	"github.com/rafrombrc/gomock/gomock"
	. "github.com/wulog/engine/pipeline"
	"github.com/wulog/engine/pipelinemock"
)

type InputTestHelper struct {
	Msg                *message.Message
	Pack               *PipelinePack
	AddrStr            string
	ResolvedAddrStr    string
	MockHelper         *pipelinemock.MockPluginHelper
	MockInputRunner    *pipelinemock.MockInputRunner
	MockDeliverer      *pipelinemock.MockDeliverer
	MockSplitterRunner *pipelinemock.MockSplitterRunner
	PackSupply         chan *PipelinePack
	// Decoder         DecoderRunner
	// DecodeChan      chan *PipelinePack
}

type OutputTestHelper struct {
	MockHelper       *pipelinemock.MockPluginHelper
	MockOutputRunner *pipelinemock.MockOutputRunner
}

func NewOutputTestHelper(ctrl *gomock.Controller) (oth *OutputTestHelper) {
	oth = new(OutputTestHelper)
	oth.MockHelper = pipelinemock.NewMockPluginHelper(ctrl)
	oth.MockOutputRunner = pipelinemock.NewMockOutputRunner(ctrl)
	return
}
