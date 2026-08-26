// Code generated for linux/amd64 by 'ccgo -I/tmp/monocypher_sources /tmp/monocypher_sources/monocypher.c', DO NOT EDIT.

//go:build linux && amd64

package monocypher

import (
	"math"
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var _ = math.Pi
var _ reflect.Type
var _ unsafe.Pointer

const B_W_WIDTH = 5
const CRYPTO_ARGON2_D = 0
const CRYPTO_ARGON2_I = 1
const CRYPTO_ARGON2_ID = 2
const INT16_MAX = 0x7fff
const INT32_MAX = 0x7fffffff
const INT64_MAX = 0x7fffffffffffffff
const INT8_MAX = 0x7f
const INTMAX_MAX = "INT64_MAX"
const INTMAX_MIN = "INT64_MIN"
const INTPTR_MAX = "INT64_MAX"
const INTPTR_MIN = "INT64_MIN"
const INT_FAST16_MAX = "INT32_MAX"
const INT_FAST16_MIN = "INT32_MIN"
const INT_FAST32_MAX = "INT32_MAX"
const INT_FAST32_MIN = "INT32_MIN"
const INT_FAST64_MAX = "INT64_MAX"
const INT_FAST64_MIN = "INT64_MIN"
const INT_FAST8_MAX = "INT8_MAX"
const INT_FAST8_MIN = "INT8_MIN"
const INT_LEAST16_MAX = "INT16_MAX"
const INT_LEAST16_MIN = "INT16_MIN"
const INT_LEAST32_MAX = "INT32_MAX"
const INT_LEAST32_MIN = "INT32_MIN"
const INT_LEAST64_MAX = "INT64_MAX"
const INT_LEAST64_MIN = "INT64_MIN"
const INT_LEAST8_MAX = "INT8_MAX"
const INT_LEAST8_MIN = "INT8_MIN"
const PTRDIFF_MAX = "INT64_MAX"
const PTRDIFF_MIN = "INT64_MIN"
const P_W_WIDTH = 3
const SIG_ATOMIC_MAX = "INT32_MAX"
const SIG_ATOMIC_MIN = "INT32_MIN"
const SIZE_MAX = "UINT64_MAX"
const UINT16_MAX = 0xffff
const UINT32_MAX = "0xffffffffu"
const UINT64_MAX = "0xffffffffffffffffu"
const UINT8_MAX = 0xff
const UINTMAX_MAX = "UINT64_MAX"
const UINTPTR_MAX = "UINT64_MAX"
const UINT_FAST16_MAX = "UINT32_MAX"
const UINT_FAST32_MAX = "UINT32_MAX"
const UINT_FAST64_MAX = "UINT64_MAX"
const UINT_FAST8_MAX = "UINT8_MAX"
const UINT_LEAST16_MAX = "UINT16_MAX"
const UINT_LEAST32_MAX = "UINT32_MAX"
const UINT_LEAST64_MAX = "UINT64_MAX"
const UINT_LEAST8_MAX = "UINT8_MAX"
const WINT_MAX = "UINT32_MAX"
const WINT_MIN = 0
const _GNU_SOURCE = 1
const _LP64 = 1
const _STDC_PREDEF_H = 1
const __ATOMIC_ACQUIRE = 2
const __ATOMIC_ACQ_REL = 4
const __ATOMIC_CONSUME = 1
const __ATOMIC_HLE_ACQUIRE = 65536
const __ATOMIC_HLE_RELEASE = 131072
const __ATOMIC_RELAXED = 0
const __ATOMIC_RELEASE = 3
const __ATOMIC_SEQ_CST = 5
const __BFLT16_DECIMAL_DIG__ = 4
const __BFLT16_DENORM_MIN__ = "9.18354961579912115600575419704879436e-41B"
const __BFLT16_DIG__ = 2
const __BFLT16_EPSILON__ = "7.81250000000000000000000000000000000e-3B"
const __BFLT16_HAS_DENORM__ = 1
const __BFLT16_HAS_INFINITY__ = 1
const __BFLT16_HAS_QUIET_NAN__ = 1
const __BFLT16_IS_IEC_60559__ = 0
const __BFLT16_MANT_DIG__ = 8
const __BFLT16_MAX_10_EXP__ = 38
const __BFLT16_MAX_EXP__ = 128
const __BFLT16_MAX__ = "3.38953138925153547590470800371487867e+38B"
const __BFLT16_MIN__ = "1.17549435082228750796873653722224568e-38B"
const __BFLT16_NORM_MAX__ = "3.38953138925153547590470800371487867e+38B"
const __BIGGEST_ALIGNMENT__ = 16
const __BIG_ENDIAN = 4321
const __BYTE_ORDER = 1234
const __BYTE_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const __CCGO__ = 1
const __CET__ = 3
const __CHAR_BIT__ = 8
const __DBL_DECIMAL_DIG__ = 17
const __DBL_DIG__ = 15
const __DBL_HAS_DENORM__ = 1
const __DBL_HAS_INFINITY__ = 1
const __DBL_HAS_QUIET_NAN__ = 1
const __DBL_IS_IEC_60559__ = 1
const __DBL_MANT_DIG__ = 53
const __DBL_MAX_10_EXP__ = 308
const __DBL_MAX_EXP__ = 1024
const __DEC128_EPSILON__ = 1e-33
const __DEC128_MANT_DIG__ = 34
const __DEC128_MAX_EXP__ = 6145
const __DEC128_MAX__ = "9.999999999999999999999999999999999E6144"
const __DEC128_MIN__ = 1e-6143
const __DEC128_SUBNORMAL_MIN__ = 0.000000000000000000000000000000001e-6143
const __DEC32_EPSILON__ = 1e-6
const __DEC32_MANT_DIG__ = 7
const __DEC32_MAX_EXP__ = 97
const __DEC32_MAX__ = 9.999999e96
const __DEC32_MIN__ = 1e-95
const __DEC32_SUBNORMAL_MIN__ = 0.000001e-95
const __DEC64_EPSILON__ = 1e-15
const __DEC64_MANT_DIG__ = 16
const __DEC64_MAX_EXP__ = 385
const __DEC64_MAX__ = "9.999999999999999E384"
const __DEC64_MIN__ = 1e-383
const __DEC64_SUBNORMAL_MIN__ = 0.000000000000001e-383
const __DECIMAL_BID_FORMAT__ = 1
const __DECIMAL_DIG__ = 17
const __DEC_EVAL_METHOD__ = 2
const __ELF__ = 1
const __FINITE_MATH_ONLY__ = 0
const __FLOAT_WORD_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const __FLT128_DECIMAL_DIG__ = 36
const __FLT128_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966
const __FLT128_DIG__ = 33
const __FLT128_EPSILON__ = 1.92592994438723585305597794258492732e-34
const __FLT128_HAS_DENORM__ = 1
const __FLT128_HAS_INFINITY__ = 1
const __FLT128_HAS_QUIET_NAN__ = 1
const __FLT128_IS_IEC_60559__ = 1
const __FLT128_MANT_DIG__ = 113
const __FLT128_MAX_10_EXP__ = 4932
const __FLT128_MAX_EXP__ = 16384
const __FLT128_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT128_MIN__ = 3.36210314311209350626267781732175260e-4932
const __FLT128_NORM_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT16_DECIMAL_DIG__ = 5
const __FLT16_DENORM_MIN__ = 5.96046447753906250000000000000000000e-8
const __FLT16_DIG__ = 3
const __FLT16_EPSILON__ = 9.76562500000000000000000000000000000e-4
const __FLT16_HAS_DENORM__ = 1
const __FLT16_HAS_INFINITY__ = 1
const __FLT16_HAS_QUIET_NAN__ = 1
const __FLT16_IS_IEC_60559__ = 1
const __FLT16_MANT_DIG__ = 11
const __FLT16_MAX_10_EXP__ = 4
const __FLT16_MAX_EXP__ = 16
const __FLT16_MAX__ = 6.55040000000000000000000000000000000e+4
const __FLT16_MIN__ = 6.10351562500000000000000000000000000e-5
const __FLT16_NORM_MAX__ = 6.55040000000000000000000000000000000e+4
const __FLT32X_DECIMAL_DIG__ = 17
const __FLT32X_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const __FLT32X_DIG__ = 15
const __FLT32X_EPSILON__ = 2.22044604925031308084726333618164062e-16
const __FLT32X_HAS_DENORM__ = 1
const __FLT32X_HAS_INFINITY__ = 1
const __FLT32X_HAS_QUIET_NAN__ = 1
const __FLT32X_IS_IEC_60559__ = 1
const __FLT32X_MANT_DIG__ = 53
const __FLT32X_MAX_10_EXP__ = 308
const __FLT32X_MAX_EXP__ = 1024
const __FLT32X_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT32X_MIN__ = 2.22507385850720138309023271733240406e-308
const __FLT32X_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT32_DECIMAL_DIG__ = 9
const __FLT32_DENORM_MIN__ = 1.40129846432481707092372958328991613e-45
const __FLT32_DIG__ = 6
const __FLT32_EPSILON__ = 1.19209289550781250000000000000000000e-7
const __FLT32_HAS_DENORM__ = 1
const __FLT32_HAS_INFINITY__ = 1
const __FLT32_HAS_QUIET_NAN__ = 1
const __FLT32_IS_IEC_60559__ = 1
const __FLT32_MANT_DIG__ = 24
const __FLT32_MAX_10_EXP__ = 38
const __FLT32_MAX_EXP__ = 128
const __FLT32_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT32_MIN__ = 1.17549435082228750796873653722224568e-38
const __FLT32_NORM_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT64X_DECIMAL_DIG__ = 36
const __FLT64X_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966
const __FLT64X_DIG__ = 33
const __FLT64X_EPSILON__ = 1.92592994438723585305597794258492732e-34
const __FLT64X_HAS_DENORM__ = 1
const __FLT64X_HAS_INFINITY__ = 1
const __FLT64X_HAS_QUIET_NAN__ = 1
const __FLT64X_IS_IEC_60559__ = 1
const __FLT64X_MANT_DIG__ = 113
const __FLT64X_MAX_10_EXP__ = 4932
const __FLT64X_MAX_EXP__ = 16384
const __FLT64X_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT64X_MIN__ = 3.36210314311209350626267781732175260e-4932
const __FLT64X_NORM_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT64_DECIMAL_DIG__ = 17
const __FLT64_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const __FLT64_DIG__ = 15
const __FLT64_EPSILON__ = 2.22044604925031308084726333618164062e-16
const __FLT64_HAS_DENORM__ = 1
const __FLT64_HAS_INFINITY__ = 1
const __FLT64_HAS_QUIET_NAN__ = 1
const __FLT64_IS_IEC_60559__ = 1
const __FLT64_MANT_DIG__ = 53
const __FLT64_MAX_10_EXP__ = 308
const __FLT64_MAX_EXP__ = 1024
const __FLT64_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT64_MIN__ = 2.22507385850720138309023271733240406e-308
const __FLT64_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT_DECIMAL_DIG__ = 9
const __FLT_DENORM_MIN__ = 1.40129846432481707092372958328991613e-45
const __FLT_DIG__ = 6
const __FLT_EPSILON__ = 1.19209289550781250000000000000000000e-7
const __FLT_EVAL_METHOD_TS_18661_3__ = 0
const __FLT_EVAL_METHOD__ = 0
const __FLT_HAS_DENORM__ = 1
const __FLT_HAS_INFINITY__ = 1
const __FLT_HAS_QUIET_NAN__ = 1
const __FLT_IS_IEC_60559__ = 1
const __FLT_MANT_DIG__ = 24
const __FLT_MAX_10_EXP__ = 38
const __FLT_MAX_EXP__ = 128
const __FLT_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT_MIN__ = 1.17549435082228750796873653722224568e-38
const __FLT_NORM_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT_RADIX__ = 2
const __FUNCTION__ = "__func__"
const __FXSR__ = 1
const __GCC_ASM_FLAG_OUTPUTS__ = 1
const __GCC_ATOMIC_BOOL_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR16_T_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR32_T_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR_LOCK_FREE = 2
const __GCC_ATOMIC_INT_LOCK_FREE = 2
const __GCC_ATOMIC_LLONG_LOCK_FREE = 2
const __GCC_ATOMIC_LONG_LOCK_FREE = 2
const __GCC_ATOMIC_POINTER_LOCK_FREE = 2
const __GCC_ATOMIC_SHORT_LOCK_FREE = 2
const __GCC_ATOMIC_TEST_AND_SET_TRUEVAL = 1
const __GCC_ATOMIC_WCHAR_T_LOCK_FREE = 2
const __GCC_CONSTRUCTIVE_SIZE = 64
const __GCC_DESTRUCTIVE_SIZE = 64
const __GCC_HAVE_DWARF2_CFI_ASM = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_1 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_2 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_4 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_8 = 1
const __GCC_IEC_559 = 2
const __GCC_IEC_559_COMPLEX = 2
const __GNUC_EXECUTION_CHARSET_NAME = "UTF-8"
const __GNUC_MINOR__ = 3
const __GNUC_PATCHLEVEL__ = 0
const __GNUC_STDC_INLINE__ = 1
const __GNUC_WIDE_EXECUTION_CHARSET_NAME = "UTF-32LE"
const __GNUC__ = 13
const __GXX_ABI_VERSION = 1018
const __HAVE_SPECULATION_SAFE_VALUE = 1
const __INT16_MAX__ = 0x7fff
const __INT32_MAX__ = 0x7fffffff
const __INT32_TYPE__ = "int"
const __INT64_MAX__ = 0x7fffffffffffffff
const __INT8_MAX__ = 0x7f
const __INTMAX_MAX__ = 0x7fffffffffffffff
const __INTMAX_WIDTH__ = 64
const __INTPTR_MAX__ = 0x7fffffffffffffff
const __INTPTR_WIDTH__ = 64
const __INT_FAST16_MAX__ = 0x7fffffffffffffff
const __INT_FAST16_WIDTH__ = 64
const __INT_FAST32_MAX__ = 0x7fffffffffffffff
const __INT_FAST32_WIDTH__ = 64
const __INT_FAST64_MAX__ = 0x7fffffffffffffff
const __INT_FAST64_WIDTH__ = 64
const __INT_FAST8_MAX__ = 0x7f
const __INT_FAST8_WIDTH__ = 8
const __INT_LEAST16_MAX__ = 0x7fff
const __INT_LEAST16_WIDTH__ = 16
const __INT_LEAST32_MAX__ = 0x7fffffff
const __INT_LEAST32_TYPE__ = "int"
const __INT_LEAST32_WIDTH__ = 32
const __INT_LEAST64_MAX__ = 0x7fffffffffffffff
const __INT_LEAST64_WIDTH__ = 64
const __INT_LEAST8_MAX__ = 0x7f
const __INT_LEAST8_WIDTH__ = 8
const __INT_MAX__ = 0x7fffffff
const __INT_WIDTH__ = 32
const __LDBL_DECIMAL_DIG__ = 17
const __LDBL_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const __LDBL_DIG__ = 15
const __LDBL_EPSILON__ = 2.22044604925031308084726333618164062e-16
const __LDBL_HAS_DENORM__ = 1
const __LDBL_HAS_INFINITY__ = 1
const __LDBL_HAS_QUIET_NAN__ = 1
const __LDBL_IS_IEC_60559__ = 1
const __LDBL_MANT_DIG__ = 53
const __LDBL_MAX_10_EXP__ = 308
const __LDBL_MAX_EXP__ = 1024
const __LDBL_MAX__ = 1.79769313486231570814527423731704357e+308
const __LDBL_MIN__ = 2.22507385850720138309023271733240406e-308
const __LDBL_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const __LITTLE_ENDIAN = 1234
const __LONG_DOUBLE_64__ = 1
const __LONG_LONG_MAX__ = 0x7fffffffffffffff
const __LONG_LONG_WIDTH__ = 64
const __LONG_MAX = 0x7fffffffffffffff
const __LONG_MAX__ = 0x7fffffffffffffff
const __LONG_WIDTH__ = 64
const __LP64__ = 1
const __MMX_WITH_SSE__ = 1
const __MMX__ = 1
const __NO_INLINE__ = 1
const __ORDER_BIG_ENDIAN__ = 4321
const __ORDER_LITTLE_ENDIAN__ = 1234
const __ORDER_PDP_ENDIAN__ = 3412
const __PIC__ = 2
const __PIE__ = 2
const __PRAGMA_REDEFINE_EXTNAME = 1
const __PRETTY_FUNCTION__ = "__func__"
const __PTRDIFF_MAX__ = 0x7fffffffffffffff
const __PTRDIFF_WIDTH__ = 64
const __SCHAR_MAX__ = 0x7f
const __SCHAR_WIDTH__ = 8
const __SEG_FS = 1
const __SEG_GS = 1
const __SHRT_MAX__ = 0x7fff
const __SHRT_WIDTH__ = 16
const __SIG_ATOMIC_MAX__ = 0x7fffffff
const __SIG_ATOMIC_TYPE__ = "int"
const __SIG_ATOMIC_WIDTH__ = 32
const __SIZEOF_DOUBLE__ = 8
const __SIZEOF_FLOAT128__ = 16
const __SIZEOF_FLOAT80__ = 16
const __SIZEOF_FLOAT__ = 4
const __SIZEOF_INT128__ = 16
const __SIZEOF_INT__ = 4
const __SIZEOF_LONG_DOUBLE__ = 8
const __SIZEOF_LONG_LONG__ = 8
const __SIZEOF_LONG__ = 8
const __SIZEOF_POINTER__ = 8
const __SIZEOF_PTRDIFF_T__ = 8
const __SIZEOF_SHORT__ = 2
const __SIZEOF_SIZE_T__ = 8
const __SIZEOF_WCHAR_T__ = 4
const __SIZEOF_WINT_T__ = 4
const __SIZE_MAX__ = 0xffffffffffffffff
const __SIZE_WIDTH__ = 64
const __SSE2_MATH__ = 1
const __SSE2__ = 1
const __SSE_MATH__ = 1
const __SSE__ = 1
const __SSP_STRONG__ = 3
const __STDC_HOSTED__ = 1
const __STDC_IEC_559_COMPLEX__ = 1
const __STDC_IEC_559__ = 1
const __STDC_IEC_60559_BFP__ = 201404
const __STDC_IEC_60559_COMPLEX__ = 201404
const __STDC_ISO_10646__ = 201706
const __STDC_UTF_16__ = 1
const __STDC_UTF_32__ = 1
const __STDC_VERSION__ = 201710
const __STDC__ = 1
const __UINT16_MAX__ = 0xffff
const __UINT32_MAX__ = 0xffffffff
const __UINT64_MAX__ = 0xffffffffffffffff
const __UINT8_MAX__ = 0xff
const __UINTMAX_MAX__ = 0xffffffffffffffff
const __UINTPTR_MAX__ = 0xffffffffffffffff
const __UINT_FAST16_MAX__ = 0xffffffffffffffff
const __UINT_FAST32_MAX__ = 0xffffffffffffffff
const __UINT_FAST64_MAX__ = 0xffffffffffffffff
const __UINT_FAST8_MAX__ = 0xff
const __UINT_LEAST16_MAX__ = 0xffff
const __UINT_LEAST32_MAX__ = 0xffffffff
const __UINT_LEAST64_MAX__ = 0xffffffffffffffff
const __UINT_LEAST8_MAX__ = 0xff
const __USE_TIME_BITS64 = 1
const __VERSION__ = "13.3.0"
const __WCHAR_MAX__ = 0x7fffffff
const __WCHAR_TYPE__ = "int"
const __WCHAR_WIDTH__ = 32
const __WINT_MAX__ = 0xffffffff
const __WINT_MIN__ = 0
const __WINT_WIDTH__ = 32
const __amd64 = 1
const __amd64__ = 1
const __code_model_small__ = 1
const __gnu_linux__ = 1
const __k8 = 1
const __k8__ = 1
const __linux = 1
const __linux__ = 1
const __pic__ = 2
const __pie__ = 2
const __restrict_arr = "restrict"
const __unix = 1
const __unix__ = 1
const __x86_64 = 1
const __x86_64__ = 1
const linux = 1
const unix = 1

type __builtin_va_list = uintptr

type __predefined_size_t = uint64

type __predefined_wchar_t = int32

type __predefined_ptrdiff_t = int64

type wchar_t = int32

type max_align_t = struct {
	F__ll int64
	F__ld float64
}

type size_t = uint64

type ptrdiff_t = int64

type uintptr_t = uint64

type intptr_t = int64

type int8_t = int8

type int16_t = int16

type int32_t = int32

type int64_t = int64

type intmax_t = int64

type uint8_t = uint8

type uint16_t = uint16

type uint32_t = uint32

type uint64_t = uint64

type uintmax_t = uint64

type int_fast8_t = int8

type int_fast64_t = int64

type int_least8_t = int8

type int_least16_t = int16

type int_least32_t = int32

type int_least64_t = int64

type uint_fast8_t = uint8

type uint_fast64_t = uint64

type uint_least8_t = uint8

type uint_least16_t = uint16

type uint_least32_t = uint32

type uint_least64_t = uint64

type int_fast16_t = int32

type int_fast32_t = int32

type uint_fast16_t = uint32

type uint_fast32_t = uint32

type crypto_aead_ctx = struct {
	Fcounter uint64_t
	Fkey     [32]uint8_t
	Fnonce   [8]uint8_t
}

type crypto_blake2b_ctx = struct {
	Fhash         [8]uint64_t
	Finput_offset [2]uint64_t
	Finput        [16]uint64_t
	Finput_idx    size_t
	Fhash_size    size_t
}

type crypto_argon2_config = struct {
	Falgorithm uint32_t
	Fnb_blocks uint32_t
	Fnb_passes uint32_t
	Fnb_lanes  uint32_t
}

type crypto_argon2_inputs = struct {
	Fpass      uintptr
	Fsalt      uintptr
	Fpass_size uint32_t
	Fsalt_size uint32_t
}

type crypto_argon2_extras = struct {
	Fkey      uintptr
	Fad       uintptr
	Fkey_size uint32_t
	Fad_size  uint32_t
}

type crypto_poly1305_ctx = struct {
	Fc     [16]uint8_t
	Fc_idx size_t
	Fr     [4]uint32_t
	Fpad   [4]uint32_t
	Fh     [5]uint32_t
}

/////////////////
/// Utilities ///
/////////////////

type i8 = int8

type u8 = uint8

type i16 = int16

type u32 = uint32

type i32 = int32

type i64 = int64

type u64 = uint64

var zero = [128]u8{}

// C documentation
//
//	// returns the smallest positive integer y such that
//	// (x + y) % pow_2  == 0
//	// Basically, y is the "gap" missing to align x.
//	// Only works when pow_2 is a power of 2.
//	// Note: we use ~x+1 instead of -x to avoid compiler warnings
func gap(tls *libc.TLS, x size_t, pow_2 size_t) (r size_t) {
	return (^x + uint64(1)) & (pow_2 - uint64(1))
}

func load24_le(tls *libc.TLS, s uintptr) (r u32) {
	return uint32(**(**u8)(__ccgo_up(s)))<<libc.Int32FromInt32(0) | uint32(**(**u8)(__ccgo_up(s + 1)))<<libc.Int32FromInt32(8) | uint32(**(**u8)(__ccgo_up(s + 2)))<<libc.Int32FromInt32(16)
}

func load32_le(tls *libc.TLS, s uintptr) (r u32) {
	return uint32(**(**u8)(__ccgo_up(s)))<<libc.Int32FromInt32(0) | uint32(**(**u8)(__ccgo_up(s + 1)))<<libc.Int32FromInt32(8) | uint32(**(**u8)(__ccgo_up(s + 2)))<<libc.Int32FromInt32(16) | uint32(**(**u8)(__ccgo_up(s + 3)))<<libc.Int32FromInt32(24)
}

func load64_le(tls *libc.TLS, s uintptr) (r u64) {
	return uint64(load32_le(tls, s)) | uint64(load32_le(tls, s+uintptr(4)))<<libc.Int32FromInt32(32)
}

func store32_le(tls *libc.TLS, out uintptr, in u32) {
	**(**u8)(__ccgo_up(out)) = uint8(in)
	**(**u8)(__ccgo_up(out + 1)) = uint8(in >> libc.Int32FromInt32(8))
	**(**u8)(__ccgo_up(out + 2)) = uint8(in >> libc.Int32FromInt32(16))
	**(**u8)(__ccgo_up(out + 3)) = uint8(in >> libc.Int32FromInt32(24))
}

func store64_le(tls *libc.TLS, out uintptr, in u64) {
	store32_le(tls, out, uint32(in))
	store32_le(tls, out+uintptr(4), uint32(in>>libc.Int32FromInt32(32)))
}

func load32_le_buf(tls *libc.TLS, dst uintptr, src uintptr, size size_t) {
	var i size_t
	_ = i
	i = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(i < size) {
			break
		}
		**(**u32)(__ccgo_up(dst + uintptr(i)*4)) = load32_le(tls, src+uintptr(i*uint64(4)))
		goto _1
	_1:
		;
		i = i + 1
	}
}

func load64_le_buf(tls *libc.TLS, dst uintptr, src uintptr, size size_t) {
	var i size_t
	_ = i
	i = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(i < size) {
			break
		}
		**(**u64)(__ccgo_up(dst + uintptr(i)*8)) = load64_le(tls, src+uintptr(i*uint64(8)))
		goto _1
	_1:
		;
		i = i + 1
	}
}

func store32_le_buf(tls *libc.TLS, dst uintptr, src uintptr, size size_t) {
	var i size_t
	_ = i
	i = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(i < size) {
			break
		}
		store32_le(tls, dst+uintptr(i*uint64(4)), **(**u32)(__ccgo_up(src + uintptr(i)*4)))
		goto _1
	_1:
		;
		i = i + 1
	}
}

func store64_le_buf(tls *libc.TLS, dst uintptr, src uintptr, size size_t) {
	var i size_t
	_ = i
	i = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(i < size) {
			break
		}
		store64_le(tls, dst+uintptr(i*uint64(8)), **(**u64)(__ccgo_up(src + uintptr(i)*8)))
		goto _1
	_1:
		;
		i = i + 1
	}
}

func rotr64(tls *libc.TLS, x u64, n u64) (r u64) {
	return x>>n ^ x<<(libc.Uint64FromInt32(64)-n)
}

func rotl32(tls *libc.TLS, x u32, n u32) (r u32) {
	return x<<n ^ x>>(libc.Uint32FromInt32(32)-n)
}

func neq0(tls *libc.TLS, diff u64) (r int32) {
	var eq0, half u64
	_, _ = eq0, half
	// constant time comparison to zero
	// return diff != 0 ? -1 : 0
	half = diff>>libc.Int32FromInt32(32) | uint64(uint32(diff)) // half < 2^32
	eq0 = uint64(1) & ((half - uint64(1)) >> int32(32))         // half == 0 ? 1 : 0
	return libc.Int32FromUint64(eq0) - int32(1)                 // half == 0 ? 0 : -1
}

func x16(tls *libc.TLS, a uintptr, b uintptr) (r u64) {
	return load64_le(tls, a+uintptr(0)) ^ load64_le(tls, b+uintptr(0)) | (load64_le(tls, a+uintptr(8)) ^ load64_le(tls, b+uintptr(8)))
}

func x32(tls *libc.TLS, a uintptr, b uintptr) (r u64) {
	return x16(tls, a, b) | x16(tls, a+uintptr(16), b+uintptr(16))
}

func x64(tls *libc.TLS, a uintptr, b uintptr) (r u64) {
	return x32(tls, a, b) | x32(tls, a+uintptr(32), b+uintptr(32))
}

func crypto_verify16(tls *libc.TLS, a uintptr, b uintptr) (r int32) {
	return neq0(tls, x16(tls, a, b))
}

func crypto_verify32(tls *libc.TLS, a uintptr, b uintptr) (r int32) {
	return neq0(tls, x32(tls, a, b))
}

func crypto_verify64(tls *libc.TLS, a uintptr, b uintptr) (r int32) {
	return neq0(tls, x64(tls, a, b))
}

func crypto_wipe(tls *libc.TLS, secret uintptr, size size_t) {
	var _i_ size_t
	var v_secret uintptr
	_, _ = _i_, v_secret
	v_secret = secret
	_i_ = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(_i_ < size) {
			break
		}
		libc.AtomicStorePUint8(v_secret+uintptr(_i_), uint8(0))
		goto _1
	_1:
		;
		_i_ = _i_ + 1
	}
}

/////////////////
/// Chacha 20 ///
/////////////////

func chacha20_rounds(tls *libc.TLS, out uintptr, in uintptr) {
	var i size_t
	var t0, t1, t10, t11, t12, t13, t14, t15, t2, t3, t4, t5, t6, t7, t8, t9 u32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = i, t0, t1, t10, t11, t12, t13, t14, t15, t2, t3, t4, t5, t6, t7, t8, t9
	// The temporary variables make Chacha20 10% faster.
	t0 = **(**u32)(__ccgo_up(in))
	t1 = **(**u32)(__ccgo_up(in + 1*4))
	t2 = **(**u32)(__ccgo_up(in + 2*4))
	t3 = **(**u32)(__ccgo_up(in + 3*4))
	t4 = **(**u32)(__ccgo_up(in + 4*4))
	t5 = **(**u32)(__ccgo_up(in + 5*4))
	t6 = **(**u32)(__ccgo_up(in + 6*4))
	t7 = **(**u32)(__ccgo_up(in + 7*4))
	t8 = **(**u32)(__ccgo_up(in + 8*4))
	t9 = **(**u32)(__ccgo_up(in + 9*4))
	t10 = **(**u32)(__ccgo_up(in + 10*4))
	t11 = **(**u32)(__ccgo_up(in + 11*4))
	t12 = **(**u32)(__ccgo_up(in + 12*4))
	t13 = **(**u32)(__ccgo_up(in + 13*4))
	t14 = **(**u32)(__ccgo_up(in + 14*4))
	t15 = **(**u32)(__ccgo_up(in + 15*4))
	i = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(i < libc.Uint64FromInt32(libc.Int32FromInt32(10))) {
			break
		} // 20 rounds, 2 rounds per loop.
		t0 = t0 + t4
		t12 = rotl32(tls, t12^t0, uint32(16))
		t8 = t8 + t12
		t4 = rotl32(tls, t4^t8, uint32(12))
		t0 = t0 + t4
		t12 = rotl32(tls, t12^t0, uint32(8))
		t8 = t8 + t12
		t4 = rotl32(tls, t4^t8, uint32(7)) // column 0
		t1 = t1 + t5
		t13 = rotl32(tls, t13^t1, uint32(16))
		t9 = t9 + t13
		t5 = rotl32(tls, t5^t9, uint32(12))
		t1 = t1 + t5
		t13 = rotl32(tls, t13^t1, uint32(8))
		t9 = t9 + t13
		t5 = rotl32(tls, t5^t9, uint32(7)) // column 1
		t2 = t2 + t6
		t14 = rotl32(tls, t14^t2, uint32(16))
		t10 = t10 + t14
		t6 = rotl32(tls, t6^t10, uint32(12))
		t2 = t2 + t6
		t14 = rotl32(tls, t14^t2, uint32(8))
		t10 = t10 + t14
		t6 = rotl32(tls, t6^t10, uint32(7)) // column 2
		t3 = t3 + t7
		t15 = rotl32(tls, t15^t3, uint32(16))
		t11 = t11 + t15
		t7 = rotl32(tls, t7^t11, uint32(12))
		t3 = t3 + t7
		t15 = rotl32(tls, t15^t3, uint32(8))
		t11 = t11 + t15
		t7 = rotl32(tls, t7^t11, uint32(7)) // column 3
		t0 = t0 + t5
		t15 = rotl32(tls, t15^t0, uint32(16))
		t10 = t10 + t15
		t5 = rotl32(tls, t5^t10, uint32(12))
		t0 = t0 + t5
		t15 = rotl32(tls, t15^t0, uint32(8))
		t10 = t10 + t15
		t5 = rotl32(tls, t5^t10, uint32(7)) // diagonal 0
		t1 = t1 + t6
		t12 = rotl32(tls, t12^t1, uint32(16))
		t11 = t11 + t12
		t6 = rotl32(tls, t6^t11, uint32(12))
		t1 = t1 + t6
		t12 = rotl32(tls, t12^t1, uint32(8))
		t11 = t11 + t12
		t6 = rotl32(tls, t6^t11, uint32(7)) // diagonal 1
		t2 = t2 + t7
		t13 = rotl32(tls, t13^t2, uint32(16))
		t8 = t8 + t13
		t7 = rotl32(tls, t7^t8, uint32(12))
		t2 = t2 + t7
		t13 = rotl32(tls, t13^t2, uint32(8))
		t8 = t8 + t13
		t7 = rotl32(tls, t7^t8, uint32(7)) // diagonal 2
		t3 = t3 + t4
		t14 = rotl32(tls, t14^t3, uint32(16))
		t9 = t9 + t14
		t4 = rotl32(tls, t4^t9, uint32(12))
		t3 = t3 + t4
		t14 = rotl32(tls, t14^t3, uint32(8))
		t9 = t9 + t14
		t4 = rotl32(tls, t4^t9, uint32(7)) // diagonal 3
		goto _1
	_1:
		;
		i = i + 1
	}
	**(**u32)(__ccgo_up(out)) = t0
	**(**u32)(__ccgo_up(out + 1*4)) = t1
	**(**u32)(__ccgo_up(out + 2*4)) = t2
	**(**u32)(__ccgo_up(out + 3*4)) = t3
	**(**u32)(__ccgo_up(out + 4*4)) = t4
	**(**u32)(__ccgo_up(out + 5*4)) = t5
	**(**u32)(__ccgo_up(out + 6*4)) = t6
	**(**u32)(__ccgo_up(out + 7*4)) = t7
	**(**u32)(__ccgo_up(out + 8*4)) = t8
	**(**u32)(__ccgo_up(out + 9*4)) = t9
	**(**u32)(__ccgo_up(out + 10*4)) = t10
	**(**u32)(__ccgo_up(out + 11*4)) = t11
	**(**u32)(__ccgo_up(out + 12*4)) = t12
	**(**u32)(__ccgo_up(out + 13*4)) = t13
	**(**u32)(__ccgo_up(out + 14*4)) = t14
	**(**u32)(__ccgo_up(out + 15*4)) = t15
}

var chacha20_constant = __ccgo_ts // 16 bytes

func crypto_chacha20_h(tls *libc.TLS, out uintptr, key uintptr, in uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var _ /* block at bp+0 */ [16]u32
	load32_le_buf(tls, bp, chacha20_constant, uint64(4))
	load32_le_buf(tls, bp+uintptr(4)*4, key, uint64(8))
	load32_le_buf(tls, bp+uintptr(12)*4, in, uint64(4))
	chacha20_rounds(tls, bp, bp)
	// prevent reversal of the rounds by revealing only half of the buffer.
	store32_le_buf(tls, out, bp, uint64(4))                           // constant
	store32_le_buf(tls, out+uintptr(16), bp+uintptr(12)*4, uint64(4)) // counter and nonce
	crypto_wipe(tls, bp, uint64(64))
}

func crypto_chacha20_djb(tls *libc.TLS, cipher_text uintptr, plain_text uintptr, text_size size_t, key uintptr, nonce uintptr, ctr u64) (r u64) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var i, i1, i2, j, j1, nb_blocks size_t
	var p, p1 u32
	var _ /* input at bp+0 */ [16]u32
	var _ /* pool at bp+64 */ [16]u32
	var _ /* tmp at bp+128 */ [64]u8
	_, _, _, _, _, _, _, _ = i, i1, i2, j, j1, nb_blocks, p, p1
	load32_le_buf(tls, bp, chacha20_constant, uint64(4))
	load32_le_buf(tls, bp+uintptr(4)*4, key, uint64(8))
	load32_le_buf(tls, bp+uintptr(14)*4, nonce, uint64(2))
	(**(**[16]u32)(__ccgo_up(bp)))[int32(12)] = uint32(ctr)
	(**(**[16]u32)(__ccgo_up(bp)))[int32(13)] = uint32(ctr >> libc.Int32FromInt32(32))
	nb_blocks = text_size >> int32(6)
	i = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(i < nb_blocks) {
			break
		}
		chacha20_rounds(tls, bp+64, bp)
		if plain_text != libc.UintptrFromInt32(0) {
			j = libc.Uint64FromInt32(libc.Int32FromInt32(0))
			for {
				if !(j < libc.Uint64FromInt32(libc.Int32FromInt32(16))) {
					break
				}
				p = (**(**[16]u32)(__ccgo_up(bp + 64)))[j] + (**(**[16]u32)(__ccgo_up(bp)))[j]
				store32_le(tls, cipher_text, p^load32_le(tls, plain_text))
				cipher_text = cipher_text + uintptr(4)
				plain_text = plain_text + uintptr(4)
				goto _2
			_2:
				;
				j = j + 1
			}
		} else {
			j1 = libc.Uint64FromInt32(libc.Int32FromInt32(0))
			for {
				if !(j1 < libc.Uint64FromInt32(libc.Int32FromInt32(16))) {
					break
				}
				p1 = (**(**[16]u32)(__ccgo_up(bp + 64)))[j1] + (**(**[16]u32)(__ccgo_up(bp)))[j1]
				store32_le(tls, cipher_text, p1)
				cipher_text = cipher_text + uintptr(4)
				goto _3
			_3:
				;
				j1 = j1 + 1
			}
		}
		(**(**[16]u32)(__ccgo_up(bp)))[int32(12)] = (**(**[16]u32)(__ccgo_up(bp)))[int32(12)] + 1
		if (**(**[16]u32)(__ccgo_up(bp)))[int32(12)] == uint32(0) {
			(**(**[16]u32)(__ccgo_up(bp)))[int32(13)] = (**(**[16]u32)(__ccgo_up(bp)))[int32(13)] + 1
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	text_size = text_size & uint64(63)
	// Last (incomplete) block
	if text_size > uint64(0) {
		if plain_text == libc.UintptrFromInt32(0) {
			plain_text = uintptr(unsafe.Pointer(&zero))
		}
		chacha20_rounds(tls, bp+64, bp)
		i1 = libc.Uint64FromInt32(libc.Int32FromInt32(0))
		for {
			if !(i1 < libc.Uint64FromInt32(libc.Int32FromInt32(16))) {
				break
			}
			store32_le(tls, bp+128+uintptr(i1*uint64(4)), (**(**[16]u32)(__ccgo_up(bp + 64)))[i1]+(**(**[16]u32)(__ccgo_up(bp)))[i1])
			goto _4
		_4:
			;
			i1 = i1 + 1
		}
		i2 = libc.Uint64FromInt32(libc.Int32FromInt32(0))
		for {
			if !(i2 < text_size) {
				break
			}
			**(**u8)(__ccgo_up(cipher_text + uintptr(i2))) = libc.Uint8FromInt32(libc.Int32FromUint8((**(**[64]u8)(__ccgo_up(bp + 128)))[i2]) ^ libc.Int32FromUint8(**(**u8)(__ccgo_up(plain_text + uintptr(i2)))))
			goto _5
		_5:
			;
			i2 = i2 + 1
		}
		crypto_wipe(tls, bp+128, uint64(64))
	}
	ctr = uint64((**(**[16]u32)(__ccgo_up(bp)))[int32(12)]) + uint64((**(**[16]u32)(__ccgo_up(bp)))[int32(13)])<<libc.Int32FromInt32(32) + libc.BoolUint64(text_size > libc.Uint64FromInt32(0))
	crypto_wipe(tls, bp+64, uint64(64))
	crypto_wipe(tls, bp, uint64(64))
	return ctr
}

