// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: protos/model/v1/file.proto

#include "protos/model/v1/file.pb.h"

#include <algorithm>

#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/extension_set.h>
#include <google/protobuf/wire_format_lite.h>
#include <google/protobuf/descriptor.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/reflection_ops.h>
#include <google/protobuf/wire_format.h>
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>

PROTOBUF_PRAGMA_INIT_SEG

namespace _pb = ::PROTOBUF_NAMESPACE_ID;
namespace _pbi = _pb::internal;

namespace v1 {
namespace model {
PROTOBUF_CONSTEXPR File::File(
    ::_pbi::ConstantInitialized)
  : filename_(&::_pbi::fixed_address_empty_string, ::_pbi::ConstantInitialized{})
  , _oneof_case_{}{}
struct FileDefaultTypeInternal {
  PROTOBUF_CONSTEXPR FileDefaultTypeInternal()
      : _instance(::_pbi::ConstantInitialized{}) {}
  ~FileDefaultTypeInternal() {}
  union {
    File _instance;
  };
};
PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 FileDefaultTypeInternal _File_default_instance_;
}  // namespace model
}  // namespace v1
static ::_pb::Metadata file_level_metadata_protos_2fmodel_2fv1_2ffile_2eproto[1];
static constexpr ::_pb::EnumDescriptor const** file_level_enum_descriptors_protos_2fmodel_2fv1_2ffile_2eproto = nullptr;
static constexpr ::_pb::ServiceDescriptor const** file_level_service_descriptors_protos_2fmodel_2fv1_2ffile_2eproto = nullptr;

const uint32_t TableStruct_protos_2fmodel_2fv1_2ffile_2eproto::offsets[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  ~0u,  // no _has_bits_
  PROTOBUF_FIELD_OFFSET(::v1::model::File, _internal_metadata_),
  ~0u,  // no _extensions_
  PROTOBUF_FIELD_OFFSET(::v1::model::File, _oneof_case_[0]),
  ~0u,  // no _weak_field_map_
  ~0u,  // no _inlined_string_donated_
  ::_pbi::kInvalidFieldOffsetTag,
  ::_pbi::kInvalidFieldOffsetTag,
  PROTOBUF_FIELD_OFFSET(::v1::model::File, filename_),
  PROTOBUF_FIELD_OFFSET(::v1::model::File, data_),
};
static const ::_pbi::MigrationSchema schemas[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  { 0, -1, -1, sizeof(::v1::model::File)},
};

static const ::_pb::Message* const file_default_instances[] = {
  &::v1::model::_File_default_instance_._instance,
};

const char descriptor_table_protodef_protos_2fmodel_2fv1_2ffile_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) =
  "\n\032protos/model/v1/file.proto\022\010v1.model\"X"
  "\n\004File\022\022\n\003url\030\001 \001(\tH\000R\003url\022\022\n\003raw\030\002 \001(\014H"
  "\000R\003raw\022\032\n\010filename\030\003 \001(\tR\010filenameB\006\n\004da"
  "taJ\004\010\004\020\005B+Z)github.com/FormantIO/genprot"
  "o/go/v1/modelb\006proto3"
  ;
static ::_pbi::once_flag descriptor_table_protos_2fmodel_2fv1_2ffile_2eproto_once;
const ::_pbi::DescriptorTable descriptor_table_protos_2fmodel_2fv1_2ffile_2eproto = {
    false, false, 181, descriptor_table_protodef_protos_2fmodel_2fv1_2ffile_2eproto,
    "protos/model/v1/file.proto",
    &descriptor_table_protos_2fmodel_2fv1_2ffile_2eproto_once, nullptr, 0, 1,
    schemas, file_default_instances, TableStruct_protos_2fmodel_2fv1_2ffile_2eproto::offsets,
    file_level_metadata_protos_2fmodel_2fv1_2ffile_2eproto, file_level_enum_descriptors_protos_2fmodel_2fv1_2ffile_2eproto,
    file_level_service_descriptors_protos_2fmodel_2fv1_2ffile_2eproto,
};
PROTOBUF_ATTRIBUTE_WEAK const ::_pbi::DescriptorTable* descriptor_table_protos_2fmodel_2fv1_2ffile_2eproto_getter() {
  return &descriptor_table_protos_2fmodel_2fv1_2ffile_2eproto;
}

