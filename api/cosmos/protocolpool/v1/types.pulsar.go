// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package protocolpoolv1

import (
	v1beta1 "cosmossdk.io/api/cosmos/base/v1beta1"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_Budget                   protoreflect.MessageDescriptor
	fd_Budget_recipient_address protoreflect.FieldDescriptor
	fd_Budget_total_budget      protoreflect.FieldDescriptor
	fd_Budget_start_time        protoreflect.FieldDescriptor
	fd_Budget_tranches          protoreflect.FieldDescriptor
	fd_Budget_period            protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_protocolpool_v1_types_proto_init()
	md_Budget = File_cosmos_protocolpool_v1_types_proto.Messages().ByName("Budget")
	fd_Budget_recipient_address = md_Budget.Fields().ByName("recipient_address")
	fd_Budget_total_budget = md_Budget.Fields().ByName("total_budget")
	fd_Budget_start_time = md_Budget.Fields().ByName("start_time")
	fd_Budget_tranches = md_Budget.Fields().ByName("tranches")
	fd_Budget_period = md_Budget.Fields().ByName("period")
}

var _ protoreflect.Message = (*fastReflection_Budget)(nil)

type fastReflection_Budget Budget

func (x *Budget) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Budget)(x)
}

func (x *Budget) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_protocolpool_v1_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Budget_messageType fastReflection_Budget_messageType
var _ protoreflect.MessageType = fastReflection_Budget_messageType{}

type fastReflection_Budget_messageType struct{}

