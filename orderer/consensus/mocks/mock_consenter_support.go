// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"justledger/common/channelconfig"
	"justledger/orderer/common/blockcutter"
	"justledger/orderer/common/msgprocessor"
	"justledger/orderer/consensus"
	cb "justledger/protos/common"
)

type FakeConsenterSupport struct {
	NewSignatureHeaderStub        func() (*cb.SignatureHeader, error)
	newSignatureHeaderMutex       sync.RWMutex
	newSignatureHeaderArgsForCall []struct{}
	newSignatureHeaderReturns     struct {
		result1 *cb.SignatureHeader
		result2 error
	}
	newSignatureHeaderReturnsOnCall map[int]struct {
		result1 *cb.SignatureHeader
		result2 error
	}
	SignStub        func(message []byte) ([]byte, error)
	signMutex       sync.RWMutex
	signArgsForCall []struct {
		message []byte
	}
	signReturns struct {
		result1 []byte
		result2 error
	}
	signReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	ClassifyMsgStub        func(chdr *cb.ChannelHeader) msgprocessor.Classification
	classifyMsgMutex       sync.RWMutex
	classifyMsgArgsForCall []struct {
		chdr *cb.ChannelHeader
	}
	classifyMsgReturns struct {
		result1 msgprocessor.Classification
	}
	classifyMsgReturnsOnCall map[int]struct {
		result1 msgprocessor.Classification
	}
	ProcessNormalMsgStub        func(env *cb.Envelope) (configSeq uint64, err error)
	processNormalMsgMutex       sync.RWMutex
	processNormalMsgArgsForCall []struct {
		env *cb.Envelope
	}
	processNormalMsgReturns struct {
		result1 uint64
		result2 error
	}
	processNormalMsgReturnsOnCall map[int]struct {
		result1 uint64
		result2 error
	}
	ProcessConfigUpdateMsgStub        func(env *cb.Envelope) (config *cb.Envelope, configSeq uint64, err error)
	processConfigUpdateMsgMutex       sync.RWMutex
	processConfigUpdateMsgArgsForCall []struct {
		env *cb.Envelope
	}
	processConfigUpdateMsgReturns struct {
		result1 *cb.Envelope
		result2 uint64
		result3 error
	}
	processConfigUpdateMsgReturnsOnCall map[int]struct {
		result1 *cb.Envelope
		result2 uint64
		result3 error
	}
	ProcessConfigMsgStub        func(env *cb.Envelope) (*cb.Envelope, uint64, error)
	processConfigMsgMutex       sync.RWMutex
	processConfigMsgArgsForCall []struct {
		env *cb.Envelope
	}
	processConfigMsgReturns struct {
		result1 *cb.Envelope
		result2 uint64
		result3 error
	}
	processConfigMsgReturnsOnCall map[int]struct {
		result1 *cb.Envelope
		result2 uint64
		result3 error
	}
	BlockCutterStub        func() blockcutter.Receiver
	blockCutterMutex       sync.RWMutex
	blockCutterArgsForCall []struct{}
	blockCutterReturns     struct {
		result1 blockcutter.Receiver
	}
	blockCutterReturnsOnCall map[int]struct {
		result1 blockcutter.Receiver
	}
	SharedConfigStub        func() channelconfig.Orderer
	sharedConfigMutex       sync.RWMutex
	sharedConfigArgsForCall []struct{}
	sharedConfigReturns     struct {
		result1 channelconfig.Orderer
	}
	sharedConfigReturnsOnCall map[int]struct {
		result1 channelconfig.Orderer
	}
	CreateNextBlockStub        func(messages []*cb.Envelope) *cb.Block
	createNextBlockMutex       sync.RWMutex
	createNextBlockArgsForCall []struct {
		messages []*cb.Envelope
	}
	createNextBlockReturns struct {
		result1 *cb.Block
	}
	createNextBlockReturnsOnCall map[int]struct {
		result1 *cb.Block
	}
	WriteBlockStub        func(block *cb.Block, encodedMetadataValue []byte)
	writeBlockMutex       sync.RWMutex
	writeBlockArgsForCall []struct {
		block                *cb.Block
		encodedMetadataValue []byte
	}
	WriteConfigBlockStub        func(block *cb.Block, encodedMetadataValue []byte)
	writeConfigBlockMutex       sync.RWMutex
	writeConfigBlockArgsForCall []struct {
		block                *cb.Block
		encodedMetadataValue []byte
	}
	SequenceStub        func() uint64
	sequenceMutex       sync.RWMutex
	sequenceArgsForCall []struct{}
	sequenceReturns     struct {
		result1 uint64
	}
	sequenceReturnsOnCall map[int]struct {
		result1 uint64
	}
	ChainIDStub        func() string
	chainIDMutex       sync.RWMutex
	chainIDArgsForCall []struct{}
	chainIDReturns     struct {
		result1 string
	}
	chainIDReturnsOnCall map[int]struct {
		result1 string
	}
	HeightStub        func() uint64
	heightMutex       sync.RWMutex
	heightArgsForCall []struct{}
	heightReturns     struct {
		result1 uint64
	}
	heightReturnsOnCall map[int]struct {
		result1 uint64
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConsenterSupport) NewSignatureHeader() (*cb.SignatureHeader, error) {
	fake.newSignatureHeaderMutex.Lock()
	ret, specificReturn := fake.newSignatureHeaderReturnsOnCall[len(fake.newSignatureHeaderArgsForCall)]
	fake.newSignatureHeaderArgsForCall = append(fake.newSignatureHeaderArgsForCall, struct{}{})
	fake.recordInvocation("NewSignatureHeader", []interface{}{})
	fake.newSignatureHeaderMutex.Unlock()
	if fake.NewSignatureHeaderStub != nil {
		return fake.NewSignatureHeaderStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.newSignatureHeaderReturns.result1, fake.newSignatureHeaderReturns.result2
}

func (fake *FakeConsenterSupport) NewSignatureHeaderCallCount() int {
	fake.newSignatureHeaderMutex.RLock()
	defer fake.newSignatureHeaderMutex.RUnlock()
	return len(fake.newSignatureHeaderArgsForCall)
}

func (fake *FakeConsenterSupport) NewSignatureHeaderReturns(result1 *cb.SignatureHeader, result2 error) {
	fake.NewSignatureHeaderStub = nil
	fake.newSignatureHeaderReturns = struct {
		result1 *cb.SignatureHeader
		result2 error
	}{result1, result2}
}

func (fake *FakeConsenterSupport) NewSignatureHeaderReturnsOnCall(i int, result1 *cb.SignatureHeader, result2 error) {
	fake.NewSignatureHeaderStub = nil
	if fake.newSignatureHeaderReturnsOnCall == nil {
		fake.newSignatureHeaderReturnsOnCall = make(map[int]struct {
			result1 *cb.SignatureHeader
			result2 error
		})
	}
	fake.newSignatureHeaderReturnsOnCall[i] = struct {
		result1 *cb.SignatureHeader
		result2 error
	}{result1, result2}
}

func (fake *FakeConsenterSupport) Sign(message []byte) ([]byte, error) {
	var messageCopy []byte
	if message != nil {
		messageCopy = make([]byte, len(message))
		copy(messageCopy, message)
	}
	fake.signMutex.Lock()
	ret, specificReturn := fake.signReturnsOnCall[len(fake.signArgsForCall)]
	fake.signArgsForCall = append(fake.signArgsForCall, struct {
		message []byte
	}{messageCopy})
	fake.recordInvocation("Sign", []interface{}{messageCopy})
	fake.signMutex.Unlock()
	if fake.SignStub != nil {
		return fake.SignStub(message)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.signReturns.result1, fake.signReturns.result2
}

func (fake *FakeConsenterSupport) SignCallCount() int {
	fake.signMutex.RLock()
	defer fake.signMutex.RUnlock()
	return len(fake.signArgsForCall)
}

func (fake *FakeConsenterSupport) SignArgsForCall(i int) []byte {
	fake.signMutex.RLock()
	defer fake.signMutex.RUnlock()
	return fake.signArgsForCall[i].message
}

func (fake *FakeConsenterSupport) SignReturns(result1 []byte, result2 error) {
	fake.SignStub = nil
	fake.signReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeConsenterSupport) SignReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.SignStub = nil
	if fake.signReturnsOnCall == nil {
		fake.signReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.signReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeConsenterSupport) ClassifyMsg(chdr *cb.ChannelHeader) msgprocessor.Classification {
	fake.classifyMsgMutex.Lock()
	ret, specificReturn := fake.classifyMsgReturnsOnCall[len(fake.classifyMsgArgsForCall)]
	fake.classifyMsgArgsForCall = append(fake.classifyMsgArgsForCall, struct {
		chdr *cb.ChannelHeader
	}{chdr})
	fake.recordInvocation("ClassifyMsg", []interface{}{chdr})
	fake.classifyMsgMutex.Unlock()
	if fake.ClassifyMsgStub != nil {
		return fake.ClassifyMsgStub(chdr)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.classifyMsgReturns.result1
}

func (fake *FakeConsenterSupport) ClassifyMsgCallCount() int {
	fake.classifyMsgMutex.RLock()
	defer fake.classifyMsgMutex.RUnlock()
	return len(fake.classifyMsgArgsForCall)
}

func (fake *FakeConsenterSupport) ClassifyMsgArgsForCall(i int) *cb.ChannelHeader {
	fake.classifyMsgMutex.RLock()
	defer fake.classifyMsgMutex.RUnlock()
	return fake.classifyMsgArgsForCall[i].chdr
}

func (fake *FakeConsenterSupport) ClassifyMsgReturns(result1 msgprocessor.Classification) {
	fake.ClassifyMsgStub = nil
	fake.classifyMsgReturns = struct {
		result1 msgprocessor.Classification
	}{result1}
}

func (fake *FakeConsenterSupport) ClassifyMsgReturnsOnCall(i int, result1 msgprocessor.Classification) {
	fake.ClassifyMsgStub = nil
	if fake.classifyMsgReturnsOnCall == nil {
		fake.classifyMsgReturnsOnCall = make(map[int]struct {
			result1 msgprocessor.Classification
		})
	}
	fake.classifyMsgReturnsOnCall[i] = struct {
		result1 msgprocessor.Classification
	}{result1}
}

func (fake *FakeConsenterSupport) ProcessNormalMsg(env *cb.Envelope) (configSeq uint64, err error) {
	fake.processNormalMsgMutex.Lock()
	ret, specificReturn := fake.processNormalMsgReturnsOnCall[len(fake.processNormalMsgArgsForCall)]
	fake.processNormalMsgArgsForCall = append(fake.processNormalMsgArgsForCall, struct {
		env *cb.Envelope
	}{env})
	fake.recordInvocation("ProcessNormalMsg", []interface{}{env})
	fake.processNormalMsgMutex.Unlock()
	if fake.ProcessNormalMsgStub != nil {
		return fake.ProcessNormalMsgStub(env)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.processNormalMsgReturns.result1, fake.processNormalMsgReturns.result2
}

func (fake *FakeConsenterSupport) ProcessNormalMsgCallCount() int {
	fake.processNormalMsgMutex.RLock()
	defer fake.processNormalMsgMutex.RUnlock()
	return len(fake.processNormalMsgArgsForCall)
}

func (fake *FakeConsenterSupport) ProcessNormalMsgArgsForCall(i int) *cb.Envelope {
	fake.processNormalMsgMutex.RLock()
	defer fake.processNormalMsgMutex.RUnlock()
	return fake.processNormalMsgArgsForCall[i].env
}

func (fake *FakeConsenterSupport) ProcessNormalMsgReturns(result1 uint64, result2 error) {
	fake.ProcessNormalMsgStub = nil
	fake.processNormalMsgReturns = struct {
		result1 uint64
		result2 error
	}{result1, result2}
}

func (fake *FakeConsenterSupport) ProcessNormalMsgReturnsOnCall(i int, result1 uint64, result2 error) {
	fake.ProcessNormalMsgStub = nil
	if fake.processNormalMsgReturnsOnCall == nil {
		fake.processNormalMsgReturnsOnCall = make(map[int]struct {
			result1 uint64
			result2 error
		})
	}
	fake.processNormalMsgReturnsOnCall[i] = struct {
		result1 uint64
		result2 error
	}{result1, result2}
}

func (fake *FakeConsenterSupport) ProcessConfigUpdateMsg(env *cb.Envelope) (config *cb.Envelope, configSeq uint64, err error) {
	fake.processConfigUpdateMsgMutex.Lock()
	ret, specificReturn := fake.processConfigUpdateMsgReturnsOnCall[len(fake.processConfigUpdateMsgArgsForCall)]
	fake.processConfigUpdateMsgArgsForCall = append(fake.processConfigUpdateMsgArgsForCall, struct {
		env *cb.Envelope
	}{env})
	fake.recordInvocation("ProcessConfigUpdateMsg", []interface{}{env})
	fake.processConfigUpdateMsgMutex.Unlock()
	if fake.ProcessConfigUpdateMsgStub != nil {
		return fake.ProcessConfigUpdateMsgStub(env)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.processConfigUpdateMsgReturns.result1, fake.processConfigUpdateMsgReturns.result2, fake.processConfigUpdateMsgReturns.result3
}

func (fake *FakeConsenterSupport) ProcessConfigUpdateMsgCallCount() int {
	fake.processConfigUpdateMsgMutex.RLock()
	defer fake.processConfigUpdateMsgMutex.RUnlock()
	return len(fake.processConfigUpdateMsgArgsForCall)
}

func (fake *FakeConsenterSupport) ProcessConfigUpdateMsgArgsForCall(i int) *cb.Envelope {
	fake.processConfigUpdateMsgMutex.RLock()
	defer fake.processConfigUpdateMsgMutex.RUnlock()
	return fake.processConfigUpdateMsgArgsForCall[i].env
}

func (fake *FakeConsenterSupport) ProcessConfigUpdateMsgReturns(result1 *cb.Envelope, result2 uint64, result3 error) {
	fake.ProcessConfigUpdateMsgStub = nil
	fake.processConfigUpdateMsgReturns = struct {
		result1 *cb.Envelope
		result2 uint64
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeConsenterSupport) ProcessConfigUpdateMsgReturnsOnCall(i int, result1 *cb.Envelope, result2 uint64, result3 error) {
	fake.ProcessConfigUpdateMsgStub = nil
	if fake.processConfigUpdateMsgReturnsOnCall == nil {
		fake.processConfigUpdateMsgReturnsOnCall = make(map[int]struct {
			result1 *cb.Envelope
			result2 uint64
			result3 error
		})
	}
	fake.processConfigUpdateMsgReturnsOnCall[i] = struct {
		result1 *cb.Envelope
		result2 uint64
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeConsenterSupport) ProcessConfigMsg(env *cb.Envelope) (*cb.Envelope, uint64, error) {
	fake.processConfigMsgMutex.Lock()
	ret, specificReturn := fake.processConfigMsgReturnsOnCall[len(fake.processConfigMsgArgsForCall)]
	fake.processConfigMsgArgsForCall = append(fake.processConfigMsgArgsForCall, struct {
		env *cb.Envelope
	}{env})
	fake.recordInvocation("ProcessConfigMsg", []interface{}{env})
	fake.processConfigMsgMutex.Unlock()
	if fake.ProcessConfigMsgStub != nil {
		return fake.ProcessConfigMsgStub(env)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.processConfigMsgReturns.result1, fake.processConfigMsgReturns.result2, fake.processConfigMsgReturns.result3
}

func (fake *FakeConsenterSupport) ProcessConfigMsgCallCount() int {
	fake.processConfigMsgMutex.RLock()
	defer fake.processConfigMsgMutex.RUnlock()
	return len(fake.processConfigMsgArgsForCall)
}

func (fake *FakeConsenterSupport) ProcessConfigMsgArgsForCall(i int) *cb.Envelope {
	fake.processConfigMsgMutex.RLock()
	defer fake.processConfigMsgMutex.RUnlock()
	return fake.processConfigMsgArgsForCall[i].env
}

func (fake *FakeConsenterSupport) ProcessConfigMsgReturns(result1 *cb.Envelope, result2 uint64, result3 error) {
	fake.ProcessConfigMsgStub = nil
	fake.processConfigMsgReturns = struct {
		result1 *cb.Envelope
		result2 uint64
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeConsenterSupport) ProcessConfigMsgReturnsOnCall(i int, result1 *cb.Envelope, result2 uint64, result3 error) {
	fake.ProcessConfigMsgStub = nil
	if fake.processConfigMsgReturnsOnCall == nil {
		fake.processConfigMsgReturnsOnCall = make(map[int]struct {
			result1 *cb.Envelope
			result2 uint64
			result3 error
		})
	}
	fake.processConfigMsgReturnsOnCall[i] = struct {
		result1 *cb.Envelope
		result2 uint64
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeConsenterSupport) BlockCutter() blockcutter.Receiver {
	fake.blockCutterMutex.Lock()
	ret, specificReturn := fake.blockCutterReturnsOnCall[len(fake.blockCutterArgsForCall)]
	fake.blockCutterArgsForCall = append(fake.blockCutterArgsForCall, struct{}{})
	fake.recordInvocation("BlockCutter", []interface{}{})
	fake.blockCutterMutex.Unlock()
	if fake.BlockCutterStub != nil {
		return fake.BlockCutterStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.blockCutterReturns.result1
}

func (fake *FakeConsenterSupport) BlockCutterCallCount() int {
	fake.blockCutterMutex.RLock()
	defer fake.blockCutterMutex.RUnlock()
	return len(fake.blockCutterArgsForCall)
}

func (fake *FakeConsenterSupport) BlockCutterReturns(result1 blockcutter.Receiver) {
	fake.BlockCutterStub = nil
	fake.blockCutterReturns = struct {
		result1 blockcutter.Receiver
	}{result1}
}

func (fake *FakeConsenterSupport) BlockCutterReturnsOnCall(i int, result1 blockcutter.Receiver) {
	fake.BlockCutterStub = nil
	if fake.blockCutterReturnsOnCall == nil {
		fake.blockCutterReturnsOnCall = make(map[int]struct {
			result1 blockcutter.Receiver
		})
	}
	fake.blockCutterReturnsOnCall[i] = struct {
		result1 blockcutter.Receiver
	}{result1}
}

func (fake *FakeConsenterSupport) SharedConfig() channelconfig.Orderer {
	fake.sharedConfigMutex.Lock()
	ret, specificReturn := fake.sharedConfigReturnsOnCall[len(fake.sharedConfigArgsForCall)]
	fake.sharedConfigArgsForCall = append(fake.sharedConfigArgsForCall, struct{}{})
	fake.recordInvocation("SharedConfig", []interface{}{})
	fake.sharedConfigMutex.Unlock()
	if fake.SharedConfigStub != nil {
		return fake.SharedConfigStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.sharedConfigReturns.result1
}

func (fake *FakeConsenterSupport) SharedConfigCallCount() int {
	fake.sharedConfigMutex.RLock()
	defer fake.sharedConfigMutex.RUnlock()
	return len(fake.sharedConfigArgsForCall)
}

func (fake *FakeConsenterSupport) SharedConfigReturns(result1 channelconfig.Orderer) {
	fake.SharedConfigStub = nil
	fake.sharedConfigReturns = struct {
		result1 channelconfig.Orderer
	}{result1}
}

func (fake *FakeConsenterSupport) SharedConfigReturnsOnCall(i int, result1 channelconfig.Orderer) {
	fake.SharedConfigStub = nil
	if fake.sharedConfigReturnsOnCall == nil {
		fake.sharedConfigReturnsOnCall = make(map[int]struct {
			result1 channelconfig.Orderer
		})
	}
	fake.sharedConfigReturnsOnCall[i] = struct {
		result1 channelconfig.Orderer
	}{result1}
}

func (fake *FakeConsenterSupport) CreateNextBlock(messages []*cb.Envelope) *cb.Block {
	var messagesCopy []*cb.Envelope
	if messages != nil {
		messagesCopy = make([]*cb.Envelope, len(messages))
		copy(messagesCopy, messages)
	}
	fake.createNextBlockMutex.Lock()
	ret, specificReturn := fake.createNextBlockReturnsOnCall[len(fake.createNextBlockArgsForCall)]
	fake.createNextBlockArgsForCall = append(fake.createNextBlockArgsForCall, struct {
		messages []*cb.Envelope
	}{messagesCopy})
	fake.recordInvocation("CreateNextBlock", []interface{}{messagesCopy})
	fake.createNextBlockMutex.Unlock()
	if fake.CreateNextBlockStub != nil {
		return fake.CreateNextBlockStub(messages)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createNextBlockReturns.result1
}

func (fake *FakeConsenterSupport) CreateNextBlockCallCount() int {
	fake.createNextBlockMutex.RLock()
	defer fake.createNextBlockMutex.RUnlock()
	return len(fake.createNextBlockArgsForCall)
}

func (fake *FakeConsenterSupport) CreateNextBlockArgsForCall(i int) []*cb.Envelope {
	fake.createNextBlockMutex.RLock()
	defer fake.createNextBlockMutex.RUnlock()
	return fake.createNextBlockArgsForCall[i].messages
}

func (fake *FakeConsenterSupport) CreateNextBlockReturns(result1 *cb.Block) {
	fake.CreateNextBlockStub = nil
	fake.createNextBlockReturns = struct {
		result1 *cb.Block
	}{result1}
}

func (fake *FakeConsenterSupport) CreateNextBlockReturnsOnCall(i int, result1 *cb.Block) {
	fake.CreateNextBlockStub = nil
	if fake.createNextBlockReturnsOnCall == nil {
		fake.createNextBlockReturnsOnCall = make(map[int]struct {
			result1 *cb.Block
		})
	}
	fake.createNextBlockReturnsOnCall[i] = struct {
		result1 *cb.Block
	}{result1}
}

func (fake *FakeConsenterSupport) WriteBlock(block *cb.Block, encodedMetadataValue []byte) {
	var encodedMetadataValueCopy []byte
	if encodedMetadataValue != nil {
		encodedMetadataValueCopy = make([]byte, len(encodedMetadataValue))
		copy(encodedMetadataValueCopy, encodedMetadataValue)
	}
	fake.writeBlockMutex.Lock()
	fake.writeBlockArgsForCall = append(fake.writeBlockArgsForCall, struct {
		block                *cb.Block
		encodedMetadataValue []byte
	}{block, encodedMetadataValueCopy})
	fake.recordInvocation("WriteBlock", []interface{}{block, encodedMetadataValueCopy})
	fake.writeBlockMutex.Unlock()
	if fake.WriteBlockStub != nil {
		fake.WriteBlockStub(block, encodedMetadataValue)
	}
}

func (fake *FakeConsenterSupport) WriteBlockCallCount() int {
	fake.writeBlockMutex.RLock()
	defer fake.writeBlockMutex.RUnlock()
	return len(fake.writeBlockArgsForCall)
}

func (fake *FakeConsenterSupport) WriteBlockArgsForCall(i int) (*cb.Block, []byte) {
	fake.writeBlockMutex.RLock()
	defer fake.writeBlockMutex.RUnlock()
	return fake.writeBlockArgsForCall[i].block, fake.writeBlockArgsForCall[i].encodedMetadataValue
}

func (fake *FakeConsenterSupport) WriteConfigBlock(block *cb.Block, encodedMetadataValue []byte) {
	var encodedMetadataValueCopy []byte
	if encodedMetadataValue != nil {
		encodedMetadataValueCopy = make([]byte, len(encodedMetadataValue))
		copy(encodedMetadataValueCopy, encodedMetadataValue)
	}
	fake.writeConfigBlockMutex.Lock()
	fake.writeConfigBlockArgsForCall = append(fake.writeConfigBlockArgsForCall, struct {
		block                *cb.Block
		encodedMetadataValue []byte
	}{block, encodedMetadataValueCopy})
	fake.recordInvocation("WriteConfigBlock", []interface{}{block, encodedMetadataValueCopy})
	fake.writeConfigBlockMutex.Unlock()
	if fake.WriteConfigBlockStub != nil {
		fake.WriteConfigBlockStub(block, encodedMetadataValue)
	}
}

func (fake *FakeConsenterSupport) WriteConfigBlockCallCount() int {
	fake.writeConfigBlockMutex.RLock()
	defer fake.writeConfigBlockMutex.RUnlock()
	return len(fake.writeConfigBlockArgsForCall)
}

func (fake *FakeConsenterSupport) WriteConfigBlockArgsForCall(i int) (*cb.Block, []byte) {
	fake.writeConfigBlockMutex.RLock()
	defer fake.writeConfigBlockMutex.RUnlock()
	return fake.writeConfigBlockArgsForCall[i].block, fake.writeConfigBlockArgsForCall[i].encodedMetadataValue
}

func (fake *FakeConsenterSupport) Sequence() uint64 {
	fake.sequenceMutex.Lock()
	ret, specificReturn := fake.sequenceReturnsOnCall[len(fake.sequenceArgsForCall)]
	fake.sequenceArgsForCall = append(fake.sequenceArgsForCall, struct{}{})
	fake.recordInvocation("Sequence", []interface{}{})
	fake.sequenceMutex.Unlock()
	if fake.SequenceStub != nil {
		return fake.SequenceStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.sequenceReturns.result1
}

func (fake *FakeConsenterSupport) SequenceCallCount() int {
	fake.sequenceMutex.RLock()
	defer fake.sequenceMutex.RUnlock()
	return len(fake.sequenceArgsForCall)
}

func (fake *FakeConsenterSupport) SequenceReturns(result1 uint64) {
	fake.SequenceStub = nil
	fake.sequenceReturns = struct {
		result1 uint64
	}{result1}
}

func (fake *FakeConsenterSupport) SequenceReturnsOnCall(i int, result1 uint64) {
	fake.SequenceStub = nil
	if fake.sequenceReturnsOnCall == nil {
		fake.sequenceReturnsOnCall = make(map[int]struct {
			result1 uint64
		})
	}
	fake.sequenceReturnsOnCall[i] = struct {
		result1 uint64
	}{result1}
}

func (fake *FakeConsenterSupport) ChainID() string {
	fake.chainIDMutex.Lock()
	ret, specificReturn := fake.chainIDReturnsOnCall[len(fake.chainIDArgsForCall)]
	fake.chainIDArgsForCall = append(fake.chainIDArgsForCall, struct{}{})
	fake.recordInvocation("ChainID", []interface{}{})
	fake.chainIDMutex.Unlock()
	if fake.ChainIDStub != nil {
		return fake.ChainIDStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.chainIDReturns.result1
}

func (fake *FakeConsenterSupport) ChainIDCallCount() int {
	fake.chainIDMutex.RLock()
	defer fake.chainIDMutex.RUnlock()
	return len(fake.chainIDArgsForCall)
}

func (fake *FakeConsenterSupport) ChainIDReturns(result1 string) {
	fake.ChainIDStub = nil
	fake.chainIDReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConsenterSupport) ChainIDReturnsOnCall(i int, result1 string) {
	fake.ChainIDStub = nil
	if fake.chainIDReturnsOnCall == nil {
		fake.chainIDReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.chainIDReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeConsenterSupport) Height() uint64 {
	fake.heightMutex.Lock()
	ret, specificReturn := fake.heightReturnsOnCall[len(fake.heightArgsForCall)]
	fake.heightArgsForCall = append(fake.heightArgsForCall, struct{}{})
	fake.recordInvocation("Height", []interface{}{})
	fake.heightMutex.Unlock()
	if fake.HeightStub != nil {
		return fake.HeightStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.heightReturns.result1
}

func (fake *FakeConsenterSupport) HeightCallCount() int {
	fake.heightMutex.RLock()
	defer fake.heightMutex.RUnlock()
	return len(fake.heightArgsForCall)
}

func (fake *FakeConsenterSupport) HeightReturns(result1 uint64) {
	fake.HeightStub = nil
	fake.heightReturns = struct {
		result1 uint64
	}{result1}
}

func (fake *FakeConsenterSupport) HeightReturnsOnCall(i int, result1 uint64) {
	fake.HeightStub = nil
	if fake.heightReturnsOnCall == nil {
		fake.heightReturnsOnCall = make(map[int]struct {
			result1 uint64
		})
	}
	fake.heightReturnsOnCall[i] = struct {
		result1 uint64
	}{result1}
}

func (fake *FakeConsenterSupport) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newSignatureHeaderMutex.RLock()
	defer fake.newSignatureHeaderMutex.RUnlock()
	fake.signMutex.RLock()
	defer fake.signMutex.RUnlock()
	fake.classifyMsgMutex.RLock()
	defer fake.classifyMsgMutex.RUnlock()
	fake.processNormalMsgMutex.RLock()
	defer fake.processNormalMsgMutex.RUnlock()
	fake.processConfigUpdateMsgMutex.RLock()
	defer fake.processConfigUpdateMsgMutex.RUnlock()
	fake.processConfigMsgMutex.RLock()
	defer fake.processConfigMsgMutex.RUnlock()
	fake.blockCutterMutex.RLock()
	defer fake.blockCutterMutex.RUnlock()
	fake.sharedConfigMutex.RLock()
	defer fake.sharedConfigMutex.RUnlock()
	fake.createNextBlockMutex.RLock()
	defer fake.createNextBlockMutex.RUnlock()
	fake.writeBlockMutex.RLock()
	defer fake.writeBlockMutex.RUnlock()
	fake.writeConfigBlockMutex.RLock()
	defer fake.writeConfigBlockMutex.RUnlock()
	fake.sequenceMutex.RLock()
	defer fake.sequenceMutex.RUnlock()
	fake.chainIDMutex.RLock()
	defer fake.chainIDMutex.RUnlock()
	fake.heightMutex.RLock()
	defer fake.heightMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeConsenterSupport) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ consensus.ConsenterSupport = new(FakeConsenterSupport)