// Force running AddDescriptors() at dynamic initialization time.
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2 static ::_pbi::AddDescriptorsRunner dynamic_init_dummy_protos_2fmodel_2fv1_2ffile_2eproto(&descriptor_table_protos_2fmodel_2fv1_2ffile_2eproto);
namespace v1 {
namespace model {

// ===================================================================

class File::_Internal {
 public:
};

File::File(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                         bool is_message_owned)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena, is_message_owned) {
  SharedCtor();
  // @@protoc_insertion_point(arena_constructor:v1.model.File)
}
File::File(const File& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  filename_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
    filename_.Set("", GetArenaForAllocation());
  #endif // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  if (!from._internal_filename().empty()) {
    filename_.Set(from._internal_filename(), 
      GetArenaForAllocation());
  }
  clear_has_data();
  switch (from.data_case()) {
    case kUrl: {
      _internal_set_url(from._internal_url());
      break;
    }
    case kRaw: {
      _internal_set_raw(from._internal_raw());
      break;
    }
    case DATA_NOT_SET: {
      break;
    }
  }
  // @@protoc_insertion_point(copy_constructor:v1.model.File)
}

inline void File::SharedCtor() {
filename_.InitDefault();
#ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
  filename_.Set("", GetArenaForAllocation());
#endif // PROTOBUF_FORCE_COPY_DEFAULT_STRING
clear_has_data();
}

File::~File() {
  // @@protoc_insertion_point(destructor:v1.model.File)
  if (auto *arena = _internal_metadata_.DeleteReturnArena<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>()) {
  (void)arena;
    return;
  }
  SharedDtor();
}

inline void File::SharedDtor() {
  GOOGLE_DCHECK(GetArenaForAllocation() == nullptr);
  filename_.Destroy();
  if (has_data()) {
    clear_data();
  }
}

void File::SetCachedSize(int size) const {
  _cached_size_.Set(size);
}

void File::clear_data() {
// @@protoc_insertion_point(one_of_clear_start:v1.model.File)
  switch (data_case()) {
    case kUrl: {
      data_.url_.Destroy();
      break;
    }
    case kRaw: {
      data_.raw_.Destroy();
      break;
    }
    case DATA_NOT_SET: {
      break;
    }
  }
  _oneof_case_[0] = DATA_NOT_SET;
}


void File::Clear() {
// @@protoc_insertion_point(message_clear_start:v1.model.File)
  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  filename_.ClearToEmpty();
  clear_data();
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* File::_InternalParse(const char* ptr, ::_pbi::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  while (!ctx->Done(&ptr)) {
    uint32_t tag;
    ptr = ::_pbi::ReadTag(ptr, &tag);
    switch (tag >> 3) {
      // string url = 1 [json_name = "url"];
      case 1:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 10)) {
          auto str = _internal_mutable_url();
          ptr = ::_pbi::InlineGreedyStringParser(str, ptr, ctx);
          CHK_(ptr);
          CHK_(::_pbi::VerifyUTF8(str, "v1.model.File.url"));
        } else
          goto handle_unusual;
        continue;
      // bytes raw = 2 [json_name = "raw"];
      case 2:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 18)) {
          auto str = _internal_mutable_raw();
          ptr = ::_pbi::InlineGreedyStringParser(str, ptr, ctx);
          CHK_(ptr);
        } else
          goto handle_unusual;
        continue;
      // string filename = 3 [json_name = "filename"];
      case 3:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 26)) {
          auto str = _internal_mutable_filename();
          ptr = ::_pbi::InlineGreedyStringParser(str, ptr, ctx);
          CHK_(ptr);
          CHK_(::_pbi::VerifyUTF8(str, "v1.model.File.filename"));
        } else
          goto handle_unusual;
        continue;
      default:
        goto handle_unusual;
    }  // switch
  handle_unusual:
    if ((tag == 0) || ((tag & 7) == 4)) {
      CHK_(ptr);
      ctx->SetLastTag(tag);
      goto message_done;
    }
    ptr = UnknownFieldParse(
        tag,
        _internal_metadata_.mutable_unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(),
        ptr, ctx);
    CHK_(ptr != nullptr);
  }  // while
message_done:
  return ptr;