func (x fastReflection_Budget_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Budget)(nil)
}
func (x fastReflection_Budget_messageType) New() protoreflect.Message {
	return new(fastReflection_Budget)
}
func (x fastReflection_Budget_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Budget
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Budget) Descriptor() protoreflect.MessageDescriptor {
	return md_Budget
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Budget) Type() protoreflect.MessageType {
	return _fastReflection_Budget_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Budget) New() protoreflect.Message {
	return new(fastReflection_Budget)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Budget) Interface() protoreflect.ProtoMessage {
	return (*Budget)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Budget) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.RecipientAddress != "" {
		value := protoreflect.ValueOfString(x.RecipientAddress)
		if !f(fd_Budget_recipient_address, value) {
			return
		}
	}
	if x.TotalBudget != nil {
		value := protoreflect.ValueOfMessage(x.TotalBudget.ProtoReflect())
		if !f(fd_Budget_total_budget, value) {
			return
		}
	}
	if x.StartTime != int64(0) {
		value := protoreflect.ValueOfInt64(x.StartTime)
		if !f(fd_Budget_start_time, value) {
			return
		}
	}
	if x.Tranches != int64(0) {
		value := protoreflect.ValueOfInt64(x.Tranches)
		if !f(fd_Budget_tranches, value) {
			return
		}
	}
	if x.Period != int64(0) {
		value := protoreflect.ValueOfInt64(x.Period)
		if !f(fd_Budget_period, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Budget) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.protocolpool.v1.Budget.recipient_address":
		return x.RecipientAddress != ""
	case "cosmos.protocolpool.v1.Budget.total_budget":
		return x.TotalBudget != nil
	case "cosmos.protocolpool.v1.Budget.start_time":
		return x.StartTime != int64(0)
	case "cosmos.protocolpool.v1.Budget.tranches":
		return x.Tranches != int64(0)
	case "cosmos.protocolpool.v1.Budget.period":
		return x.Period != int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.protocolpool.v1.Budget"))
		}
		panic(fmt.Errorf("message cosmos.protocolpool.v1.Budget does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Budget) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.protocolpool.v1.Budget.recipient_address":
		x.RecipientAddress = ""
	case "cosmos.protocolpool.v1.Budget.total_budget":
		x.TotalBudget = nil
	case "cosmos.protocolpool.v1.Budget.start_time":
		x.StartTime = int64(0)
	case "cosmos.protocolpool.v1.Budget.tranches":
		x.Tranches = int64(0)
	case "cosmos.protocolpool.v1.Budget.period":
		x.Period = int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.protocolpool.v1.Budget"))
		}
		panic(fmt.Errorf("message cosmos.protocolpool.v1.Budget does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Budget) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.protocolpool.v1.Budget.recipient_address":
		value := x.RecipientAddress
		return protoreflect.ValueOfString(value)
	case "cosmos.protocolpool.v1.Budget.total_budget":
		value := x.TotalBudget
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.protocolpool.v1.Budget.start_time":
		value := x.StartTime
		return protoreflect.ValueOfInt64(value)
	case "cosmos.protocolpool.v1.Budget.tranches":
		value := x.Tranches
		return protoreflect.ValueOfInt64(value)
	case "cosmos.protocolpool.v1.Budget.period":
		value := x.Period
		return protoreflect.ValueOfInt64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.protocolpool.v1.Budget"))
		}
		panic(fmt.Errorf("message cosmos.protocolpool.v1.Budget does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Budget) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.protocolpool.v1.Budget.recipient_address":
		x.RecipientAddress = value.Interface().(string)
	case "cosmos.protocolpool.v1.Budget.total_budget":
		x.TotalBudget = value.Message().Interface().(*v1beta1.Coin)
	case "cosmos.protocolpool.v1.Budget.start_time":
		x.StartTime = value.Int()
	case "cosmos.protocolpool.v1.Budget.tranches":
		x.Tranches = value.Int()
	case "cosmos.protocolpool.v1.Budget.period":
		x.Period = value.Int()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.protocolpool.v1.Budget"))
		}
		panic(fmt.Errorf("message cosmos.protocolpool.v1.Budget does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Budget) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.protocolpool.v1.Budget.total_budget":
		if x.TotalBudget == nil {
			x.TotalBudget = new(v1beta1.Coin)
		}
		return protoreflect.ValueOfMessage(x.TotalBudget.ProtoReflect())
	case "cosmos.protocolpool.v1.Budget.recipient_address":
		panic(fmt.Errorf("field recipient_address of message cosmos.protocolpool.v1.Budget is not mutable"))
	case "cosmos.protocolpool.v1.Budget.start_time":
		panic(fmt.Errorf("field start_time of message cosmos.protocolpool.v1.Budget is not mutable"))
	case "cosmos.protocolpool.v1.Budget.tranches":
		panic(fmt.Errorf("field tranches of message cosmos.protocolpool.v1.Budget is not mutable"))
	case "cosmos.protocolpool.v1.Budget.period":
		panic(fmt.Errorf("field period of message cosmos.protocolpool.v1.Budget is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.protocolpool.v1.Budget"))
		}
		panic(fmt.Errorf("message cosmos.protocolpool.v1.Budget does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Budget) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.protocolpool.v1.Budget.recipient_address":
		return protoreflect.ValueOfString("")
	case "cosmos.protocolpool.v1.Budget.total_budget":
		m := new(v1beta1.Coin)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.protocolpool.v1.Budget.start_time":
		return protoreflect.ValueOfInt64(int64(0))
	case "cosmos.protocolpool.v1.Budget.tranches":
		return protoreflect.ValueOfInt64(int64(0))
	case "cosmos.protocolpool.v1.Budget.period":
		return protoreflect.ValueOfInt64(int64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.protocolpool.v1.Budget"))
		}
		panic(fmt.Errorf("message cosmos.protocolpool.v1.Budget does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Budget) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.protocolpool.v1.Budget", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Budget) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Budget) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Budget) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Budget) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Budget)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.RecipientAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.TotalBudget != nil {
			l = options.Size(x.TotalBudget)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.StartTime != 0 {
			n += 1 + runtime.Sov(uint64(x.StartTime))
		}
		if x.Tranches != 0 {
			n += 1 + runtime.Sov(uint64(x.Tranches))
		}
		if x.Period != 0 {
			n += 1 + runtime.Sov(uint64(x.Period))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Budget)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Period != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Period))
			i--
			dAtA[i] = 0x28
		}
		if x.Tranches != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Tranches))
			i--
			dAtA[i] = 0x20
		}
		if x.StartTime != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.StartTime))
			i--
			dAtA[i] = 0x18
		}
		if x.TotalBudget != nil {
			encoded, err := options.Marshal(x.TotalBudget)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.RecipientAddress) > 0 {
			i -= len(x.RecipientAddress)
			copy(dAtA[i:], x.RecipientAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RecipientAddress)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Budget)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Budget: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Budget: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RecipientAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.RecipientAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TotalBudget", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.TotalBudget == nil {
					x.TotalBudget = &v1beta1.Coin{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.TotalBudget); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
				}
				x.StartTime = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.StartTime |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Tranches", wireType)
				}
				x.Tranches = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Tranches |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 5:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Period", wireType)
				}
				x.Period = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Period |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: cosmos/protocolpool/v1/types.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Budget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// recipient_address is the address of the recipient who can claim the budget.
	RecipientAddress string `protobuf:"bytes,1,opt,name=recipient_address,json=recipientAddress,proto3" json:"recipient_address,omitempty"`
	// total_budget is the total amount allocated for the budget.
	TotalBudget *v1beta1.Coin `protobuf:"bytes,2,opt,name=total_budget,json=totalBudget,proto3" json:"total_budget,omitempty"`
	// start_time is the time when the budget becomes claimable, represented in seconds since the Unix epoch.
	StartTime int64 `protobuf:"varint,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// tranches is the number of times the total budget amount is to be distributed.
	Tranches int64 `protobuf:"varint,4,opt,name=tranches,proto3" json:"tranches,omitempty"`
	// period is the interval (in blocks) between each distribution of funds.
	Period int64 `protobuf:"varint,5,opt,name=period,proto3" json:"period,omitempty"`
}

func (x *Budget) Reset() {
	*x = Budget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_protocolpool_v1_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Budget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Budget) ProtoMessage() {}

// Deprecated: Use Budget.ProtoReflect.Descriptor instead.
func (*Budget) Descriptor() ([]byte, []int) {
	return file_cosmos_protocolpool_v1_types_proto_rawDescGZIP(), []int{0}
}

func (x *Budget) GetRecipientAddress() string {
	if x != nil {
		return x.RecipientAddress
	}
	return ""
}

func (x *Budget) GetTotalBudget() *v1beta1.Coin {
	if x != nil {
		return x.TotalBudget
	}
	return nil
}

func (x *Budget) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *Budget) GetTranches() int64 {
	if x != nil {
		return x.Tranches
	}
	return 0
}

func (x *Budget) GetPeriod() int64 {
	if x != nil {
		return x.Period
	}
	return 0
}

var File_cosmos_protocolpool_v1_types_proto protoreflect.FileDescriptor

var file_cosmos_protocolpool_v1_types_proto_rawDesc = []byte{
	0x0a, 0x22, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x69,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x01, 0x0a, 0x06, 0x42, 0x75, 0x64, 0x67,
	0x65, 0x74, 0x12, 0x45, 0x0a, 0x11, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2,
	0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x10, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65,
	0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3c, 0x0a, 0x0c, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x42, 0xda, 0x01, 0x0a, 0x1a, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73,
	0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x76, 0x31,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x70, 0x6f, 0x6f, 0x6c, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x43, 0x50, 0x58, 0xaa, 0x02, 0x16, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x16, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x70, 0x6f, 0x6f, 0x6c, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x22, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x5c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x70, 0x6f, 0x6f, 0x6c, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x18, 0x43,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x70,
	0x6f, 0x6f, 0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_protocolpool_v1_types_proto_rawDescOnce sync.Once
	file_cosmos_protocolpool_v1_types_proto_rawDescData = file_cosmos_protocolpool_v1_types_proto_rawDesc
)

func file_cosmos_protocolpool_v1_types_proto_rawDescGZIP() []byte {
	file_cosmos_protocolpool_v1_types_proto_rawDescOnce.Do(func() {
		file_cosmos_protocolpool_v1_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_protocolpool_v1_types_proto_rawDescData)
	})
	return file_cosmos_protocolpool_v1_types_proto_rawDescData
}

var file_cosmos_protocolpool_v1_types_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cosmos_protocolpool_v1_types_proto_goTypes = []interface{}{
	(*Budget)(nil),       // 0: cosmos.protocolpool.v1.Budget
	(*v1beta1.Coin)(nil), // 1: cosmos.base.v1beta1.Coin
}
var file_cosmos_protocolpool_v1_types_proto_depIdxs = []int32{
	1, // 0: cosmos.protocolpool.v1.Budget.total_budget:type_name -> cosmos.base.v1beta1.Coin
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cosmos_protocolpool_v1_types_proto_init() }
func file_cosmos_protocolpool_v1_types_proto_init() {
	if File_cosmos_protocolpool_v1_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_protocolpool_v1_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Budget); i {
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
			RawDescriptor: file_cosmos_protocolpool_v1_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_protocolpool_v1_types_proto_goTypes,
		DependencyIndexes: file_cosmos_protocolpool_v1_types_proto_depIdxs,
		MessageInfos:      file_cosmos_protocolpool_v1_types_proto_msgTypes,
	}.Build()
	File_cosmos_protocolpool_v1_types_proto = out.File
	file_cosmos_protocolpool_v1_types_proto_rawDesc = nil
	file_cosmos_protocolpool_v1_types_proto_goTypes = nil
	file_cosmos_protocolpool_v1_types_proto_depIdxs = nil
}