func crypto_chacha20_ietf(tls *libc.TLS, cipher_text uintptr, plain_text uintptr, text_size size_t, key uintptr, nonce uintptr, ctr u32) (r u32) {
	var big_ctr u64
	_ = big_ctr
	big_ctr = uint64(ctr) + uint64(load32_le(tls, nonce))<<libc.Int32FromInt32(32)
	return uint32(crypto_chacha20_djb(tls, cipher_text, plain_text, text_size, key, nonce+uintptr(4), big_ctr))
}

func crypto_chacha20_x(tls *libc.TLS, cipher_text uintptr, plain_text uintptr, text_size size_t, key uintptr, nonce uintptr, ctr u64) (r u64) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var _ /* sub_key at bp+0 */ [32]u8
	crypto_chacha20_h(tls, bp, key, nonce)
	ctr = crypto_chacha20_djb(tls, cipher_text, plain_text, text_size, bp, nonce+uintptr(16), ctr)
	crypto_wipe(tls, bp, uint64(32))
	return ctr
}

/////////////////
/// Poly 1305 ///
/////////////////

// C documentation
//
//	// h = (h + c) * r
//	// preconditions:
//	//   ctx->h <= 4_ffffffff_ffffffff_ffffffff_ffffffff
//	//   ctx->r <=   0ffffffc_0ffffffc_0ffffffc_0fffffff
//	//   end    <= 1
//	// Postcondition:
//	//   ctx->h <= 4_ffffffff_ffffffff_ffffffff_ffffffff
func poly_blocks(tls *libc.TLS, ctx uintptr, in uintptr, nb_blocks size_t, end uint32) {
	var h0, h1, h2, h3, h4, r0, r1, r2, r3, rr0, rr1, rr2, rr3, rr4, s4, u4, u5, x4 u32
	var i size_t
	var s0, s1, s2, s3, u0, u1, u2, u3, x0, x1, x2, x3 u64
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = h0, h1, h2, h3, h4, i, r0, r1, r2, r3, rr0, rr1, rr2, rr3, rr4, s0, s1, s2, s3, s4, u0, u1, u2, u3, u4, u5, x0, x1, x2, x3, x4
	// Local all the things!
	r0 = **(**uint32_t)(__ccgo_up(ctx + 24))
	r1 = **(**uint32_t)(__ccgo_up(ctx + 24 + 1*4))
	r2 = **(**uint32_t)(__ccgo_up(ctx + 24 + 2*4))
	r3 = **(**uint32_t)(__ccgo_up(ctx + 24 + 3*4))
	rr0 = r0 >> libc.Int32FromInt32(2) * uint32(5) // lose 2 bits...
	rr1 = r1>>libc.Int32FromInt32(2) + r1          // rr1 == (r1 >> 2) * 5
	rr2 = r2>>libc.Int32FromInt32(2) + r2          // rr1 == (r2 >> 2) * 5
	rr3 = r3>>libc.Int32FromInt32(2) + r3          // rr1 == (r3 >> 2) * 5
	rr4 = r0 & uint32(3)                           // ...recover 2 bits
	h0 = **(**uint32_t)(__ccgo_up(ctx + 56))
	h1 = **(**uint32_t)(__ccgo_up(ctx + 56 + 1*4))
	h2 = **(**uint32_t)(__ccgo_up(ctx + 56 + 2*4))
	h3 = **(**uint32_t)(__ccgo_up(ctx + 56 + 3*4))
	h4 = **(**uint32_t)(__ccgo_up(ctx + 56 + 4*4))
	i = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(i < nb_blocks) {
			break
		}
		// h + c, without carry propagation
		s0 = uint64(h0) + uint64(load32_le(tls, in))
		in = in + uintptr(4)
		s1 = uint64(h1) + uint64(load32_le(tls, in))
		in = in + uintptr(4)
		s2 = uint64(h2) + uint64(load32_le(tls, in))
		in = in + uintptr(4)
		s3 = uint64(h3) + uint64(load32_le(tls, in))
		in = in + uintptr(4)
		s4 = h4 + end
		// (h + c) * r, without carry propagation
		x0 = s0*uint64(r0) + s1*uint64(rr3) + s2*uint64(rr2) + s3*uint64(rr1) + uint64(s4*rr0)
		x1 = s0*uint64(r1) + s1*uint64(r0) + s2*uint64(rr3) + s3*uint64(rr2) + uint64(s4*rr1)
		x2 = s0*uint64(r2) + s1*uint64(r1) + s2*uint64(r0) + s3*uint64(rr3) + uint64(s4*rr2)
		x3 = s0*uint64(r3) + s1*uint64(r2) + s2*uint64(r1) + s3*uint64(r0) + uint64(s4*rr3)
		x4 = s4 * rr4
		// partial reduction modulo 2^130 - 5
		u5 = uint32(x3>>libc.Int32FromInt32(32)) + x4 // u5 <= 7ffffff5
		u0 = uint64(u5>>libc.Int32FromInt32(2)*uint32(5)) + x0&uint64(0xffffffff)
		u1 = uint64(uint32(u0>>libc.Int32FromInt32(32))) + x1&uint64(0xffffffff) + x0>>libc.Int32FromInt32(32)
		u2 = uint64(uint32(u1>>libc.Int32FromInt32(32))) + x2&uint64(0xffffffff) + x1>>libc.Int32FromInt32(32)
		u3 = uint64(uint32(u2>>libc.Int32FromInt32(32))) + x3&uint64(0xffffffff) + x2>>libc.Int32FromInt32(32)
		u4 = uint32(u3>>libc.Int32FromInt32(32)) + u5&uint32(3) // u4 <= 4
		// Update the hash
		h0 = uint32(u0 & uint64(0xffffffff))
		h1 = uint32(u1 & uint64(0xffffffff))
		h2 = uint32(u2 & uint64(0xffffffff))
		h3 = uint32(u3 & uint64(0xffffffff))
		h4 = u4
		goto _1
	_1:
		;
		i = i + 1
	}
	**(**uint32_t)(__ccgo_up(ctx + 56)) = h0
	**(**uint32_t)(__ccgo_up(ctx + 56 + 1*4)) = h1
	**(**uint32_t)(__ccgo_up(ctx + 56 + 2*4)) = h2
	**(**uint32_t)(__ccgo_up(ctx + 56 + 3*4)) = h3
	**(**uint32_t)(__ccgo_up(ctx + 56 + 4*4)) = h4
}

func crypto_poly1305_init(tls *libc.TLS, ctx uintptr, key uintptr) {
	var _i_, i, i1 size_t
	_, _, _ = _i_, i, i1
	_i_ = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(_i_ < libc.Uint64FromInt32(libc.Int32FromInt32(5))) {
			break
		}
		**(**uint32_t)(__ccgo_up(ctx + 56 + uintptr(_i_)*4)) = uint32(0)
		goto _1
	_1:
		;
		_i_ = _i_ + 1
	} // Initial hash is zero
	(*crypto_poly1305_ctx)(unsafe.Pointer(ctx)).Fc_idx = uint64(0)
	// load r and pad (r has some of its bits cleared)
	load32_le_buf(tls, ctx+24, key, uint64(4))
	load32_le_buf(tls, ctx+40, key+uintptr(16), uint64(4))
	i = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(i < libc.Uint64FromInt32(libc.Int32FromInt32(1))) {
			break
		}
		**(**uint32_t)(__ccgo_up(ctx + 24 + uintptr(i)*4)) &= uint32(0x0fffffff)
		goto _2
	_2:
		;
		i = i + 1
	}
	i1 = libc.Uint64FromInt32(libc.Int32FromInt32(1))
	for {
		if !(i1 < libc.Uint64FromInt32(libc.Int32FromInt32(4))) {
			break
		}
		**(**uint32_t)(__ccgo_up(ctx + 24 + uintptr(i1)*4)) &= uint32(0x0ffffffc)
		goto _3
	_3:
		;
		i1 = i1 + 1
	}
}

func crypto_poly1305_update(tls *libc.TLS, ctx uintptr, message uintptr, message_size size_t) {
	var aligned, i, i1, nb_blocks size_t
	var v1 uint64
	_, _, _, _, _ = aligned, i, i1, nb_blocks, v1
	// Avoid undefined NULL pointer increments with empty messages
	if message_size == uint64(0) {
		return
	}
	if gap(tls, (*crypto_poly1305_ctx)(unsafe.Pointer(ctx)).Fc_idx, uint64(16)) <= message_size {
		v1 = gap(tls, (*crypto_poly1305_ctx)(unsafe.Pointer(ctx)).Fc_idx, uint64(16))
	} else {
		v1 = message_size
	}
	// Align ourselves with block boundaries
	aligned = v1
	i = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(i < aligned) {
			break
		}
		**(**uint8_t)(__ccgo_up(ctx + uintptr((*crypto_poly1305_ctx)(unsafe.Pointer(ctx)).Fc_idx))) = **(**u8)(__ccgo_up(message))
		(*crypto_poly1305_ctx)(unsafe.Pointer(ctx)).Fc_idx = (*crypto_poly1305_ctx)(unsafe.Pointer(ctx)).Fc_idx + 1
		message = message + 1
		message_size = message_size - 1
		goto _2
	_2:
		;
		i = i + 1
	}
	// If block is complete, process it
	if (*crypto_poly1305_ctx)(unsafe.Pointer(ctx)).Fc_idx == uint64(16) {
		poly_blocks(tls, ctx, ctx, uint64(1), uint32(1))
		(*crypto_poly1305_ctx)(unsafe.Pointer(ctx)).Fc_idx = uint64(0)
	}
	// Process the message block by block
	nb_blocks = message_size >> int32(4)
	poly_blocks(tls, ctx, message, nb_blocks, uint32(1))
	message = message + uintptr(nb_blocks<<int32(4))
	message_size = message_size & uint64(15)
	// remaining bytes (we never complete a block here)
	i1 = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(i1 < message_size) {
			break
		}
		**(**uint8_t)(__ccgo_up(ctx + uintptr((*crypto_poly1305_ctx)(unsafe.Pointer(ctx)).Fc_idx))) = **(**u8)(__ccgo_up(message + uintptr(i1)))
		(*crypto_poly1305_ctx)(unsafe.Pointer(ctx)).Fc_idx = (*crypto_poly1305_ctx)(unsafe.Pointer(ctx)).Fc_idx + 1
		goto _3
	_3:
		;
		i1 = i1 + 1
	}
}

func crypto_poly1305_final(tls *libc.TLS, ctx uintptr, mac uintptr) {
	var _i_, i, i1 size_t
	var c u64
	_, _, _, _ = _i_, c, i, i1
	// Process the last block (if any)
	// We move the final 1 according to remaining input length
	// (this will add less than 2^130 to the last input block)
	if (*crypto_poly1305_ctx)(unsafe.Pointer(ctx)).Fc_idx != uint64(0) {
		_i_ = libc.Uint64FromInt32(libc.Int32FromInt32(0))
		for {
			if !(_i_ < uint64(16)-(*crypto_poly1305_ctx)(unsafe.Pointer(ctx)).Fc_idx) {
				break
			}
			**(**uint8_t)(__ccgo_up(ctx + uintptr((*crypto_poly1305_ctx)(unsafe.Pointer(ctx)).Fc_idx) + uintptr(_i_))) = uint8(0)
			goto _1
		_1:
			;
			_i_ = _i_ + 1
		}
		**(**uint8_t)(__ccgo_up(ctx + uintptr((*crypto_poly1305_ctx)(unsafe.Pointer(ctx)).Fc_idx))) = uint8(1)
		poly_blocks(tls, ctx, ctx, uint64(1), uint32(0))
	}
	// check if we should subtract 2^130-5 by performing the
	// corresponding carry propagation.
	c = uint64(5)
	i = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(i < libc.Uint64FromInt32(libc.Int32FromInt32(4))) {
			break
		}
		c = c + uint64(**(**uint32_t)(__ccgo_up(ctx + 56 + uintptr(i)*4)))
		c = c >> uint64(32)
		goto _2
	_2:
		;
		i = i + 1
	}
	c = c + uint64(**(**uint32_t)(__ccgo_up(ctx + 56 + 4*4)))
	c = c >> libc.Int32FromInt32(2) * uint64(5) // shift the carry back to the beginning
	// c now indicates how many times we should subtract 2^130-5 (0 or 1)
	i1 = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(i1 < libc.Uint64FromInt32(libc.Int32FromInt32(4))) {
			break
		}
		c = c + (uint64(**(**uint32_t)(__ccgo_up(ctx + 56 + uintptr(i1)*4))) + uint64(**(**uint32_t)(__ccgo_up(ctx + 40 + uintptr(i1)*4))))
		store32_le(tls, mac+uintptr(i1*uint64(4)), uint32(c))
		c = c >> int32(32)
		goto _3
	_3:
		;
		i1 = i1 + 1
	}
	crypto_wipe(tls, ctx, uint64(80))
}

func crypto_poly1305(tls *libc.TLS, mac uintptr, message uintptr, message_size size_t, key uintptr) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var _ /* ctx at bp+0 */ crypto_poly1305_ctx
	crypto_poly1305_init(tls, bp, key)
	crypto_poly1305_update(tls, bp, message, message_size)
	crypto_poly1305_final(tls, bp, mac)
}

// C documentation
//
//	////////////////
//	/// BLAKE2 b ///
//	////////////////
var iv = [8]u64{
	0: uint64(0x6a09e667f3bcc908),
	1: uint64(0xbb67ae8584caa73b),
	2: uint64(0x3c6ef372fe94f82b),
	3: uint64(0xa54ff53a5f1d36f1),
	4: uint64(0x510e527fade682d1),
	5: uint64(0x9b05688c2b3e6c1f),
	6: uint64(0x1f83d9abfb41bd6b),
	7: uint64(0x5be0cd19137e2179),
}

func blake2b_compress(tls *libc.TLS, ctx uintptr, is_last_block int32) {
	var input, x uintptr
	var v0, v1, v10, v11, v12, v13, v14, v15, v2, v3, v4, v5, v6, v7, v8, v9 u64
	var y size_t
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = input, v0, v1, v10, v11, v12, v13, v14, v15, v2, v3, v4, v5, v6, v7, v8, v9, x, y
	// increment input offset
	x = ctx + 64
	y = (*crypto_blake2b_ctx)(unsafe.Pointer(ctx)).Finput_idx
	**(**u64)(__ccgo_up(x)) += y
	if **(**u64)(__ccgo_up(x)) < y {
		**(**u64)(__ccgo_up(x + 1*8)) = **(**u64)(__ccgo_up(x + 1*8)) + 1
	}
	// init work vector
	v0 = **(**uint64_t)(__ccgo_up(ctx))
	v8 = iv[0]
	v1 = **(**uint64_t)(__ccgo_up(ctx + 1*8))
	v9 = iv[int32(1)]
	v2 = **(**uint64_t)(__ccgo_up(ctx + 2*8))
	v10 = iv[int32(2)]
	v3 = **(**uint64_t)(__ccgo_up(ctx + 3*8))
	v11 = iv[int32(3)]
	v4 = **(**uint64_t)(__ccgo_up(ctx + 4*8))
	v12 = iv[int32(4)] ^ **(**uint64_t)(__ccgo_up(ctx + 64))
	v5 = **(**uint64_t)(__ccgo_up(ctx + 5*8))
	v13 = iv[int32(5)] ^ **(**uint64_t)(__ccgo_up(ctx + 64 + 1*8))
	v6 = **(**uint64_t)(__ccgo_up(ctx + 6*8))
	v14 = iv[int32(6)] ^ libc.Uint64FromInt32(^(is_last_block - libc.Int32FromInt32(1)))
	v7 = **(**uint64_t)(__ccgo_up(ctx + 7*8))
	v15 = iv[int32(7)]
	// mangle work vector
	input = ctx + 80
	v0 = v0 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)))))*8)))
	v12 = rotr64(tls, v12^v0, uint64(32))
	v8 = v8 + v12
	v4 = rotr64(tls, v4^v8, uint64(24))
	v0 = v0 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 1)))*8)))
	v12 = rotr64(tls, v12^v0, uint64(16))
	v8 = v8 + v12
	v4 = rotr64(tls, v4^v8, uint64(63))
	v1 = v1 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 2)))*8)))
	v13 = rotr64(tls, v13^v1, uint64(32))
	v9 = v9 + v13
	v5 = rotr64(tls, v5^v9, uint64(24))
	v1 = v1 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 3)))*8)))
	v13 = rotr64(tls, v13^v1, uint64(16))
	v9 = v9 + v13
	v5 = rotr64(tls, v5^v9, uint64(63))
	v2 = v2 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 4)))*8)))
	v14 = rotr64(tls, v14^v2, uint64(32))
	v10 = v10 + v14
	v6 = rotr64(tls, v6^v10, uint64(24))
	v2 = v2 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 5)))*8)))
	v14 = rotr64(tls, v14^v2, uint64(16))
	v10 = v10 + v14
	v6 = rotr64(tls, v6^v10, uint64(63))
	v3 = v3 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 6)))*8)))
	v15 = rotr64(tls, v15^v3, uint64(32))
	v11 = v11 + v15
	v7 = rotr64(tls, v7^v11, uint64(24))
	v3 = v3 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 7)))*8)))
	v15 = rotr64(tls, v15^v3, uint64(16))
	v11 = v11 + v15
	v7 = rotr64(tls, v7^v11, uint64(63))
	v0 = v0 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 8)))*8)))
	v15 = rotr64(tls, v15^v0, uint64(32))
	v10 = v10 + v15
	v5 = rotr64(tls, v5^v10, uint64(24))
	v0 = v0 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 9)))*8)))
	v15 = rotr64(tls, v15^v0, uint64(16))
	v10 = v10 + v15
	v5 = rotr64(tls, v5^v10, uint64(63))
	v1 = v1 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 10)))*8)))
	v12 = rotr64(tls, v12^v1, uint64(32))
	v11 = v11 + v12
	v6 = rotr64(tls, v6^v11, uint64(24))
	v1 = v1 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 11)))*8)))
	v12 = rotr64(tls, v12^v1, uint64(16))
	v11 = v11 + v12
	v6 = rotr64(tls, v6^v11, uint64(63))
	v2 = v2 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 12)))*8)))
	v13 = rotr64(tls, v13^v2, uint64(32))
	v8 = v8 + v13
	v7 = rotr64(tls, v7^v8, uint64(24))
	v2 = v2 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 13)))*8)))
	v13 = rotr64(tls, v13^v2, uint64(16))
	v8 = v8 + v13
	v7 = rotr64(tls, v7^v8, uint64(63))
	v3 = v3 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 14)))*8)))
	v14 = rotr64(tls, v14^v3, uint64(32))
	v9 = v9 + v14
	v4 = rotr64(tls, v4^v9, uint64(24))
	v3 = v3 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 15)))*8)))
	v14 = rotr64(tls, v14^v3, uint64(16))
	v9 = v9 + v14
	v4 = rotr64(tls, v4^v9, uint64(63))
	v0 = v0 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 1*16)))*8)))
	v12 = rotr64(tls, v12^v0, uint64(32))
	v8 = v8 + v12
	v4 = rotr64(tls, v4^v8, uint64(24))
	v0 = v0 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 1*16 + 1)))*8)))
	v12 = rotr64(tls, v12^v0, uint64(16))
	v8 = v8 + v12
	v4 = rotr64(tls, v4^v8, uint64(63))
	v1 = v1 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 1*16 + 2)))*8)))
	v13 = rotr64(tls, v13^v1, uint64(32))
	v9 = v9 + v13
	v5 = rotr64(tls, v5^v9, uint64(24))
	v1 = v1 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 1*16 + 3)))*8)))
	v13 = rotr64(tls, v13^v1, uint64(16))
	v9 = v9 + v13
	v5 = rotr64(tls, v5^v9, uint64(63))
	v2 = v2 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 1*16 + 4)))*8)))
	v14 = rotr64(tls, v14^v2, uint64(32))
	v10 = v10 + v14
	v6 = rotr64(tls, v6^v10, uint64(24))
	v2 = v2 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 1*16 + 5)))*8)))
	v14 = rotr64(tls, v14^v2, uint64(16))
	v10 = v10 + v14
	v6 = rotr64(tls, v6^v10, uint64(63))
	v3 = v3 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 1*16 + 6)))*8)))
	v15 = rotr64(tls, v15^v3, uint64(32))
	v11 = v11 + v15
	v7 = rotr64(tls, v7^v11, uint64(24))
	v3 = v3 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 1*16 + 7)))*8)))
	v15 = rotr64(tls, v15^v3, uint64(16))
	v11 = v11 + v15
	v7 = rotr64(tls, v7^v11, uint64(63))
	v0 = v0 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 1*16 + 8)))*8)))
	v15 = rotr64(tls, v15^v0, uint64(32))
	v10 = v10 + v15
	v5 = rotr64(tls, v5^v10, uint64(24))
	v0 = v0 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 1*16 + 9)))*8)))
	v15 = rotr64(tls, v15^v0, uint64(16))
	v10 = v10 + v15
	v5 = rotr64(tls, v5^v10, uint64(63))
	v1 = v1 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 1*16 + 10)))*8)))
	v12 = rotr64(tls, v12^v1, uint64(32))
	v11 = v11 + v12
	v6 = rotr64(tls, v6^v11, uint64(24))
	v1 = v1 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 1*16 + 11)))*8)))
	v12 = rotr64(tls, v12^v1, uint64(16))
	v11 = v11 + v12
	v6 = rotr64(tls, v6^v11, uint64(63))
	v2 = v2 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 1*16 + 12)))*8)))
	v13 = rotr64(tls, v13^v2, uint64(32))
	v8 = v8 + v13
	v7 = rotr64(tls, v7^v8, uint64(24))
	v2 = v2 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 1*16 + 13)))*8)))
	v13 = rotr64(tls, v13^v2, uint64(16))
	v8 = v8 + v13
	v7 = rotr64(tls, v7^v8, uint64(63))
	v3 = v3 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 1*16 + 14)))*8)))
	v14 = rotr64(tls, v14^v3, uint64(32))
	v9 = v9 + v14
	v4 = rotr64(tls, v4^v9, uint64(24))
	v3 = v3 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 1*16 + 15)))*8)))
	v14 = rotr64(tls, v14^v3, uint64(16))
	v9 = v9 + v14
	v4 = rotr64(tls, v4^v9, uint64(63))
	v0 = v0 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 2*16)))*8)))
	v12 = rotr64(tls, v12^v0, uint64(32))
	v8 = v8 + v12
	v4 = rotr64(tls, v4^v8, uint64(24))
	v0 = v0 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 2*16 + 1)))*8)))
	v12 = rotr64(tls, v12^v0, uint64(16))
	v8 = v8 + v12
	v4 = rotr64(tls, v4^v8, uint64(63))
	v1 = v1 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 2*16 + 2)))*8)))
	v13 = rotr64(tls, v13^v1, uint64(32))
	v9 = v9 + v13
	v5 = rotr64(tls, v5^v9, uint64(24))
	v1 = v1 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 2*16 + 3)))*8)))
	v13 = rotr64(tls, v13^v1, uint64(16))
	v9 = v9 + v13
	v5 = rotr64(tls, v5^v9, uint64(63))
	v2 = v2 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 2*16 + 4)))*8)))
	v14 = rotr64(tls, v14^v2, uint64(32))
	v10 = v10 + v14
	v6 = rotr64(tls, v6^v10, uint64(24))
	v2 = v2 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 2*16 + 5)))*8)))
	v14 = rotr64(tls, v14^v2, uint64(16))
	v10 = v10 + v14
	v6 = rotr64(tls, v6^v10, uint64(63))
	v3 = v3 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 2*16 + 6)))*8)))
	v15 = rotr64(tls, v15^v3, uint64(32))
	v11 = v11 + v15
	v7 = rotr64(tls, v7^v11, uint64(24))
	v3 = v3 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 2*16 + 7)))*8)))
	v15 = rotr64(tls, v15^v3, uint64(16))
	v11 = v11 + v15
	v7 = rotr64(tls, v7^v11, uint64(63))
	v0 = v0 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 2*16 + 8)))*8)))
	v15 = rotr64(tls, v15^v0, uint64(32))
	v10 = v10 + v15
	v5 = rotr64(tls, v5^v10, uint64(24))
	v0 = v0 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 2*16 + 9)))*8)))
	v15 = rotr64(tls, v15^v0, uint64(16))
	v10 = v10 + v15
	v5 = rotr64(tls, v5^v10, uint64(63))
	v1 = v1 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 2*16 + 10)))*8)))
	v12 = rotr64(tls, v12^v1, uint64(32))
	v11 = v11 + v12
	v6 = rotr64(tls, v6^v11, uint64(24))
	v1 = v1 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 2*16 + 11)))*8)))
	v12 = rotr64(tls, v12^v1, uint64(16))
	v11 = v11 + v12
	v6 = rotr64(tls, v6^v11, uint64(63))
	v2 = v2 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 2*16 + 12)))*8)))
	v13 = rotr64(tls, v13^v2, uint64(32))
	v8 = v8 + v13
	v7 = rotr64(tls, v7^v8, uint64(24))
	v2 = v2 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 2*16 + 13)))*8)))
	v13 = rotr64(tls, v13^v2, uint64(16))
	v8 = v8 + v13
	v7 = rotr64(tls, v7^v8, uint64(63))
	v3 = v3 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 2*16 + 14)))*8)))
	v14 = rotr64(tls, v14^v3, uint64(32))
	v9 = v9 + v14
	v4 = rotr64(tls, v4^v9, uint64(24))
	v3 = v3 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 2*16 + 15)))*8)))
	v14 = rotr64(tls, v14^v3, uint64(16))
	v9 = v9 + v14
	v4 = rotr64(tls, v4^v9, uint64(63))
	v0 = v0 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 3*16)))*8)))
	v12 = rotr64(tls, v12^v0, uint64(32))
	v8 = v8 + v12
	v4 = rotr64(tls, v4^v8, uint64(24))
	v0 = v0 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 3*16 + 1)))*8)))
	v12 = rotr64(tls, v12^v0, uint64(16))
	v8 = v8 + v12
	v4 = rotr64(tls, v4^v8, uint64(63))
	v1 = v1 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 3*16 + 2)))*8)))
	v13 = rotr64(tls, v13^v1, uint64(32))
	v9 = v9 + v13
	v5 = rotr64(tls, v5^v9, uint64(24))
	v1 = v1 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 3*16 + 3)))*8)))
	v13 = rotr64(tls, v13^v1, uint64(16))
	v9 = v9 + v13
	v5 = rotr64(tls, v5^v9, uint64(63))
	v2 = v2 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 3*16 + 4)))*8)))
	v14 = rotr64(tls, v14^v2, uint64(32))
	v10 = v10 + v14
	v6 = rotr64(tls, v6^v10, uint64(24))
	v2 = v2 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 3*16 + 5)))*8)))
	v14 = rotr64(tls, v14^v2, uint64(16))
	v10 = v10 + v14
	v6 = rotr64(tls, v6^v10, uint64(63))
	v3 = v3 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 3*16 + 6)))*8)))
	v15 = rotr64(tls, v15^v3, uint64(32))
	v11 = v11 + v15
	v7 = rotr64(tls, v7^v11, uint64(24))
	v3 = v3 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 3*16 + 7)))*8)))
	v15 = rotr64(tls, v15^v3, uint64(16))
	v11 = v11 + v15
	v7 = rotr64(tls, v7^v11, uint64(63))
	v0 = v0 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 3*16 + 8)))*8)))
	v15 = rotr64(tls, v15^v0, uint64(32))
	v10 = v10 + v15
	v5 = rotr64(tls, v5^v10, uint64(24))
	v0 = v0 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 3*16 + 9)))*8)))
	v15 = rotr64(tls, v15^v0, uint64(16))
	v10 = v10 + v15
	v5 = rotr64(tls, v5^v10, uint64(63))
	v1 = v1 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 3*16 + 10)))*8)))
	v12 = rotr64(tls, v12^v1, uint64(32))
	v11 = v11 + v12
	v6 = rotr64(tls, v6^v11, uint64(24))
	v1 = v1 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 3*16 + 11)))*8)))
	v12 = rotr64(tls, v12^v1, uint64(16))
	v11 = v11 + v12
	v6 = rotr64(tls, v6^v11, uint64(63))
	v2 = v2 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 3*16 + 12)))*8)))
	v13 = rotr64(tls, v13^v2, uint64(32))
	v8 = v8 + v13
	v7 = rotr64(tls, v7^v8, uint64(24))
	v2 = v2 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 3*16 + 13)))*8)))
	v13 = rotr64(tls, v13^v2, uint64(16))
	v8 = v8 + v13
	v7 = rotr64(tls, v7^v8, uint64(63))
	v3 = v3 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 3*16 + 14)))*8)))
	v14 = rotr64(tls, v14^v3, uint64(32))
	v9 = v9 + v14
	v4 = rotr64(tls, v4^v9, uint64(24))
	v3 = v3 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 3*16 + 15)))*8)))
	v14 = rotr64(tls, v14^v3, uint64(16))
	v9 = v9 + v14
	v4 = rotr64(tls, v4^v9, uint64(63))
	v0 = v0 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 4*16)))*8)))
	v12 = rotr64(tls, v12^v0, uint64(32))
	v8 = v8 + v12
	v4 = rotr64(tls, v4^v8, uint64(24))
	v0 = v0 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 4*16 + 1)))*8)))
	v12 = rotr64(tls, v12^v0, uint64(16))
	v8 = v8 + v12
	v4 = rotr64(tls, v4^v8, uint64(63))
	v1 = v1 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 4*16 + 2)))*8)))
	v13 = rotr64(tls, v13^v1, uint64(32))
	v9 = v9 + v13
	v5 = rotr64(tls, v5^v9, uint64(24))
	v1 = v1 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 4*16 + 3)))*8)))
	v13 = rotr64(tls, v13^v1, uint64(16))
	v9 = v9 + v13
	v5 = rotr64(tls, v5^v9, uint64(63))
	v2 = v2 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 4*16 + 4)))*8)))
	v14 = rotr64(tls, v14^v2, uint64(32))
	v10 = v10 + v14
	v6 = rotr64(tls, v6^v10, uint64(24))
	v2 = v2 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 4*16 + 5)))*8)))
	v14 = rotr64(tls, v14^v2, uint64(16))
	v10 = v10 + v14
	v6 = rotr64(tls, v6^v10, uint64(63))
	v3 = v3 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 4*16 + 6)))*8)))
	v15 = rotr64(tls, v15^v3, uint64(32))
	v11 = v11 + v15
	v7 = rotr64(tls, v7^v11, uint64(24))
	v3 = v3 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 4*16 + 7)))*8)))
	v15 = rotr64(tls, v15^v3, uint64(16))
	v11 = v11 + v15
	v7 = rotr64(tls, v7^v11, uint64(63))
	v0 = v0 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 4*16 + 8)))*8)))
	v15 = rotr64(tls, v15^v0, uint64(32))
	v10 = v10 + v15
	v5 = rotr64(tls, v5^v10, uint64(24))
	v0 = v0 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 4*16 + 9)))*8)))
	v15 = rotr64(tls, v15^v0, uint64(16))
	v10 = v10 + v15
	v5 = rotr64(tls, v5^v10, uint64(63))
	v1 = v1 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 4*16 + 10)))*8)))
	v12 = rotr64(tls, v12^v1, uint64(32))
	v11 = v11 + v12
	v6 = rotr64(tls, v6^v11, uint64(24))
	v1 = v1 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 4*16 + 11)))*8)))
	v12 = rotr64(tls, v12^v1, uint64(16))
	v11 = v11 + v12
	v6 = rotr64(tls, v6^v11, uint64(63))
	v2 = v2 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 4*16 + 12)))*8)))
	v13 = rotr64(tls, v13^v2, uint64(32))
	v8 = v8 + v13
	v7 = rotr64(tls, v7^v8, uint64(24))
	v2 = v2 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 4*16 + 13)))*8)))
	v13 = rotr64(tls, v13^v2, uint64(16))
	v8 = v8 + v13
	v7 = rotr64(tls, v7^v8, uint64(63))
	v3 = v3 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 4*16 + 14)))*8)))
	v14 = rotr64(tls, v14^v3, uint64(32))
	v9 = v9 + v14
	v4 = rotr64(tls, v4^v9, uint64(24))
	v3 = v3 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 4*16 + 15)))*8)))
	v14 = rotr64(tls, v14^v3, uint64(16))
	v9 = v9 + v14
	v4 = rotr64(tls, v4^v9, uint64(63))
	v0 = v0 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 5*16)))*8)))
	v12 = rotr64(tls, v12^v0, uint64(32))
	v8 = v8 + v12
	v4 = rotr64(tls, v4^v8, uint64(24))
	v0 = v0 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 5*16 + 1)))*8)))
	v12 = rotr64(tls, v12^v0, uint64(16))
	v8 = v8 + v12
	v4 = rotr64(tls, v4^v8, uint64(63))
	v1 = v1 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 5*16 + 2)))*8)))
	v13 = rotr64(tls, v13^v1, uint64(32))
	v9 = v9 + v13
	v5 = rotr64(tls, v5^v9, uint64(24))
	v1 = v1 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 5*16 + 3)))*8)))
	v13 = rotr64(tls, v13^v1, uint64(16))
	v9 = v9 + v13
	v5 = rotr64(tls, v5^v9, uint64(63))
	v2 = v2 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 5*16 + 4)))*8)))
	v14 = rotr64(tls, v14^v2, uint64(32))
	v10 = v10 + v14
	v6 = rotr64(tls, v6^v10, uint64(24))
	v2 = v2 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 5*16 + 5)))*8)))
	v14 = rotr64(tls, v14^v2, uint64(16))
	v10 = v10 + v14
	v6 = rotr64(tls, v6^v10, uint64(63))
	v3 = v3 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 5*16 + 6)))*8)))
	v15 = rotr64(tls, v15^v3, uint64(32))
	v11 = v11 + v15
	v7 = rotr64(tls, v7^v11, uint64(24))
	v3 = v3 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 5*16 + 7)))*8)))
	v15 = rotr64(tls, v15^v3, uint64(16))
	v11 = v11 + v15
	v7 = rotr64(tls, v7^v11, uint64(63))
	v0 = v0 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 5*16 + 8)))*8)))
	v15 = rotr64(tls, v15^v0, uint64(32))
	v10 = v10 + v15
	v5 = rotr64(tls, v5^v10, uint64(24))
	v0 = v0 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 5*16 + 9)))*8)))
	v15 = rotr64(tls, v15^v0, uint64(16))
	v10 = v10 + v15
	v5 = rotr64(tls, v5^v10, uint64(63))
	v1 = v1 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 5*16 + 10)))*8)))
	v12 = rotr64(tls, v12^v1, uint64(32))
	v11 = v11 + v12
	v6 = rotr64(tls, v6^v11, uint64(24))
	v1 = v1 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 5*16 + 11)))*8)))
	v12 = rotr64(tls, v12^v1, uint64(16))
	v11 = v11 + v12
	v6 = rotr64(tls, v6^v11, uint64(63))
	v2 = v2 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 5*16 + 12)))*8)))
	v13 = rotr64(tls, v13^v2, uint64(32))
	v8 = v8 + v13
	v7 = rotr64(tls, v7^v8, uint64(24))
	v2 = v2 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 5*16 + 13)))*8)))
	v13 = rotr64(tls, v13^v2, uint64(16))
	v8 = v8 + v13
	v7 = rotr64(tls, v7^v8, uint64(63))
	v3 = v3 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 5*16 + 14)))*8)))
	v14 = rotr64(tls, v14^v3, uint64(32))
	v9 = v9 + v14
	v4 = rotr64(tls, v4^v9, uint64(24))
	v3 = v3 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 5*16 + 15)))*8)))
	v14 = rotr64(tls, v14^v3, uint64(16))
	v9 = v9 + v14
	v4 = rotr64(tls, v4^v9, uint64(63))
	v0 = v0 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 6*16)))*8)))
	v12 = rotr64(tls, v12^v0, uint64(32))
	v8 = v8 + v12
	v4 = rotr64(tls, v4^v8, uint64(24))
	v0 = v0 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 6*16 + 1)))*8)))
	v12 = rotr64(tls, v12^v0, uint64(16))
	v8 = v8 + v12
	v4 = rotr64(tls, v4^v8, uint64(63))
	v1 = v1 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 6*16 + 2)))*8)))
	v13 = rotr64(tls, v13^v1, uint64(32))
	v9 = v9 + v13
	v5 = rotr64(tls, v5^v9, uint64(24))
	v1 = v1 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 6*16 + 3)))*8)))
	v13 = rotr64(tls, v13^v1, uint64(16))
	v9 = v9 + v13
	v5 = rotr64(tls, v5^v9, uint64(63))
	v2 = v2 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 6*16 + 4)))*8)))
	v14 = rotr64(tls, v14^v2, uint64(32))
	v10 = v10 + v14
	v6 = rotr64(tls, v6^v10, uint64(24))
	v2 = v2 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 6*16 + 5)))*8)))
	v14 = rotr64(tls, v14^v2, uint64(16))
	v10 = v10 + v14
	v6 = rotr64(tls, v6^v10, uint64(63))
	v3 = v3 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 6*16 + 6)))*8)))
	v15 = rotr64(tls, v15^v3, uint64(32))
	v11 = v11 + v15
	v7 = rotr64(tls, v7^v11, uint64(24))
	v3 = v3 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 6*16 + 7)))*8)))
	v15 = rotr64(tls, v15^v3, uint64(16))
	v11 = v11 + v15
	v7 = rotr64(tls, v7^v11, uint64(63))
	v0 = v0 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 6*16 + 8)))*8)))
	v15 = rotr64(tls, v15^v0, uint64(32))
	v10 = v10 + v15
	v5 = rotr64(tls, v5^v10, uint64(24))
	v0 = v0 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 6*16 + 9)))*8)))
	v15 = rotr64(tls, v15^v0, uint64(16))
	v10 = v10 + v15
	v5 = rotr64(tls, v5^v10, uint64(63))
	v1 = v1 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 6*16 + 10)))*8)))
	v12 = rotr64(tls, v12^v1, uint64(32))
	v11 = v11 + v12
	v6 = rotr64(tls, v6^v11, uint64(24))
	v1 = v1 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 6*16 + 11)))*8)))
	v12 = rotr64(tls, v12^v1, uint64(16))
	v11 = v11 + v12
	v6 = rotr64(tls, v6^v11, uint64(63))
	v2 = v2 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 6*16 + 12)))*8)))
	v13 = rotr64(tls, v13^v2, uint64(32))
	v8 = v8 + v13
	v7 = rotr64(tls, v7^v8, uint64(24))
	v2 = v2 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 6*16 + 13)))*8)))
	v13 = rotr64(tls, v13^v2, uint64(16))
	v8 = v8 + v13
	v7 = rotr64(tls, v7^v8, uint64(63))
	v3 = v3 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 6*16 + 14)))*8)))
	v14 = rotr64(tls, v14^v3, uint64(32))
	v9 = v9 + v14
	v4 = rotr64(tls, v4^v9, uint64(24))
	v3 = v3 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 6*16 + 15)))*8)))
	v14 = rotr64(tls, v14^v3, uint64(16))
	v9 = v9 + v14
	v4 = rotr64(tls, v4^v9, uint64(63))
	v0 = v0 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 7*16)))*8)))
	v12 = rotr64(tls, v12^v0, uint64(32))
	v8 = v8 + v12
	v4 = rotr64(tls, v4^v8, uint64(24))
	v0 = v0 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 7*16 + 1)))*8)))
	v12 = rotr64(tls, v12^v0, uint64(16))
	v8 = v8 + v12
	v4 = rotr64(tls, v4^v8, uint64(63))
	v1 = v1 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 7*16 + 2)))*8)))
	v13 = rotr64(tls, v13^v1, uint64(32))
	v9 = v9 + v13
	v5 = rotr64(tls, v5^v9, uint64(24))
	v1 = v1 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 7*16 + 3)))*8)))
	v13 = rotr64(tls, v13^v1, uint64(16))
	v9 = v9 + v13
	v5 = rotr64(tls, v5^v9, uint64(63))
	v2 = v2 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 7*16 + 4)))*8)))
	v14 = rotr64(tls, v14^v2, uint64(32))
	v10 = v10 + v14
	v6 = rotr64(tls, v6^v10, uint64(24))
	v2 = v2 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 7*16 + 5)))*8)))
	v14 = rotr64(tls, v14^v2, uint64(16))
	v10 = v10 + v14
	v6 = rotr64(tls, v6^v10, uint64(63))
	v3 = v3 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 7*16 + 6)))*8)))
	v15 = rotr64(tls, v15^v3, uint64(32))
	v11 = v11 + v15
	v7 = rotr64(tls, v7^v11, uint64(24))
	v3 = v3 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 7*16 + 7)))*8)))
	v15 = rotr64(tls, v15^v3, uint64(16))
	v11 = v11 + v15
	v7 = rotr64(tls, v7^v11, uint64(63))
	v0 = v0 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 7*16 + 8)))*8)))
	v15 = rotr64(tls, v15^v0, uint64(32))
	v10 = v10 + v15
	v5 = rotr64(tls, v5^v10, uint64(24))
	v0 = v0 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 7*16 + 9)))*8)))
	v15 = rotr64(tls, v15^v0, uint64(16))
	v10 = v10 + v15
	v5 = rotr64(tls, v5^v10, uint64(63))
	v1 = v1 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 7*16 + 10)))*8)))
	v12 = rotr64(tls, v12^v1, uint64(32))
	v11 = v11 + v12
	v6 = rotr64(tls, v6^v11, uint64(24))
	v1 = v1 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 7*16 + 11)))*8)))
	v12 = rotr64(tls, v12^v1, uint64(16))
	v11 = v11 + v12
	v6 = rotr64(tls, v6^v11, uint64(63))
	v2 = v2 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 7*16 + 12)))*8)))
	v13 = rotr64(tls, v13^v2, uint64(32))
	v8 = v8 + v13
	v7 = rotr64(tls, v7^v8, uint64(24))
	v2 = v2 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 7*16 + 13)))*8)))
	v13 = rotr64(tls, v13^v2, uint64(16))
	v8 = v8 + v13
	v7 = rotr64(tls, v7^v8, uint64(63))
	v3 = v3 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 7*16 + 14)))*8)))
	v14 = rotr64(tls, v14^v3, uint64(32))
	v9 = v9 + v14
	v4 = rotr64(tls, v4^v9, uint64(24))
	v3 = v3 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 7*16 + 15)))*8)))
	v14 = rotr64(tls, v14^v3, uint64(16))
	v9 = v9 + v14
	v4 = rotr64(tls, v4^v9, uint64(63))
	v0 = v0 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 8*16)))*8)))
	v12 = rotr64(tls, v12^v0, uint64(32))
	v8 = v8 + v12
	v4 = rotr64(tls, v4^v8, uint64(24))
	v0 = v0 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 8*16 + 1)))*8)))
	v12 = rotr64(tls, v12^v0, uint64(16))
	v8 = v8 + v12
	v4 = rotr64(tls, v4^v8, uint64(63))
	v1 = v1 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 8*16 + 2)))*8)))
	v13 = rotr64(tls, v13^v1, uint64(32))
	v9 = v9 + v13
	v5 = rotr64(tls, v5^v9, uint64(24))
	v1 = v1 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 8*16 + 3)))*8)))
	v13 = rotr64(tls, v13^v1, uint64(16))
	v9 = v9 + v13
	v5 = rotr64(tls, v5^v9, uint64(63))
	v2 = v2 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 8*16 + 4)))*8)))
	v14 = rotr64(tls, v14^v2, uint64(32))
	v10 = v10 + v14
	v6 = rotr64(tls, v6^v10, uint64(24))
	v2 = v2 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 8*16 + 5)))*8)))
	v14 = rotr64(tls, v14^v2, uint64(16))
	v10 = v10 + v14
	v6 = rotr64(tls, v6^v10, uint64(63))
	v3 = v3 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 8*16 + 6)))*8)))
	v15 = rotr64(tls, v15^v3, uint64(32))
	v11 = v11 + v15
	v7 = rotr64(tls, v7^v11, uint64(24))
	v3 = v3 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 8*16 + 7)))*8)))
	v15 = rotr64(tls, v15^v3, uint64(16))
	v11 = v11 + v15
	v7 = rotr64(tls, v7^v11, uint64(63))
	v0 = v0 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 8*16 + 8)))*8)))
	v15 = rotr64(tls, v15^v0, uint64(32))
	v10 = v10 + v15
	v5 = rotr64(tls, v5^v10, uint64(24))
	v0 = v0 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 8*16 + 9)))*8)))
	v15 = rotr64(tls, v15^v0, uint64(16))
	v10 = v10 + v15
	v5 = rotr64(tls, v5^v10, uint64(63))
	v1 = v1 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 8*16 + 10)))*8)))
	v12 = rotr64(tls, v12^v1, uint64(32))
	v11 = v11 + v12
	v6 = rotr64(tls, v6^v11, uint64(24))
	v1 = v1 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 8*16 + 11)))*8)))
	v12 = rotr64(tls, v12^v1, uint64(16))
	v11 = v11 + v12
	v6 = rotr64(tls, v6^v11, uint64(63))
	v2 = v2 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 8*16 + 12)))*8)))
	v13 = rotr64(tls, v13^v2, uint64(32))
	v8 = v8 + v13
	v7 = rotr64(tls, v7^v8, uint64(24))
	v2 = v2 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 8*16 + 13)))*8)))
	v13 = rotr64(tls, v13^v2, uint64(16))
	v8 = v8 + v13
	v7 = rotr64(tls, v7^v8, uint64(63))
	v3 = v3 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 8*16 + 14)))*8)))
	v14 = rotr64(tls, v14^v3, uint64(32))
	v9 = v9 + v14
	v4 = rotr64(tls, v4^v9, uint64(24))
	v3 = v3 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 8*16 + 15)))*8)))
	v14 = rotr64(tls, v14^v3, uint64(16))
	v9 = v9 + v14
	v4 = rotr64(tls, v4^v9, uint64(63))
	v0 = v0 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 9*16)))*8)))
	v12 = rotr64(tls, v12^v0, uint64(32))
	v8 = v8 + v12
	v4 = rotr64(tls, v4^v8, uint64(24))
	v0 = v0 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 9*16 + 1)))*8)))
	v12 = rotr64(tls, v12^v0, uint64(16))
	v8 = v8 + v12
	v4 = rotr64(tls, v4^v8, uint64(63))
	v1 = v1 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 9*16 + 2)))*8)))
	v13 = rotr64(tls, v13^v1, uint64(32))
	v9 = v9 + v13
	v5 = rotr64(tls, v5^v9, uint64(24))
	v1 = v1 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 9*16 + 3)))*8)))
	v13 = rotr64(tls, v13^v1, uint64(16))
	v9 = v9 + v13
	v5 = rotr64(tls, v5^v9, uint64(63))
	v2 = v2 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 9*16 + 4)))*8)))
	v14 = rotr64(tls, v14^v2, uint64(32))
	v10 = v10 + v14
	v6 = rotr64(tls, v6^v10, uint64(24))
	v2 = v2 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 9*16 + 5)))*8)))
	v14 = rotr64(tls, v14^v2, uint64(16))
	v10 = v10 + v14
	v6 = rotr64(tls, v6^v10, uint64(63))
	v3 = v3 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 9*16 + 6)))*8)))
	v15 = rotr64(tls, v15^v3, uint64(32))
	v11 = v11 + v15
	v7 = rotr64(tls, v7^v11, uint64(24))
	v3 = v3 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 9*16 + 7)))*8)))
	v15 = rotr64(tls, v15^v3, uint64(16))
	v11 = v11 + v15
	v7 = rotr64(tls, v7^v11, uint64(63))
	v0 = v0 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 9*16 + 8)))*8)))
	v15 = rotr64(tls, v15^v0, uint64(32))
	v10 = v10 + v15
	v5 = rotr64(tls, v5^v10, uint64(24))
	v0 = v0 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 9*16 + 9)))*8)))
	v15 = rotr64(tls, v15^v0, uint64(16))
	v10 = v10 + v15
	v5 = rotr64(tls, v5^v10, uint64(63))
	v1 = v1 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 9*16 + 10)))*8)))
	v12 = rotr64(tls, v12^v1, uint64(32))
	v11 = v11 + v12
	v6 = rotr64(tls, v6^v11, uint64(24))
	v1 = v1 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 9*16 + 11)))*8)))
	v12 = rotr64(tls, v12^v1, uint64(16))
	v11 = v11 + v12
	v6 = rotr64(tls, v6^v11, uint64(63))
	v2 = v2 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 9*16 + 12)))*8)))
	v13 = rotr64(tls, v13^v2, uint64(32))
	v8 = v8 + v13
	v7 = rotr64(tls, v7^v8, uint64(24))
	v2 = v2 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 9*16 + 13)))*8)))
	v13 = rotr64(tls, v13^v2, uint64(16))
	v8 = v8 + v13
	v7 = rotr64(tls, v7^v8, uint64(63))
	v3 = v3 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 9*16 + 14)))*8)))
	v14 = rotr64(tls, v14^v3, uint64(32))
	v9 = v9 + v14
	v4 = rotr64(tls, v4^v9, uint64(24))
	v3 = v3 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 9*16 + 15)))*8)))
	v14 = rotr64(tls, v14^v3, uint64(16))
	v9 = v9 + v14
	v4 = rotr64(tls, v4^v9, uint64(63))
	v0 = v0 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 10*16)))*8)))
	v12 = rotr64(tls, v12^v0, uint64(32))
	v8 = v8 + v12
	v4 = rotr64(tls, v4^v8, uint64(24))
	v0 = v0 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 10*16 + 1)))*8)))
	v12 = rotr64(tls, v12^v0, uint64(16))
	v8 = v8 + v12
	v4 = rotr64(tls, v4^v8, uint64(63))
	v1 = v1 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 10*16 + 2)))*8)))
	v13 = rotr64(tls, v13^v1, uint64(32))
	v9 = v9 + v13
	v5 = rotr64(tls, v5^v9, uint64(24))
	v1 = v1 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 10*16 + 3)))*8)))
	v13 = rotr64(tls, v13^v1, uint64(16))
	v9 = v9 + v13
	v5 = rotr64(tls, v5^v9, uint64(63))
	v2 = v2 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 10*16 + 4)))*8)))
	v14 = rotr64(tls, v14^v2, uint64(32))
	v10 = v10 + v14
	v6 = rotr64(tls, v6^v10, uint64(24))
	v2 = v2 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 10*16 + 5)))*8)))
	v14 = rotr64(tls, v14^v2, uint64(16))
	v10 = v10 + v14
	v6 = rotr64(tls, v6^v10, uint64(63))
	v3 = v3 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 10*16 + 6)))*8)))
	v15 = rotr64(tls, v15^v3, uint64(32))
	v11 = v11 + v15
	v7 = rotr64(tls, v7^v11, uint64(24))
	v3 = v3 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 10*16 + 7)))*8)))
	v15 = rotr64(tls, v15^v3, uint64(16))
	v11 = v11 + v15
	v7 = rotr64(tls, v7^v11, uint64(63))
	v0 = v0 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 10*16 + 8)))*8)))
	v15 = rotr64(tls, v15^v0, uint64(32))
	v10 = v10 + v15
	v5 = rotr64(tls, v5^v10, uint64(24))
	v0 = v0 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 10*16 + 9)))*8)))
	v15 = rotr64(tls, v15^v0, uint64(16))
	v10 = v10 + v15
	v5 = rotr64(tls, v5^v10, uint64(63))
	v1 = v1 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 10*16 + 10)))*8)))
	v12 = rotr64(tls, v12^v1, uint64(32))
	v11 = v11 + v12
	v6 = rotr64(tls, v6^v11, uint64(24))
	v1 = v1 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 10*16 + 11)))*8)))
	v12 = rotr64(tls, v12^v1, uint64(16))
	v11 = v11 + v12
	v6 = rotr64(tls, v6^v11, uint64(63))
	v2 = v2 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 10*16 + 12)))*8)))
	v13 = rotr64(tls, v13^v2, uint64(32))
	v8 = v8 + v13
	v7 = rotr64(tls, v7^v8, uint64(24))
	v2 = v2 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 10*16 + 13)))*8)))
	v13 = rotr64(tls, v13^v2, uint64(16))
	v8 = v8 + v13
	v7 = rotr64(tls, v7^v8, uint64(63))
	v3 = v3 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 10*16 + 14)))*8)))
	v14 = rotr64(tls, v14^v3, uint64(32))
	v9 = v9 + v14
	v4 = rotr64(tls, v4^v9, uint64(24))
	v3 = v3 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 10*16 + 15)))*8)))
	v14 = rotr64(tls, v14^v3, uint64(16))
	v9 = v9 + v14
	v4 = rotr64(tls, v4^v9, uint64(63))
	v0 = v0 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 11*16)))*8)))
	v12 = rotr64(tls, v12^v0, uint64(32))
	v8 = v8 + v12
	v4 = rotr64(tls, v4^v8, uint64(24))
	v0 = v0 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 11*16 + 1)))*8)))
	v12 = rotr64(tls, v12^v0, uint64(16))
	v8 = v8 + v12
	v4 = rotr64(tls, v4^v8, uint64(63))
	v1 = v1 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 11*16 + 2)))*8)))
	v13 = rotr64(tls, v13^v1, uint64(32))
	v9 = v9 + v13
	v5 = rotr64(tls, v5^v9, uint64(24))
	v1 = v1 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 11*16 + 3)))*8)))
	v13 = rotr64(tls, v13^v1, uint64(16))
	v9 = v9 + v13
	v5 = rotr64(tls, v5^v9, uint64(63))
	v2 = v2 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 11*16 + 4)))*8)))
	v14 = rotr64(tls, v14^v2, uint64(32))
	v10 = v10 + v14
	v6 = rotr64(tls, v6^v10, uint64(24))
	v2 = v2 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 11*16 + 5)))*8)))
	v14 = rotr64(tls, v14^v2, uint64(16))
	v10 = v10 + v14
	v6 = rotr64(tls, v6^v10, uint64(63))
	v3 = v3 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 11*16 + 6)))*8)))
	v15 = rotr64(tls, v15^v3, uint64(32))
	v11 = v11 + v15
	v7 = rotr64(tls, v7^v11, uint64(24))
	v3 = v3 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 11*16 + 7)))*8)))
	v15 = rotr64(tls, v15^v3, uint64(16))
	v11 = v11 + v15
	v7 = rotr64(tls, v7^v11, uint64(63))
	v0 = v0 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 11*16 + 8)))*8)))
	v15 = rotr64(tls, v15^v0, uint64(32))
	v10 = v10 + v15
	v5 = rotr64(tls, v5^v10, uint64(24))
	v0 = v0 + (v5 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 11*16 + 9)))*8)))
	v15 = rotr64(tls, v15^v0, uint64(16))
	v10 = v10 + v15
	v5 = rotr64(tls, v5^v10, uint64(63))
	v1 = v1 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 11*16 + 10)))*8)))
	v12 = rotr64(tls, v12^v1, uint64(32))
	v11 = v11 + v12
	v6 = rotr64(tls, v6^v11, uint64(24))
	v1 = v1 + (v6 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 11*16 + 11)))*8)))
	v12 = rotr64(tls, v12^v1, uint64(16))
	v11 = v11 + v12
	v6 = rotr64(tls, v6^v11, uint64(63))
	v2 = v2 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 11*16 + 12)))*8)))
	v13 = rotr64(tls, v13^v2, uint64(32))
	v8 = v8 + v13
	v7 = rotr64(tls, v7^v8, uint64(24))
	v2 = v2 + (v7 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 11*16 + 13)))*8)))
	v13 = rotr64(tls, v13^v2, uint64(16))
	v8 = v8 + v13
	v7 = rotr64(tls, v7^v8, uint64(63))
	v3 = v3 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 11*16 + 14)))*8)))
	v14 = rotr64(tls, v14^v3, uint64(32))
	v9 = v9 + v14
	v4 = rotr64(tls, v4^v9, uint64(24))
	v3 = v3 + (v4 + **(**u64)(__ccgo_up(input + uintptr(**(**u8)(__ccgo_up(uintptr(unsafe.Pointer(&sigma)) + 11*16 + 15)))*8)))
	v14 = rotr64(tls, v14^v3, uint64(16))
	v9 = v9 + v14
	v4 = rotr64(tls, v4^v9, uint64(63))
	// update hash
	**(**uint64_t)(__ccgo_up(ctx)) ^= v0 ^ v8
	**(**uint64_t)(__ccgo_up(ctx + 1*8)) ^= v1 ^ v9
	**(**uint64_t)(__ccgo_up(ctx + 2*8)) ^= v2 ^ v10
	**(**uint64_t)(__ccgo_up(ctx + 3*8)) ^= v3 ^ v11
	**(**uint64_t)(__ccgo_up(ctx + 4*8)) ^= v4 ^ v12
	**(**uint64_t)(__ccgo_up(ctx + 5*8)) ^= v5 ^ v13
	**(**uint64_t)(__ccgo_up(ctx + 6*8)) ^= v6 ^ v14
	**(**uint64_t)(__ccgo_up(ctx + 7*8)) ^= v7 ^ v15
}

