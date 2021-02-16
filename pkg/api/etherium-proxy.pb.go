// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: api/etherium-proxy.proto

package api

import (
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetBlockByNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId string `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
}

func (x *GetBlockByNumberRequest) Reset() {
	*x = GetBlockByNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_etherium_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockByNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockByNumberRequest) ProtoMessage() {}

func (x *GetBlockByNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_etherium_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockByNumberRequest.ProtoReflect.Descriptor instead.
func (*GetBlockByNumberRequest) Descriptor() ([]byte, []int) {
	return file_api_etherium_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *GetBlockByNumberRequest) GetBlockId() string {
	if x != nil {
		return x.BlockId
	}
	return ""
}

type GetBlockByNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block *Block `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *GetBlockByNumberResponse) Reset() {
	*x = GetBlockByNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_etherium_proxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockByNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockByNumberResponse) ProtoMessage() {}

func (x *GetBlockByNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_etherium_proxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockByNumberResponse.ProtoReflect.Descriptor instead.
func (*GetBlockByNumberResponse) Descriptor() ([]byte, []int) {
	return file_api_etherium_proxy_proto_rawDescGZIP(), []int{1}
}

func (x *GetBlockByNumberResponse) GetBlock() *Block {
	if x != nil {
		return x.Block
	}
	return nil
}

type GetTxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId string `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	TxId    string `protobuf:"bytes,2,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
}

func (x *GetTxRequest) Reset() {
	*x = GetTxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_etherium_proxy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTxRequest) ProtoMessage() {}