failure:
  ptr = nullptr;
  goto message_done;
#undef CHK_
}

uint8_t* File::_InternalSerialize(
    uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:v1.model.File)
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  // string url = 1 [json_name = "url"];
  if (_internal_has_url()) {
    ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::VerifyUtf8String(
      this->_internal_url().data(), static_cast<int>(this->_internal_url().length()),
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::SERIALIZE,
      "v1.model.File.url");
    target = stream->WriteStringMaybeAliased(
        1, this->_internal_url(), target);
  }

  // bytes raw = 2 [json_name = "raw"];
  if (_internal_has_raw()) {
    target = stream->WriteBytesMaybeAliased(
        2, this->_internal_raw(), target);
  }

  // string filename = 3 [json_name = "filename"];
  if (!this->_internal_filename().empty()) {
    ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::VerifyUtf8String(
      this->_internal_filename().data(), static_cast<int>(this->_internal_filename().length()),
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::SERIALIZE,
      "v1.model.File.filename");
    target = stream->WriteStringMaybeAliased(
        3, this->_internal_filename(), target);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:v1.model.File)
  return target;
}

size_t File::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:v1.model.File)
  size_t total_size = 0;

  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // string filename = 3 [json_name = "filename"];
  if (!this->_internal_filename().empty()) {
    total_size += 1 +
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::StringSize(
        this->_internal_filename());
  }

  switch (data_case()) {
    // string url = 1 [json_name = "url"];
    case kUrl: {
      total_size += 1 +
        ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::StringSize(
          this->_internal_url());
      break;
    }
    // bytes raw = 2 [json_name = "raw"];
    case kRaw: {
      total_size += 1 +
        ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::BytesSize(
          this->_internal_raw());
      break;
    }
    case DATA_NOT_SET: {
      break;
    }
  }
  return MaybeComputeUnknownFieldsSize(total_size, &_cached_size_);
}

const ::PROTOBUF_NAMESPACE_ID::Message::ClassData File::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::Message::CopyWithSizeCheck,
    File::MergeImpl
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*File::GetClassData() const { return &_class_data_; }

void File::MergeImpl(::PROTOBUF_NAMESPACE_ID::Message* to,
                      const ::PROTOBUF_NAMESPACE_ID::Message& from) {
  static_cast<File *>(to)->MergeFrom(
      static_cast<const File &>(from));
}


void File::MergeFrom(const File& from) {
// @@protoc_insertion_point(class_specific_merge_from_start:v1.model.File)
  GOOGLE_DCHECK_NE(&from, this);
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  if (!from._internal_filename().empty()) {
    _internal_set_filename(from._internal_filename());
  }
  switch (from.data_case()) {
    case kUrl: {
      _internal_set_url(from._internal_url());
      break;
    }
    case kRaw: {
      _internal_set_raw(from._internal_raw());
      break;
    }
    case DATA_NOT_SET: {
      break;
    }
  }
  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
}

void File::CopyFrom(const File& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:v1.model.File)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool File::IsInitialized() const {
  return true;
}

void File::InternalSwap(File* other) {
  using std::swap;
  auto* lhs_arena = GetArenaForAllocation();
  auto* rhs_arena = other->GetArenaForAllocation();
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  ::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr::InternalSwap(
      &filename_, lhs_arena,
      &other->filename_, rhs_arena
  );
  swap(data_, other->data_);
  swap(_oneof_case_[0], other->_oneof_case_[0]);
}

::PROTOBUF_NAMESPACE_ID::Metadata File::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_protos_2fmodel_2fv1_2ffile_2eproto_getter, &descriptor_table_protos_2fmodel_2fv1_2ffile_2eproto_once,
      file_level_metadata_protos_2fmodel_2fv1_2ffile_2eproto[0]);
}

// @@protoc_insertion_point(namespace_scope)
}  // namespace model
}  // namespace v1
PROTOBUF_NAMESPACE_OPEN
template<> PROTOBUF_NOINLINE ::v1::model::File*
Arena::CreateMaybeMessage< ::v1::model::File >(Arena* arena) {
  return Arena::CreateMessageInternal< ::v1::model::File >(arena);
}
PROTOBUF_NAMESPACE_CLOSE

// @@protoc_insertion_point(global_scope)
#include <google/protobuf/port_undef.inc>