var sigma = [12][16]u8{
	0: {
		1:  uint8(1),
		2:  uint8(2),
		3:  uint8(3),
		4:  uint8(4),
		5:  uint8(5),
		6:  uint8(6),
		7:  uint8(7),
		8:  uint8(8),
		9:  uint8(9),
		10: uint8(10),
		11: uint8(11),
		12: uint8(12),
		13: uint8(13),
		14: uint8(14),
		15: uint8(15),
	},
	1: {
		0:  uint8(14),
		1:  uint8(10),
		2:  uint8(4),
		3:  uint8(8),
		4:  uint8(9),
		5:  uint8(15),
		6:  uint8(13),
		7:  uint8(6),
		8:  uint8(1),
		9:  uint8(12),
		11: uint8(2),
		12: uint8(11),
		13: uint8(7),
		14: uint8(5),
		15: uint8(3),
	},
	2: {
		0:  uint8(11),
		1:  uint8(8),
		2:  uint8(12),
		4:  uint8(5),
		5:  uint8(2),
		6:  uint8(15),
		7:  uint8(13),
		8:  uint8(10),
		9:  uint8(14),
		10: uint8(3),
		11: uint8(6),
		12: uint8(7),
		13: uint8(1),
		14: uint8(9),
		15: uint8(4),
	},
	3: {
		0:  uint8(7),
		1:  uint8(9),
		2:  uint8(3),
		3:  uint8(1),
		4:  uint8(13),
		5:  uint8(12),
		6:  uint8(11),
		7:  uint8(14),
		8:  uint8(2),
		9:  uint8(6),
		10: uint8(5),
		11: uint8(10),
		12: uint8(4),
		14: uint8(15),
		15: uint8(8),
	},
	4: {
		0:  uint8(9),
		2:  uint8(5),
		3:  uint8(7),
		4:  uint8(2),
		5:  uint8(4),
		6:  uint8(10),
		7:  uint8(15),
		8:  uint8(14),
		9:  uint8(1),
		10: uint8(11),
		11: uint8(12),
		12: uint8(6),
		13: uint8(8),
		14: uint8(3),
		15: uint8(13),
	},
	5: {
		0:  uint8(2),
		1:  uint8(12),
		2:  uint8(6),
		3:  uint8(10),
		5:  uint8(11),
		6:  uint8(8),
		7:  uint8(3),
		8:  uint8(4),
		9:  uint8(13),
		10: uint8(7),
		11: uint8(5),
		12: uint8(15),
		13: uint8(14),
		14: uint8(1),
		15: uint8(9),
	},
	6: {
		0:  uint8(12),
		1:  uint8(5),
		2:  uint8(1),
		3:  uint8(15),
		4:  uint8(14),
		5:  uint8(13),
		6:  uint8(4),
		7:  uint8(10),
		9:  uint8(7),
		10: uint8(6),
		11: uint8(3),
		12: uint8(9),
		13: uint8(2),
		14: uint8(8),
		15: uint8(11),
	},
	7: {
		0:  uint8(13),
		1:  uint8(11),
		2:  uint8(7),
		3:  uint8(14),
		4:  uint8(12),
		5:  uint8(1),
		6:  uint8(3),
		7:  uint8(9),
		8:  uint8(5),
		10: uint8(15),
		11: uint8(4),
		12: uint8(8),
		13: uint8(6),
		14: uint8(2),
		15: uint8(10),
	},
	8: {
		0:  uint8(6),
		1:  uint8(15),
		2:  uint8(14),
		3:  uint8(9),
		4:  uint8(11),
		5:  uint8(3),
		7:  uint8(8),
		8:  uint8(12),
		9:  uint8(2),
		10: uint8(13),
		11: uint8(7),
		12: uint8(1),
		13: uint8(4),
		14: uint8(10),
		15: uint8(5),
	},
	9: {
		0:  uint8(10),
		1:  uint8(2),
		2:  uint8(8),
		3:  uint8(4),
		4:  uint8(7),
		5:  uint8(6),
		6:  uint8(1),
		7:  uint8(5),
		8:  uint8(15),
		9:  uint8(11),
		10: uint8(9),
		11: uint8(14),
		12: uint8(3),
		13: uint8(12),
		14: uint8(13),
	},
	10: {
		1:  uint8(1),
		2:  uint8(2),
		3:  uint8(3),
		4:  uint8(4),
		5:  uint8(5),
		6:  uint8(6),
		7:  uint8(7),
		8:  uint8(8),
		9:  uint8(9),
		10: uint8(10),
		11: uint8(11),
		12: uint8(12),
		13: uint8(13),
		14: uint8(14),
		15: uint8(15),
	},
	11: {
		0:  uint8(14),
		1:  uint8(10),
		2:  uint8(4),
		3:  uint8(8),
		4:  uint8(9),
		5:  uint8(15),
		6:  uint8(13),
		7:  uint8(6),
		8:  uint8(1),
		9:  uint8(12),
		11: uint8(2),
		12: uint8(11),
		13: uint8(7),
		14: uint8(5),
		15: uint8(3),
	},
}

func crypto_blake2b_keyed_init(tls *libc.TLS, ctx uintptr, hash_size size_t, key uintptr, key_size size_t) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var _i_, _i_1, _i_2 size_t
	var _ /* key_block at bp+0 */ [128]u8
	_, _, _ = _i_, _i_1, _i_2
	// initial hash
	_i_ = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(_i_ < libc.Uint64FromInt32(libc.Int32FromInt32(8))) {
			break
		}
		**(**uint64_t)(__ccgo_up(ctx + uintptr(_i_)*8)) = iv[_i_]
		goto _1
	_1:
		;
		_i_ = _i_ + 1
	}
	**(**uint64_t)(__ccgo_up(ctx)) ^= uint64(0x01010000) ^ key_size<<libc.Int32FromInt32(8) ^ hash_size
	**(**uint64_t)(__ccgo_up(ctx + 64)) = uint64(0)       // beginning of the input, no offset
	**(**uint64_t)(__ccgo_up(ctx + 64 + 1*8)) = uint64(0) // beginning of the input, no offset
	(*crypto_blake2b_ctx)(unsafe.Pointer(ctx)).Fhash_size = hash_size
	(*crypto_blake2b_ctx)(unsafe.Pointer(ctx)).Finput_idx = uint64(0)
	_i_1 = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(_i_1 < libc.Uint64FromInt32(libc.Int32FromInt32(16))) {
			break
		}
		**(**uint64_t)(__ccgo_up(ctx + 80 + uintptr(_i_1)*8)) = uint64(0)
		goto _2
	_2:
		;
		_i_1 = _i_1 + 1
	}
	// if there is a key, the first block is that key (padded with zeroes)
	if key_size > uint64(0) {
		**(**[128]u8)(__ccgo_up(bp)) = [128]u8{}
		_i_2 = libc.Uint64FromInt32(libc.Int32FromInt32(0))
		for {
			if !(_i_2 < key_size) {
				break
			}
			(**(**[128]u8)(__ccgo_up(bp)))[_i_2] = **(**u8)(__ccgo_up(key + uintptr(_i_2)))
			goto _3
		_3:
			;
			_i_2 = _i_2 + 1
		}
		// same as calling crypto_blake2b_update(ctx, key_block , 128)
		load64_le_buf(tls, ctx+80, bp, uint64(16))
		(*crypto_blake2b_ctx)(unsafe.Pointer(ctx)).Finput_idx = uint64(128)
		crypto_wipe(tls, bp, uint64(128))
	}
}

func crypto_blake2b_init(tls *libc.TLS, ctx uintptr, hash_size size_t) {
	crypto_blake2b_keyed_init(tls, ctx, hash_size, uintptr(0), uint64(0))
}

func crypto_blake2b_update(tls *libc.TLS, ctx uintptr, message uintptr, message_size size_t) {
	var _i_, byte1, byte11, i, i1, i2, nb_blocks, nb_bytes, nb_words, nb_words1, word, word1 size_t
	var v1 uint64
	_, _, _, _, _, _, _, _, _, _, _, _, _ = _i_, byte1, byte11, i, i1, i2, nb_blocks, nb_bytes, nb_words, nb_words1, word, word1, v1
	// Avoid undefined NULL pointer increments with empty messages
	if message_size == uint64(0) {
		return
	}
	// Align with word boundaries
	if (*crypto_blake2b_ctx)(unsafe.Pointer(ctx)).Finput_idx&uint64(7) != uint64(0) {
		if gap(tls, (*crypto_blake2b_ctx)(unsafe.Pointer(ctx)).Finput_idx, uint64(8)) <= message_size {
			v1 = gap(tls, (*crypto_blake2b_ctx)(unsafe.Pointer(ctx)).Finput_idx, uint64(8))
		} else {
			v1 = message_size
		}
		nb_bytes = v1
		word = (*crypto_blake2b_ctx)(unsafe.Pointer(ctx)).Finput_idx >> int32(3)
		byte1 = (*crypto_blake2b_ctx)(unsafe.Pointer(ctx)).Finput_idx & uint64(7)
		i = libc.Uint64FromInt32(libc.Int32FromInt32(0))
		for {
			if !(i < nb_bytes) {
				break
			}
			**(**uint64_t)(__ccgo_up(ctx + 80 + uintptr(word)*8)) |= uint64(**(**u8)(__ccgo_up(message + uintptr(i)))) << ((byte1 + i) << int32(3))
			goto _2
		_2:
			;
			i = i + 1
		}
		**(**size_t)(__ccgo_up(ctx + 208)) += nb_bytes
		message = message + uintptr(nb_bytes)
		message_size = message_size - nb_bytes
	}
	// Align with block boundaries (faster than byte by byte)
	if (*crypto_blake2b_ctx)(unsafe.Pointer(ctx)).Finput_idx&uint64(127) != uint64(0) {
		if gap(tls, (*crypto_blake2b_ctx)(unsafe.Pointer(ctx)).Finput_idx, uint64(128)) <= message_size {
			v1 = gap(tls, (*crypto_blake2b_ctx)(unsafe.Pointer(ctx)).Finput_idx, uint64(128))
		} else {
			v1 = message_size
		}
		nb_words = v1 >> int32(3)
		load64_le_buf(tls, ctx+80+uintptr((*crypto_blake2b_ctx)(unsafe.Pointer(ctx)).Finput_idx>>libc.Int32FromInt32(3))*8, message, nb_words)
		**(**size_t)(__ccgo_up(ctx + 208)) += nb_words << int32(3)
		message = message + uintptr(nb_words<<int32(3))
		message_size = message_size - nb_words<<int32(3)
	}
	// Process block by block
	nb_blocks = message_size >> int32(7)
	i1 = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(i1 < nb_blocks) {
			break
		}
		if (*crypto_blake2b_ctx)(unsafe.Pointer(ctx)).Finput_idx == uint64(128) {
			blake2b_compress(tls, ctx, 0)
		}
		load64_le_buf(tls, ctx+80, message, uint64(16))
		message = message + uintptr(128)
		(*crypto_blake2b_ctx)(unsafe.Pointer(ctx)).Finput_idx = uint64(128)
		goto _4
	_4:
		;
		i1 = i1 + 1
	}
	message_size = message_size & uint64(127)
	if message_size != uint64(0) {
		// Compress block & flush input buffer as needed
		if (*crypto_blake2b_ctx)(unsafe.Pointer(ctx)).Finput_idx == uint64(128) {
			blake2b_compress(tls, ctx, 0)
			(*crypto_blake2b_ctx)(unsafe.Pointer(ctx)).Finput_idx = uint64(0)
		}
		if (*crypto_blake2b_ctx)(unsafe.Pointer(ctx)).Finput_idx == uint64(0) {
			_i_ = libc.Uint64FromInt32(libc.Int32FromInt32(0))
			for {
				if !(_i_ < libc.Uint64FromInt32(libc.Int32FromInt32(16))) {
					break
				}
				**(**uint64_t)(__ccgo_up(ctx + 80 + uintptr(_i_)*8)) = uint64(0)
				goto _5
			_5:
				;
				_i_ = _i_ + 1
			}
		}
		// Fill remaining words (faster than byte by byte)
		nb_words1 = message_size >> int32(3)
		load64_le_buf(tls, ctx+80, message, nb_words1)
		**(**size_t)(__ccgo_up(ctx + 208)) += nb_words1 << int32(3)
		message = message + uintptr(nb_words1<<int32(3))
		message_size = message_size - nb_words1<<int32(3)
		// Fill remaining bytes
		i2 = libc.Uint64FromInt32(libc.Int32FromInt32(0))
		for {
			if !(i2 < message_size) {
				break
			}
			word1 = (*crypto_blake2b_ctx)(unsafe.Pointer(ctx)).Finput_idx >> int32(3)
			byte11 = (*crypto_blake2b_ctx)(unsafe.Pointer(ctx)).Finput_idx & uint64(7)
			**(**uint64_t)(__ccgo_up(ctx + 80 + uintptr(word1)*8)) |= uint64(**(**u8)(__ccgo_up(message + uintptr(i2)))) << (byte11 << int32(3))
			(*crypto_blake2b_ctx)(unsafe.Pointer(ctx)).Finput_idx = (*crypto_blake2b_ctx)(unsafe.Pointer(ctx)).Finput_idx + 1
			goto _6
		_6:
			;
			i2 = i2 + 1
		}
	}
}

func crypto_blake2b_final(tls *libc.TLS, ctx uintptr, hash uintptr) {
	var hash_size, i, nb_words size_t
	var v1 uint64
	_, _, _, _ = hash_size, i, nb_words, v1
	blake2b_compress(tls, ctx, int32(1))
	if (*crypto_blake2b_ctx)(unsafe.Pointer(ctx)).Fhash_size <= libc.Uint64FromInt32(libc.Int32FromInt32(64)) {
		v1 = (*crypto_blake2b_ctx)(unsafe.Pointer(ctx)).Fhash_size
	} else {
		v1 = libc.Uint64FromInt32(libc.Int32FromInt32(64))
	} // compress the last block
	hash_size = v1
	nb_words = hash_size >> int32(3)
	store64_le_buf(tls, hash, ctx, nb_words)
	i = nb_words << int32(3)
	for {
		if !(i < hash_size) {
			break
		}
		**(**u8)(__ccgo_up(hash + uintptr(i))) = uint8(**(**uint64_t)(__ccgo_up(ctx + uintptr(i>>int32(3))*8)) >> (libc.Uint64FromInt32(8) * (i & libc.Uint64FromInt32(7))) & uint64(0xff))
		goto _2
	_2:
		;
		i = i + 1
	}
	crypto_wipe(tls, ctx, uint64(224))
}

func crypto_blake2b_keyed(tls *libc.TLS, hash uintptr, hash_size size_t, key uintptr, key_size size_t, message uintptr, message_size size_t) {
	bp := tls.Alloc(224)
	defer tls.Free(224)
	var _ /* ctx at bp+0 */ crypto_blake2b_ctx
	crypto_blake2b_keyed_init(tls, bp, hash_size, key, key_size)
	crypto_blake2b_update(tls, bp, message, message_size)
	crypto_blake2b_final(tls, bp, hash)
}

func crypto_blake2b(tls *libc.TLS, hash uintptr, hash_size size_t, msg uintptr, msg_size size_t) {
	crypto_blake2b_keyed(tls, hash, hash_size, uintptr(0), uint64(0), msg, msg_size)
}

//////////////
/// Argon2 ///
//////////////
// references to R, Z, Q etc. come from the spec

// C documentation
//
//	// Argon2 operates on 1024 byte blocks.
type blk = struct {
	Fa [128]u64
}

// C documentation
//
//	// updates a BLAKE2 hash with a 32 bit word, little endian.
func blake_update_32(tls *libc.TLS, ctx uintptr, input u32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* buf at bp+0 */ [4]u8
	store32_le(tls, bp, input)
	crypto_blake2b_update(tls, ctx, bp, uint64(4))
	crypto_wipe(tls, bp, uint64(4))
}

func blake_update_32_buf(tls *libc.TLS, ctx uintptr, buf uintptr, size u32) {
	blake_update_32(tls, ctx, size)
	crypto_blake2b_update(tls, ctx, buf, uint64(size))
}

func copy_block(tls *libc.TLS, o uintptr, in uintptr) {
	var i size_t
	_ = i
	i = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(i < libc.Uint64FromInt32(libc.Int32FromInt32(128))) {
			break
		}
		**(**u64)(__ccgo_up(o + uintptr(i)*8)) = **(**u64)(__ccgo_up(in + uintptr(i)*8))
		goto _1
	_1:
		;
		i = i + 1
	}
}

func xor_block(tls *libc.TLS, o uintptr, in uintptr) {
	var i size_t
	_ = i
	i = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(i < libc.Uint64FromInt32(libc.Int32FromInt32(128))) {
			break
		}
		**(**u64)(__ccgo_up(o + uintptr(i)*8)) ^= **(**u64)(__ccgo_up(in + uintptr(i)*8))
		goto _1
	_1:
		;
		i = i + 1
	}
}

// C documentation
//
//	// Hash with a virtually unlimited digest size.
//	// Doesn't extract more entropy than the base hash function.
//	// Mainly used for filling a whole kilobyte block with pseudo-random bytes.
//	// (One could use a stream cipher with a seed hash as the key, but
//	//  this would introduce another dependency —and point of failure.)
func extended_hash(tls *libc.TLS, digest uintptr, digest_size u32, input uintptr, input_size u32) {
	bp := tls.Alloc(224)
	defer tls.Free(224)
	var i, in, out, r u32
	var v1 uint32
	var _ /* ctx at bp+0 */ crypto_blake2b_ctx
	_, _, _, _, _ = i, in, out, r, v1
	if digest_size <= libc.Uint32FromInt32(libc.Int32FromInt32(64)) {
		v1 = digest_size
	} else {
		v1 = libc.Uint32FromInt32(libc.Int32FromInt32(64))
	}
	crypto_blake2b_init(tls, bp, uint64(v1))
	blake_update_32(tls, bp, digest_size)
	crypto_blake2b_update(tls, bp, input, uint64(input_size))
	crypto_blake2b_final(tls, bp, digest)
	if digest_size > uint32(64) {
		// the conversion to u64 avoids integer overflow on
		// ludicrously big hash sizes.
		r = uint32((uint64(digest_size)+libc.Uint64FromInt32(31))>>libc.Int32FromInt32(5)) - uint32(2)
		i = uint32(1)
		in = uint32(0)
		out = uint32(32)
		for i < r {
			// Input and output overlap. This is intentional
			crypto_blake2b(tls, digest+uintptr(out), uint64(64), digest+uintptr(in), uint64(64))
			i = i + uint32(1)
			in = in + uint32(32)
			out = out + uint32(32)
		}
		crypto_blake2b(tls, digest+uintptr(out), uint64(digest_size-uint32(32)*r), digest+uintptr(in), uint64(64))
	}
}