func (x *GetTxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_etherium_proxy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTxRequest.ProtoReflect.Descriptor instead.
func (*GetTxRequest) Descriptor() ([]byte, []int) {
	return file_api_etherium_proxy_proto_rawDescGZIP(), []int{2}
}

func (x *GetTxRequest) GetBlockId() string {
	if x != nil {
		return x.BlockId
	}
	return ""
}

func (x *GetTxRequest) GetTxId() string {
	if x != nil {
		return x.TxId
	}
	return ""
}

type GetTxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (x *GetTxResponse) Reset() {
	*x = GetTxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_etherium_proxy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTxResponse) ProtoMessage() {}

func (x *GetTxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_etherium_proxy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTxResponse.ProtoReflect.Descriptor instead.
func (*GetTxResponse) Descriptor() ([]byte, []int) {
	return file_api_etherium_proxy_proto_rawDescGZIP(), []int{3}
}

func (x *GetTxResponse) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Difficulty       string         `protobuf:"bytes,1,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	ExtraData        string         `protobuf:"bytes,2,opt,name=extra_data,json=extraData,proto3" json:"extra_data,omitempty"`
	GasLimit         string         `protobuf:"bytes,3,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	GasUsed          string         `protobuf:"bytes,4,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
	Hash             string         `protobuf:"bytes,5,opt,name=hash,proto3" json:"hash,omitempty"`
	LogsBloom        string         `protobuf:"bytes,6,opt,name=logs_bloom,json=logsBloom,proto3" json:"logs_bloom,omitempty"`
	Miner            string         `protobuf:"bytes,7,opt,name=miner,proto3" json:"miner,omitempty"`
	MixHash          string         `protobuf:"bytes,8,opt,name=mix_hash,json=mixHash,proto3" json:"mix_hash,omitempty"`
	Nonce            string         `protobuf:"bytes,9,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Number           string         `protobuf:"bytes,10,opt,name=number,proto3" json:"number,omitempty"`
	ParentHash       string         `protobuf:"bytes,11,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty"`
	ReceiptsRoot     string         `protobuf:"bytes,12,opt,name=receipts_root,json=receiptsRoot,proto3" json:"receipts_root,omitempty"`
	Sha3Uncles       string         `protobuf:"bytes,13,opt,name=sha3_uncles,json=sha3Uncles,proto3" json:"sha3_uncles,omitempty"`
	Size             string         `protobuf:"bytes,14,opt,name=size,proto3" json:"size,omitempty"`
	StateRoot        string         `protobuf:"bytes,15,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty"`
	Timestamp        string         `protobuf:"bytes,16,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	TotalDifficulty  string         `protobuf:"bytes,17,opt,name=total_difficulty,json=totalDifficulty,proto3" json:"total_difficulty,omitempty"`
	Transactions     []*Transaction `protobuf:"bytes,18,rep,name=transactions,proto3" json:"transactions,omitempty"`
	TransactionsRoot string         `protobuf:"bytes,19,opt,name=transactions_root,json=transactionsRoot,proto3" json:"transactions_root,omitempty"`
	Uncles           []string       `protobuf:"bytes,20,rep,name=uncles,proto3" json:"uncles,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_etherium_proxy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_api_etherium_proxy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_api_etherium_proxy_proto_rawDescGZIP(), []int{4}
}

func (x *Block) GetDifficulty() string {
	if x != nil {
		return x.Difficulty
	}
	return ""
}

func (x *Block) GetExtraData() string {
	if x != nil {
		return x.ExtraData
	}
	return ""
}

func (x *Block) GetGasLimit() string {
	if x != nil {
		return x.GasLimit
	}
	return ""
}

func (x *Block) GetGasUsed() string {
	if x != nil {
		return x.GasUsed
	}
	return ""
}

func (x *Block) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Block) GetLogsBloom() string {
	if x != nil {
		return x.LogsBloom
	}
	return ""
}

func (x *Block) GetMiner() string {
	if x != nil {
		return x.Miner
	}
	return ""
}

func (x *Block) GetMixHash() string {
	if x != nil {
		return x.MixHash
	}
	return ""
}

func (x *Block) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *Block) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *Block) GetParentHash() string {
	if x != nil {
		return x.ParentHash
	}
	return ""
}

func (x *Block) GetReceiptsRoot() string {
	if x != nil {
		return x.ReceiptsRoot
	}
	return ""
}

func (x *Block) GetSha3Uncles() string {
	if x != nil {
		return x.Sha3Uncles
	}
	return ""
}

func (x *Block) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

func (x *Block) GetStateRoot() string {
	if x != nil {
		return x.StateRoot
	}
	return ""
}

func (x *Block) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *Block) GetTotalDifficulty() string {
	if x != nil {
		return x.TotalDifficulty
	}
	return ""
}

func (x *Block) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *Block) GetTransactionsRoot() string {
	if x != nil {
		return x.TransactionsRoot
	}
	return ""
}

func (x *Block) GetUncles() []string {
	if x != nil {
		return x.Uncles
	}
	return nil
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockHash        string `protobuf:"bytes,1,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	BlockNumber      string `protobuf:"bytes,2,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	From             string `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	Gas              string `protobuf:"bytes,4,opt,name=gas,proto3" json:"gas,omitempty"`
	GasPrice         string `protobuf:"bytes,5,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	Hash             string `protobuf:"bytes,6,opt,name=hash,proto3" json:"hash,omitempty"`
	Input            string `protobuf:"bytes,7,opt,name=input,proto3" json:"input,omitempty"`
	Nonce            string `protobuf:"bytes,8,opt,name=nonce,proto3" json:"nonce,omitempty"`
	To               string `protobuf:"bytes,9,opt,name=to,proto3" json:"to,omitempty"`
	TransactionIndex string `protobuf:"bytes,10,opt,name=transaction_index,json=transactionIndex,proto3" json:"transaction_index,omitempty"`
	Value            string `protobuf:"bytes,11,opt,name=value,proto3" json:"value,omitempty"`
	V                string `protobuf:"bytes,12,opt,name=v,proto3" json:"v,omitempty"`
	R                string `protobuf:"bytes,13,opt,name=r,proto3" json:"r,omitempty"`
	S                string `protobuf:"bytes,14,opt,name=s,proto3" json:"s,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_etherium_proxy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_api_etherium_proxy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_api_etherium_proxy_proto_rawDescGZIP(), []int{5}
}

func (x *Transaction) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *Transaction) GetBlockNumber() string {
	if x != nil {
		return x.BlockNumber
	}
	return ""
}

func (x *Transaction) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Transaction) GetGas() string {
	if x != nil {
		return x.Gas
	}
	return ""
}

func (x *Transaction) GetGasPrice() string {
	if x != nil {
		return x.GasPrice
	}
	return ""
}

func (x *Transaction) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Transaction) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *Transaction) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *Transaction) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Transaction) GetTransactionIndex() string {
	if x != nil {
		return x.TransactionIndex
	}
	return ""
}

func (x *Transaction) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Transaction) GetV() string {
	if x != nil {
		return x.V
	}
	return ""
}

func (x *Transaction) GetR() string {
	if x != nil {
		return x.R
	}
	return ""
}

func (x *Transaction) GetS() string {
	if x != nil {
		return x.S
	}
	return ""
}

var File_api_etherium_proxy_proto protoreflect.FileDescriptor

var file_api_etherium_proxy_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x69, 0x75, 0x6d, 0x2d, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x65, 0x74, 0x68, 0x5f,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42,
	0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x18, 0x47, 0x65, 0x74,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x74, 0x68, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x3e, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x13, 0x0a, 0x05, 0x74, 0x78, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x49, 0x64, 0x22, 0x49, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x54, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38,
	0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x74, 0x68, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xf4, 0x04, 0x0a, 0x05, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c,
	0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x74, 0x72, 0x61, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x67, 0x61, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x67, 0x61, 0x73, 0x55, 0x73, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x1d, 0x0a,
	0x0a, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x62, 0x6c, 0x6f, 0x6f, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x73, 0x42, 0x6c, 0x6f, 0x6f, 0x6d, 0x12, 0x14, 0x0a, 0x05,
	0x6d, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x69, 0x6e,
	0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x69, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x69, 0x78, 0x48, 0x61, 0x73, 0x68, 0x12, 0x14, 0x0a,
	0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f,
	0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x73, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x73, 0x52, 0x6f, 0x6f,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x61, 0x33, 0x5f, 0x75, 0x6e, 0x63, 0x6c, 0x65, 0x73,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x68, 0x61, 0x33, 0x55, 0x6e, 0x63, 0x6c,
	0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f,
	0x72, 0x6f, 0x6f, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x66,
	0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x44, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x3a,
	0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x12,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x74, 0x68, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x6e, 0x63, 0x6c, 0x65,
	0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x75, 0x6e, 0x63, 0x6c, 0x65, 0x73, 0x22,
	0xcf, 0x02, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x21,
	0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x61, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x67, 0x61, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x73, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x61, 0x73, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e,
	0x6f, 0x6e, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x74, 0x6f, 0x12, 0x2b, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x01, 0x76, 0x12, 0x0c, 0x0a, 0x01, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x01, 0x72, 0x12, 0x0c, 0x0a, 0x01, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01,
	0x73, 0x32, 0xed, 0x01, 0x0a, 0x0a, 0x45, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x79, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x79, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x65, 0x74, 0x68, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x65, 0x74, 0x68, 0x5f, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x79, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x2f, 0x7b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x64, 0x0a, 0x05, 0x47,
	0x65, 0x74, 0x54, 0x78, 0x12, 0x17, 0x2e, 0x65, 0x74, 0x68, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x65, 0x74, 0x68, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x78, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12,
	0x20, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x7b, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x74, 0x78, 0x73, 0x2f, 0x7b, 0x74, 0x78, 0x5f, 0x69, 0x64,
	0x7d, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_etherium_proxy_proto_rawDescOnce sync.Once
	file_api_etherium_proxy_proto_rawDescData = file_api_etherium_proxy_proto_rawDesc
)

func file_api_etherium_proxy_proto_rawDescGZIP() []byte {
	file_api_etherium_proxy_proto_rawDescOnce.Do(func() {
		file_api_etherium_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_etherium_proxy_proto_rawDescData)
	})
	return file_api_etherium_proxy_proto_rawDescData
}

var file_api_etherium_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_etherium_proxy_proto_goTypes = []interface{}{
	(*GetBlockByNumberRequest)(nil),  // 0: eth_proxy.GetBlockByNumberRequest
	(*GetBlockByNumberResponse)(nil), // 1: eth_proxy.GetBlockByNumberResponse
	(*GetTxRequest)(nil),             // 2: eth_proxy.GetTxRequest
	(*GetTxResponse)(nil),            // 3: eth_proxy.GetTxResponse
	(*Block)(nil),                    // 4: eth_proxy.Block
	(*Transaction)(nil),              // 5: eth_proxy.Transaction
}
var file_api_etherium_proxy_proto_depIdxs = []int32{
	4, // 0: eth_proxy.GetBlockByNumberResponse.block:type_name -> eth_proxy.Block
	5, // 1: eth_proxy.GetTxResponse.transaction:type_name -> eth_proxy.Transaction
	5, // 2: eth_proxy.Block.transactions:type_name -> eth_proxy.Transaction
	0, // 3: eth_proxy.EthService.GetBlockByNumber:input_type -> eth_proxy.GetBlockByNumberRequest
	2, // 4: eth_proxy.EthService.GetTx:input_type -> eth_proxy.GetTxRequest
	1, // 5: eth_proxy.EthService.GetBlockByNumber:output_type -> eth_proxy.GetBlockByNumberResponse
	3, // 6: eth_proxy.EthService.GetTx:output_type -> eth_proxy.GetTxResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_etherium_proxy_proto_init() }
func file_api_etherium_proxy_proto_init() {
	if File_api_etherium_proxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_etherium_proxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockByNumberRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_etherium_proxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockByNumberResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_etherium_proxy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTxRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_etherium_proxy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTxResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_etherium_proxy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_etherium_proxy_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_etherium_proxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_etherium_proxy_proto_goTypes,
		DependencyIndexes: file_api_etherium_proxy_proto_depIdxs,
		MessageInfos:      file_api_etherium_proxy_proto_msgTypes,
	}.Build()
	File_api_etherium_proxy_proto = out.File
	file_api_etherium_proxy_proto_rawDesc = nil
	file_api_etherium_proxy_proto_goTypes = nil
	file_api_etherium_proxy_proto_depIdxs = nil
}