// C documentation
//
//	// Core of the compression function G.  Computes Z from R in place.
func g_rounds(tls *libc.TLS, b uintptr) {
	var i, i1 int32
	_, _ = i, i1
	// column rounds (work_block = Q)
	i = 0
	for {
		if !(i < int32(128)) {
			break
		}
		**(**u64)(__ccgo_up(b + uintptr(i)*8)) += **(**u64)(__ccgo_up(b + uintptr(i+int32(4))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i)*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(4))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i+int32(12))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i)*8))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(12))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i+int32(12))*8)), uint64(32))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(8))*8)) += **(**u64)(__ccgo_up(b + uintptr(i+int32(12))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(8))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(12))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i+int32(4))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i+int32(8))*8))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(4))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i+int32(4))*8)), uint64(24))
		**(**u64)(__ccgo_up(b + uintptr(i)*8)) += **(**u64)(__ccgo_up(b + uintptr(i+int32(4))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i)*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(4))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i+int32(12))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i)*8))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(12))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i+int32(12))*8)), uint64(16))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(8))*8)) += **(**u64)(__ccgo_up(b + uintptr(i+int32(12))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(8))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(12))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i+int32(4))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i+int32(8))*8))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(4))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i+int32(4))*8)), uint64(63))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(1))*8)) += **(**u64)(__ccgo_up(b + uintptr(i+int32(5))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(1))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(5))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i+int32(13))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i+int32(1))*8))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(13))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i+int32(13))*8)), uint64(32))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(9))*8)) += **(**u64)(__ccgo_up(b + uintptr(i+int32(13))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(9))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(13))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i+int32(5))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i+int32(9))*8))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(5))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i+int32(5))*8)), uint64(24))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(1))*8)) += **(**u64)(__ccgo_up(b + uintptr(i+int32(5))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(1))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(5))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i+int32(13))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i+int32(1))*8))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(13))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i+int32(13))*8)), uint64(16))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(9))*8)) += **(**u64)(__ccgo_up(b + uintptr(i+int32(13))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(9))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(13))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i+int32(5))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i+int32(9))*8))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(5))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i+int32(5))*8)), uint64(63))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(2))*8)) += **(**u64)(__ccgo_up(b + uintptr(i+int32(6))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(2))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(6))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i+int32(14))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i+int32(2))*8))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(14))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i+int32(14))*8)), uint64(32))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(10))*8)) += **(**u64)(__ccgo_up(b + uintptr(i+int32(14))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(10))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(14))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i+int32(6))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i+int32(10))*8))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(6))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i+int32(6))*8)), uint64(24))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(2))*8)) += **(**u64)(__ccgo_up(b + uintptr(i+int32(6))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(2))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(6))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i+int32(14))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i+int32(2))*8))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(14))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i+int32(14))*8)), uint64(16))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(10))*8)) += **(**u64)(__ccgo_up(b + uintptr(i+int32(14))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(10))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(14))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i+int32(6))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i+int32(10))*8))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(6))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i+int32(6))*8)), uint64(63))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(3))*8)) += **(**u64)(__ccgo_up(b + uintptr(i+int32(7))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(3))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(7))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i+int32(15))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i+int32(3))*8))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(15))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i+int32(15))*8)), uint64(32))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(11))*8)) += **(**u64)(__ccgo_up(b + uintptr(i+int32(15))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(11))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(15))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i+int32(7))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i+int32(11))*8))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(7))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i+int32(7))*8)), uint64(24))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(3))*8)) += **(**u64)(__ccgo_up(b + uintptr(i+int32(7))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(3))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(7))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i+int32(15))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i+int32(3))*8))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(15))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i+int32(15))*8)), uint64(16))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(11))*8)) += **(**u64)(__ccgo_up(b + uintptr(i+int32(15))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(11))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(15))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i+int32(7))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i+int32(11))*8))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(7))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i+int32(7))*8)), uint64(63))
		**(**u64)(__ccgo_up(b + uintptr(i)*8)) += **(**u64)(__ccgo_up(b + uintptr(i+int32(5))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i)*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(5))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i+int32(15))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i)*8))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(15))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i+int32(15))*8)), uint64(32))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(10))*8)) += **(**u64)(__ccgo_up(b + uintptr(i+int32(15))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(10))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(15))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i+int32(5))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i+int32(10))*8))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(5))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i+int32(5))*8)), uint64(24))
		**(**u64)(__ccgo_up(b + uintptr(i)*8)) += **(**u64)(__ccgo_up(b + uintptr(i+int32(5))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i)*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(5))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i+int32(15))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i)*8))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(15))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i+int32(15))*8)), uint64(16))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(10))*8)) += **(**u64)(__ccgo_up(b + uintptr(i+int32(15))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(10))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(15))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i+int32(5))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i+int32(10))*8))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(5))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i+int32(5))*8)), uint64(63))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(1))*8)) += **(**u64)(__ccgo_up(b + uintptr(i+int32(6))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(1))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(6))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i+int32(12))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i+int32(1))*8))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(12))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i+int32(12))*8)), uint64(32))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(11))*8)) += **(**u64)(__ccgo_up(b + uintptr(i+int32(12))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(11))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(12))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i+int32(6))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i+int32(11))*8))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(6))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i+int32(6))*8)), uint64(24))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(1))*8)) += **(**u64)(__ccgo_up(b + uintptr(i+int32(6))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(1))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(6))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i+int32(12))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i+int32(1))*8))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(12))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i+int32(12))*8)), uint64(16))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(11))*8)) += **(**u64)(__ccgo_up(b + uintptr(i+int32(12))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(11))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(12))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i+int32(6))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i+int32(11))*8))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(6))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i+int32(6))*8)), uint64(63))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(2))*8)) += **(**u64)(__ccgo_up(b + uintptr(i+int32(7))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(2))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(7))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i+int32(13))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i+int32(2))*8))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(13))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i+int32(13))*8)), uint64(32))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(8))*8)) += **(**u64)(__ccgo_up(b + uintptr(i+int32(13))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(8))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(13))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i+int32(7))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i+int32(8))*8))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(7))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i+int32(7))*8)), uint64(24))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(2))*8)) += **(**u64)(__ccgo_up(b + uintptr(i+int32(7))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(2))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(7))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i+int32(13))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i+int32(2))*8))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(13))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i+int32(13))*8)), uint64(16))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(8))*8)) += **(**u64)(__ccgo_up(b + uintptr(i+int32(13))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(8))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(13))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i+int32(7))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i+int32(8))*8))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(7))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i+int32(7))*8)), uint64(63))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(3))*8)) += **(**u64)(__ccgo_up(b + uintptr(i+int32(4))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(3))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(4))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i+int32(14))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i+int32(3))*8))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(14))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i+int32(14))*8)), uint64(32))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(9))*8)) += **(**u64)(__ccgo_up(b + uintptr(i+int32(14))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(9))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(14))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i+int32(4))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i+int32(9))*8))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(4))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i+int32(4))*8)), uint64(24))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(3))*8)) += **(**u64)(__ccgo_up(b + uintptr(i+int32(4))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(3))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(4))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i+int32(14))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i+int32(3))*8))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(14))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i+int32(14))*8)), uint64(16))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(9))*8)) += **(**u64)(__ccgo_up(b + uintptr(i+int32(14))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(9))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i+int32(14))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i+int32(4))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i+int32(9))*8))
		**(**u64)(__ccgo_up(b + uintptr(i+int32(4))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i+int32(4))*8)), uint64(63))
		goto _1
	_1:
		;
		i = i + int32(16)
	}
	// row rounds (b = Z)
	i1 = 0
	for {
		if !(i1 < int32(16)) {
			break
		}
		**(**u64)(__ccgo_up(b + uintptr(i1)*8)) += **(**u64)(__ccgo_up(b + uintptr(i1+int32(32))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1)*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(32))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(96))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i1)*8))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(96))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i1+int32(96))*8)), uint64(32))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(64))*8)) += **(**u64)(__ccgo_up(b + uintptr(i1+int32(96))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(64))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(96))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(32))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i1+int32(64))*8))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(32))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i1+int32(32))*8)), uint64(24))
		**(**u64)(__ccgo_up(b + uintptr(i1)*8)) += **(**u64)(__ccgo_up(b + uintptr(i1+int32(32))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1)*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(32))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(96))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i1)*8))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(96))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i1+int32(96))*8)), uint64(16))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(64))*8)) += **(**u64)(__ccgo_up(b + uintptr(i1+int32(96))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(64))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(96))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(32))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i1+int32(64))*8))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(32))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i1+int32(32))*8)), uint64(63))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(1))*8)) += **(**u64)(__ccgo_up(b + uintptr(i1+int32(33))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(1))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(33))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(97))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i1+int32(1))*8))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(97))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i1+int32(97))*8)), uint64(32))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(65))*8)) += **(**u64)(__ccgo_up(b + uintptr(i1+int32(97))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(65))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(97))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(33))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i1+int32(65))*8))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(33))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i1+int32(33))*8)), uint64(24))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(1))*8)) += **(**u64)(__ccgo_up(b + uintptr(i1+int32(33))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(1))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(33))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(97))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i1+int32(1))*8))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(97))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i1+int32(97))*8)), uint64(16))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(65))*8)) += **(**u64)(__ccgo_up(b + uintptr(i1+int32(97))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(65))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(97))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(33))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i1+int32(65))*8))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(33))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i1+int32(33))*8)), uint64(63))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(16))*8)) += **(**u64)(__ccgo_up(b + uintptr(i1+int32(48))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(16))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(48))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(112))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i1+int32(16))*8))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(112))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i1+int32(112))*8)), uint64(32))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(80))*8)) += **(**u64)(__ccgo_up(b + uintptr(i1+int32(112))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(80))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(112))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(48))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i1+int32(80))*8))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(48))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i1+int32(48))*8)), uint64(24))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(16))*8)) += **(**u64)(__ccgo_up(b + uintptr(i1+int32(48))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(16))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(48))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(112))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i1+int32(16))*8))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(112))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i1+int32(112))*8)), uint64(16))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(80))*8)) += **(**u64)(__ccgo_up(b + uintptr(i1+int32(112))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(80))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(112))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(48))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i1+int32(80))*8))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(48))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i1+int32(48))*8)), uint64(63))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(17))*8)) += **(**u64)(__ccgo_up(b + uintptr(i1+int32(49))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(17))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(49))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(113))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i1+int32(17))*8))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(113))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i1+int32(113))*8)), uint64(32))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(81))*8)) += **(**u64)(__ccgo_up(b + uintptr(i1+int32(113))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(81))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(113))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(49))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i1+int32(81))*8))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(49))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i1+int32(49))*8)), uint64(24))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(17))*8)) += **(**u64)(__ccgo_up(b + uintptr(i1+int32(49))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(17))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(49))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(113))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i1+int32(17))*8))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(113))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i1+int32(113))*8)), uint64(16))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(81))*8)) += **(**u64)(__ccgo_up(b + uintptr(i1+int32(113))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(81))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(113))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(49))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i1+int32(81))*8))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(49))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i1+int32(49))*8)), uint64(63))
		**(**u64)(__ccgo_up(b + uintptr(i1)*8)) += **(**u64)(__ccgo_up(b + uintptr(i1+int32(33))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1)*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(33))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(113))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i1)*8))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(113))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i1+int32(113))*8)), uint64(32))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(80))*8)) += **(**u64)(__ccgo_up(b + uintptr(i1+int32(113))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(80))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(113))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(33))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i1+int32(80))*8))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(33))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i1+int32(33))*8)), uint64(24))
		**(**u64)(__ccgo_up(b + uintptr(i1)*8)) += **(**u64)(__ccgo_up(b + uintptr(i1+int32(33))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1)*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(33))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(113))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i1)*8))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(113))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i1+int32(113))*8)), uint64(16))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(80))*8)) += **(**u64)(__ccgo_up(b + uintptr(i1+int32(113))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(80))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(113))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(33))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i1+int32(80))*8))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(33))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i1+int32(33))*8)), uint64(63))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(1))*8)) += **(**u64)(__ccgo_up(b + uintptr(i1+int32(48))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(1))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(48))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(96))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i1+int32(1))*8))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(96))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i1+int32(96))*8)), uint64(32))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(81))*8)) += **(**u64)(__ccgo_up(b + uintptr(i1+int32(96))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(81))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(96))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(48))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i1+int32(81))*8))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(48))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i1+int32(48))*8)), uint64(24))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(1))*8)) += **(**u64)(__ccgo_up(b + uintptr(i1+int32(48))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(1))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(48))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(96))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i1+int32(1))*8))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(96))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i1+int32(96))*8)), uint64(16))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(81))*8)) += **(**u64)(__ccgo_up(b + uintptr(i1+int32(96))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(81))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(96))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(48))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i1+int32(81))*8))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(48))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i1+int32(48))*8)), uint64(63))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(16))*8)) += **(**u64)(__ccgo_up(b + uintptr(i1+int32(49))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(16))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(49))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(97))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i1+int32(16))*8))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(97))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i1+int32(97))*8)), uint64(32))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(64))*8)) += **(**u64)(__ccgo_up(b + uintptr(i1+int32(97))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(64))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(97))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(49))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i1+int32(64))*8))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(49))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i1+int32(49))*8)), uint64(24))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(16))*8)) += **(**u64)(__ccgo_up(b + uintptr(i1+int32(49))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(16))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(49))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(97))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i1+int32(16))*8))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(97))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i1+int32(97))*8)), uint64(16))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(64))*8)) += **(**u64)(__ccgo_up(b + uintptr(i1+int32(97))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(64))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(97))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(49))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i1+int32(64))*8))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(49))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i1+int32(49))*8)), uint64(63))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(17))*8)) += **(**u64)(__ccgo_up(b + uintptr(i1+int32(32))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(17))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(32))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(112))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i1+int32(17))*8))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(112))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i1+int32(112))*8)), uint64(32))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(65))*8)) += **(**u64)(__ccgo_up(b + uintptr(i1+int32(112))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(65))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(112))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(32))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i1+int32(65))*8))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(32))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i1+int32(32))*8)), uint64(24))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(17))*8)) += **(**u64)(__ccgo_up(b + uintptr(i1+int32(32))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(17))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(32))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(112))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i1+int32(17))*8))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(112))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i1+int32(112))*8)), uint64(16))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(65))*8)) += **(**u64)(__ccgo_up(b + uintptr(i1+int32(112))*8)) + uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(65))*8))))*uint64(uint32(**(**u64)(__ccgo_up(b + uintptr(i1+int32(112))*8))))<<int32(1)
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(32))*8)) ^= **(**u64)(__ccgo_up(b + uintptr(i1+int32(65))*8))
		**(**u64)(__ccgo_up(b + uintptr(i1+int32(32))*8)) = rotr64(tls, **(**u64)(__ccgo_up(b + uintptr(i1+int32(32))*8)), uint64(63))
		goto _2
	_2:
		;
		i1 = i1 + int32(2)
	}
}

func crypto_argon2(tls *libc.TLS, hash uintptr, hash_size u32, work_area uintptr, config crypto_argon2_config, inputs crypto_argon2_inputs, extras crypto_argon2_extras) {
	bp := tls.Alloc(4400)
	defer tls.Free(4400)
	var _i_, _i_1, _i_2 size_t
	var block, i, index, index_ctr, l, lane, lane1, lane_offset, lane_size, nb_blocks, nb_segments, next_slice, pass, pass_offset, ref, segment, segment_size, slice, slice_offset, window_size, window_start u32
	var blocks, current, last_block, next_block, p, previous, reference, segment_start, v8 uintptr
	var constant_time, v5 int32
	var index_seed, j1, x, y, z u64
	var v10, v11, v12, v13, v14 uint32
	var _ /* ctx at bp+72 */ crypto_blake2b_ctx
	var _ /* final_block at bp+3368 */ [1024]u8
	var _ /* hash_area at bp+296 */ [1024]u8
	var _ /* index_block at bp+2344 */ blk
	var _ /* initial_hash at bp+0 */ [72]u8
	var _ /* tmp at bp+1320 */ blk
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = _i_, _i_1, _i_2, block, blocks, constant_time, current, i, index, index_ctr, index_seed, j1, l, lane, lane1, lane_offset, lane_size, last_block, nb_blocks, nb_segments, next_block, next_slice, p, pass, pass_offset, previous, ref, reference, segment, segment_size, segment_start, slice, slice_offset, window_size, window_start, x, y, z, v10, v11, v12, v13, v14, v5, v8
	segment_size = config.Fnb_blocks / config.Fnb_lanes / uint32(4)
	lane_size = segment_size * uint32(4)
	nb_blocks = lane_size * config.Fnb_lanes // rounding down
	// work area seen as blocks (must be suitably aligned)
	blocks = work_area
	crypto_blake2b_init(tls, bp+72, uint64(64))
	blake_update_32(tls, bp+72, config.Fnb_lanes) // p: number of "threads"
	blake_update_32(tls, bp+72, hash_size)
	blake_update_32(tls, bp+72, config.Fnb_blocks)
	blake_update_32(tls, bp+72, config.Fnb_passes)
	blake_update_32(tls, bp+72, uint32(0x13))      // v: version number
	blake_update_32(tls, bp+72, config.Falgorithm) // y: Argon2i, Argon2d...
	blake_update_32_buf(tls, bp+72, inputs.Fpass, inputs.Fpass_size)
	blake_update_32_buf(tls, bp+72, inputs.Fsalt, inputs.Fsalt_size)
	blake_update_32_buf(tls, bp+72, extras.Fkey, extras.Fkey_size)
	blake_update_32_buf(tls, bp+72, extras.Fad, extras.Fad_size)
	crypto_blake2b_final(tls, bp+72, bp)
	l = libc.Uint32FromInt32(libc.Int32FromInt32(0))
	for {
		if !(l < config.Fnb_lanes) {
			break
		}
		i = libc.Uint32FromInt32(libc.Int32FromInt32(0))
		for {
			if !(i < libc.Uint32FromInt32(libc.Int32FromInt32(2))) {
				break
			}
			store32_le(tls, bp+uintptr(64), i) // first  additional word
			store32_le(tls, bp+uintptr(68), l) // second additional word
			extended_hash(tls, bp+296, uint32(1024), bp, uint32(72))
			load64_le_buf(tls, blocks+uintptr(l*lane_size+i)*1024, bp+296, uint64(128))
			goto _2
		_2:
			;
			i = i + 1
		}
		goto _1
	_1:
		;
		l = l + 1
	}
	crypto_wipe(tls, bp, uint64(72))
	crypto_wipe(tls, bp+296, uint64(1024))
	// Argon2i and Argon2id start with constant time indexing
	constant_time = libc.BoolInt32(config.Falgorithm != uint32(CRYPTO_ARGON2_D))
	pass = libc.Uint32FromInt32(libc.Int32FromInt32(0))
	for {
		if !(pass < config.Fnb_passes) {
			break
		}
		slice = libc.Uint32FromInt32(libc.Int32FromInt32(0))
		for {
			if !(slice < libc.Uint32FromInt32(libc.Int32FromInt32(4))) {
				break
			}
			if pass == uint32(0) && slice == uint32(0) {
				v5 = int32(2)
			} else {
				v5 = 0
			}
			// On the first slice of the first pass,
			// blocks 0 and 1 are already filled, hence pass_offset.
			pass_offset = libc.Uint32FromInt32(v5)
			slice_offset = slice * segment_size
			// Argon2id switches back to non-constant time indexing
			// after the first two slices of the first pass
			if slice == uint32(2) && config.Falgorithm == uint32(CRYPTO_ARGON2_ID) {
				constant_time = 0
			}
			// Each iteration of the following loop may be performed in
			// a separate thread.  All segments must be fully completed
			// before we start filling the next slice.
			segment = libc.Uint32FromInt32(libc.Int32FromInt32(0))
			for {
				if !(segment < config.Fnb_lanes) {
					break
				}
				index_ctr = uint32(1)
				block = pass_offset
				for {
					if !(block < segment_size) {
						break
					}
					// Current and previous blocks
					lane_offset = segment * lane_size
					segment_start = blocks + uintptr(lane_offset)*1024 + uintptr(slice_offset)*1024
					current = segment_start + uintptr(block)*1024
					if block == uint32(0) && slice_offset == uint32(0) {
						v8 = segment_start + uintptr(lane_size)*1024 - uintptr(1)*1024
					} else {
						v8 = segment_start + uintptr(block)*1024 - uintptr(1)*1024
					}
					previous = v8
					if constant_time != 0 {
						if block == pass_offset || block%uint32(128) == uint32(0) {
							// Fill or refresh deterministic indices block
							// seed the beginning of the block...
							_i_ = libc.Uint64FromInt32(libc.Int32FromInt32(0))
							for {
								if !(_i_ < libc.Uint64FromInt32(libc.Int32FromInt32(128))) {
									break
								}
								**(**u64)(__ccgo_up(bp + 2344 + uintptr(_i_)*8)) = uint64(0)
								goto _9
							_9:
								;
								_i_ = _i_ + 1
							}
							**(**u64)(__ccgo_up(bp + 2344)) = uint64(pass)
							**(**u64)(__ccgo_up(bp + 2344 + 1*8)) = uint64(segment)
							**(**u64)(__ccgo_up(bp + 2344 + 2*8)) = uint64(slice)
							**(**u64)(__ccgo_up(bp + 2344 + 3*8)) = uint64(nb_blocks)
							**(**u64)(__ccgo_up(bp + 2344 + 4*8)) = uint64(config.Fnb_passes)
							**(**u64)(__ccgo_up(bp + 2344 + 5*8)) = uint64(config.Falgorithm)
							**(**u64)(__ccgo_up(bp + 2344 + 6*8)) = uint64(index_ctr)
							index_ctr = index_ctr + 1
							// ... then shuffle it
							copy_block(tls, bp+1320, bp+2344)
							g_rounds(tls, bp+2344)
							xor_block(tls, bp+2344, bp+1320)
							copy_block(tls, bp+1320, bp+2344)
							g_rounds(tls, bp+2344)
							xor_block(tls, bp+2344, bp+1320)
						}
						index_seed = **(**u64)(__ccgo_up(bp + 2344 + uintptr(block%uint32(128))*8))
					} else {
						index_seed = **(**u64)(__ccgo_up(previous))
					}
					// Establish the reference set.  *Approximately* comprises:
					// - The last 3 slices (if they exist yet)
					// - The already constructed blocks in the current segment
					next_slice = (slice + uint32(1)) % uint32(4) * segment_size
					if pass == uint32(0) {
						v10 = uint32(0)
					} else {
						v10 = next_slice
					}
					window_start = v10
					if pass == uint32(0) {
						v11 = slice
					} else {
						v11 = uint32(3)
					}
					nb_segments = v11
					if pass == uint32(0) && slice == uint32(0) {
						v12 = segment
					} else {
						v12 = uint32(index_seed>>libc.Int32FromInt32(32)) % config.Fnb_lanes
					}
					lane = v12
					if lane == segment {
						v13 = block - uint32(1)
					} else {
						if block == uint32(0) {
							v14 = libc.Uint32FromInt32(-libc.Int32FromInt32(1))
						} else {
							v14 = uint32(0)
						}
						v13 = v14
					}
					window_size = nb_segments*segment_size + v13
					// Find reference block
					j1 = index_seed & uint64(0xffffffff) // block selector
					x = j1 * j1 >> int32(32)
					y = uint64(window_size) * x >> int32(32)
					z = uint64(window_size-libc.Uint32FromInt32(1)) - y
					ref = uint32((uint64(window_start) + z) % uint64(lane_size))
					index = lane*lane_size + ref
					reference = blocks + uintptr(index)*1024
					// Shuffle the previous & reference block
					// into the current block
					copy_block(tls, bp+1320, previous)
					xor_block(tls, bp+1320, reference)
					if pass == uint32(0) {
						copy_block(tls, current, bp+1320)
					} else {
						xor_block(tls, current, bp+1320)
					}
					g_rounds(tls, bp+1320)
					xor_block(tls, current, bp+1320)
					goto _7
				_7:
					;
					block = block + 1
				}
				goto _6
			_6:
				;
				segment = segment + 1
			}
			goto _4
		_4:
			;
			slice = slice + 1
		}
		goto _3
	_3:
		;
		pass = pass + 1
	}
	// Wipe temporary block
	p = bp + 1320
	_i_1 = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(_i_1 < libc.Uint64FromInt32(libc.Int32FromInt32(128))) {
			break
		}
		libc.AtomicStorePUint64(p+uintptr(_i_1)*8, uint64(0))
		goto _15
	_15:
		;
		_i_1 = _i_1 + 1
	}
	// XOR last blocks of each lane
	last_block = blocks + uintptr(lane_size)*1024 - uintptr(1)*1024
	lane1 = libc.Uint32FromInt32(libc.Int32FromInt32(1))
	for {
		if !(lane1 < config.Fnb_lanes) {
			break
		}
		next_block = last_block + uintptr(lane_size)*1024
		xor_block(tls, next_block, last_block)
		last_block = next_block
		goto _16
	_16:
		;
		lane1 = lane1 + 1
	}
	store64_le_buf(tls, bp+3368, last_block, uint64(128))
	// Wipe work area
	p = work_area
	_i_2 = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(_i_2 < uint64(libc.Uint32FromInt32(128)*nb_blocks)) {
			break
		}
		libc.AtomicStorePUint64(p+uintptr(_i_2)*8, uint64(0))
		goto _17
	_17:
		;
		_i_2 = _i_2 + 1
	}
	// Hash the very last block with H' into the output hash
	extended_hash(tls, hash, hash_size, bp+3368, uint32(1024))
	crypto_wipe(tls, bp+3368, uint64(1024))
}

////////////////////////////////////
/// Arithmetic modulo 2^255 - 19 ///
////////////////////////////////////
//  Originally taken from SUPERCOP's ref10 implementation.
//  A bit bigger than TweetNaCl, over 4 times faster.

// C documentation
//
//	// field element
type fe = [10]i32

// C documentation
//
//	// field constants
//	//
//	// fe_one      : 1
//	// sqrtm1      : sqrt(-1)
//	// d           :     -121665 / 121666
//	// D2          : 2 * -121665 / 121666
//	// lop_x, lop_y: low order point in Edwards coordinates
//	// ufactor     : -sqrt(-1) * 2
//	// A2          : 486662^2  (A squared)
var fe_one = fe{
	0: int32(1),
}
var sqrtm1 = fe{
	0: -int32(32595792),
	1: -int32(7943725),
	2: int32(9377950),
	3: int32(3500415),
	4: int32(12389472),
	5: -int32(272473),
	6: -int32(25146209),
	7: -int32(2005654),
	8: int32(326686),
	9: int32(11406482),
}
var d = fe{
	0: -int32(10913610),
	1: int32(13857413),
	2: -int32(15372611),
	3: int32(6949391),
	4: int32(114729),
	5: -int32(8787816),
	6: -int32(6275908),
	7: -int32(3247719),
	8: -int32(18696448),
	9: -int32(12055116),
}
var D2 = fe{
	0: -int32(21827239),
	1: -int32(5839606),
	2: -int32(30745221),
	3: int32(13898782),
	4: int32(229458),
	5: int32(15978800),
	6: -int32(12551817),
	7: -int32(6495438),
	8: int32(29715968),
	9: int32(9444199),
}
var lop_x = fe{
	0: int32(21352778),
	1: int32(5345713),
	2: int32(4660180),
	3: -int32(8347857),
	4: int32(24143090),
	5: int32(14568123),
	6: int32(30185756),
	7: -int32(12247770),
	8: -int32(33528939),
	9: int32(8345319),
}
var lop_y = fe{
	0: -int32(6952922),
	1: -int32(1265500),
	2: int32(6862341),
	3: -int32(7057498),
	4: -int32(4037696),
	5: -int32(5447722),
	6: int32(31680899),
	7: -int32(15325402),
	8: -int32(19365852),
	9: int32(1569102),
}
var ufactor = fe{
	0: -int32(1917299),
	1: int32(15887451),
	2: -int32(18755900),
	3: -int32(7000830),
	4: -int32(24778944),
	5: int32(544946),
	6: -int32(16816446),
	7: int32(4011309),
	8: -int32(653372),
	9: int32(10741468),
}
var A2 = fe{
	0: int32(12721188),
	1: int32(3529),
}

func fe_0(tls *libc.TLS, h uintptr) {
	var _i_ size_t
	_ = _i_
	_i_ = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(_i_ < libc.Uint64FromInt32(libc.Int32FromInt32(10))) {
			break
		}
		**(**i32)(__ccgo_up(h + uintptr(_i_)*4)) = 0
		goto _1
	_1:
		;
		_i_ = _i_ + 1
	}
}

func fe_1(tls *libc.TLS, h uintptr) {
	var _i_ size_t
	_ = _i_
	**(**i32)(__ccgo_up(h)) = int32(1)
	_i_ = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(_i_ < libc.Uint64FromInt32(libc.Int32FromInt32(9))) {
			break
		}
		**(**i32)(__ccgo_up(h + libc.UintptrFromInt32(1)*4 + uintptr(_i_)*4)) = 0
		goto _1
	_1:
		;
		_i_ = _i_ + 1
	}
}

func fe_copy(tls *libc.TLS, h uintptr, f uintptr) {
	var i size_t
	_ = i
	i = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(i < libc.Uint64FromInt32(libc.Int32FromInt32(10))) {
			break
		}
		**(**i32)(__ccgo_up(h + uintptr(i)*4)) = **(**i32)(__ccgo_up(f + uintptr(i)*4))
		goto _1
	_1:
		;
		i = i + 1
	}
}

func fe_neg(tls *libc.TLS, h uintptr, f uintptr) {
	var i size_t
	_ = i
	i = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(i < libc.Uint64FromInt32(libc.Int32FromInt32(10))) {
			break
		}
		**(**i32)(__ccgo_up(h + uintptr(i)*4)) = -**(**i32)(__ccgo_up(f + uintptr(i)*4))
		goto _1
	_1:
		;
		i = i + 1
	}
}

func fe_add(tls *libc.TLS, h uintptr, f uintptr, g uintptr) {
	var i size_t
	_ = i
	i = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(i < libc.Uint64FromInt32(libc.Int32FromInt32(10))) {
			break
		}
		**(**i32)(__ccgo_up(h + uintptr(i)*4)) = **(**i32)(__ccgo_up(f + uintptr(i)*4)) + **(**i32)(__ccgo_up(g + uintptr(i)*4))
		goto _1
	_1:
		;
		i = i + 1
	}
}

func fe_sub(tls *libc.TLS, h uintptr, f uintptr, g uintptr) {
	var i size_t
	_ = i
	i = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(i < libc.Uint64FromInt32(libc.Int32FromInt32(10))) {
			break
		}
		**(**i32)(__ccgo_up(h + uintptr(i)*4)) = **(**i32)(__ccgo_up(f + uintptr(i)*4)) - **(**i32)(__ccgo_up(g + uintptr(i)*4))
		goto _1
	_1:
		;
		i = i + 1
	}
}

// C documentation
//
//	// Some compilers, when inlining fe_cswap() or fe_ccopy(), may introduce
//	// a timing leak.  It happens when it notices `b` has only 2 possible
//	// values, and either replace the arithmetic by a secret dependent
//	// branch, or (as has been observed), swap pointers instead of values,
//	// which intruduces a secret dependent index.
//	//
//	// We apply two mitigations here:
//	// - Add `volatile` in the mask declaration.
//	// - Unroll the copy loop (costs couple hundred bytes of binary code).
//	//
//	// As of June 2026, those mitigation work when applied separately or
//	// together.  Applying them both is currently overkill, but may help
//	// delay the day compilers grow clever enough to defeat it.  (The true
//	// fix is in the semantics of the language itself: C currently has no
//	// way to specify constant time code).
//	//
//	// Note (Loup): as of June 2026, the problem has yet to surface in
//	// fe_cswap(), but since this is almost the same code as fe_ccopy() I
//	// believe it is more prudent to apply the precaution there too.
func fe_cswap(tls *libc.TLS, f uintptr, g uintptr, b int32) {
	var mask, x0, x1, x2, x3, x4, x5, x6, x7, x8, x9 i32
	_, _, _, _, _, _, _, _, _, _, _ = mask, x0, x1, x2, x3, x4, x5, x6, x7, x8, x9
	mask = -b // -1 = 0xffffffff
	x0 = (**(**i32)(__ccgo_up(f)) ^ **(**i32)(__ccgo_up(g))) & mask
	**(**i32)(__ccgo_up(f)) = **(**i32)(__ccgo_up(f)) ^ x0
	**(**i32)(__ccgo_up(g)) = **(**i32)(__ccgo_up(g)) ^ x0
	x1 = (**(**i32)(__ccgo_up(f + 1*4)) ^ **(**i32)(__ccgo_up(g + 1*4))) & mask
	**(**i32)(__ccgo_up(f + 1*4)) = **(**i32)(__ccgo_up(f + 1*4)) ^ x1
	**(**i32)(__ccgo_up(g + 1*4)) = **(**i32)(__ccgo_up(g + 1*4)) ^ x1
	x2 = (**(**i32)(__ccgo_up(f + 2*4)) ^ **(**i32)(__ccgo_up(g + 2*4))) & mask
	**(**i32)(__ccgo_up(f + 2*4)) = **(**i32)(__ccgo_up(f + 2*4)) ^ x2
	**(**i32)(__ccgo_up(g + 2*4)) = **(**i32)(__ccgo_up(g + 2*4)) ^ x2
	x3 = (**(**i32)(__ccgo_up(f + 3*4)) ^ **(**i32)(__ccgo_up(g + 3*4))) & mask
	**(**i32)(__ccgo_up(f + 3*4)) = **(**i32)(__ccgo_up(f + 3*4)) ^ x3
	**(**i32)(__ccgo_up(g + 3*4)) = **(**i32)(__ccgo_up(g + 3*4)) ^ x3
	x4 = (**(**i32)(__ccgo_up(f + 4*4)) ^ **(**i32)(__ccgo_up(g + 4*4))) & mask
	**(**i32)(__ccgo_up(f + 4*4)) = **(**i32)(__ccgo_up(f + 4*4)) ^ x4
	**(**i32)(__ccgo_up(g + 4*4)) = **(**i32)(__ccgo_up(g + 4*4)) ^ x4
	x5 = (**(**i32)(__ccgo_up(f + 5*4)) ^ **(**i32)(__ccgo_up(g + 5*4))) & mask
	**(**i32)(__ccgo_up(f + 5*4)) = **(**i32)(__ccgo_up(f + 5*4)) ^ x5
	**(**i32)(__ccgo_up(g + 5*4)) = **(**i32)(__ccgo_up(g + 5*4)) ^ x5
	x6 = (**(**i32)(__ccgo_up(f + 6*4)) ^ **(**i32)(__ccgo_up(g + 6*4))) & mask
	**(**i32)(__ccgo_up(f + 6*4)) = **(**i32)(__ccgo_up(f + 6*4)) ^ x6
	**(**i32)(__ccgo_up(g + 6*4)) = **(**i32)(__ccgo_up(g + 6*4)) ^ x6
	x7 = (**(**i32)(__ccgo_up(f + 7*4)) ^ **(**i32)(__ccgo_up(g + 7*4))) & mask
	**(**i32)(__ccgo_up(f + 7*4)) = **(**i32)(__ccgo_up(f + 7*4)) ^ x7
	**(**i32)(__ccgo_up(g + 7*4)) = **(**i32)(__ccgo_up(g + 7*4)) ^ x7
	x8 = (**(**i32)(__ccgo_up(f + 8*4)) ^ **(**i32)(__ccgo_up(g + 8*4))) & mask
	**(**i32)(__ccgo_up(f + 8*4)) = **(**i32)(__ccgo_up(f + 8*4)) ^ x8
	**(**i32)(__ccgo_up(g + 8*4)) = **(**i32)(__ccgo_up(g + 8*4)) ^ x8
	x9 = (**(**i32)(__ccgo_up(f + 9*4)) ^ **(**i32)(__ccgo_up(g + 9*4))) & mask
	**(**i32)(__ccgo_up(f + 9*4)) = **(**i32)(__ccgo_up(f + 9*4)) ^ x9
	**(**i32)(__ccgo_up(g + 9*4)) = **(**i32)(__ccgo_up(g + 9*4)) ^ x9
}

func fe_ccopy(tls *libc.TLS, f uintptr, g uintptr, b int32) {
	var mask, x0, x1, x2, x3, x4, x5, x6, x7, x8, x9 i32
	_, _, _, _, _, _, _, _, _, _, _ = mask, x0, x1, x2, x3, x4, x5, x6, x7, x8, x9
	mask = -b // -1 = 0xffffffff
	x0 = (**(**i32)(__ccgo_up(f)) ^ **(**i32)(__ccgo_up(g))) & mask
	**(**i32)(__ccgo_up(f)) = **(**i32)(__ccgo_up(f)) ^ x0
	x1 = (**(**i32)(__ccgo_up(f + 1*4)) ^ **(**i32)(__ccgo_up(g + 1*4))) & mask
	**(**i32)(__ccgo_up(f + 1*4)) = **(**i32)(__ccgo_up(f + 1*4)) ^ x1
	x2 = (**(**i32)(__ccgo_up(f + 2*4)) ^ **(**i32)(__ccgo_up(g + 2*4))) & mask
	**(**i32)(__ccgo_up(f + 2*4)) = **(**i32)(__ccgo_up(f + 2*4)) ^ x2
	x3 = (**(**i32)(__ccgo_up(f + 3*4)) ^ **(**i32)(__ccgo_up(g + 3*4))) & mask
	**(**i32)(__ccgo_up(f + 3*4)) = **(**i32)(__ccgo_up(f + 3*4)) ^ x3
	x4 = (**(**i32)(__ccgo_up(f + 4*4)) ^ **(**i32)(__ccgo_up(g + 4*4))) & mask
	**(**i32)(__ccgo_up(f + 4*4)) = **(**i32)(__ccgo_up(f + 4*4)) ^ x4
	x5 = (**(**i32)(__ccgo_up(f + 5*4)) ^ **(**i32)(__ccgo_up(g + 5*4))) & mask
	**(**i32)(__ccgo_up(f + 5*4)) = **(**i32)(__ccgo_up(f + 5*4)) ^ x5
	x6 = (**(**i32)(__ccgo_up(f + 6*4)) ^ **(**i32)(__ccgo_up(g + 6*4))) & mask
	**(**i32)(__ccgo_up(f + 6*4)) = **(**i32)(__ccgo_up(f + 6*4)) ^ x6
	x7 = (**(**i32)(__ccgo_up(f + 7*4)) ^ **(**i32)(__ccgo_up(g + 7*4))) & mask
	**(**i32)(__ccgo_up(f + 7*4)) = **(**i32)(__ccgo_up(f + 7*4)) ^ x7
	x8 = (**(**i32)(__ccgo_up(f + 8*4)) ^ **(**i32)(__ccgo_up(g + 8*4))) & mask
	**(**i32)(__ccgo_up(f + 8*4)) = **(**i32)(__ccgo_up(f + 8*4)) ^ x8
	x9 = (**(**i32)(__ccgo_up(f + 9*4)) ^ **(**i32)(__ccgo_up(g + 9*4))) & mask
	**(**i32)(__ccgo_up(f + 9*4)) = **(**i32)(__ccgo_up(f + 9*4)) ^ x9
}

// Signed carry propagation
// ------------------------
//
// Let t be a number.  It can be uniquely decomposed thus:
//
//    t = h*2^26 + l
//    such that -2^25 <= l < 2^25
//
// Let c = (t + 2^25) / 2^26            (rounded down)
//     c = (h*2^26 + l + 2^25) / 2^26   (rounded down)
//     c =  h   +   (l + 2^25) / 2^26   (rounded down)
//     c =  h                           (exactly)
// Because 0 <= l + 2^25 < 2^26
//
// Let u = t          - c*2^26
//     u = h*2^26 + l - h*2^26
//     u = l
// Therefore, -2^25 <= u < 2^25
//
// Additionally, if |t| < x, then |h| < x/2^26 (rounded down)
//
// Notations:
// - In C, 1<<25 means 2^25.
// - In C, x>>25 means floor(x / (2^25)).
// - All of the above applies with 25 & 24 as well as 26 & 25.
//
//
// Note on negative right shifts
// -----------------------------
//
// In C, x >> n, where x is a negative integer, is implementation
// defined.  In practice, all platforms do arithmetic shift, which is
// equivalent to division by 2^26, rounded down.  Some compilers, like
// GCC, even guarantee it.
//
// If we ever stumble upon a platform that does not propagate the sign
// bit (we won't), visible failures will show at the slightest test, and
// the signed shifts can be replaced by the following:
//
//     typedef struct { i64 x:39; } s25;
//     typedef struct { i64 x:38; } s26;
//     i64 shift25(i64 x) { s25 s; s.x = ((u64)x)>>25; return s.x; }
//     i64 shift26(i64 x) { s26 s; s.x = ((u64)x)>>26; return s.x; }
//
// Current compilers cannot optimise this, causing a 30% drop in
// performance.  Fairly expensive for something that never happens.
//
//
// Precondition
// ------------
//
// |t0|       < 2^63
// |t1|..|t9| < 2^62
//
// Algorithm
// ---------
// c   = t0 + 2^25 / 2^26   -- |c|  <= 2^36
// t0 -= c * 2^26           -- |t0| <= 2^25
// t1 += c                  -- |t1| <= 2^63
//
// c   = t4 + 2^25 / 2^26   -- |c|  <= 2^36
// t4 -= c * 2^26           -- |t4| <= 2^25
// t5 += c                  -- |t5| <= 2^63
//
// c   = t1 + 2^24 / 2^25   -- |c|  <= 2^38
// t1 -= c * 2^25           -- |t1| <= 2^24
// t2 += c                  -- |t2| <= 2^63
//
// c   = t5 + 2^24 / 2^25   -- |c|  <= 2^38
// t5 -= c * 2^25           -- |t5| <= 2^24
// t6 += c                  -- |t6| <= 2^63
//
// c   = t2 + 2^25 / 2^26   -- |c|  <= 2^37
// t2 -= c * 2^26           -- |t2| <= 2^25        < 1.1 * 2^25  (final t2)
// t3 += c                  -- |t3| <= 2^63
//
// c   = t6 + 2^25 / 2^26   -- |c|  <= 2^37
// t6 -= c * 2^26           -- |t6| <= 2^25        < 1.1 * 2^25  (final t6)
// t7 += c                  -- |t7| <= 2^63
//
// c   = t3 + 2^24 / 2^25   -- |c|  <= 2^38
// t3 -= c * 2^25           -- |t3| <= 2^24        < 1.1 * 2^24  (final t3)
// t4 += c                  -- |t4| <= 2^25 + 2^38 < 2^39
//
// c   = t7 + 2^24 / 2^25   -- |c|  <= 2^38
// t7 -= c * 2^25           -- |t7| <= 2^24        < 1.1 * 2^24  (final t7)
// t8 += c                  -- |t8| <= 2^63
//
// c   = t4 + 2^25 / 2^26   -- |c|  <= 2^13
// t4 -= c * 2^26           -- |t4| <= 2^25        < 1.1 * 2^25  (final t4)
// t5 += c                  -- |t5| <= 2^24 + 2^13 < 1.1 * 2^24  (final t5)
//
// c   = t8 + 2^25 / 2^26   -- |c|  <= 2^37
// t8 -= c * 2^26           -- |t8| <= 2^25        < 1.1 * 2^25  (final t8)
// t9 += c                  -- |t9| <= 2^63
//
// c   = t9 + 2^24 / 2^25   -- |c|  <= 2^38
// t9 -= c * 2^25           -- |t9| <= 2^24        < 1.1 * 2^24  (final t9)
// t0 += c * 19             -- |t0| <= 2^25 + 2^38*19 < 2^44
//
// c   = t0 + 2^25 / 2^26   -- |c|  <= 2^18
// t0 -= c * 2^26           -- |t0| <= 2^25        < 1.1 * 2^25  (final t0)
// t1 += c                  -- |t1| <= 2^24 + 2^18 < 1.1 * 2^24  (final t1)
//
// Postcondition
// -------------
//   |t0|, |t2|, |t4|, |t6|, |t8|  <  1.1 * 2^25
//   |t1|, |t3|, |t5|, |t7|, |t9|  <  1.1 * 2^24

// C documentation
//
//	// Decodes a field element from a byte buffer.
//	// mask specifies how many bits we ignore.
//	// Traditionally we ignore 1. It's useful for EdDSA,
//	// which uses that bit to denote the sign of x.
//	// Elligator however uses positive representatives,
//	// which means ignoring 2 bits instead.
func fe_frombytes_mask(tls *libc.TLS, h uintptr, s uintptr, nb_mask uint32) {
	var c, t0, t1, t2, t3, t4, t5, t6, t7, t8, t9 i64
	var mask u32
	_, _, _, _, _, _, _, _, _, _, _, _ = c, mask, t0, t1, t2, t3, t4, t5, t6, t7, t8, t9
	mask = libc.Uint32FromInt32(int32(0xffffff) >> nb_mask)
	t0 = libc.Int64FromUint32(load32_le(tls, s))                         // t0 < 2^32
	t1 = libc.Int64FromUint32(load24_le(tls, s+uintptr(4)) << int32(6))  // t1 < 2^30
	t2 = libc.Int64FromUint32(load24_le(tls, s+uintptr(7)) << int32(5))  // t2 < 2^29
	t3 = libc.Int64FromUint32(load24_le(tls, s+uintptr(10)) << int32(3)) // t3 < 2^27
	t4 = libc.Int64FromUint32(load24_le(tls, s+uintptr(13)) << int32(2)) // t4 < 2^26
	t5 = libc.Int64FromUint32(load32_le(tls, s+uintptr(16)))             // t5 < 2^32
	t6 = libc.Int64FromUint32(load24_le(tls, s+uintptr(20)) << int32(7)) // t6 < 2^31
	t7 = libc.Int64FromUint32(load24_le(tls, s+uintptr(23)) << int32(5)) // t7 < 2^29
	t8 = libc.Int64FromUint32(load24_le(tls, s+uintptr(26)) << int32(4)) // t8 < 2^28
	t9 = libc.Int64FromUint32(load24_le(tls, s+uintptr(29)) & mask << int32(2))
	c = (t0 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(25)) >> int32(26)
	t0 = t0 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(26))
	t1 = t1 + c
	c = (t4 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(25)) >> int32(26)
	t4 = t4 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(26))
	t5 = t5 + c
	c = (t1 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(24)) >> int32(25)
	t1 = t1 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(25))
	t2 = t2 + c
	c = (t5 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(24)) >> int32(25)
	t5 = t5 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(25))
	t6 = t6 + c
	c = (t2 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(25)) >> int32(26)
	t2 = t2 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(26))
	t3 = t3 + c
	c = (t6 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(25)) >> int32(26)
	t6 = t6 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(26))
	t7 = t7 + c
	c = (t3 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(24)) >> int32(25)
	t3 = t3 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(25))
	t4 = t4 + c
	c = (t7 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(24)) >> int32(25)
	t7 = t7 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(25))
	t8 = t8 + c
	c = (t4 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(25)) >> int32(26)
	t4 = t4 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(26))
	t5 = t5 + c
	c = (t8 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(25)) >> int32(26)
	t8 = t8 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(26))
	t9 = t9 + c
	c = (t9 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(24)) >> int32(25)
	t9 = t9 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(25))
	t0 = t0 + c*int64(19)
	c = (t0 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(25)) >> int32(26)
	t0 = t0 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(26))
	t1 = t1 + c
	**(**i32)(__ccgo_up(h)) = int32(t0)
	**(**i32)(__ccgo_up(h + 1*4)) = int32(t1)
	**(**i32)(__ccgo_up(h + 2*4)) = int32(t2)
	**(**i32)(__ccgo_up(h + 3*4)) = int32(t3)
	**(**i32)(__ccgo_up(h + 4*4)) = int32(t4)
	**(**i32)(__ccgo_up(h + 5*4)) = int32(t5)
	**(**i32)(__ccgo_up(h + 6*4)) = int32(t6)
	**(**i32)(__ccgo_up(h + 7*4)) = int32(t7)
	**(**i32)(__ccgo_up(h + 8*4)) = int32(t8)
	**(**i32)(__ccgo_up(h + 9*4)) = int32(t9) // Carry precondition OK
}

func fe_frombytes(tls *libc.TLS, h uintptr, s uintptr) {
	fe_frombytes_mask(tls, h, s, uint32(1))
}

// C documentation
//
//	// Precondition
//	//   |h[0]|, |h[2]|, |h[4]|, |h[6]|, |h[8]|  <  1.1 * 2^25
//	//   |h[1]|, |h[3]|, |h[5]|, |h[7]|, |h[9]|  <  1.1 * 2^24
//	//
//	// Therefore, |h| < 2^255-19
//	// There are two possibilities:
//	//
//	// - If h is positive, all we need to do is reduce its individual
//	//   limbs down to their tight positive range.
//	// - If h is negative, we also need to add 2^255-19 to it.
//	//   Or just remove 19 and chop off any excess bit.
func fe_tobytes(tls *libc.TLS, s uintptr, h uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var _i_, i, i1 size_t
	var q i32
	var _ /* t at bp+0 */ [10]i32
	_, _, _, _ = _i_, i, i1, q
	_i_ = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(_i_ < libc.Uint64FromInt32(libc.Int32FromInt32(10))) {
			break
		}
		(**(**[10]i32)(__ccgo_up(bp)))[_i_] = **(**i32)(__ccgo_up(h + uintptr(_i_)*4))
		goto _1
	_1:
		;
		_i_ = _i_ + 1
	}
	q = (int32(19)*(**(**[10]i32)(__ccgo_up(bp)))[int32(9)] + libc.Int32FromInt32(1)<<libc.Int32FromInt32(24)) >> int32(25)
	//                 |t9|                    < 1.1 * 2^24
	//  -1.1 * 2^24  <  t9                     < 1.1 * 2^24
	//  -21  * 2^24  <  19 * t9                < 21  * 2^24
	//  -2^29        <  19 * t9 + 2^24         < 2^29
	//  -2^29 / 2^25 < (19 * t9 + 2^24) / 2^25 < 2^29 / 2^25
	//  -16          < (19 * t9 + 2^24) / 2^25 < 16
	i = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(i < libc.Uint64FromInt32(libc.Int32FromInt32(5))) {
			break
		}
		q = q + (**(**[10]i32)(__ccgo_up(bp)))[uint64(2)*i]
		q = q >> int32(26) // q = 0 or -1
		q = q + (**(**[10]i32)(__ccgo_up(bp)))[uint64(2)*i+uint64(1)]
		q = q >> int32(25) // q = 0 or -1
		goto _2
	_2:
		;
		i = i + 1
	}
	// q =  0 iff h >= 0
	// q = -1 iff h <  0
	// Adding q * 19 to h reduces h to its proper range.
	q = q * int32(19) // Shift carry back to the beginning
	i1 = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(i1 < libc.Uint64FromInt32(libc.Int32FromInt32(5))) {
			break
		}
		**(**i32)(__ccgo_up(bp + uintptr(i1*uint64(2))*4)) += q
		q = (**(**[10]i32)(__ccgo_up(bp)))[i1*uint64(2)] >> int32(26)
		**(**i32)(__ccgo_up(bp + uintptr(i1*uint64(2))*4)) -= q * (libc.Int32FromInt32(1) << libc.Int32FromInt32(26))
		**(**i32)(__ccgo_up(bp + uintptr(i1*uint64(2)+uint64(1))*4)) += q
		q = (**(**[10]i32)(__ccgo_up(bp)))[i1*uint64(2)+uint64(1)] >> int32(25)
		**(**i32)(__ccgo_up(bp + uintptr(i1*uint64(2)+uint64(1))*4)) -= q * (libc.Int32FromInt32(1) << libc.Int32FromInt32(25))
		goto _3
	_3:
		;
		i1 = i1 + 1
	}
	// h is now fully reduced, and q represents the excess bit.
	store32_le(tls, s+uintptr(0), libc.Uint32FromInt32((**(**[10]i32)(__ccgo_up(bp)))[0])>>libc.Int32FromInt32(0)|libc.Uint32FromInt32((**(**[10]i32)(__ccgo_up(bp)))[int32(1)])<<libc.Int32FromInt32(26))
	store32_le(tls, s+uintptr(4), libc.Uint32FromInt32((**(**[10]i32)(__ccgo_up(bp)))[int32(1)])>>libc.Int32FromInt32(6)|libc.Uint32FromInt32((**(**[10]i32)(__ccgo_up(bp)))[int32(2)])<<libc.Int32FromInt32(19))
	store32_le(tls, s+uintptr(8), libc.Uint32FromInt32((**(**[10]i32)(__ccgo_up(bp)))[int32(2)])>>libc.Int32FromInt32(13)|libc.Uint32FromInt32((**(**[10]i32)(__ccgo_up(bp)))[int32(3)])<<libc.Int32FromInt32(13))
	store32_le(tls, s+uintptr(12), libc.Uint32FromInt32((**(**[10]i32)(__ccgo_up(bp)))[int32(3)])>>libc.Int32FromInt32(19)|libc.Uint32FromInt32((**(**[10]i32)(__ccgo_up(bp)))[int32(4)])<<libc.Int32FromInt32(6))
	store32_le(tls, s+uintptr(16), libc.Uint32FromInt32((**(**[10]i32)(__ccgo_up(bp)))[int32(5)])>>libc.Int32FromInt32(0)|libc.Uint32FromInt32((**(**[10]i32)(__ccgo_up(bp)))[int32(6)])<<libc.Int32FromInt32(25))
	store32_le(tls, s+uintptr(20), libc.Uint32FromInt32((**(**[10]i32)(__ccgo_up(bp)))[int32(6)])>>libc.Int32FromInt32(7)|libc.Uint32FromInt32((**(**[10]i32)(__ccgo_up(bp)))[int32(7)])<<libc.Int32FromInt32(19))
	store32_le(tls, s+uintptr(24), libc.Uint32FromInt32((**(**[10]i32)(__ccgo_up(bp)))[int32(7)])>>libc.Int32FromInt32(13)|libc.Uint32FromInt32((**(**[10]i32)(__ccgo_up(bp)))[int32(8)])<<libc.Int32FromInt32(12))
	store32_le(tls, s+uintptr(28), libc.Uint32FromInt32((**(**[10]i32)(__ccgo_up(bp)))[int32(8)])>>libc.Int32FromInt32(20)|libc.Uint32FromInt32((**(**[10]i32)(__ccgo_up(bp)))[int32(9)])<<libc.Int32FromInt32(6))
	crypto_wipe(tls, bp, uint64(40))
}

// C documentation
//
//	// Precondition
//	// -------------
//	//   |f0|, |f2|, |f4|, |f6|, |f8|  <  1.65 * 2^26
//	//   |f1|, |f3|, |f5|, |f7|, |f9|  <  1.65 * 2^25
//	//
//	//   |g0|, |g2|, |g4|, |g6|, |g8|  <  1.65 * 2^26
//	//   |g1|, |g3|, |g5|, |g7|, |g9|  <  1.65 * 2^25
func fe_mul_small(tls *libc.TLS, h uintptr, f uintptr, g i32) {
	var c, t0, t1, t2, t3, t4, t5, t6, t7, t8, t9 i64
	_, _, _, _, _, _, _, _, _, _, _ = c, t0, t1, t2, t3, t4, t5, t6, t7, t8, t9
	t0 = int64(**(**i32)(__ccgo_up(f))) * int64(g)
	t1 = int64(**(**i32)(__ccgo_up(f + 1*4))) * int64(g)
	t2 = int64(**(**i32)(__ccgo_up(f + 2*4))) * int64(g)
	t3 = int64(**(**i32)(__ccgo_up(f + 3*4))) * int64(g)
	t4 = int64(**(**i32)(__ccgo_up(f + 4*4))) * int64(g)
	t5 = int64(**(**i32)(__ccgo_up(f + 5*4))) * int64(g)
	t6 = int64(**(**i32)(__ccgo_up(f + 6*4))) * int64(g)
	t7 = int64(**(**i32)(__ccgo_up(f + 7*4))) * int64(g)
	t8 = int64(**(**i32)(__ccgo_up(f + 8*4))) * int64(g)
	t9 = int64(**(**i32)(__ccgo_up(f + 9*4))) * int64(g)
	c = (t0 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(25)) >> int32(26)
	t0 = t0 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(26))
	t1 = t1 + c
	c = (t4 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(25)) >> int32(26)
	t4 = t4 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(26))
	t5 = t5 + c
	c = (t1 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(24)) >> int32(25)
	t1 = t1 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(25))
	t2 = t2 + c
	c = (t5 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(24)) >> int32(25)
	t5 = t5 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(25))
	t6 = t6 + c
	c = (t2 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(25)) >> int32(26)
	t2 = t2 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(26))
	t3 = t3 + c
	c = (t6 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(25)) >> int32(26)
	t6 = t6 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(26))
	t7 = t7 + c
	c = (t3 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(24)) >> int32(25)
	t3 = t3 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(25))
	t4 = t4 + c
	c = (t7 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(24)) >> int32(25)
	t7 = t7 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(25))
	t8 = t8 + c
	c = (t4 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(25)) >> int32(26)
	t4 = t4 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(26))
	t5 = t5 + c
	c = (t8 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(25)) >> int32(26)
	t8 = t8 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(26))
	t9 = t9 + c
	c = (t9 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(24)) >> int32(25)
	t9 = t9 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(25))
	t0 = t0 + c*int64(19)
	c = (t0 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(25)) >> int32(26)
	t0 = t0 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(26))
	t1 = t1 + c
	**(**i32)(__ccgo_up(h)) = int32(t0)
	**(**i32)(__ccgo_up(h + 1*4)) = int32(t1)
	**(**i32)(__ccgo_up(h + 2*4)) = int32(t2)
	**(**i32)(__ccgo_up(h + 3*4)) = int32(t3)
	**(**i32)(__ccgo_up(h + 4*4)) = int32(t4)
	**(**i32)(__ccgo_up(h + 5*4)) = int32(t5)
	**(**i32)(__ccgo_up(h + 6*4)) = int32(t6)
	**(**i32)(__ccgo_up(h + 7*4)) = int32(t7)
	**(**i32)(__ccgo_up(h + 8*4)) = int32(t8)
	**(**i32)(__ccgo_up(h + 9*4)) = int32(t9) // Carry precondition OK
}

// C documentation
//
//	// Precondition
//	// -------------
//	//   |f0|, |f2|, |f4|, |f6|, |f8|  <  1.65 * 2^26
//	//   |f1|, |f3|, |f5|, |f7|, |f9|  <  1.65 * 2^25
//	//
//	//   |g0|, |g2|, |g4|, |g6|, |g8|  <  1.65 * 2^26
//	//   |g1|, |g3|, |g5|, |g7|, |g9|  <  1.65 * 2^25
func fe_mul(tls *libc.TLS, h uintptr, f uintptr, g uintptr) {
	var F1, F3, F5, F7, F9, G1, G2, G3, G4, G5, G6, G7, G8, G9, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, g0, g1, g2, g3, g4, g5, g6, g7, g8, g9 i32
	var c, t0, t1, t2, t3, t4, t5, t6, t7, t8, t9 i64
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = F1, F3, F5, F7, F9, G1, G2, G3, G4, G5, G6, G7, G8, G9, c, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, g0, g1, g2, g3, g4, g5, g6, g7, g8, g9, t0, t1, t2, t3, t4, t5, t6, t7, t8, t9
	// Everything is unrolled and put in temporary variables.
	// We could roll the loop, but that would make curve25519 twice as slow.
	f0 = **(**i32)(__ccgo_up(f))
	f1 = **(**i32)(__ccgo_up(f + 1*4))
	f2 = **(**i32)(__ccgo_up(f + 2*4))
	f3 = **(**i32)(__ccgo_up(f + 3*4))
	f4 = **(**i32)(__ccgo_up(f + 4*4))
	f5 = **(**i32)(__ccgo_up(f + 5*4))
	f6 = **(**i32)(__ccgo_up(f + 6*4))
	f7 = **(**i32)(__ccgo_up(f + 7*4))
	f8 = **(**i32)(__ccgo_up(f + 8*4))
	f9 = **(**i32)(__ccgo_up(f + 9*4))
	g0 = **(**i32)(__ccgo_up(g))
	g1 = **(**i32)(__ccgo_up(g + 1*4))
	g2 = **(**i32)(__ccgo_up(g + 2*4))
	g3 = **(**i32)(__ccgo_up(g + 3*4))
	g4 = **(**i32)(__ccgo_up(g + 4*4))
	g5 = **(**i32)(__ccgo_up(g + 5*4))
	g6 = **(**i32)(__ccgo_up(g + 6*4))
	g7 = **(**i32)(__ccgo_up(g + 7*4))
	g8 = **(**i32)(__ccgo_up(g + 8*4))
	g9 = **(**i32)(__ccgo_up(g + 9*4))
	F1 = f1 * int32(2)
	F3 = f3 * int32(2)
	F5 = f5 * int32(2)
	F7 = f7 * int32(2)
	F9 = f9 * int32(2)
	G1 = g1 * int32(19)
	G2 = g2 * int32(19)
	G3 = g3 * int32(19)
	G4 = g4 * int32(19)
	G5 = g5 * int32(19)
	G6 = g6 * int32(19)
	G7 = g7 * int32(19)
	G8 = g8 * int32(19)
	G9 = g9 * int32(19)
	// |F1|, |F3|, |F5|, |F7|, |F9|  <  1.65 * 2^26
	// |G0|, |G2|, |G4|, |G6|, |G8|  <  2^31
	// |G1|, |G3|, |G5|, |G7|, |G9|  <  2^30
	t0 = int64(f0)*int64(g0) + int64(F1)*int64(G9) + int64(f2)*int64(G8) + int64(F3)*int64(G7) + int64(f4)*int64(G6) + int64(F5)*int64(G5) + int64(f6)*int64(G4) + int64(F7)*int64(G3) + int64(f8)*int64(G2) + int64(F9)*int64(G1)
	t1 = int64(f0)*int64(g1) + int64(f1)*int64(g0) + int64(f2)*int64(G9) + int64(f3)*int64(G8) + int64(f4)*int64(G7) + int64(f5)*int64(G6) + int64(f6)*int64(G5) + int64(f7)*int64(G4) + int64(f8)*int64(G3) + int64(f9)*int64(G2)
	t2 = int64(f0)*int64(g2) + int64(F1)*int64(g1) + int64(f2)*int64(g0) + int64(F3)*int64(G9) + int64(f4)*int64(G8) + int64(F5)*int64(G7) + int64(f6)*int64(G6) + int64(F7)*int64(G5) + int64(f8)*int64(G4) + int64(F9)*int64(G3)
	t3 = int64(f0)*int64(g3) + int64(f1)*int64(g2) + int64(f2)*int64(g1) + int64(f3)*int64(g0) + int64(f4)*int64(G9) + int64(f5)*int64(G8) + int64(f6)*int64(G7) + int64(f7)*int64(G6) + int64(f8)*int64(G5) + int64(f9)*int64(G4)
	t4 = int64(f0)*int64(g4) + int64(F1)*int64(g3) + int64(f2)*int64(g2) + int64(F3)*int64(g1) + int64(f4)*int64(g0) + int64(F5)*int64(G9) + int64(f6)*int64(G8) + int64(F7)*int64(G7) + int64(f8)*int64(G6) + int64(F9)*int64(G5)
	t5 = int64(f0)*int64(g5) + int64(f1)*int64(g4) + int64(f2)*int64(g3) + int64(f3)*int64(g2) + int64(f4)*int64(g1) + int64(f5)*int64(g0) + int64(f6)*int64(G9) + int64(f7)*int64(G8) + int64(f8)*int64(G7) + int64(f9)*int64(G6)
	t6 = int64(f0)*int64(g6) + int64(F1)*int64(g5) + int64(f2)*int64(g4) + int64(F3)*int64(g3) + int64(f4)*int64(g2) + int64(F5)*int64(g1) + int64(f6)*int64(g0) + int64(F7)*int64(G9) + int64(f8)*int64(G8) + int64(F9)*int64(G7)
	t7 = int64(f0)*int64(g7) + int64(f1)*int64(g6) + int64(f2)*int64(g5) + int64(f3)*int64(g4) + int64(f4)*int64(g3) + int64(f5)*int64(g2) + int64(f6)*int64(g1) + int64(f7)*int64(g0) + int64(f8)*int64(G9) + int64(f9)*int64(G8)
	t8 = int64(f0)*int64(g8) + int64(F1)*int64(g7) + int64(f2)*int64(g6) + int64(F3)*int64(g5) + int64(f4)*int64(g4) + int64(F5)*int64(g3) + int64(f6)*int64(g2) + int64(F7)*int64(g1) + int64(f8)*int64(g0) + int64(F9)*int64(G9)
	t9 = int64(f0)*int64(g9) + int64(f1)*int64(g8) + int64(f2)*int64(g7) + int64(f3)*int64(g6) + int64(f4)*int64(g5) + int64(f5)*int64(g4) + int64(f6)*int64(g3) + int64(f7)*int64(g2) + int64(f8)*int64(g1) + int64(f9)*int64(g0)
	c = (t0 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(25)) >> int32(26)
	t0 = t0 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(26))
	t1 = t1 + c
	c = (t4 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(25)) >> int32(26)
	t4 = t4 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(26))
	t5 = t5 + c
	c = (t1 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(24)) >> int32(25)
	t1 = t1 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(25))
	t2 = t2 + c
	c = (t5 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(24)) >> int32(25)
	t5 = t5 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(25))
	t6 = t6 + c
	c = (t2 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(25)) >> int32(26)
	t2 = t2 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(26))
	t3 = t3 + c
	c = (t6 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(25)) >> int32(26)
	t6 = t6 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(26))
	t7 = t7 + c
	c = (t3 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(24)) >> int32(25)
	t3 = t3 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(25))
	t4 = t4 + c
	c = (t7 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(24)) >> int32(25)
	t7 = t7 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(25))
	t8 = t8 + c
	c = (t4 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(25)) >> int32(26)
	t4 = t4 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(26))
	t5 = t5 + c
	c = (t8 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(25)) >> int32(26)
	t8 = t8 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(26))
	t9 = t9 + c
	c = (t9 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(24)) >> int32(25)
	t9 = t9 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(25))
	t0 = t0 + c*int64(19)
	c = (t0 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(25)) >> int32(26)
	t0 = t0 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(26))
	t1 = t1 + c
	**(**i32)(__ccgo_up(h)) = int32(t0)
	**(**i32)(__ccgo_up(h + 1*4)) = int32(t1)
	**(**i32)(__ccgo_up(h + 2*4)) = int32(t2)
	**(**i32)(__ccgo_up(h + 3*4)) = int32(t3)
	**(**i32)(__ccgo_up(h + 4*4)) = int32(t4)
	**(**i32)(__ccgo_up(h + 5*4)) = int32(t5)
	**(**i32)(__ccgo_up(h + 6*4)) = int32(t6)
	**(**i32)(__ccgo_up(h + 7*4)) = int32(t7)
	**(**i32)(__ccgo_up(h + 8*4)) = int32(t8)
	**(**i32)(__ccgo_up(h + 9*4)) = int32(t9) // Everything below 2^62, Carry precondition OK
}

// C documentation
//
//	// Precondition
//	// -------------
//	//   |f0|, |f2|, |f4|, |f6|, |f8|  <  1.65 * 2^26
//	//   |f1|, |f3|, |f5|, |f7|, |f9|  <  1.65 * 2^25
//	//
//	// Note: we could use fe_mul() for this, but this is significantly faster
func fe_sq(tls *libc.TLS, h uintptr, f uintptr) {
	var c, t0, t1, t2, t3, t4, t5, t6, t7, t8, t9 i64
	var f0, f0_2, f1, f1_2, f2, f2_2, f3, f3_2, f4, f4_2, f5, f5_2, f5_38, f6, f6_19, f6_2, f7, f7_2, f7_38, f8, f8_19, f9, f9_38 i32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, f0, f0_2, f1, f1_2, f2, f2_2, f3, f3_2, f4, f4_2, f5, f5_2, f5_38, f6, f6_19, f6_2, f7, f7_2, f7_38, f8, f8_19, f9, f9_38, t0, t1, t2, t3, t4, t5, t6, t7, t8, t9
	f0 = **(**i32)(__ccgo_up(f))
	f1 = **(**i32)(__ccgo_up(f + 1*4))
	f2 = **(**i32)(__ccgo_up(f + 2*4))
	f3 = **(**i32)(__ccgo_up(f + 3*4))
	f4 = **(**i32)(__ccgo_up(f + 4*4))
	f5 = **(**i32)(__ccgo_up(f + 5*4))
	f6 = **(**i32)(__ccgo_up(f + 6*4))
	f7 = **(**i32)(__ccgo_up(f + 7*4))
	f8 = **(**i32)(__ccgo_up(f + 8*4))
	f9 = **(**i32)(__ccgo_up(f + 9*4))
	f0_2 = f0 * int32(2)
	f1_2 = f1 * int32(2)
	f2_2 = f2 * int32(2)
	f3_2 = f3 * int32(2)
	f4_2 = f4 * int32(2)
	f5_2 = f5 * int32(2)
	f6_2 = f6 * int32(2)
	f7_2 = f7 * int32(2)
	f5_38 = f5 * int32(38)
	f6_19 = f6 * int32(19)
	f7_38 = f7 * int32(38)
	f8_19 = f8 * int32(19)
	f9_38 = f9 * int32(38)
	// |f0_2| , |f2_2| , |f4_2| , |f6_2| , |f8_2|  <  1.65 * 2^27
	// |f1_2| , |f3_2| , |f5_2| , |f7_2| , |f9_2|  <  1.65 * 2^26
	// |f5_38|, |f6_19|, |f7_38|, |f8_19|, |f9_38| <  2^31
	t0 = int64(f0)*int64(f0) + int64(f1_2)*int64(f9_38) + int64(f2_2)*int64(f8_19) + int64(f3_2)*int64(f7_38) + int64(f4_2)*int64(f6_19) + int64(f5)*int64(f5_38)
	t1 = int64(f0_2)*int64(f1) + int64(f2)*int64(f9_38) + int64(f3_2)*int64(f8_19) + int64(f4)*int64(f7_38) + int64(f5_2)*int64(f6_19)
	t2 = int64(f0_2)*int64(f2) + int64(f1_2)*int64(f1) + int64(f3_2)*int64(f9_38) + int64(f4_2)*int64(f8_19) + int64(f5_2)*int64(f7_38) + int64(f6)*int64(f6_19)
	t3 = int64(f0_2)*int64(f3) + int64(f1_2)*int64(f2) + int64(f4)*int64(f9_38) + int64(f5_2)*int64(f8_19) + int64(f6)*int64(f7_38)
	t4 = int64(f0_2)*int64(f4) + int64(f1_2)*int64(f3_2) + int64(f2)*int64(f2) + int64(f5_2)*int64(f9_38) + int64(f6_2)*int64(f8_19) + int64(f7)*int64(f7_38)
	t5 = int64(f0_2)*int64(f5) + int64(f1_2)*int64(f4) + int64(f2_2)*int64(f3) + int64(f6)*int64(f9_38) + int64(f7_2)*int64(f8_19)
	t6 = int64(f0_2)*int64(f6) + int64(f1_2)*int64(f5_2) + int64(f2_2)*int64(f4) + int64(f3_2)*int64(f3) + int64(f7_2)*int64(f9_38) + int64(f8)*int64(f8_19)
	t7 = int64(f0_2)*int64(f7) + int64(f1_2)*int64(f6) + int64(f2_2)*int64(f5) + int64(f3_2)*int64(f4) + int64(f8)*int64(f9_38)
	t8 = int64(f0_2)*int64(f8) + int64(f1_2)*int64(f7_2) + int64(f2_2)*int64(f6) + int64(f3_2)*int64(f5_2) + int64(f4)*int64(f4) + int64(f9)*int64(f9_38)
	t9 = int64(f0_2)*int64(f9) + int64(f1_2)*int64(f8) + int64(f2_2)*int64(f7) + int64(f3_2)*int64(f6) + int64(f4)*int64(f5_2)
	c = (t0 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(25)) >> int32(26)
	t0 = t0 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(26))
	t1 = t1 + c
	c = (t4 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(25)) >> int32(26)
	t4 = t4 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(26))
	t5 = t5 + c
	c = (t1 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(24)) >> int32(25)
	t1 = t1 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(25))
	t2 = t2 + c
	c = (t5 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(24)) >> int32(25)
	t5 = t5 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(25))
	t6 = t6 + c
	c = (t2 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(25)) >> int32(26)
	t2 = t2 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(26))
	t3 = t3 + c
	c = (t6 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(25)) >> int32(26)
	t6 = t6 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(26))
	t7 = t7 + c
	c = (t3 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(24)) >> int32(25)
	t3 = t3 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(25))
	t4 = t4 + c
	c = (t7 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(24)) >> int32(25)
	t7 = t7 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(25))
	t8 = t8 + c
	c = (t4 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(25)) >> int32(26)
	t4 = t4 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(26))
	t5 = t5 + c
	c = (t8 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(25)) >> int32(26)
	t8 = t8 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(26))
	t9 = t9 + c
	c = (t9 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(24)) >> int32(25)
	t9 = t9 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(25))
	t0 = t0 + c*int64(19)
	c = (t0 + libc.Int64FromInt32(1)<<libc.Int32FromInt32(25)) >> int32(26)
	t0 = t0 - c*(libc.Int64FromInt32(1)<<libc.Int32FromInt32(26))
	t1 = t1 + c
	**(**i32)(__ccgo_up(h)) = int32(t0)
	**(**i32)(__ccgo_up(h + 1*4)) = int32(t1)
	**(**i32)(__ccgo_up(h + 2*4)) = int32(t2)
	**(**i32)(__ccgo_up(h + 3*4)) = int32(t3)
	**(**i32)(__ccgo_up(h + 4*4)) = int32(t4)
	**(**i32)(__ccgo_up(h + 5*4)) = int32(t5)
	**(**i32)(__ccgo_up(h + 6*4)) = int32(t6)
	**(**i32)(__ccgo_up(h + 7*4)) = int32(t7)
	**(**i32)(__ccgo_up(h + 8*4)) = int32(t8)
	**(**i32)(__ccgo_up(h + 9*4)) = int32(t9)
}

// C documentation
//
//	//  Parity check.  Returns 0 if even, 1 if odd
func fe_isodd(tls *libc.TLS, f uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var isodd u8
	var _ /* s at bp+0 */ [32]u8
	_ = isodd
	fe_tobytes(tls, bp, f)
	isodd = libc.Uint8FromInt32(libc.Int32FromUint8((**(**[32]u8)(__ccgo_up(bp)))[0]) & int32(1))
	crypto_wipe(tls, bp, uint64(32))
	return libc.Int32FromUint8(isodd)
}

// C documentation
//
//	// Returns 1 if equal, 0 if not equal
func fe_isequal(tls *libc.TLS, f uintptr, g uintptr) (r int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var isdifferent int32
	var _ /* fs at bp+0 */ [32]u8
	var _ /* gs at bp+32 */ [32]u8
	_ = isdifferent
	fe_tobytes(tls, bp, f)
	fe_tobytes(tls, bp+32, g)
	isdifferent = crypto_verify32(tls, bp, bp+32)
	crypto_wipe(tls, bp, uint64(32))
	crypto_wipe(tls, bp+32, uint64(32))
	return int32(1) + isdifferent
}

// C documentation
//
//	// Inverse square root.
//	// Returns true if x is a square, false otherwise.
//	// After the call:
//	//   isr = sqrt(1/x)        if x is a non-zero square.
//	//   isr = sqrt(sqrt(-1)/x) if x is not a square.
//	//   isr = 0                if x is zero.
//	// We do not guarantee the sign of the square root.
//	//
//	// Notes:
//	// Let quartic = x^((p-1)/4)
//	//
//	// x^((p-1)/2) = chi(x)
//	// quartic^2   = chi(x)
//	// quartic     = sqrt(chi(x))
//	// quartic     = 1 or -1 or sqrt(-1) or -sqrt(-1)
//	//
//	// Note that x is a square if quartic is 1 or -1
//	// There are 4 cases to consider:
//	//
//	// if   quartic         = 1  (x is a square)
//	// then x^((p-1)/4)     = 1
//	//      x^((p-5)/4) * x = 1
//	//      x^((p-5)/4)     = 1/x
//	//      x^((p-5)/8)     = sqrt(1/x) or -sqrt(1/x)
//	//
//	// if   quartic                = -1  (x is a square)
//	// then x^((p-1)/4)            = -1
//	//      x^((p-5)/4) * x        = -1
//	//      x^((p-5)/4)            = -1/x
//	//      x^((p-5)/8)            = sqrt(-1)   / sqrt(x)
//	//      x^((p-5)/8) * sqrt(-1) = sqrt(-1)^2 / sqrt(x)
//	//      x^((p-5)/8) * sqrt(-1) = -1/sqrt(x)
//	//      x^((p-5)/8) * sqrt(-1) = -sqrt(1/x) or sqrt(1/x)
//	//
//	// if   quartic         = sqrt(-1)  (x is not a square)
//	// then x^((p-1)/4)     = sqrt(-1)
//	//      x^((p-5)/4) * x = sqrt(-1)
//	//      x^((p-5)/4)     = sqrt(-1)/x
//	//      x^((p-5)/8)     = sqrt(sqrt(-1)/x) or -sqrt(sqrt(-1)/x)
//	//
//	// Note that the product of two non-squares is always a square:
//	//   For any non-squares a and b, chi(a) = -1 and chi(b) = -1.
//	//   Since chi(x) = x^((p-1)/2), chi(a)*chi(b) = chi(a*b) = 1.
//	//   Therefore a*b is a square.
//	//
//	//   Since sqrt(-1) and x are both non-squares, their product is a
//	//   square, and we can compute their square root.
//	//
//	// if   quartic                = -sqrt(-1)  (x is not a square)
//	// then x^((p-1)/4)            = -sqrt(-1)
//	//      x^((p-5)/4) * x        = -sqrt(-1)
//	//      x^((p-5)/4)            = -sqrt(-1)/x
//	//      x^((p-5)/8)            = sqrt(-sqrt(-1)/x)
//	//      x^((p-5)/8)            = sqrt( sqrt(-1)/x) * sqrt(-1)
//	//      x^((p-5)/8) * sqrt(-1) = sqrt( sqrt(-1)/x) * sqrt(-1)^2
//	//      x^((p-5)/8) * sqrt(-1) = sqrt( sqrt(-1)/x) * -1
//	//      x^((p-5)/8) * sqrt(-1) = -sqrt(sqrt(-1)/x) or sqrt(sqrt(-1)/x)
func invsqrt(tls *libc.TLS, isr uintptr, x uintptr) (r int32) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var check, quartic uintptr
	var i, i1, i2, i3, i4, i5, i6, i7 size_t
	var m1, ms, p1, z0 int32
	var _ /* t0 at bp+0 */ fe
	var _ /* t1 at bp+40 */ fe
	var _ /* t2 at bp+80 */ fe
	_, _, _, _, _, _, _, _, _, _, _, _, _, _ = check, i, i1, i2, i3, i4, i5, i6, i7, m1, ms, p1, quartic, z0
	// t0 = x^((p-5)/8)
	// Can be achieved with a simple double & add ladder,
	// but it would be slower.
	fe_sq(tls, bp, x)
	fe_sq(tls, bp+40, bp)
	fe_sq(tls, bp+40, bp+40)
	fe_mul(tls, bp+40, x, bp+40)
	fe_mul(tls, bp, bp, bp+40)
	fe_sq(tls, bp, bp)
	fe_mul(tls, bp, bp+40, bp)
	fe_sq(tls, bp+40, bp)
	i = libc.Uint64FromInt32(libc.Int32FromInt32(1))
	for {
		if !(i < libc.Uint64FromInt32(libc.Int32FromInt32(5))) {
			break
		}
		fe_sq(tls, bp+40, bp+40)
		goto _1
	_1:
		;
		i = i + 1
	}
	fe_mul(tls, bp, bp+40, bp)
	fe_sq(tls, bp+40, bp)
	i1 = libc.Uint64FromInt32(libc.Int32FromInt32(1))
	for {
		if !(i1 < libc.Uint64FromInt32(libc.Int32FromInt32(10))) {
			break
		}
		fe_sq(tls, bp+40, bp+40)
		goto _2
	_2:
		;
		i1 = i1 + 1
	}
	fe_mul(tls, bp+40, bp+40, bp)
	fe_sq(tls, bp+80, bp+40)
	i2 = libc.Uint64FromInt32(libc.Int32FromInt32(1))
	for {
		if !(i2 < libc.Uint64FromInt32(libc.Int32FromInt32(20))) {
			break
		}
		fe_sq(tls, bp+80, bp+80)
		goto _3
	_3:
		;
		i2 = i2 + 1
	}
	fe_mul(tls, bp+40, bp+80, bp+40)
	fe_sq(tls, bp+40, bp+40)
	i3 = libc.Uint64FromInt32(libc.Int32FromInt32(1))
	for {
		if !(i3 < libc.Uint64FromInt32(libc.Int32FromInt32(10))) {
			break
		}
		fe_sq(tls, bp+40, bp+40)
		goto _4
	_4:
		;
		i3 = i3 + 1
	}
	fe_mul(tls, bp, bp+40, bp)
	fe_sq(tls, bp+40, bp)
	i4 = libc.Uint64FromInt32(libc.Int32FromInt32(1))
	for {
		if !(i4 < libc.Uint64FromInt32(libc.Int32FromInt32(50))) {
			break
		}
		fe_sq(tls, bp+40, bp+40)
		goto _5
	_5:
		;
		i4 = i4 + 1
	}
	fe_mul(tls, bp+40, bp+40, bp)
	fe_sq(tls, bp+80, bp+40)
	i5 = libc.Uint64FromInt32(libc.Int32FromInt32(1))
	for {
		if !(i5 < libc.Uint64FromInt32(libc.Int32FromInt32(100))) {
			break
		}
		fe_sq(tls, bp+80, bp+80)
		goto _6
	_6:
		;
		i5 = i5 + 1
	}
	fe_mul(tls, bp+40, bp+80, bp+40)
	fe_sq(tls, bp+40, bp+40)
	i6 = libc.Uint64FromInt32(libc.Int32FromInt32(1))
	for {
		if !(i6 < libc.Uint64FromInt32(libc.Int32FromInt32(50))) {
			break
		}
		fe_sq(tls, bp+40, bp+40)
		goto _7
	_7:
		;
		i6 = i6 + 1
	}
	fe_mul(tls, bp, bp+40, bp)
	fe_sq(tls, bp, bp)
	i7 = libc.Uint64FromInt32(libc.Int32FromInt32(1))
	for {
		if !(i7 < libc.Uint64FromInt32(libc.Int32FromInt32(2))) {
			break
		}
		fe_sq(tls, bp, bp)
		goto _8
	_8:
		;
		i7 = i7 + 1
	}
	fe_mul(tls, bp, bp, x)
	// quartic = x^((p-1)/4)
	quartic = bp + 40
	fe_sq(tls, quartic, bp)
	fe_mul(tls, quartic, quartic, x)
	check = bp + 80
	fe_0(tls, check)
	z0 = fe_isequal(tls, x, check)
	fe_1(tls, check)
	p1 = fe_isequal(tls, quartic, check)
	fe_neg(tls, check, check)
	m1 = fe_isequal(tls, quartic, check)
	fe_neg(tls, check, uintptr(unsafe.Pointer(&sqrtm1)))
	ms = fe_isequal(tls, quartic, check)
	// if quartic == -1 or sqrt(-1)
	// then  isr = x^((p-1)/4) * sqrt(-1)
	// else  isr = x^((p-1)/4)
	fe_mul(tls, isr, bp, uintptr(unsafe.Pointer(&sqrtm1)))
	fe_ccopy(tls, isr, bp, int32(1)-(m1|ms))
	crypto_wipe(tls, bp, uint64(40))
	crypto_wipe(tls, bp+40, uint64(40))
	crypto_wipe(tls, bp+80, uint64(40))
	return p1 | m1 | z0
}

// C documentation
//
//	// Inverse in terms of inverse square root.
//	// Requires two additional squarings to get rid of the sign.
//	//
//	//   1/x = x * (+invsqrt(x^2))^2
//	//       = x * (-invsqrt(x^2))^2
//	//
//	// A fully optimised exponentiation by p-1 would save 6 field
//	// multiplications, but it would require more code.
func fe_invert(tls *libc.TLS, out uintptr, x uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var _ /* tmp at bp+0 */ fe
	fe_sq(tls, bp, x)
	invsqrt(tls, bp, bp)
	fe_sq(tls, bp, bp)
	fe_mul(tls, out, bp, x)
	crypto_wipe(tls, bp, uint64(40))
}

// C documentation
//
//	// trim a scalar for scalar multiplication
func crypto_eddsa_trim_scalar(tls *libc.TLS, out uintptr, in uintptr) {
	var _i_ size_t
	var v2 uintptr
	_, _ = _i_, v2
	_i_ = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(_i_ < libc.Uint64FromInt32(libc.Int32FromInt32(32))) {
			break
		}
		**(**u8)(__ccgo_up(out + uintptr(_i_))) = **(**u8)(__ccgo_up(in + uintptr(_i_)))
		goto _1
	_1:
		;
		_i_ = _i_ + 1
	}
	v2 = out
	*(*u8)(unsafe.Pointer(v2)) = u8(int32(*(*u8)(unsafe.Pointer(v2))) & libc.Int32FromInt32(248))
	v2 = out + 31
	*(*u8)(unsafe.Pointer(v2)) = u8(int32(*(*u8)(unsafe.Pointer(v2))) & libc.Int32FromInt32(127))
	v2 = out + 31
	*(*u8)(unsafe.Pointer(v2)) = u8(int32(*(*u8)(unsafe.Pointer(v2))) | libc.Int32FromInt32(64))
}

// C documentation
//
//	// get bit from scalar at position i
func scalar_bit(tls *libc.TLS, s uintptr, i int32) (r int32) {
	if i < 0 {
		return 0
	} // handle -1 for sliding windows
	return libc.Int32FromUint8(**(**u8)(__ccgo_up(s + uintptr(i>>int32(3))))) >> (i & int32(7)) & int32(1)
}

// C documentation
//
//	///////////////
//	/// X-25519 /// Taken from SUPERCOP's ref10 implementation.
//	///////////////
func scalarmult(tls *libc.TLS, q uintptr, scalar uintptr, p uintptr, nb_bits int32) {
	bp := tls.Alloc(288)
	defer tls.Free(288)
	var b, pos, swap int32
	var _ /* t0 at bp+200 */ fe
	var _ /* t1 at bp+240 */ fe
	var _ /* x1 at bp+0 */ fe
	var _ /* x2 at bp+40 */ fe
	var _ /* x3 at bp+120 */ fe
	var _ /* z2 at bp+80 */ fe
	var _ /* z3 at bp+160 */ fe
	_, _, _ = b, pos, swap
	fe_frombytes(tls, bp, p)
	// Montgomery ladder
	// In projective coordinates, to avoid divisions: x = X / Z
	// We don't care about the y coordinate, it's only 1 bit of information
	fe_1(tls, bp+40)
	fe_0(tls, bp+80) // "zero" point
	fe_copy(tls, bp+120, bp)
	fe_1(tls, bp+160) // "one"  point
	swap = 0
	pos = nb_bits - int32(1)
	for {
		if !(pos >= 0) {
			break
		}
		// constant time conditional swap before ladder step
		b = scalar_bit(tls, scalar, pos)
		swap = swap ^ b // xor trick avoids swapping at the end of the loop
		fe_cswap(tls, bp+40, bp+120, swap)
		fe_cswap(tls, bp+80, bp+160, swap)
		swap = b // anticipates one last swap after the loop
		// Montgomery ladder step: replaces (P2, P3) by (P2*2, P2+P3)
		// with differential addition
		fe_sub(tls, bp+200, bp+120, bp+160)
		fe_sub(tls, bp+240, bp+40, bp+80)
		fe_add(tls, bp+40, bp+40, bp+80)
		fe_add(tls, bp+80, bp+120, bp+160)
		fe_mul(tls, bp+160, bp+200, bp+40)
		fe_mul(tls, bp+80, bp+80, bp+240)
		fe_sq(tls, bp+200, bp+240)
		fe_sq(tls, bp+240, bp+40)
		fe_add(tls, bp+120, bp+160, bp+80)
		fe_sub(tls, bp+80, bp+160, bp+80)
		fe_mul(tls, bp+40, bp+240, bp+200)
		fe_sub(tls, bp+240, bp+240, bp+200)
		fe_sq(tls, bp+80, bp+80)
		fe_mul_small(tls, bp+160, bp+240, int32(121666))
		fe_sq(tls, bp+120, bp+120)
		fe_add(tls, bp+200, bp+200, bp+160)
		fe_mul(tls, bp+160, bp, bp+80)
		fe_mul(tls, bp+80, bp+240, bp+200)
		goto _1
	_1:
		;
		pos = pos - 1
	}
	// last swap is necessary to compensate for the xor trick
	// Note: after this swap, P3 == P2 + P1.
	fe_cswap(tls, bp+40, bp+120, swap)
	fe_cswap(tls, bp+80, bp+160, swap)
	// normalises the coordinates: x == X / Z
	fe_invert(tls, bp+80, bp+80)
	fe_mul(tls, bp+40, bp+40, bp+80)
	fe_tobytes(tls, q, bp+40)
	crypto_wipe(tls, bp, uint64(40))
	crypto_wipe(tls, bp+40, uint64(40))
	crypto_wipe(tls, bp+80, uint64(40))
	crypto_wipe(tls, bp+200, uint64(40))
	crypto_wipe(tls, bp+120, uint64(40))
	crypto_wipe(tls, bp+160, uint64(40))
	crypto_wipe(tls, bp+240, uint64(40))
}

func crypto_x25519(tls *libc.TLS, raw_shared_secret uintptr, your_secret_key uintptr, their_public_key uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var _ /* e at bp+0 */ [32]u8
	crypto_eddsa_trim_scalar(tls, bp, your_secret_key)
	scalarmult(tls, raw_shared_secret, bp, their_public_key, int32(255))
	crypto_wipe(tls, bp, uint64(32))
}

func crypto_x25519_public_key(tls *libc.TLS, public_key uintptr, secret_key uintptr) {
	crypto_x25519(tls, public_key, secret_key, uintptr(unsafe.Pointer(&base_point)))
}

var base_point = [32]u8{
	0: uint8(9),
}

// C documentation
//
//	///////////////////////////
//	/// Arithmetic modulo L ///
//	///////////////////////////
var L = [8]u32{
	0: uint32(0x5cf5d3ed),
	1: uint32(0x5812631a),
	2: uint32(0xa2f79cd6),
	3: uint32(0x14def9de),
	7: uint32(0x10000000),
}

// C documentation
//
//	//  p = a*b + p
func multiply(tls *libc.TLS, p uintptr, a uintptr, b uintptr) {
	var carry u64
	var i, j size_t
	_, _, _ = carry, i, j
	i = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(i < libc.Uint64FromInt32(libc.Int32FromInt32(8))) {
			break
		}
		carry = uint64(0)
		j = libc.Uint64FromInt32(libc.Int32FromInt32(0))
		for {
			if !(j < libc.Uint64FromInt32(libc.Int32FromInt32(8))) {
				break
			}
			carry = carry + (uint64(**(**u32)(__ccgo_up(p + uintptr(i+j)*4))) + uint64(**(**u32)(__ccgo_up(a + uintptr(i)*4)))*uint64(**(**u32)(__ccgo_up(b + uintptr(j)*4))))
			**(**u32)(__ccgo_up(p + uintptr(i+j)*4)) = uint32(carry)
			carry = carry >> uint64(32)
			goto _2
		_2:
			;
			j = j + 1
		}
		**(**u32)(__ccgo_up(p + uintptr(i+uint64(8))*4)) = uint32(carry)
		goto _1
	_1:
		;
		i = i + 1
	}
}

func is_above_l(tls *libc.TLS, x uintptr) (r int32) {
	var carry u64
	var i size_t
	_, _ = carry, i
	// We work with L directly, in a 2's complement encoding
	// (-L == ~L + 1)
	carry = uint64(1)
	i = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(i < libc.Uint64FromInt32(libc.Int32FromInt32(8))) {
			break
		}
		carry = carry + (uint64(**(**u32)(__ccgo_up(x + uintptr(i)*4))) + uint64(^L[i]&libc.Uint32FromUint32(0xffffffff)))
		carry = carry >> uint64(32)
		goto _1
	_1:
		;
		i = i + 1
	}
	return libc.Int32FromUint64(carry) // carry is either 0 or 1
}

// C documentation
//
//	// Final reduction modulo L, by conditionally removing L.
//	// if x < l     , then r = x
//	// if l <= x 2*l, then r = x-l
//	// otherwise the result will be wrong
func remove_l(tls *libc.TLS, r uintptr, x uintptr) {
	var carry u64
	var i size_t
	var mask u32
	_, _, _ = carry, i, mask
	carry = libc.Uint64FromInt32(is_above_l(tls, x))
	mask = ^uint32(carry) + uint32(1) // carry == 0 or 1
	i = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(i < libc.Uint64FromInt32(libc.Int32FromInt32(8))) {
			break
		}
		carry = carry + (uint64(**(**u32)(__ccgo_up(x + uintptr(i)*4))) + uint64(^L[i]&mask))
		**(**u32)(__ccgo_up(r + uintptr(i)*4)) = uint32(carry)
		carry = carry >> uint64(32)
		goto _1
	_1:
		;
		i = i + 1
	}
}

// C documentation
//
//	// Full reduction modulo L (Barrett reduction)
func mod_l(tls *libc.TLS, reduced uintptr, x uintptr) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var _i_, i, i1, i2, j, j1 size_t
	var carry, carry1, carry2 u64
	var _ /* xr at bp+0 */ [25]u32
	_, _, _, _, _, _, _, _, _ = _i_, carry, carry1, carry2, i, i1, i2, j, j1
	// xr = x * r
	**(**[25]u32)(__ccgo_up(bp)) = [25]u32{}
	i = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(i < libc.Uint64FromInt32(libc.Int32FromInt32(9))) {
			break
		}
		carry = uint64(0)
		j = libc.Uint64FromInt32(libc.Int32FromInt32(0))
		for {
			if !(j < libc.Uint64FromInt32(libc.Int32FromInt32(16))) {
				break
			}
			carry = carry + (uint64((**(**[25]u32)(__ccgo_up(bp)))[i+j]) + uint64(r[i])*uint64(**(**u32)(__ccgo_up(x + uintptr(j)*4))))
			(**(**[25]u32)(__ccgo_up(bp)))[i+j] = uint32(carry)
			carry = carry >> uint64(32)
			goto _2
		_2:
			;
			j = j + 1
		}
		(**(**[25]u32)(__ccgo_up(bp)))[i+uint64(16)] = uint32(carry)
		goto _1
	_1:
		;
		i = i + 1
	}
	// xr = floor(xr / 2^512) * L
	// Since the result is guaranteed to be below 2*L,
	// it is enough to only compute the first 256 bits.
	// The division is performed by saying xr[i+16]. (16 * 32 = 512)
	_i_ = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(_i_ < libc.Uint64FromInt32(libc.Int32FromInt32(8))) {
			break
		}
		(**(**[25]u32)(__ccgo_up(bp)))[_i_] = uint32(0)
		goto _3
	_3:
		;
		_i_ = _i_ + 1
	}
	i1 = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(i1 < libc.Uint64FromInt32(libc.Int32FromInt32(8))) {
			break
		}
		carry1 = uint64(0)
		j1 = libc.Uint64FromInt32(libc.Int32FromInt32(0))
		for {
			if !(j1 < uint64(8)-i1) {
				break
			}
			carry1 = carry1 + (uint64((**(**[25]u32)(__ccgo_up(bp)))[i1+j1]) + uint64((**(**[25]u32)(__ccgo_up(bp)))[i1+uint64(16)])*uint64(L[j1]))
			(**(**[25]u32)(__ccgo_up(bp)))[i1+j1] = uint32(carry1)
			carry1 = carry1 >> uint64(32)
			goto _5
		_5:
			;
			j1 = j1 + 1
		}
		goto _4
	_4:
		;
		i1 = i1 + 1
	}
	// xr = x - xr
	carry2 = uint64(1)
	i2 = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(i2 < libc.Uint64FromInt32(libc.Int32FromInt32(8))) {
			break
		}
		carry2 = carry2 + (uint64(**(**u32)(__ccgo_up(x + uintptr(i2)*4))) + uint64(^(**(**[25]u32)(__ccgo_up(bp)))[i2]&libc.Uint32FromUint32(0xffffffff)))
		(**(**[25]u32)(__ccgo_up(bp)))[i2] = uint32(carry2)
		carry2 = carry2 >> uint64(32)
		goto _6
	_6:
		;
		i2 = i2 + 1
	}
	// Final reduction modulo L (conditional subtraction)
	remove_l(tls, bp, bp)
	store32_le_buf(tls, reduced, bp, uint64(8))
	crypto_wipe(tls, bp, uint64(100))
}

var r = [9]u32{
	0: uint32(0x0a2c131b),
	1: uint32(0xed9ce5a3),
	2: uint32(0x086329a7),
	3: uint32(0x2106215d),
	4: uint32(0xffffffeb),
	5: uint32(0xffffffff),
	6: uint32(0xffffffff),
	7: uint32(0xffffffff),
	8: uint32(0xf),
}

func crypto_eddsa_reduce(tls *libc.TLS, reduced uintptr, expanded uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var _ /* x at bp+0 */ [16]u32
	load32_le_buf(tls, bp, expanded, uint64(16))
	mod_l(tls, reduced, bp)
	crypto_wipe(tls, bp, uint64(64))
}

// C documentation
//
//	// r = (a * b) + c
func crypto_eddsa_mul_add(tls *libc.TLS, r uintptr, a uintptr, b uintptr, c uintptr) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var _i_ size_t
	var _ /* A at bp+0 */ [8]u32
	var _ /* B at bp+32 */ [8]u32
	var _ /* p at bp+64 */ [16]u32
	_ = _i_
	load32_le_buf(tls, bp, a, uint64(8))
	load32_le_buf(tls, bp+32, b, uint64(8))
	load32_le_buf(tls, bp+64, c, uint64(8))
	_i_ = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(_i_ < libc.Uint64FromInt32(libc.Int32FromInt32(8))) {
			break
		}
		**(**u32)(__ccgo_up(bp + 64 + libc.UintptrFromInt32(8)*4 + uintptr(_i_)*4)) = uint32(0)
		goto _1
	_1:
		;
		_i_ = _i_ + 1
	}
	multiply(tls, bp+64, bp, bp+32)
	mod_l(tls, r, bp+64)
	crypto_wipe(tls, bp+64, uint64(64))
	crypto_wipe(tls, bp, uint64(32))
	crypto_wipe(tls, bp+32, uint64(32))
}

///////////////
/// Ed25519 ///
///////////////

// C documentation
//
//	// Point (group element, ge) in a twisted Edwards curve,
//	// in extended projective coordinates.
//	// ge        : x  = X/Z, y  = Y/Z, T  = XY/Z
//	// ge_cached : Yp = X+Y, Ym = X-Y, T2 = T*D2
//	// ge_precomp: Z  = 1
type ge = struct {
	FX fe
	FY fe
	FZ fe
	FT fe
}

type ge_cached = struct {
	FYp fe
	FYm fe
	FZ  fe
	FT2 fe
}

type ge_precomp = struct {
	FYp fe
	FYm fe
	FT2 fe
}

func ge_zero(tls *libc.TLS, p uintptr) {
	fe_0(tls, p)
	fe_1(tls, p+40)
	fe_1(tls, p+80)
	fe_0(tls, p+120)
}

func ge_tobytes(tls *libc.TLS, s uintptr, h uintptr) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var v1 uintptr
	var _ /* recip at bp+0 */ fe
	var _ /* x at bp+40 */ fe
	var _ /* y at bp+80 */ fe
	_ = v1
	fe_invert(tls, bp, h+80)
	fe_mul(tls, bp+40, h, bp)
	fe_mul(tls, bp+80, h+40, bp)
	fe_tobytes(tls, s, bp+80)
	v1 = s + 31
	*(*u8)(unsafe.Pointer(v1)) = u8(int32(*(*u8)(unsafe.Pointer(v1))) ^ libc.Int32FromUint8(libc.Uint8FromInt32(fe_isodd(tls, bp+40)))<<libc.Int32FromInt32(7))
	crypto_wipe(tls, bp, uint64(40))
	crypto_wipe(tls, bp+40, uint64(40))
	crypto_wipe(tls, bp+80, uint64(40))
}

// C documentation
//
//	// h = -s, where s is a point encoded in 32 bytes
//	//
//	// Variable time!  Inputs must not be secret!
//	// => Use only to *check* signatures.
//	//
//	// From the specifications:
//	//   The encoding of s contains y and the sign of x
//	//   x = sqrt((y^2 - 1) / (d*y^2 + 1))
//	// In extended coordinates:
//	//   X = x, Y = y, Z = 1, T = x*y
//	//
//	//    Note that num * den is a square iff num / den is a square
//	//    If num * den is not a square, the point was not on the curve.
//	// From the above:
//	//   Let num =   y^2 - 1
//	//   Let den = d*y^2 + 1
//	//   x = sqrt((y^2 - 1) / (d*y^2 + 1))
//	//   x = sqrt(num / den)
//	//   x = sqrt(num^2 / (num * den))
//	//   x = num * sqrt(1 / (num * den))
//	//
//	// Therefore, we can just compute:
//	//   num =   y^2 - 1
//	//   den = d*y^2 + 1
//	//   isr = invsqrt(num * den)  // abort if not square
//	//   x   = num * isr
//	// Finally, negate x if its sign is not as specified.
func ge_frombytes_neg_vartime(tls *libc.TLS, h uintptr, s uintptr) (r int32) {
	var is_square int32
	_ = is_square
	fe_frombytes(tls, h+40, s)
	fe_1(tls, h+80)
	fe_sq(tls, h+120, h+40)                            // t =   y^2
	fe_mul(tls, h, h+120, uintptr(unsafe.Pointer(&d))) // x = d*y^2
	fe_sub(tls, h+120, h+120, h+80)                    // t =   y^2 - 1
	fe_add(tls, h, h, h+80)                            // x = d*y^2 + 1
	fe_mul(tls, h, h+120, h)                           // x = (y^2 - 1) * (d*y^2 + 1)
	is_square = invsqrt(tls, h, h)
	if !(is_square != 0) {
		return -int32(1) // Not on the curve, abort
	}
	fe_mul(tls, h, h+120, h) // x = sqrt((y^2 - 1) / (d*y^2 + 1))
	if fe_isodd(tls, h) == libc.Int32FromUint8(**(**u8)(__ccgo_up(s + 31)))>>int32(7) {
		fe_neg(tls, h, h)
	}
	fe_mul(tls, h+120, h, h+40)
	return 0
}

func ge_cache(tls *libc.TLS, c uintptr, p uintptr) {
	fe_add(tls, c, p+40, p)
	fe_sub(tls, c+40, p+40, p)
	fe_copy(tls, c+80, p+80)
	fe_mul(tls, c+120, p+120, uintptr(unsafe.Pointer(&D2)))
}

// C documentation
//
//	// Internal buffers are not wiped! Inputs must not be secret!
//	// => Use only to *check* signatures.
func ge_add(tls *libc.TLS, s uintptr, p uintptr, q uintptr) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var _ /* a at bp+0 */ fe
	var _ /* b at bp+40 */ fe
	fe_add(tls, bp, p+40, p)
	fe_sub(tls, bp+40, p+40, p)
	fe_mul(tls, bp, bp, q)
	fe_mul(tls, bp+40, bp+40, q+40)
	fe_add(tls, s+40, bp, bp+40)
	fe_sub(tls, s, bp, bp+40)
	fe_add(tls, s+80, p+80, p+80)
	fe_mul(tls, s+80, s+80, q+80)
	fe_mul(tls, s+120, p+120, q+120)
	fe_add(tls, bp, s+80, s+120)
	fe_sub(tls, bp+40, s+80, s+120)
	fe_mul(tls, s+120, s, s+40)
	fe_mul(tls, s, s, bp+40)
	fe_mul(tls, s+40, s+40, bp)
	fe_mul(tls, s+80, bp, bp+40)
}

// C documentation
//
//	// Internal buffers are not wiped! Inputs must not be secret!
//	// => Use only to *check* signatures.
func ge_sub(tls *libc.TLS, s uintptr, p uintptr, q uintptr) {
	bp := tls.Alloc(160)
	defer tls.Free(160)
	var _ /* neg at bp+0 */ ge_cached
	fe_copy(tls, bp+40, q)
	fe_copy(tls, bp, q+40)
	fe_copy(tls, bp+80, q+80)
	fe_neg(tls, bp+120, q+120)
	ge_add(tls, s, p, bp)
}

func ge_madd(tls *libc.TLS, s uintptr, p uintptr, q uintptr, a uintptr, b uintptr) {
	fe_add(tls, a, p+40, p)
	fe_sub(tls, b, p+40, p)
	fe_mul(tls, a, a, q)
	fe_mul(tls, b, b, q+40)
	fe_add(tls, s+40, a, b)
	fe_sub(tls, s, a, b)
	fe_add(tls, s+80, p+80, p+80)
	fe_mul(tls, s+120, p+120, q+80)
	fe_add(tls, a, s+80, s+120)
	fe_sub(tls, b, s+80, s+120)
	fe_mul(tls, s+120, s, s+40)
	fe_mul(tls, s, s, b)
	fe_mul(tls, s+40, s+40, a)
	fe_mul(tls, s+80, a, b)
}

// C documentation
//
//	// Internal buffers are not wiped! Inputs must not be secret!
//	// => Use only to *check* signatures.
func ge_msub(tls *libc.TLS, s uintptr, p uintptr, q uintptr, a uintptr, b uintptr) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var _ /* neg at bp+0 */ ge_precomp
	fe_copy(tls, bp+40, q)
	fe_copy(tls, bp, q+40)
	fe_neg(tls, bp+80, q+80)
	ge_madd(tls, s, p, bp, a, b)
}

func ge_double(tls *libc.TLS, s uintptr, p uintptr, q uintptr) {
	fe_sq(tls, q, p)
	fe_sq(tls, q+40, p+40)
	fe_sq(tls, q+80, p+80)                  // qZ = pZ^2
	fe_mul_small(tls, q+80, q+80, int32(2)) // qZ = pZ^2 * 2
	fe_add(tls, q+120, p, p+40)
	fe_sq(tls, s+120, q+120)
	fe_add(tls, q+120, q+40, q)
	fe_sub(tls, q+40, q+40, q)
	fe_sub(tls, q, s+120, q+120)
	fe_sub(tls, q+80, q+80, q+40)
	fe_mul(tls, s, q, q+80)
	fe_mul(tls, s+40, q+120, q+40)
	fe_mul(tls, s+80, q+40, q+80)
	fe_mul(tls, s+120, q, q+120)
}

// C documentation
//
//	// 5-bit signed window in cached format (Niels coordinates, Z=1)
var b_window = [8]ge_precomp{
	0: {
		FYp: fe{
			0: int32(25967493),
			1: -int32(14356035),
			2: int32(29566456),
			3: int32(3660896),
			4: -int32(12694345),
			5: int32(4014787),
			6: int32(27544626),
			7: -int32(11754271),
			8: -int32(6079156),
			9: int32(2047605),
		},
		FYm: fe{
			0: -int32(12545711),
			1: int32(934262),
			2: -int32(2722910),
			3: int32(3049990),
			4: -int32(727428),
			5: int32(9406986),
			6: int32(12720692),
			7: int32(5043384),
			8: int32(19500929),
			9: -int32(15469378),
		},
		FT2: fe{
			0: -int32(8738181),
			1: int32(4489570),
			2: int32(9688441),
			3: -int32(14785194),
			4: int32(10184609),
			5: -int32(12363380),
			6: int32(29287919),
			7: int32(11864899),
			8: -int32(24514362),
			9: -int32(4438546),
		},
	},
	1: {
		FYp: fe{
			0: int32(15636291),
			1: -int32(9688557),
			2: int32(24204773),
			3: -int32(7912398),
			4: int32(616977),
			5: -int32(16685262),
			6: int32(27787600),
			7: -int32(14772189),
			8: int32(28944400),
			9: -int32(1550024),
		},
		FYm: fe{
			0: int32(16568933),
			1: int32(4717097),
			2: -int32(11556148),
			3: -int32(1102322),
			4: int32(15682896),
			5: -int32(11807043),
			6: int32(16354577),
			7: -int32(11775962),
			8: int32(7689662),
			9: int32(11199574),
		},
		FT2: fe{
			0: int32(30464156),
			1: -int32(5976125),
			2: -int32(11779434),
			3: -int32(15670865),
			4: int32(23220365),
			5: int32(15915852),
			6: int32(7512774),
			7: int32(10017326),
			8: -int32(17749093),
			9: -int32(9920357),
		},
	},
	2: {
		FYp: fe{
			0: int32(10861363),
			1: int32(11473154),
			2: int32(27284546),
			3: int32(1981175),
			4: -int32(30064349),
			5: int32(12577861),
			6: int32(32867885),
			7: int32(14515107),
			8: -int32(15438304),
			9: int32(10819380),
		},
		FYm: fe{
			0: int32(4708026),
			1: int32(6336745),
			2: int32(20377586),
			3: int32(9066809),
			4: -int32(11272109),
			5: int32(6594696),
			6: -int32(25653668),
			7: int32(12483688),
			8: -int32(12668491),
			9: int32(5581306),
		},
		FT2: fe{
			0: int32(19563160),
			1: int32(16186464),
			2: -int32(29386857),
			3: int32(4097519),
			4: int32(10237984),
			5: -int32(4348115),
			6: int32(28542350),
			7: int32(13850243),
			8: -int32(23678021),
			9: -int32(15815942),
		},
	},
	3: {
		FYp: fe{
			0: int32(5153746),
			1: int32(9909285),
			2: int32(1723747),
			3: -int32(2777874),
			4: int32(30523605),
			5: int32(5516873),
			6: int32(19480852),
			7: int32(5230134),
			8: -int32(23952439),
			9: -int32(15175766),
		},
		FYm: fe{
			0: -int32(30269007),
			1: -int32(3463509),
			2: int32(7665486),
			3: int32(10083793),
			4: int32(28475525),
			5: int32(1649722),
			6: int32(20654025),
			7: int32(16520125),
			8: int32(30598449),
			9: int32(7715701),
		},
		FT2: fe{
			0: int32(28881845),
			1: int32(14381568),
			2: int32(9657904),
			3: int32(3680757),
			4: -int32(20181635),
			5: int32(7843316),
			6: -int32(31400660),
			7: int32(1370708),
			8: int32(29794553),
			9: -int32(1409300),
		},
	},
	4: {
		FYp: fe{
			0: -int32(22518993),
			1: -int32(6692182),
			2: int32(14201702),
			3: -int32(8745502),
			4: -int32(23510406),
			5: int32(8844726),
			6: int32(18474211),
			7: -int32(1361450),
			8: -int32(13062696),
			9: int32(13821877),
		},
		FYm: fe{
			0: -int32(6455177),
			1: -int32(7839871),
			2: int32(3374702),
			3: -int32(4740862),
			4: -int32(27098617),
			5: -int32(10571707),
			6: int32(31655028),
			7: -int32(7212327),
			8: int32(18853322),
			9: -int32(14220951),
		},
		FT2: fe{
			0: int32(4566830),
			1: -int32(12963868),
			2: -int32(28974889),
			3: -int32(12240689),
			4: -int32(7602672),
			5: -int32(2830569),
			6: -int32(8514358),
			7: -int32(10431137),
			8: int32(2207753),
			9: -int32(3209784),
		},
	},
	5: {
		FYp: fe{
			0: -int32(25154831),
			1: -int32(4185821),
			2: int32(29681144),
			3: int32(7868801),
			4: -int32(6854661),
			5: -int32(9423865),
			6: -int32(12437364),
			7: -int32(663000),
			8: -int32(31111463),
			9: -int32(16132436),
		},
		FYm: fe{
			0: int32(25576264),
			1: -int32(2703214),
			2: int32(7349804),
			3: -int32(11814844),
			4: int32(16472782),
			5: int32(9300885),
			6: int32(3844789),
			7: int32(15725684),
			8: int32(171356),
			9: int32(6466918),
		},
		FT2: fe{
			0: int32(23103977),
			1: int32(13316479),
			2: int32(9739013),
			3: -int32(16149481),
			4: int32(817875),
			5: -int32(15038942),
			6: int32(8965339),
			7: -int32(14088058),
			8: -int32(30714912),
			9: int32(16193877),
		},
	},
	6: {
		FYp: fe{
			0: -int32(33521811),
			1: int32(3180713),
			2: -int32(2394130),
			3: int32(14003687),
			4: -int32(16903474),
			5: -int32(16270840),
			6: int32(17238398),
			7: int32(4729455),
			8: -int32(18074513),
			9: int32(9256800),
		},
		FYm: fe{
			0: -int32(25182317),
			1: -int32(4174131),
			2: int32(32336398),
			3: int32(5036987),
			4: -int32(21236817),
			5: int32(11360617),
			6: int32(22616405),
			7: int32(9761698),
			8: -int32(19827198),
			9: int32(630305),
		},
		FT2: fe{
			0: -int32(13720693),
			1: int32(2639453),
			2: -int32(24237460),
			3: -int32(7406481),
			4: int32(9494427),
			5: -int32(5774029),
			6: -int32(6554551),
			7: -int32(15960994),
			8: -int32(2449256),
			9: -int32(14291300),
		},
	},
	7: {
		FYp: fe{
			0: -int32(3151181),
			1: -int32(5046075),
			2: int32(9282714),
			3: int32(6866145),
			4: -int32(31907062),
			5: -int32(863023),
			6: -int32(18940575),
			7: int32(15033784),
			8: int32(25105118),
			9: -int32(7894876),
		},
		FYm: fe{
			0: -int32(24326370),
			1: int32(15950226),
			2: -int32(31801215),
			3: -int32(14592823),
			4: -int32(11662737),
			5: -int32(5090925),
			6: int32(1573892),
			7: -int32(2625887),
			8: int32(2198790),
			9: -int32(15804619),
		},
		FT2: fe{
			0: -int32(3099351),
			1: int32(10324967),
			2: -int32(2241613),
			3: int32(7453183),
			4: -int32(5446979),
			5: -int32(2735503),
			6: -int32(13812022),
			7: -int32(16236442),
			8: -int32(32461234),
			9: -int32(12290683),
		},
	},
}

// C documentation
//
//	// Incremental sliding windows (left to right)
//	// Based on Roberto Maria Avanzi[2005]
type slide_ctx = struct {
	Fnext_index i16
	Fnext_digit i8
	Fnext_check u8
}

func slide_init(tls *libc.TLS, ctx uintptr, scalar uintptr) {
	var i int32
	_ = i
	// scalar is guaranteed to be below L, either because we checked (s),
	// or because we reduced it modulo L (h_ram). L is under 2^253, so
	// so bits 253 to 255 are guaranteed to be zero. No need to test them.
	//
	// Note however that L is very close to 2^252, so bit 252 is almost
	// always zero.  If we were to start at bit 251, the tests wouldn't
	// catch the off-by-one error (constructing one that does would be
	// prohibitively expensive).
	//
	// We should still check bit 252, though.
	i = int32(252)
	for i > 0 && scalar_bit(tls, scalar, i) == 0 {
		i = i - 1
	}
	(*slide_ctx)(unsafe.Pointer(ctx)).Fnext_check = libc.Uint8FromInt32(i + libc.Int32FromInt32(1))
	(*slide_ctx)(unsafe.Pointer(ctx)).Fnext_index = int16(-int32(1))
	(*slide_ctx)(unsafe.Pointer(ctx)).Fnext_digit = int8(-int32(1))
}

func slide_step(tls *libc.TLS, ctx uintptr, width int32, i int32, scalar uintptr) (r int32) {
	var j, lsb, s, v, w, v1 int32
	var v3 uintptr
	_, _, _, _, _, _, _ = j, lsb, s, v, w, v1, v3
	if i == libc.Int32FromUint8((*slide_ctx)(unsafe.Pointer(ctx)).Fnext_check) {
		if scalar_bit(tls, scalar, i) == scalar_bit(tls, scalar, i-int32(1)) {
			(*slide_ctx)(unsafe.Pointer(ctx)).Fnext_check = (*slide_ctx)(unsafe.Pointer(ctx)).Fnext_check - 1
		} else {
			if width <= i+int32(1) {
				v1 = width
			} else {
				v1 = i + int32(1)
			}
			// compute digit of next window
			w = v1
			v = -(scalar_bit(tls, scalar, i) << (w - int32(1)))
			j = 0
			for {
				if !(j < w-int32(1)) {
					break
				}
				v = v + scalar_bit(tls, scalar, i-(w-int32(1))+j)<<j
				goto _2
			_2:
				;
				j = j + 1
			}
			v = v + scalar_bit(tls, scalar, i-w)
			lsb = v & (^v + int32(1)) // smallest bit of v
			s = libc.BoolInt32(lsb&int32(0xAA) != 0)<<0 | libc.BoolInt32(lsb&int32(0xCC) != 0)<<int32(1) | libc.BoolInt32(lsb&int32(0xF0) != 0)<<int32(2)
			(*slide_ctx)(unsafe.Pointer(ctx)).Fnext_index = int16(i - (w - libc.Int32FromInt32(1)) + s)
			(*slide_ctx)(unsafe.Pointer(ctx)).Fnext_digit = int8(v >> s)
			v3 = ctx + 3
			*(*u8)(unsafe.Pointer(v3)) = u8(int32(*(*u8)(unsafe.Pointer(v3))) - libc.Int32FromUint8(libc.Uint8FromInt32(w)))
		}
	}
	if i == int32((*slide_ctx)(unsafe.Pointer(ctx)).Fnext_index) {
		v1 = int32((*slide_ctx)(unsafe.Pointer(ctx)).Fnext_digit)
	} else {
		v1 = 0
	}
	return v1
}

func crypto_eddsa_check_equation(tls *libc.TLS, signature uintptr, public_key uintptr, h uintptr) (r int32) {
	bp := tls.Alloc(1440)
	defer tls.Free(1440)
	var h_digit, i1, s_digit, v2 int32
	var i size_t
	var s, sum uintptr
	var _ /* cached at bp+1240 */ ge_cached
	var _ /* check at bp+1400 */ [32]u8
	var _ /* h_slide at bp+992 */ slide_ctx
	var _ /* lutA at bp+352 */ [2]ge_cached
	var _ /* minus_A at bp+0 */ ge
	var _ /* minus_A2 at bp+672 */ ge
	var _ /* minus_R at bp+160 */ ge
	var _ /* s32 at bp+320 */ [8]u32
	var _ /* s_slide at bp+996 */ slide_ctx
	var _ /* t1 at bp+1160 */ fe
	var _ /* t2 at bp+1200 */ fe
	var _ /* tmp at bp+1000 */ ge
	var _ /* tmp at bp+832 */ ge
	_, _, _, _, _, _, _ = h_digit, i, i1, s, s_digit, sum, v2 // -first_half_of_signature
	s = signature + uintptr(32)
	// Check that A and R are on the curve
	// Check that 0 <= S < L (prevents malleability)
	// *Allow* non-cannonical encoding for A and R
	load32_le_buf(tls, bp+320, s, uint64(8))
	if ge_frombytes_neg_vartime(tls, bp, public_key) != 0 || ge_frombytes_neg_vartime(tls, bp+160, signature) != 0 || is_above_l(tls, bp+320) != 0 {
		return -int32(1)
	}
	ge_double(tls, bp+672, bp, bp+832)
	ge_cache(tls, bp+352, bp)
	i = libc.Uint64FromInt32(libc.Int32FromInt32(1))
	for {
		if !(i < libc.Uint64FromInt32(libc.Int32FromInt32(1)<<(libc.Int32FromInt32(P_W_WIDTH)-libc.Int32FromInt32(2)))) {
			break
		}
		ge_add(tls, bp+832, bp+672, bp+352+uintptr(i-uint64(1))*160)
		ge_cache(tls, bp+352+uintptr(i)*160, bp+832)
		goto _1
	_1:
		;
		i = i + 1
	}
	slide_init(tls, bp+992, h)
	slide_init(tls, bp+996, s)
	if libc.Int32FromUint8((**(**slide_ctx)(__ccgo_up(bp + 992))).Fnext_check) >= libc.Int32FromUint8((**(**slide_ctx)(__ccgo_up(bp + 996))).Fnext_check) {
		v2 = libc.Int32FromUint8((**(**slide_ctx)(__ccgo_up(bp + 992))).Fnext_check)
	} else {
		v2 = libc.Int32FromUint8((**(**slide_ctx)(__ccgo_up(bp + 996))).Fnext_check)
	}
	i1 = v2
	sum = bp // reuse minus_A for the sum
	ge_zero(tls, sum)
	for i1 >= 0 {
		ge_double(tls, sum, sum, bp+1000)
		h_digit = slide_step(tls, bp+992, int32(P_W_WIDTH), i1, h)
		s_digit = slide_step(tls, bp+996, int32(B_W_WIDTH), i1, s)
		if h_digit > 0 {
			ge_add(tls, sum, sum, bp+352+uintptr(h_digit/int32(2))*160)
		}
		if h_digit < 0 {
			ge_sub(tls, sum, sum, bp+352+uintptr(-h_digit/int32(2))*160)
		}
		if s_digit > 0 {
			ge_madd(tls, sum, sum, uintptr(unsafe.Pointer(&b_window))+uintptr(s_digit/int32(2))*120, bp+1160, bp+1200)
		}
		if s_digit < 0 {
			ge_msub(tls, sum, sum, uintptr(unsafe.Pointer(&b_window))+uintptr(-s_digit/int32(2))*120, bp+1160, bp+1200)
		}
		i1 = i1 - 1
	}
	ge_cache(tls, bp+1240, bp+160)
	ge_add(tls, sum, sum, bp+1240)
	ge_double(tls, sum, sum, bp+160) // reuse minus_R as temporary
	ge_double(tls, sum, sum, bp+160) // reuse minus_R as temporary
	ge_double(tls, sum, sum, bp+160) // reuse minus_R as temporary
	ge_tobytes(tls, bp+1400, sum)
	return crypto_verify32(tls, bp+1400, uintptr(unsafe.Pointer(&zero_point)))
}

var zero_point = [32]u8{
	0: uint8(1),
} // Point of order 1

// C documentation
//
//	// 5-bit signed comb in cached format (Niels coordinates, Z=1)
var b_comb_low = [8]ge_precomp{
	0: {
		FYp: fe{
			0: -int32(6816601),
			1: -int32(2324159),
			2: -int32(22559413),
			3: int32(124364),
			4: int32(18015490),
			5: int32(8373481),
			6: int32(19993724),
			7: int32(1979872),
			8: -int32(18549925),
			9: int32(9085059),
		},
		FYm: fe{
			0: int32(10306321),
			1: int32(403248),
			2: int32(14839893),
			3: int32(9633706),
			4: int32(8463310),
			5: -int32(8354981),
			6: -int32(14305673),
			7: int32(14668847),
			8: int32(26301366),
			9: int32(2818560),
		},
		FT2: fe{
			0: -int32(22701500),
			1: -int32(3210264),
			2: -int32(13831292),
			3: -int32(2927732),
			4: -int32(16326337),
			5: -int32(14016360),
			6: int32(12940910),
			7: int32(177905),
			8: int32(12165515),
			9: -int32(2397893),
		},
	},
	1: {
		FYp: fe{
			0: -int32(12282262),
			1: -int32(7022066),
			2: int32(9920413),
			3: -int32(3064358),
			4: -int32(32147467),
			5: int32(2927790),
			6: int32(22392436),
			7: -int32(14852487),
			8: int32(2719975),
			9: int32(16402117),
		},
		FYm: fe{
			0: -int32(7236961),
			1: -int32(4729776),
			2: int32(2685954),
			3: -int32(6525055),
			4: -int32(24242706),
			5: -int32(15940211),
			6: -int32(6238521),
			7: int32(14082855),
			8: int32(10047669),
			9: int32(12228189),
		},
		FT2: fe{
			0: -int32(30495588),
			1: -int32(12893761),
			2: -int32(11161261),
			3: int32(3539405),
			4: -int32(11502464),
			5: int32(16491580),
			6: -int32(27286798),
			7: -int32(15030530),
			8: -int32(7272871),
			9: -int32(15934455),
		},
	},
	2: {
		FYp: fe{
			0: int32(17650926),
			1: int32(582297),
			2: -int32(860412),
			3: -int32(187745),
			4: -int32(12072900),
			5: -int32(10683391),
			6: -int32(20352381),
			7: int32(15557840),
			8: -int32(31072141),
			9: -int32(5019061),
		},
		FYm: fe{
			0: -int32(6283632),
			1: -int32(2259834),
			2: -int32(4674247),
			3: -int32(4598977),
			4: -int32(4089240),
			5: int32(12435688),
			6: -int32(31278303),
			7: int32(1060251),
			8: int32(6256175),
			9: int32(10480726),
		},
		FT2: fe{
			0: -int32(13871026),
			1: int32(2026300),
			2: -int32(21928428),
			3: -int32(2741605),
			4: -int32(2406664),
			5: -int32(8034988),
			6: int32(7355518),
			7: int32(15733500),
			8: -int32(23379862),
			9: int32(7489131),
		},
	},
	3: {
		FYp: fe{
			0: int32(6883359),
			1: int32(695140),
			2: int32(23196907),
			3: int32(9644202),
			4: -int32(33430614),
			5: int32(11354760),
			6: -int32(20134606),
			7: int32(6388313),
			8: -int32(8263585),
			9: -int32(8491918),
		},
		FYm: fe{
			0: -int32(7716174),
			1: -int32(13605463),
			2: -int32(13646110),
			3: int32(14757414),
			4: -int32(19430591),
			5: -int32(14967316),
			6: int32(10359532),
			7: -int32(11059670),
			8: -int32(21935259),
			9: int32(12082603),
		},
		FT2: fe{
			0: -int32(11253345),
			1: -int32(15943946),
			2: int32(10046784),
			3: int32(5414629),
			4: int32(24840771),
			5: int32(8086951),
			6: -int32(6694742),
			7: int32(9868723),
			8: int32(15842692),
			9: -int32(16224787),
		},
	},
	4: {
		FYp: fe{
			0: int32(9639399),
			1: int32(11810955),
			2: -int32(24007778),
			3: -int32(9320054),
			4: int32(3912937),
			5: -int32(9856959),
			6: int32(996125),
			7: -int32(8727907),
			8: -int32(8919186),
			9: -int32(14097242),
		},
		FYm: fe{
			0: int32(7248867),
			1: int32(14468564),
			2: int32(25228636),
			3: -int32(8795035),
			4: int32(14346339),
			5: int32(8224790),
			6: int32(6388427),
			7: -int32(7181107),
			8: int32(6468218),
			9: -int32(8720783),
		},
		FT2: fe{
			0: int32(15513115),
			1: int32(15439095),
			2: int32(7342322),
			3: -int32(10157390),
			4: int32(18005294),
			5: -int32(7265713),
			6: int32(2186239),
			7: int32(4884640),
			8: int32(10826567),
			9: int32(7135781),
		},
	},
	5: {
		FYp: fe{
			0: -int32(14204238),
			1: int32(5297536),
			2: -int32(5862318),
			3: -int32(6004934),
			4: int32(28095835),
			5: int32(4236101),
			6: -int32(14203318),
			7: int32(1958636),
			8: -int32(16816875),
			9: int32(3837147),
		},
		FYm: fe{
			0: -int32(5511166),
			1: -int32(13176782),
			2: -int32(29588215),
			3: int32(12339465),
			4: int32(15325758),
			5: -int32(15945770),
			6: -int32(8813185),
			7: int32(11075932),
			8: -int32(19608050),
			9: -int32(3776283),
		},
		FT2: fe{
			0: int32(11728032),
			1: int32(9603156),
			2: -int32(4637821),
			3: -int32(5304487),
			4: -int32(7827751),
			5: int32(2724948),
			6: int32(31236191),
			7: -int32(16760175),
			8: -int32(7268616),
			9: int32(14799772),
		},
	},
	6: {
		FYp: fe{
			0: -int32(28842672),
			1: int32(4840636),
			2: -int32(12047946),
			3: -int32(9101456),
			4: -int32(1445464),
			5: int32(381905),
			6: -int32(30977094),
			7: -int32(16523389),
			8: int32(1290540),
			9: int32(12798615),
		},
		FYm: fe{
			0: int32(27246947),
			1: -int32(10320914),
			2: int32(14792098),
			3: -int32(14518944),
			4: int32(5302070),
			5: -int32(8746152),
			6: -int32(3403974),
			7: -int32(4149637),
			8: -int32(27061213),
			9: int32(10749585),
		},
		FT2: fe{
			0: int32(25572375),
			1: -int32(6270368),
			2: -int32(15353037),
			3: int32(16037944),
			4: int32(1146292),
			5: int32(32198),
			6: int32(23487090),
			7: int32(9585613),
			8: int32(24714571),
			9: -int32(1418265),
		},
	},
	7: {
		FYp: fe{
			0: int32(19844825),
			1: int32(282124),
			2: -int32(17583147),
			3: int32(11004019),
			4: -int32(32004269),
			5: -int32(2716035),
			6: int32(6105106),
			7: -int32(1711007),
			8: -int32(21010044),
			9: int32(14338445),
		},
		FYm: fe{
			0: int32(8027505),
			1: int32(8191102),
			2: -int32(18504907),
			3: -int32(12335737),
			4: int32(25173494),
			5: -int32(5923905),
			6: int32(15446145),
			7: int32(7483684),
			8: -int32(30440441),
			9: int32(10009108),
		},
		FT2: fe{
			0: -int32(14134701),
			1: -int32(4174411),
			2: int32(10246585),
			3: -int32(14677495),
			4: int32(33553567),
			5: -int32(14012935),
			6: int32(23366126),
			7: int32(15080531),
			8: -int32(7969992),
			9: int32(7663473),
		},
	},
}

var b_comb_high = [8]ge_precomp{
	0: {
		FYp: fe{
			0: int32(33055887),
			1: -int32(4431773),
			2: -int32(521787),
			3: int32(6654165),
			4: int32(951411),
			5: -int32(6266464),
			6: -int32(5158124),
			7: int32(6995613),
			8: -int32(5397442),
			9: -int32(6985227),
		},
		FYm: fe{
			0: int32(4014062),
			1: int32(6967095),
			2: -int32(11977872),
			3: int32(3960002),
			4: int32(8001989),
			5: int32(5130302),
			6: -int32(2154812),
			7: -int32(1899602),
			8: -int32(31954493),
			9: -int32(16173976),
		},
		FT2: fe{
			0: int32(16271757),
			1: -int32(9212948),
			2: int32(23792794),
			3: int32(731486),
			4: -int32(25808309),
			5: -int32(3546396),
			6: int32(6964344),
			7: -int32(4767590),
			8: int32(10976593),
			9: int32(10050757),
		},
	},
	1: {
		FYp: fe{
			0: int32(2533007),
			1: -int32(4288439),
			2: -int32(24467768),
			3: -int32(12387405),
			4: -int32(13450051),
			5: int32(14542280),
			6: int32(12876301),
			7: int32(13893535),
			8: int32(15067764),
			9: int32(8594792),
		},
		FYm: fe{
			0: int32(20073501),
			1: -int32(11623621),
			2: int32(3165391),
			3: -int32(13119866),
			4: int32(13188608),
			5: -int32(11540496),
			6: -int32(10751437),
			7: -int32(13482671),
			8: int32(29588810),
			9: int32(2197295),
		},
		FT2: fe{
			0: -int32(1084082),
			1: int32(11831693),
			2: int32(6031797),
			3: int32(14062724),
			4: int32(14748428),
			5: -int32(8159962),
			6: -int32(20721760),
			7: int32(11742548),
			8: int32(31368706),
			9: int32(13161200),
		},
	},
	2: {
		FYp: fe{
			0: int32(2050412),
			1: -int32(6457589),
			2: int32(15321215),
			3: int32(5273360),
			4: int32(25484180),
			5: int32(124590),
			6: -int32(18187548),
			7: -int32(7097255),
			8: -int32(6691621),
			9: -int32(14604792),
		},
		FYm: fe{
			0: int32(9938196),
			1: int32(2162889),
			2: -int32(6158074),
			3: -int32(1711248),
			4: int32(4278932),
			5: -int32(2598531),
			6: -int32(22865792),
			7: -int32(7168500),
			8: -int32(24323168),
			9: int32(11746309),
		},
		FT2: fe{
			0: -int32(22691768),
			1: -int32(14268164),
			2: int32(5965485),
			3: int32(9383325),
			4: int32(20443693),
			5: int32(5854192),
			6: int32(28250679),
			7: -int32(1381811),
			8: -int32(10837134),
			9: int32(13717818),
		},
	},
	3: {
		FYp: fe{
			0: -int32(8495530),
			1: int32(16382250),
			2: int32(9548884),
			3: -int32(4971523),
			4: -int32(4491811),
			5: -int32(3902147),
			6: int32(6182256),
			7: -int32(12832479),
			8: int32(26628081),
			9: int32(10395408),
		},
		FYm: fe{
			0: int32(27329048),
			1: -int32(15853735),
			2: int32(7715764),
			3: int32(8717446),
			4: -int32(9215518),
			5: -int32(14633480),
			6: int32(28982250),
			7: -int32(5668414),
			8: int32(4227628),
			9: int32(242148),
		},
		FT2: fe{
			0: -int32(13279943),
			1: -int32(7986904),
			2: -int32(7100016),
			3: int32(8764468),
			4: -int32(27276630),
			5: int32(3096719),
			6: int32(29678419),
			7: -int32(9141299),
			8: int32(3906709),
			9: int32(11265498),
		},
	},
	4: {
		FYp: fe{
			0: int32(11918285),
			1: int32(15686328),
			2: -int32(17757323),
			3: -int32(11217300),
			4: -int32(27548967),
			5: int32(4853165),
			6: -int32(27168827),
			7: int32(6807359),
			8: int32(6871949),
			9: -int32(1075745),
		},
		FYm: fe{
			0: -int32(29002610),
			1: int32(13984323),
			2: -int32(27111812),
			3: -int32(2713442),
			4: int32(28107359),
			5: -int32(13266203),
			6: int32(6155126),
			7: int32(15104658),
			8: int32(3538727),
			9: -int32(7513788),
		},
		FT2: fe{
			0: int32(14103158),
			1: int32(11233913),
			2: -int32(33165269),
			3: int32(9279850),
			4: int32(31014152),
			5: int32(4335090),
			6: -int32(1827936),
			7: int32(4590951),
			8: int32(13960841),
			9: int32(12787712),
		},
	},
	5: {
		FYp: fe{
			0: int32(1469134),
			1: -int32(16738009),
			2: int32(33411928),
			3: int32(13942824),
			4: int32(8092558),
			5: -int32(8778224),
			6: -int32(11165065),
			7: int32(1437842),
			8: int32(22521552),
			9: -int32(2792954),
		},
		FYm: fe{
			0: int32(31352705),
			1: -int32(4807352),
			2: -int32(25327300),
			3: int32(3962447),
			4: int32(12541566),
			5: -int32(9399651),
			6: -int32(27425693),
			7: int32(7964818),
			8: -int32(23829869),
			9: int32(5541287),
		},
		FT2: fe{
			0: -int32(25732021),
			1: -int32(6864887),
			2: int32(23848984),
			3: int32(3039395),
			4: -int32(9147354),
			5: int32(6022816),
			6: -int32(27421653),
			7: int32(10590137),
			8: int32(25309915),
			9: -int32(1584678),
		},
	},
	6: {
		FYp: fe{
			0: -int32(22951376),
			1: int32(5048948),
			2: int32(31139401),
			3: -int32(190316),
			4: -int32(19542447),
			5: -int32(626310),
			6: -int32(17486305),
			7: -int32(16511925),
			8: -int32(18851313),
			9: -int32(12985140),
		},
		FYm: fe{
			0: -int32(9684890),
			1: int32(14681754),
			2: int32(30487568),
			3: int32(7717771),
			4: -int32(10829709),
			5: int32(9630497),
			6: int32(30290549),
			7: -int32(10531496),
			8: -int32(27798994),
			9: -int32(13812825),
		},
		FT2: fe{
			0: int32(5827835),
			1: int32(16097107),
			2: -int32(24501327),
			3: int32(12094619),
			4: int32(7413972),
			5: int32(11447087),
			6: int32(28057551),
			7: -int32(1793987),
			8: -int32(14056981),
			9: int32(4359312),
		},
	},
	7: {
		FYp: fe{
			0: int32(26323183),
			1: int32(2342588),
			2: -int32(21887793),
			3: -int32(1623758),
			4: -int32(6062284),
			5: int32(2107090),
			6: -int32(28724907),
			7: int32(9036464),
			8: -int32(19618351),
			9: -int32(13055189),
		},
		FYm: fe{
			0: -int32(29697200),
			1: int32(14829398),
			2: -int32(4596333),
			3: int32(14220089),
			4: -int32(30022969),
			5: int32(2955645),
			6: int32(12094100),
			7: -int32(13693652),
			8: -int32(5941445),
			9: int32(7047569),
		},
		FT2: fe{
			0: -int32(3201977),
			1: int32(14413268),
			2: -int32(12058324),
			3: -int32(16417589),
			4: -int32(9035655),
			5: -int32(7224648),
			6: int32(9258160),
			7: int32(1399236),
			8: int32(30397584),
			9: -int32(5684634),
		},
	},
}

func lookup_add(tls *libc.TLS, p uintptr, tmp_c uintptr, tmp_a uintptr, tmp_b uintptr, comb uintptr, scalar uintptr, i int32) {
	var high, index, teeth u8
	var j size_t
	var select1 i32
	_, _, _, _, _ = high, index, j, select1, teeth
	teeth = libc.Uint8FromInt32(scalar_bit(tls, scalar, i) + scalar_bit(tls, scalar, i+int32(32))<<libc.Int32FromInt32(1) + scalar_bit(tls, scalar, i+int32(64))<<libc.Int32FromInt32(2) + scalar_bit(tls, scalar, i+int32(96))<<libc.Int32FromInt32(3))
	high = libc.Uint8FromInt32(libc.Int32FromUint8(teeth) >> int32(3))
	index = libc.Uint8FromInt32((libc.Int32FromUint8(teeth) ^ (libc.Int32FromUint8(high) - int32(1))) & int32(7))
	j = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(j < libc.Uint64FromInt32(libc.Int32FromInt32(8))) {
			break
		}
		select1 = libc.Int32FromUint64(uint64(1) & ((j ^ uint64(index) - uint64(1)) >> int32(8)))
		fe_ccopy(tls, tmp_c, comb+uintptr(j)*120, select1)
		fe_ccopy(tls, tmp_c+40, comb+uintptr(j)*120+40, select1)
		fe_ccopy(tls, tmp_c+80, comb+uintptr(j)*120+80, select1)
		goto _1
	_1:
		;
		j = j + 1
	}
	fe_neg(tls, tmp_a, tmp_c+80)
	fe_cswap(tls, tmp_c+80, tmp_a, libc.Int32FromUint8(high)^int32(1))
	fe_cswap(tls, tmp_c, tmp_c+40, libc.Int32FromUint8(high)^int32(1))
	ge_madd(tls, p, p, tmp_c, tmp_a, tmp_b)
}

// C documentation
//
//	// p = [scalar]B, where B is the base point
func ge_scalarmult_base(tls *libc.TLS, p uintptr, scalar uintptr) {
	bp := tls.Alloc(400)
	defer tls.Free(400)
	var i int32
	var _ /* s_scalar at bp+0 */ [32]u8
	var _ /* tmp_a at bp+32 */ fe
	var _ /* tmp_b at bp+72 */ fe
	var _ /* tmp_c at bp+112 */ ge_precomp
	var _ /* tmp_d at bp+232 */ ge
	_ = i
	crypto_eddsa_mul_add(tls, bp, scalar, uintptr(unsafe.Pointer(&half_mod_L)), uintptr(unsafe.Pointer(&half_ones))) // temporary for doubling
	fe_1(tls, bp+112)
	fe_1(tls, bp+112+40)
	fe_0(tls, bp+112+80)
	// Save a double on the first iteration
	ge_zero(tls, p)
	lookup_add(tls, p, bp+112, bp+32, bp+72, uintptr(unsafe.Pointer(&b_comb_low)), bp, int32(31))
	lookup_add(tls, p, bp+112, bp+32, bp+72, uintptr(unsafe.Pointer(&b_comb_high)), bp, libc.Int32FromInt32(31)+libc.Int32FromInt32(128))
	// Regular double & add for the rest
	i = int32(30)
	for {
		if !(i >= 0) {
			break
		}
		ge_double(tls, p, p, bp+232)
		lookup_add(tls, p, bp+112, bp+32, bp+72, uintptr(unsafe.Pointer(&b_comb_low)), bp, i)
		lookup_add(tls, p, bp+112, bp+32, bp+72, uintptr(unsafe.Pointer(&b_comb_high)), bp, i+int32(128))
		goto _1
	_1:
		;
		i = i - 1
	}
	// Note: we could save one addition at the end if we assumed the
	// scalar fit in 252 bits.  Which it does in practice if it is
	// selected at random.  However, non-random, non-hashed scalars
	// *can* overflow 252 bits in practice.  Better account for that
	// than leaving that kind of subtle corner case.
	crypto_wipe(tls, bp+32, uint64(40))
	crypto_wipe(tls, bp+232, uint64(160))
	crypto_wipe(tls, bp+72, uint64(40))
	crypto_wipe(tls, bp+112, uint64(120))
	crypto_wipe(tls, bp, uint64(32))
}

// twin 4-bits signed combs, from Mike Hamburg's
// Fast and compact elliptic-curve cryptography (2012)
// 1 / 2 modulo L
var half_mod_L = [32]u8{
	0:  uint8(247),
	1:  uint8(233),
	2:  uint8(122),
	3:  uint8(46),
	4:  uint8(141),
	5:  uint8(49),
	6:  uint8(9),
	7:  uint8(44),
	8:  uint8(107),
	9:  uint8(206),
	10: uint8(123),
	11: uint8(81),
	12: uint8(239),
	13: uint8(124),
	14: uint8(111),
	15: uint8(10),
	31: uint8(8),
}

// (2^256 - 1) / 2 modulo L
var half_ones = [32]u8{
	0:  uint8(142),
	1:  uint8(74),
	2:  uint8(204),
	3:  uint8(70),
	4:  uint8(186),
	5:  uint8(24),
	6:  uint8(118),
	7:  uint8(107),
	8:  uint8(184),
	9:  uint8(231),
	10: uint8(190),
	11: uint8(57),
	12: uint8(250),
	13: uint8(173),
	14: uint8(119),
	15: uint8(99),
	16: uint8(255),
	17: uint8(255),
	18: uint8(255),
	19: uint8(255),
	20: uint8(255),
	21: uint8(255),
	22: uint8(255),
	23: uint8(255),
	24: uint8(255),
	25: uint8(255),
	26: uint8(255),
	27: uint8(255),
	28: uint8(255),
	29: uint8(255),
	30: uint8(255),
	31: uint8(7),
}

func crypto_eddsa_scalarbase(tls *libc.TLS, point uintptr, scalar uintptr) {
	bp := tls.Alloc(160)
	defer tls.Free(160)
	var _ /* P at bp+0 */ ge
	ge_scalarmult_base(tls, bp, scalar)
	ge_tobytes(tls, point, bp)
	crypto_wipe(tls, bp, uint64(160))
}

func crypto_eddsa_key_pair(tls *libc.TLS, secret_key uintptr, public_key uintptr, seed uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var _i_, _i_1, _i_2 size_t
	var _ /* a at bp+0 */ [64]u8
	_, _, _ = _i_, _i_1, _i_2
	_i_ = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(_i_ < libc.Uint64FromInt32(libc.Int32FromInt32(32))) {
			break
		}
		(**(**[64]u8)(__ccgo_up(bp)))[_i_] = **(**u8)(__ccgo_up(seed + uintptr(_i_)))
		goto _1
	_1:
		;
		_i_ = _i_ + 1
	}
	crypto_wipe(tls, seed, uint64(32))
	_i_1 = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(_i_1 < libc.Uint64FromInt32(libc.Int32FromInt32(32))) {
			break
		}
		**(**u8)(__ccgo_up(secret_key + uintptr(_i_1))) = (**(**[64]u8)(__ccgo_up(bp)))[_i_1]
		goto _2
	_2:
		;
		_i_1 = _i_1 + 1
	}
	crypto_blake2b(tls, bp, uint64(64), bp, uint64(32))
	crypto_eddsa_trim_scalar(tls, bp, bp)
	crypto_eddsa_scalarbase(tls, secret_key+uintptr(32), bp)
	_i_2 = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(_i_2 < libc.Uint64FromInt32(libc.Int32FromInt32(32))) {
			break
		}
		**(**u8)(__ccgo_up(public_key + uintptr(_i_2))) = **(**u8)(__ccgo_up(secret_key + libc.UintptrFromInt32(32) + uintptr(_i_2)))
		goto _3
	_3:
		;
		_i_2 = _i_2 + 1
	}
	crypto_wipe(tls, bp, uint64(64))
}

func hash_reduce(tls *libc.TLS, h uintptr, a uintptr, a_size size_t, b uintptr, b_size size_t, c uintptr, c_size size_t) {
	bp := tls.Alloc(288)
	defer tls.Free(288)
	var _ /* ctx at bp+64 */ crypto_blake2b_ctx
	var _ /* hash at bp+0 */ [64]u8
	crypto_blake2b_init(tls, bp+64, uint64(64))
	crypto_blake2b_update(tls, bp+64, a, a_size)
	crypto_blake2b_update(tls, bp+64, b, b_size)
	crypto_blake2b_update(tls, bp+64, c, c_size)
	crypto_blake2b_final(tls, bp+64, bp)
	crypto_eddsa_reduce(tls, h, bp)
}

// C documentation
//
//	// Digital signature of a message with from a secret key.
//	//
//	// The secret key comprises two parts:
//	// - The seed that generates the key (secret_key[ 0..31])
//	// - The public key                  (secret_key[32..63])
//	//
//	// The seed and the public key are bundled together to make sure users
//	// don't use mismatched seeds and public keys, which would instantly
//	// leak the secret scalar and allow forgeries (allowing this to happen
//	// has resulted in critical vulnerabilities in the wild).
//	//
//	// The seed is hashed to derive the secret scalar and a secret prefix.
//	// The sole purpose of the prefix is to generate a secret random nonce.
//	// The properties of that nonce must be as follows:
//	// - Unique: we need a different one for each message.
//	// - Secret: third parties must not be able to predict it.
//	// - Random: any detectable bias would break all security.
//	//
//	// There are two ways to achieve these properties.  The obvious one is
//	// to simply generate a random number.  Here that would be a parameter
//	// (Monocypher doesn't have an RNG).  It works, but then users may reuse
//	// the nonce by accident, which _also_ leaks the secret scalar and
//	// allows forgeries.  This has happened in the wild too.
//	//
//	// This is no good, so instead we generate that nonce deterministically
//	// by reducing modulo L a hash of the secret prefix and the message.
//	// The secret prefix makes the nonce unpredictable, the message makes it
//	// unique, and the hash/reduce removes all bias.
//	//
//	// The cost of that safety is hashing the message twice.  If that cost
//	// is unacceptable, there are two alternatives:
//	//
//	// - Signing a hash of the message instead of the message itself.  This
//	//   is fine as long as the hash is collision resistant. It is not
//	//   compatible with existing "pure" signatures, but at least it's safe.
//	//
//	// - Using a random nonce.  Please exercise **EXTREME CAUTION** if you
//	//   ever do that.  It is absolutely **critical** that the nonce is
//	//   really an unbiased random number between 0 and L-1, never reused,
//	//   and wiped immediately.
//	//
//	//   To lower the likelihood of complete catastrophe if the RNG is
//	//   either flawed or misused, you can hash the RNG output together with
//	//   the secret prefix and the beginning of the message, and use the
//	//   reduction of that hash instead of the RNG output itself.  It's not
//	//   foolproof (you'd need to hash the whole message) but it helps.
//	//
//	// Signing a message involves the following operations:
//	//
//	//   scalar, prefix = HASH(secret_key)
//	//   r              = HASH(prefix || message) % L
//	//   R              = [r]B
//	//   h              = HASH(R || public_key || message) % L
//	//   S              = ((h * a) + r) % L
//	//   signature      = R || S
func crypto_eddsa_sign(tls *libc.TLS, signature uintptr, secret_key uintptr, message uintptr, message_size size_t) {
	bp := tls.Alloc(160)
	defer tls.Free(160)
	var _i_ size_t
	var _ /* R at bp+128 */ [32]u8
	var _ /* a at bp+0 */ [64]u8
	var _ /* h at bp+96 */ [32]u8
	var _ /* r at bp+64 */ [32]u8
	_ = _i_ // first half of the signature (allows overlapping inputs)
	crypto_blake2b(tls, bp, uint64(64), secret_key, uint64(32))
	crypto_eddsa_trim_scalar(tls, bp, bp)
	hash_reduce(tls, bp+64, bp+uintptr(32), uint64(32), message, message_size, uintptr(0), uint64(0))
	crypto_eddsa_scalarbase(tls, bp+128, bp+64)
	hash_reduce(tls, bp+96, bp+128, uint64(32), secret_key+uintptr(32), uint64(32), message, message_size)
	_i_ = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(_i_ < libc.Uint64FromInt32(libc.Int32FromInt32(32))) {
			break
		}
		**(**u8)(__ccgo_up(signature + uintptr(_i_))) = (**(**[32]u8)(__ccgo_up(bp + 128)))[_i_]
		goto _1
	_1:
		;
		_i_ = _i_ + 1
	}
	crypto_eddsa_mul_add(tls, signature+uintptr(32), bp+96, bp, bp+64)
	crypto_wipe(tls, bp, uint64(64))
	crypto_wipe(tls, bp+64, uint64(32))
}

// C documentation
//
//	// To check the signature R, S of the message M with the public key A,
//	// there are 3 steps:
//	//
//	//   compute h = HASH(R || A || message) % L
//	//   check that A is on the curve.
//	//   check that R == [s]B - [h]A
//	//
//	// The last two steps are done in crypto_eddsa_check_equation()
func crypto_eddsa_check(tls *libc.TLS, signature uintptr, public_key uintptr, message uintptr, message_size size_t) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var _ /* h at bp+0 */ [32]u8
	hash_reduce(tls, bp, signature, uint64(32), public_key, uint64(32), message, message_size)
	return crypto_eddsa_check_equation(tls, signature, public_key, bp)
}

// C documentation
//
//	/////////////////////////
//	/// EdDSA <--> X25519 ///
//	/////////////////////////
func crypto_eddsa_to_x25519(tls *libc.TLS, x25519 uintptr, eddsa uintptr) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var _ /* t1 at bp+0 */ fe
	var _ /* t2 at bp+40 */ fe
	fe_frombytes(tls, bp+40, eddsa)
	fe_add(tls, bp, uintptr(unsafe.Pointer(&fe_one)), bp+40)
	fe_sub(tls, bp+40, uintptr(unsafe.Pointer(&fe_one)), bp+40)
	fe_invert(tls, bp+40, bp+40)
	fe_mul(tls, bp, bp, bp+40)
	fe_tobytes(tls, x25519, bp)
	crypto_wipe(tls, bp, uint64(40))
	crypto_wipe(tls, bp+40, uint64(40))
}

func crypto_x25519_to_eddsa(tls *libc.TLS, eddsa uintptr, x25519 uintptr) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var _ /* t1 at bp+0 */ fe
	var _ /* t2 at bp+40 */ fe
	fe_frombytes(tls, bp+40, x25519)
	fe_sub(tls, bp, bp+40, uintptr(unsafe.Pointer(&fe_one)))
	fe_add(tls, bp+40, bp+40, uintptr(unsafe.Pointer(&fe_one)))
	fe_invert(tls, bp+40, bp+40)
	fe_mul(tls, bp, bp, bp+40)
	fe_tobytes(tls, eddsa, bp)
	crypto_wipe(tls, bp, uint64(40))
	crypto_wipe(tls, bp+40, uint64(40))
}

/////////////////////////////////////////////
/// Dirty ephemeral public key generation ///
/////////////////////////////////////////////

// Those functions generates a public key, *without* clearing the
// cofactor.  Sending that key over the network leaks 3 bits of the
// private key.  Use only to generate ephemeral keys that will be hidden
// with crypto_curve_to_hidden().
//
// The public key is otherwise compatible with crypto_x25519(), which
// properly clears the cofactor.
//
// Note that the distribution of the resulting public keys is almost
// uniform.  Flipping the sign of the v coordinate (not provided by this
// function), covers the entire key space almost perfectly, where
// "almost" means a 2^-128 bias (undetectable).  This uniformity is
// needed to ensure the proper randomness of the resulting
// representatives (once we apply crypto_curve_to_hidden()).
//
// Recall that Curve25519 has order C = 2^255 + e, with e < 2^128 (not
// to be confused with the prime order of the main subgroup, L, which is
// 8 times less than that).
//
// Generating all points would require us to multiply a point of order C
// (the base point plus any point of order 8) by all scalars from 0 to
// C-1.  Clamping limits us to scalars between 2^254 and 2^255 - 1. But
// by negating the resulting point at random, we also cover scalars from
// -2^255 + 1 to -2^254 (which modulo C is congruent to e+1 to 2^254 + e).
//
// In practice:
// - Scalars from 0         to e + 1     are never generated
// - Scalars from 2^255     to 2^255 + e are never generated
// - Scalars from 2^254 + 1 to 2^254 + e are generated twice
//
// Since e < 2^128, detecting this bias requires observing over 2^100
// representatives from a given source (this will never happen), *and*
// recovering enough of the private key to determine that they do, or do
// not, belong to the biased set (this practically requires solving
// discrete logarithm, which is conjecturally intractable).
//
// In practice, this means the bias is impossible to detect.

// C documentation
//
//	// s + (x*L) % 8*L
//	// Guaranteed to fit in 256 bits iff s fits in 255 bits.
//	//   L             < 2^253
//	//   x%8           < 2^3
//	//   L * (x%8)     < 2^255
//	//   s             < 2^255
//	//   s + L * (x%8) < 2^256
func add_xl(tls *libc.TLS, s uintptr, x u8) {
	var carry, mod8 u64
	var i size_t
	_, _, _ = carry, i, mod8
	mod8 = libc.Uint64FromInt32(libc.Int32FromUint8(x) & int32(7))
	carry = uint64(0)
	i = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(i < libc.Uint64FromInt32(libc.Int32FromInt32(8))) {
			break
		}
		carry = carry + uint64(load32_le(tls, s+uintptr(uint64(4)*i))) + uint64(L[i])*mod8
		store32_le(tls, s+uintptr(uint64(4)*i), uint32(carry))
		carry = carry >> uint64(32)
		goto _1
	_1:
		;
		i = i + 1
	}
}

// C documentation
//
//	// "Small" dirty ephemeral key.
//	// Use if you need to shrink the size of the binary, and can afford to
//	// slow down by a factor of two (compared to the fast version)
//	//
//	// This version works by decoupling the cofactor from the main factor.
//	//
//	// - The trimmed scalar determines the main factor
//	// - The clamped bits of the scalar determine the cofactor.
//	//
//	// Cofactor and main factor are combined into a single scalar, which is
//	// then multiplied by a point of order 8*L (unlike the base point, which
//	// has prime order).  That "dirty" base point is the addition of the
//	// regular base point (9), and a point of order 8.
func crypto_x25519_dirty_small(tls *libc.TLS, public_key uintptr, secret_key uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var _ /* scalar at bp+0 */ [32]u8
	crypto_eddsa_trim_scalar(tls, bp, secret_key)
	// Separate the main factor and the cofactor
	//
	// The scalar is trimmed, so its cofactor is cleared.  The three
	// least significant bits however still have a main factor.  We must
	// remove it for X25519 compatibility.
	//
	//   cofactor = lsb * L            (modulo 8*L)
	//   combined = scalar + cofactor  (modulo 8*L)
	add_xl(tls, bp, **(**u8)(__ccgo_up(secret_key)))
	scalarmult(tls, public_key, bp, uintptr(unsafe.Pointer(&dirty_base_point)), int32(256))
	crypto_wipe(tls, bp, uint64(32))
}

// Base point of order 8*L
// Raw scalar multiplication with it does not clear the cofactor,
// and the resulting public key will reveal 3 bits of the scalar.
//
// The low order component of this base point  has been chosen
// to yield the same results as crypto_x25519_dirty_fast().
var dirty_base_point = [32]u8{
	0:  uint8(0xd8),
	1:  uint8(0x86),
	2:  uint8(0x1a),
	3:  uint8(0xa2),
	4:  uint8(0x78),
	5:  uint8(0x7a),
	6:  uint8(0xd9),
	7:  uint8(0x26),
	8:  uint8(0x8b),
	9:  uint8(0x74),
	10: uint8(0x74),
	11: uint8(0xb6),
	12: uint8(0x82),
	13: uint8(0xe3),
	14: uint8(0xbe),
	15: uint8(0xc3),
	16: uint8(0xce),
	17: uint8(0x36),
	18: uint8(0x9a),
	19: uint8(0x1e),
	20: uint8(0x5e),
	21: uint8(0x31),
	22: uint8(0x47),
	23: uint8(0xa2),
	24: uint8(0x6d),
	25: uint8(0x37),
	26: uint8(0x7c),
	27: uint8(0xfd),
	28: uint8(0x20),
	29: uint8(0xb5),
	30: uint8(0xdf),
	31: uint8(0x75),
}

// C documentation
//
//	// Select low order point
//	// We're computing the [cofactor]lop scalar multiplication, where:
//	//
//	//   cofactor = tweak & 7.
//	//   lop      = (lop_x, lop_y)
//	//   lop_x    = sqrt((sqrt(d + 1) + 1) / d)
//	//   lop_y    = -lop_x * sqrtm1
//	//
//	// The low order point has order 8. There are 4 such points.  We've
//	// chosen the one whose both coordinates are positive (below p/2).
//	// The 8 low order points are as follows:
//	//
//	// [0]lop = ( 0       ,  1    )
//	// [1]lop = ( lop_x   ,  lop_y)
//	// [2]lop = ( sqrt(-1), -0    )
//	// [3]lop = ( lop_x   , -lop_y)
//	// [4]lop = (-0       , -1    )
//	// [5]lop = (-lop_x   , -lop_y)
//	// [6]lop = (-sqrt(-1),  0    )
//	// [7]lop = (-lop_x   ,  lop_y)
//	//
//	// The x coordinate is either 0, sqrt(-1), lop_x, or their opposite.
//	// The y coordinate is either 0,      -1 , lop_y, or their opposite.
//	// The pattern for both is the same, except for a rotation of 2 (modulo 8)
//	//
//	// This helper function captures the pattern, and we can use it thus:
//	//
//	//    select_lop(x, lop_x, sqrtm1, cofactor);
//	//    select_lop(y, lop_y, fe_one, cofactor + 2);
//	//
//	// This is faster than an actual scalar multiplication,
//	// and requires less code than naive constant time look up.
func select_lop(tls *libc.TLS, out uintptr, x uintptr, k uintptr, cofactor u8) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var _ /* tmp at bp+0 */ fe
	fe_0(tls, out)
	fe_ccopy(tls, out, k, libc.Int32FromUint8(cofactor)>>int32(1)&int32(1)) // bit 1
	fe_ccopy(tls, out, x, libc.Int32FromUint8(cofactor)>>0&int32(1))        // bit 0
	fe_neg(tls, bp, out)
	fe_ccopy(tls, out, bp, libc.Int32FromUint8(cofactor)>>int32(2)&int32(1)) // bit 2
	crypto_wipe(tls, bp, uint64(40))
}

// C documentation
//
//	// "Fast" dirty ephemeral key
//	// We use this one by default.
//	//
//	// This version works by performing a regular scalar multiplication,
//	// then add a low order point.  The scalar multiplication is done in
//	// Edwards space for more speed (*2 compared to the "small" version).
//	// The cost is a bigger binary for programs that don't also sign messages.
func crypto_x25519_dirty_fast(tls *libc.TLS, public_key uintptr, secret_key uintptr) {
	bp := tls.Alloc(400)
	defer tls.Free(400)
	var _ /* low_order_point at bp+272 */ ge_precomp
	var _ /* pk at bp+32 */ ge
	var _ /* scalar at bp+0 */ [32]u8
	var _ /* t1 at bp+192 */ fe
	var _ /* t2 at bp+232 */ fe
	crypto_eddsa_trim_scalar(tls, bp, secret_key)
	ge_scalarmult_base(tls, bp+32, bp)
	select_lop(tls, bp+192, uintptr(unsafe.Pointer(&lop_x)), uintptr(unsafe.Pointer(&sqrtm1)), **(**u8)(__ccgo_up(secret_key)))
	select_lop(tls, bp+232, uintptr(unsafe.Pointer(&lop_y)), uintptr(unsafe.Pointer(&fe_one)), libc.Uint8FromInt32(libc.Int32FromUint8(**(**u8)(__ccgo_up(secret_key)))+int32(2)))
	fe_add(tls, bp+272, bp+232, bp+192)
	fe_sub(tls, bp+272+40, bp+232, bp+192)
	fe_mul(tls, bp+272+80, bp+232, bp+192)
	fe_mul(tls, bp+272+80, bp+272+80, uintptr(unsafe.Pointer(&D2)))
	// Add low order point to the public key
	ge_madd(tls, bp+32, bp+32, bp+272, bp+192, bp+232)
	// Convert to Montgomery u coordinate (we ignore the sign)
	fe_add(tls, bp+192, bp+32+80, bp+32+40)
	fe_sub(tls, bp+232, bp+32+80, bp+32+40)
	fe_invert(tls, bp+232, bp+232)
	fe_mul(tls, bp+192, bp+192, bp+232)
	fe_tobytes(tls, public_key, bp+192)
	crypto_wipe(tls, bp+192, uint64(40))
	crypto_wipe(tls, bp+32, uint64(160))
	crypto_wipe(tls, bp+232, uint64(40))
	crypto_wipe(tls, bp+272, uint64(120))
	crypto_wipe(tls, bp, uint64(32))
}

// C documentation
//
//	///////////////////
//	/// Elligator 2 ///
//	///////////////////
var A = fe{
	0: int32(486662),
}

// C documentation
//
//	// Elligator direct map
//	//
//	// Computes the point corresponding to a representative, encoded in 32
//	// bytes (little Endian).  Since positive representatives fits in 254
//	// bits, The two most significant bits are ignored.
//	//
//	// From the paper:
//	// w = -A / (fe(1) + non_square * r^2)
//	// e = chi(w^3 + A*w^2 + w)
//	// u = e*w - (fe(1)-e)*(A//2)
//	// v = -e * sqrt(u^3 + A*u^2 + u)
//	//
//	// We ignore v because we don't need it for X25519 (the Montgomery
//	// ladder only uses u).
//	//
//	// Note that e is either 0, 1 or -1
//	// if e = 0    u = 0  and v = 0
//	// if e = 1    u = w
//	// if e = -1   u = -w - A = w * non_square * r^2
//	//
//	// Let r1 = non_square * r^2
//	// Let r2 = 1 + r1
//	// Note that r2 cannot be zero, -1/non_square is not a square.
//	// We can (tediously) verify that:
//	//   w^3 + A*w^2 + w = (A^2*r1 - r2^2) * A / r2^3
//	// Therefore:
//	//   chi(w^3 + A*w^2 + w) = chi((A^2*r1 - r2^2) * (A / r2^3))
//	//   chi(w^3 + A*w^2 + w) = chi((A^2*r1 - r2^2) * (A / r2^3)) * 1
//	//   chi(w^3 + A*w^2 + w) = chi((A^2*r1 - r2^2) * (A / r2^3)) * chi(r2^6)
//	//   chi(w^3 + A*w^2 + w) = chi((A^2*r1 - r2^2) * (A / r2^3)  *     r2^6)
//	//   chi(w^3 + A*w^2 + w) = chi((A^2*r1 - r2^2) *  A * r2^3)
//	// Corollary:
//	//   e =  1 if (A^2*r1 - r2^2) *  A * r2^3) is a non-zero square
//	//   e = -1 if (A^2*r1 - r2^2) *  A * r2^3) is not a square
//	//   Note that w^3 + A*w^2 + w (and therefore e) can never be zero:
//	//     w^3 + A*w^2 + w = w * (w^2 + A*w + 1)
//	//     w^3 + A*w^2 + w = w * (w^2 + A*w + A^2/4 - A^2/4 + 1)
//	//     w^3 + A*w^2 + w = w * (w + A/2)^2        - A^2/4 + 1)
//	//     which is zero only if:
//	//       w = 0                   (impossible)
//	//       (w + A/2)^2 = A^2/4 - 1 (impossible, because A^2/4-1 is not a square)
//	//
//	// Let isr   = invsqrt((A^2*r1 - r2^2) *  A * r2^3)
//	//     isr   = sqrt(1        / ((A^2*r1 - r2^2) *  A * r2^3)) if e =  1
//	//     isr   = sqrt(sqrt(-1) / ((A^2*r1 - r2^2) *  A * r2^3)) if e = -1
//	//
//	// if e = 1
//	//   let u1 = -A * (A^2*r1 - r2^2) * A * r2^2 * isr^2
//	//       u1 = w
//	//       u1 = u
//	//
//	// if e = -1
//	//   let ufactor = -non_square * sqrt(-1) * r^2
//	//   let vfactor = sqrt(ufactor)
//	//   let u2 = -A * (A^2*r1 - r2^2) * A * r2^2 * isr^2 * ufactor
//	//       u2 = w * -1 * -non_square * r^2
//	//       u2 = w * non_square * r^2
//	//       u2 = u
func crypto_elligator_map(tls *libc.TLS, curve uintptr, hidden uintptr) {
	bp := tls.Alloc(208)
	defer tls.Free(208)
	var is_square int32
	var _ /* r at bp+0 */ fe
	var _ /* t1 at bp+80 */ fe
	var _ /* t2 at bp+120 */ fe
	var _ /* t3 at bp+160 */ fe
	var _ /* u at bp+40 */ fe
	_ = is_square
	fe_frombytes_mask(tls, bp, hidden, uint32(2)) // r is encoded in 254 bits.
	fe_sq(tls, bp, bp)
	fe_add(tls, bp+80, bp, bp)
	fe_add(tls, bp+40, bp+80, uintptr(unsafe.Pointer(&fe_one)))
	fe_sq(tls, bp+120, bp+40)
	fe_mul(tls, bp+160, uintptr(unsafe.Pointer(&A2)), bp+80)
	fe_sub(tls, bp+160, bp+160, bp+120)
	fe_mul(tls, bp+160, bp+160, uintptr(unsafe.Pointer(&A)))
	fe_mul(tls, bp+80, bp+120, bp+40)
	fe_mul(tls, bp+80, bp+160, bp+80)
	is_square = invsqrt(tls, bp+80, bp+80)
	fe_mul(tls, bp+40, bp, uintptr(unsafe.Pointer(&ufactor)))
	fe_ccopy(tls, bp+40, uintptr(unsafe.Pointer(&fe_one)), is_square)
	fe_sq(tls, bp+80, bp+80)
	fe_mul(tls, bp+40, bp+40, uintptr(unsafe.Pointer(&A)))
	fe_mul(tls, bp+40, bp+40, bp+160)
	fe_mul(tls, bp+40, bp+40, bp+120)
	fe_mul(tls, bp+40, bp+40, bp+80)
	fe_neg(tls, bp+40, bp+40)
	fe_tobytes(tls, curve, bp+40)
	crypto_wipe(tls, bp+80, uint64(40))
	crypto_wipe(tls, bp, uint64(40))
	crypto_wipe(tls, bp+120, uint64(40))
	crypto_wipe(tls, bp+40, uint64(40))
	crypto_wipe(tls, bp+160, uint64(40))
}

// C documentation
//
//	// Elligator inverse map
//	//
//	// Computes the representative of a point, if possible.  If not, it does
//	// nothing and returns -1.  Note that the success of the operation
//	// depends only on the point (more precisely its u coordinate).  The
//	// tweak parameter is used only upon success
//	//
//	// The tweak should be a random byte.  Beyond that, its contents are an
//	// implementation detail. Currently, the tweak comprises:
//	// - Bit  1  : sign of the v coordinate (0 if positive, 1 if negative)
//	// - Bit  2-5: not used
//	// - Bits 6-7: random padding
//	//
//	// From the paper:
//	// Let sq = -non_square * u * (u+A)
//	// if sq is not a square, or u = -A, there is no mapping
//	// Assuming there is a mapping:
//	//    if v is positive: r = sqrt(-u     / (non_square * (u+A)))
//	//    if v is negative: r = sqrt(-(u+A) / (non_square * u    ))
//	//
//	// We compute isr = invsqrt(-non_square * u * (u+A))
//	// if it wasn't a square, abort.
//	// else, isr = sqrt(-1 / (non_square * u * (u+A))
//	//
//	// If v is positive, we return isr * u:
//	//   isr * u = sqrt(-1 / (non_square * u * (u+A)) * u
//	//   isr * u = sqrt(-u / (non_square * (u+A))
//	//
//	// If v is negative, we return isr * (u+A):
//	//   isr * (u+A) = sqrt(-1     / (non_square * u * (u+A)) * (u+A)
//	//   isr * (u+A) = sqrt(-(u+A) / (non_square * u)
func crypto_elligator_rev(tls *libc.TLS, hidden uintptr, public_key uintptr, tweak u8) (r int32) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var is_square int32
	var v1 uintptr
	var _ /* t1 at bp+0 */ fe
	var _ /* t2 at bp+40 */ fe
	var _ /* t3 at bp+80 */ fe
	_, _ = is_square, v1
	fe_frombytes(tls, bp, public_key)                   // t1 = u
	fe_add(tls, bp+40, bp, uintptr(unsafe.Pointer(&A))) // t2 = u + A
	fe_mul(tls, bp+80, bp, bp+40)
	fe_mul_small(tls, bp+80, bp+80, -int32(2))
	is_square = invsqrt(tls, bp+80, bp+80) // t3 = sqrt(-1 / non_square * u * (u+A))
	if is_square != 0 {
		// The only variable time bit.  This ultimately reveals how many
		// tries it took us to find a representable key.
		// This does not affect security as long as we try keys at random.
		fe_ccopy(tls, bp, bp+40, libc.Int32FromUint8(tweak)&int32(1)) // multiply by u if v is positive,
		fe_mul(tls, bp+80, bp, bp+80)                                 // multiply by u+A otherwise
		fe_mul_small(tls, bp, bp+80, int32(2))
		fe_neg(tls, bp+40, bp+80)
		fe_ccopy(tls, bp+80, bp+40, fe_isodd(tls, bp))
		fe_tobytes(tls, hidden, bp+80)
		// Pad with two random bits
		v1 = hidden + 31
		*(*u8)(unsafe.Pointer(v1)) = u8(int32(*(*u8)(unsafe.Pointer(v1))) | libc.Int32FromUint8(tweak)&libc.Int32FromInt32(0xc0))
	}
	crypto_wipe(tls, bp, uint64(40))
	crypto_wipe(tls, bp+40, uint64(40))
	crypto_wipe(tls, bp+80, uint64(40))
	return is_square - int32(1)
}

func crypto_elligator_key_pair(tls *libc.TLS, hidden uintptr, secret_key uintptr, seed uintptr) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var _i_, _i_1, _i_2 size_t
	var _ /* buf at bp+32 */ [64]u8
	var _ /* pk at bp+0 */ [32]u8
	_, _, _ = _i_, _i_1, _i_2 // seed + representative
	_i_ = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(_i_ < libc.Uint64FromInt32(libc.Int32FromInt32(32))) {
			break
		}
		**(**u8)(__ccgo_up(bp + 32 + libc.UintptrFromInt32(32) + uintptr(_i_))) = **(**u8)(__ccgo_up(seed + uintptr(_i_)))
		goto _1
	_1:
		;
		_i_ = _i_ + 1
	}
	for cond := true; cond; cond = crypto_elligator_rev(tls, bp+32+uintptr(32), bp, (**(**[64]u8)(__ccgo_up(bp + 32)))[int32(32)]) != 0 {
		crypto_chacha20_djb(tls, bp+32, uintptr(0), uint64(64), bp+32+uintptr(32), uintptr(unsafe.Pointer(&zero)), uint64(0))
		crypto_x25519_dirty_fast(tls, bp, bp+32) // or the "small" version
	}
	// Note that the return value of crypto_elligator_rev() is
	// independent from its tweak parameter.
	// Therefore, buf[32] is not actually reused.  Either we loop one
	// more time and buf[32] is used for the new seed, or we succeeded,
	// and buf[32] becomes the tweak parameter.
	crypto_wipe(tls, seed, uint64(32))
	_i_1 = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(_i_1 < libc.Uint64FromInt32(libc.Int32FromInt32(32))) {
			break
		}
		**(**u8)(__ccgo_up(hidden + uintptr(_i_1))) = **(**u8)(__ccgo_up(bp + 32 + libc.UintptrFromInt32(32) + uintptr(_i_1)))
		goto _2
	_2:
		;
		_i_1 = _i_1 + 1
	}
	_i_2 = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(_i_2 < libc.Uint64FromInt32(libc.Int32FromInt32(32))) {
			break
		}
		**(**u8)(__ccgo_up(secret_key + uintptr(_i_2))) = (**(**[64]u8)(__ccgo_up(bp + 32)))[_i_2]
		goto _3
	_3:
		;
		_i_2 = _i_2 + 1
	}
	crypto_wipe(tls, bp+32, uint64(64))
	crypto_wipe(tls, bp, uint64(32))
}

///////////////////////
/// Scalar division ///
///////////////////////

// C documentation
//
//	// Montgomery reduction.
//	// Divides x by (2^256), and reduces the result modulo L
//	//
//	// Precondition:
//	//   x < L * 2^256
//	// Constants:
//	//   r = 2^256                 (makes division by r trivial)
//	//   k = (r * (1/r) - 1) // L  (1/r is computed modulo L   )
//	// Algorithm:
//	//   s = (x * k) % r
//	//   t = x + s*L      (t is always a multiple of r)
//	//   u = (t/r) % L    (u is always below 2*L, conditional subtraction is enough)
func redc(tls *libc.TLS, u uintptr, x uintptr) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var carry, carry1 u64
	var i, i1, j size_t
	var _ /* s at bp+0 */ [8]u32
	var _ /* t at bp+32 */ [16]u32
	_, _, _, _, _ = carry, carry1, i, i1, j
	// s = x * k (modulo 2^256)
	// This is cheaper than the full multiplication.
	**(**[8]u32)(__ccgo_up(bp)) = [8]u32{}
	i = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(i < libc.Uint64FromInt32(libc.Int32FromInt32(8))) {
			break
		}
		carry = uint64(0)
		j = libc.Uint64FromInt32(libc.Int32FromInt32(0))
		for {
			if !(j < uint64(8)-i) {
				break
			}
			carry = carry + (uint64((**(**[8]u32)(__ccgo_up(bp)))[i+j]) + uint64(**(**u32)(__ccgo_up(x + uintptr(i)*4)))*uint64(k[j]))
			(**(**[8]u32)(__ccgo_up(bp)))[i+j] = uint32(carry)
			carry = carry >> uint64(32)
			goto _2
		_2:
			;
			j = j + 1
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	**(**[16]u32)(__ccgo_up(bp + 32)) = [16]u32{}
	multiply(tls, bp+32, bp, uintptr(unsafe.Pointer(&L)))
	// t = t + x
	carry1 = uint64(0)
	i1 = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(i1 < libc.Uint64FromInt32(libc.Int32FromInt32(16))) {
			break
		}
		carry1 = carry1 + (uint64((**(**[16]u32)(__ccgo_up(bp + 32)))[i1]) + uint64(**(**u32)(__ccgo_up(x + uintptr(i1)*4))))
		(**(**[16]u32)(__ccgo_up(bp + 32)))[i1] = uint32(carry1)
		carry1 = carry1 >> uint64(32)
		goto _3
	_3:
		;
		i1 = i1 + 1
	}
	// u = (t / 2^256) % L
	// Note that t / 2^256 is always below 2*L,
	// So a constant time conditional subtraction is enough
	remove_l(tls, u, bp+32+uintptr(8)*4)
	crypto_wipe(tls, bp, uint64(32))
	crypto_wipe(tls, bp+32, uint64(64))
}

var k = [8]u32{
	0: uint32(0x12547e1b),
	1: uint32(0xd2b51da3),
	2: uint32(0xfdba84ff),
	3: uint32(0xb1a206f2),
	4: uint32(0xffa36bea),
	5: uint32(0x14e75438),
	6: uint32(0x6fe91836),
	7: uint32(0x9db6c6f2),
}

func crypto_x25519_inverse(tls *libc.TLS, blind_salt uintptr, private_key uintptr, curve_point uintptr) {
	bp := tls.Alloc(224)
	defer tls.Free(224)
	var _i_, _i_1, _i_2, _i_3, _i_4 size_t
	var i int32
	var _ /* m_inv at bp+0 */ [8]u32
	var _ /* m_scl at bp+64 */ [8]u32
	var _ /* product at bp+160 */ [16]u32
	var _ /* scalar at bp+32 */ [32]u8
	var _ /* tmp at bp+96 */ [16]u32
	_, _, _, _, _, _ = _i_, _i_1, _i_2, _i_3, _i_4, i
	// 1 in Montgomery form
	**(**[8]u32)(__ccgo_up(bp)) = [8]u32{
		0: uint32(0x8d98951d),
		1: uint32(0xd6ec3174),
		2: uint32(0x737dcf70),
		3: uint32(0xc6ef5bf4),
		4: uint32(0xfffffffe),
		5: uint32(0xffffffff),
		6: uint32(0xffffffff),
		7: uint32(0x0fffffff),
	}
	crypto_eddsa_trim_scalar(tls, bp+32, private_key)
	_i_ = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(_i_ < libc.Uint64FromInt32(libc.Int32FromInt32(8))) {
			break
		}
		(**(**[16]u32)(__ccgo_up(bp + 96)))[_i_] = uint32(0)
		goto _1
	_1:
		;
		_i_ = _i_ + 1
	}
	load32_le_buf(tls, bp+96+uintptr(8)*4, bp+32, uint64(8))
	mod_l(tls, bp+32, bp+96)
	load32_le_buf(tls, bp+64, bp+32, uint64(8))
	crypto_wipe(tls, bp+96, uint64(64)) // Wipe ASAP to save stack space
	i = int32(252)
	for {
		if !(i >= 0) {
			break
		}
		_i_1 = libc.Uint64FromInt32(libc.Int32FromInt32(0))
		for {
			if !(_i_1 < libc.Uint64FromInt32(libc.Int32FromInt32(16))) {
				break
			}
			(**(**[16]u32)(__ccgo_up(bp + 160)))[_i_1] = uint32(0)
			goto _3
		_3:
			;
			_i_1 = _i_1 + 1
		}
		multiply(tls, bp+160, bp, bp)
		redc(tls, bp, bp+160)
		if scalar_bit(tls, uintptr(unsafe.Pointer(&Lm2)), i) != 0 {
			_i_2 = libc.Uint64FromInt32(libc.Int32FromInt32(0))
			for {
				if !(_i_2 < libc.Uint64FromInt32(libc.Int32FromInt32(16))) {
					break
				}
				(**(**[16]u32)(__ccgo_up(bp + 160)))[_i_2] = uint32(0)
				goto _4
			_4:
				;
				_i_2 = _i_2 + 1
			}
			multiply(tls, bp+160, bp, bp+64)
			redc(tls, bp, bp+160)
		}
		goto _2
	_2:
		;
		i = i - 1
	}
	// Convert the inverse *out* of Montgomery form
	// scalar = m_inv / 2^256 (modulo L)
	_i_3 = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(_i_3 < libc.Uint64FromInt32(libc.Int32FromInt32(8))) {
			break
		}
		(**(**[16]u32)(__ccgo_up(bp + 160)))[_i_3] = (**(**[8]u32)(__ccgo_up(bp)))[_i_3]
		goto _5
	_5:
		;
		_i_3 = _i_3 + 1
	}
	_i_4 = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(_i_4 < libc.Uint64FromInt32(libc.Int32FromInt32(8))) {
			break
		}
		**(**u32)(__ccgo_up(bp + 160 + libc.UintptrFromInt32(8)*4 + uintptr(_i_4)*4)) = uint32(0)
		goto _6
	_6:
		;
		_i_4 = _i_4 + 1
	}
	redc(tls, bp, bp+160)
	store32_le_buf(tls, bp+32, bp, uint64(8)) // the *inverse* of the scalar
	// Clear the cofactor of scalar:
	//   cleared = scalar * (3*L + 1)      (modulo 8*L)
	//   cleared = scalar + scalar * 3 * L (modulo 8*L)
	// Note that (scalar * 3) is reduced modulo 8, so we only need the
	// first byte.
	add_xl(tls, bp+32, libc.Uint8FromInt32(libc.Int32FromUint8((**(**[32]u8)(__ccgo_up(bp + 32)))[0])*int32(3)))
	// Recall that 8*L < 2^256. However it is also very close to
	// 2^255. If we spanned the ladder over 255 bits, random tests
	// wouldn't catch the off-by-one error.
	scalarmult(tls, blind_salt, bp+32, curve_point, int32(256))
	crypto_wipe(tls, bp+32, uint64(32))
	crypto_wipe(tls, bp+64, uint64(32))
	crypto_wipe(tls, bp+160, uint64(64))
	crypto_wipe(tls, bp, uint64(32))
}

var Lm2 = [32]u8{
	0:  uint8(0xeb),
	1:  uint8(0xd3),
	2:  uint8(0xf5),
	3:  uint8(0x5c),
	4:  uint8(0x1a),
	5:  uint8(0x63),
	6:  uint8(0x12),
	7:  uint8(0x58),
	8:  uint8(0xd6),
	9:  uint8(0x9c),
	10: uint8(0xf7),
	11: uint8(0xa2),
	12: uint8(0xde),
	13: uint8(0xf9),
	14: uint8(0xde),
	15: uint8(0x14),
	31: uint8(0x10),
}

// C documentation
//
//	////////////////////////////////
//	/// Authenticated encryption ///
//	////////////////////////////////
func lock_auth(tls *libc.TLS, mac uintptr, auth_key uintptr, ad uintptr, ad_size size_t, cipher_text uintptr, text_size size_t) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var _ /* poly_ctx at bp+16 */ crypto_poly1305_ctx
	var _ /* sizes at bp+0 */ [16]u8 // Not secret, not wiped
	store64_le(tls, bp+uintptr(0), ad_size)
	store64_le(tls, bp+uintptr(8), text_size) // auto wiped...
	crypto_poly1305_init(tls, bp+16, auth_key)
	crypto_poly1305_update(tls, bp+16, ad, ad_size)
	crypto_poly1305_update(tls, bp+16, uintptr(unsafe.Pointer(&zero)), gap(tls, ad_size, uint64(16)))
	crypto_poly1305_update(tls, bp+16, cipher_text, text_size)
	crypto_poly1305_update(tls, bp+16, uintptr(unsafe.Pointer(&zero)), gap(tls, text_size, uint64(16)))
	crypto_poly1305_update(tls, bp+16, bp, uint64(16))
	crypto_poly1305_final(tls, bp+16, mac) // ...here
}

func crypto_aead_init_x(tls *libc.TLS, ctx uintptr, key uintptr, nonce uintptr) {
	var _i_ size_t
	_ = _i_
	crypto_chacha20_h(tls, ctx+8, key, nonce)
	_i_ = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(_i_ < libc.Uint64FromInt32(libc.Int32FromInt32(8))) {
			break
		}
		**(**uint8_t)(__ccgo_up(ctx + 40 + uintptr(_i_))) = **(**u8)(__ccgo_up(nonce + libc.UintptrFromInt32(16) + uintptr(_i_)))
		goto _1
	_1:
		;
		_i_ = _i_ + 1
	}
	(*crypto_aead_ctx)(unsafe.Pointer(ctx)).Fcounter = uint64(0)
}

func crypto_aead_init_djb(tls *libc.TLS, ctx uintptr, key uintptr, nonce uintptr) {
	var _i_, _i_1 size_t
	_, _ = _i_, _i_1
	_i_ = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(_i_ < libc.Uint64FromInt32(libc.Int32FromInt32(32))) {
			break
		}
		**(**uint8_t)(__ccgo_up(ctx + 8 + uintptr(_i_))) = **(**u8)(__ccgo_up(key + uintptr(_i_)))
		goto _1
	_1:
		;
		_i_ = _i_ + 1
	}
	_i_1 = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(_i_1 < libc.Uint64FromInt32(libc.Int32FromInt32(8))) {
			break
		}
		**(**uint8_t)(__ccgo_up(ctx + 40 + uintptr(_i_1))) = **(**u8)(__ccgo_up(nonce + uintptr(_i_1)))
		goto _2
	_2:
		;
		_i_1 = _i_1 + 1
	}
	(*crypto_aead_ctx)(unsafe.Pointer(ctx)).Fcounter = uint64(0)
}

func crypto_aead_init_ietf(tls *libc.TLS, ctx uintptr, key uintptr, nonce uintptr) {
	var _i_, _i_1 size_t
	_, _ = _i_, _i_1
	_i_ = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(_i_ < libc.Uint64FromInt32(libc.Int32FromInt32(32))) {
			break
		}
		**(**uint8_t)(__ccgo_up(ctx + 8 + uintptr(_i_))) = **(**u8)(__ccgo_up(key + uintptr(_i_)))
		goto _1
	_1:
		;
		_i_ = _i_ + 1
	}
	_i_1 = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(_i_1 < libc.Uint64FromInt32(libc.Int32FromInt32(8))) {
			break
		}
		**(**uint8_t)(__ccgo_up(ctx + 40 + uintptr(_i_1))) = **(**u8)(__ccgo_up(nonce + libc.UintptrFromInt32(4) + uintptr(_i_1)))
		goto _2
	_2:
		;
		_i_1 = _i_1 + 1
	}
	(*crypto_aead_ctx)(unsafe.Pointer(ctx)).Fcounter = uint64(load32_le(tls, nonce)) << int32(32)
}

func crypto_aead_write(tls *libc.TLS, ctx uintptr, cipher_text uintptr, mac uintptr, ad uintptr, ad_size size_t, plain_text uintptr, text_size size_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var _i_ size_t
	var _ /* auth_key at bp+0 */ [64]u8
	_ = _i_ // the last 32 bytes are used for rekeying.
	crypto_chacha20_djb(tls, bp, uintptr(0), uint64(64), ctx+8, ctx+40, (*crypto_aead_ctx)(unsafe.Pointer(ctx)).Fcounter)
	crypto_chacha20_djb(tls, cipher_text, plain_text, text_size, ctx+8, ctx+40, (*crypto_aead_ctx)(unsafe.Pointer(ctx)).Fcounter+uint64(1))
	lock_auth(tls, mac, bp, ad, ad_size, cipher_text, text_size)
	_i_ = libc.Uint64FromInt32(libc.Int32FromInt32(0))
	for {
		if !(_i_ < libc.Uint64FromInt32(libc.Int32FromInt32(32))) {
			break
		}
		**(**uint8_t)(__ccgo_up(ctx + 8 + uintptr(_i_))) = **(**u8)(__ccgo_up(bp + libc.UintptrFromInt32(32) + uintptr(_i_)))
		goto _1
	_1:
		;
		_i_ = _i_ + 1
	}
	crypto_wipe(tls, bp, uint64(64))
}

func crypto_aead_read(tls *libc.TLS, ctx uintptr, plain_text uintptr, mac uintptr, ad uintptr, ad_size size_t, cipher_text uintptr, text_size size_t) (r int32) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var _i_ size_t
	var mismatch int32
	var _ /* auth_key at bp+0 */ [64]u8
	var _ /* real_mac at bp+64 */ [16]u8
	_, _ = _i_, mismatch
	crypto_chacha20_djb(tls, bp, uintptr(0), uint64(64), ctx+8, ctx+40, (*crypto_aead_ctx)(unsafe.Pointer(ctx)).Fcounter)
	lock_auth(tls, bp+64, bp, ad, ad_size, cipher_text, text_size)
	mismatch = crypto_verify16(tls, mac, bp+64)
	if !(mismatch != 0) {
		crypto_chacha20_djb(tls, plain_text, cipher_text, text_size, ctx+8, ctx+40, (*crypto_aead_ctx)(unsafe.Pointer(ctx)).Fcounter+uint64(1))
		_i_ = libc.Uint64FromInt32(libc.Int32FromInt32(0))
		for {
			if !(_i_ < libc.Uint64FromInt32(libc.Int32FromInt32(32))) {
				break
			}
			**(**uint8_t)(__ccgo_up(ctx + 8 + uintptr(_i_))) = **(**u8)(__ccgo_up(bp + libc.UintptrFromInt32(32) + uintptr(_i_)))
			goto _1
		_1:
			;
			_i_ = _i_ + 1
		}
	}
	crypto_wipe(tls, bp, uint64(64))
	crypto_wipe(tls, bp+64, uint64(16))
	return mismatch
}

func crypto_aead_lock(tls *libc.TLS, cipher_text uintptr, mac uintptr, key uintptr, nonce uintptr, ad uintptr, ad_size size_t, plain_text uintptr, text_size size_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var _ /* ctx at bp+0 */ crypto_aead_ctx
	crypto_aead_init_x(tls, bp, key, nonce)
	crypto_aead_write(tls, bp, cipher_text, mac, ad, ad_size, plain_text, text_size)
	crypto_wipe(tls, bp, uint64(48))
}

func crypto_aead_unlock(tls *libc.TLS, plain_text uintptr, mac uintptr, key uintptr, nonce uintptr, ad uintptr, ad_size size_t, cipher_text uintptr, text_size size_t) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var mismatch int32
	var _ /* ctx at bp+0 */ crypto_aead_ctx
	_ = mismatch
	crypto_aead_init_x(tls, bp, key, nonce)
	mismatch = crypto_aead_read(tls, bp, plain_text, mac, ad, ad_size, cipher_text, text_size)
	crypto_wipe(tls, bp, uint64(48))
	return mismatch
}

func __ccgo_up(n uintptr) unsafe.Pointer {
	return unsafe.Pointer(&n)
}

var crypto_argon2_no_extras = crypto_argon2_extras{}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "expand 32-byte k\x00"
