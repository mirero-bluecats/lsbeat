// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by mk_audit_exit_codes - DO NOT EDIT.

package auparse

var AuditErrnoToNum = map[string]int{
	"EDEADLOCK":       0x23,
	"EWOULDBLOCK":     0xb,
	"EPERM":           0x1,
	"ENOENT":          0x2,
	"ESRCH":           0x3,
	"EINTR":           0x4,
	"EIO":             0x5,
	"ENXIO":           0x6,
	"E2BIG":           0x7,
	"ENOEXEC":         0x8,
	"EBADF":           0x9,
	"ECHILD":          0xa,
	"EAGAIN":          0xb,
	"ENOMEM":          0xc,
	"EACCES":          0xd,
	"EFAULT":          0xe,
	"ENOTBLK":         0xf,
	"EBUSY":           0x10,
	"EEXIST":          0x11,
	"EXDEV":           0x12,
	"ENODEV":          0x13,
	"ENOTDIR":         0x14,
	"EISDIR":          0x15,
	"EINVAL":          0x16,
	"ENFILE":          0x17,
	"EMFILE":          0x18,
	"ENOTTY":          0x19,
	"ETXTBSY":         0x1a,
	"EFBIG":           0x1b,
	"ENOSPC":          0x1c,
	"ESPIPE":          0x1d,
	"EROFS":           0x1e,
	"EMLINK":          0x1f,
	"EPIPE":           0x20,
	"EDOM":            0x21,
	"ERANGE":          0x22,
	"EDEADLK":         0x23,
	"ENAMETOOLONG":    0x24,
	"ENOLCK":          0x25,
	"ENOSYS":          0x26,
	"ENOTEMPTY":       0x27,
	"ELOOP":           0x28,
	"ENOMSG":          0x2a,
	"EIDRM":           0x2b,
	"ECHRNG":          0x2c,
	"EL2NSYNC":        0x2d,
	"EL3HLT":          0x2e,
	"EL3RST":          0x2f,
	"ELNRNG":          0x30,
	"EUNATCH":         0x31,
	"ENOCSI":          0x32,
	"EL2HLT":          0x33,
	"EBADE":           0x34,
	"EBADR":           0x35,
	"EXFULL":          0x36,
	"ENOANO":          0x37,
	"EBADRQC":         0x38,
	"EBADSLT":         0x39,
	"EBFONT":          0x3b,
	"ENOSTR":          0x3c,
	"ENODATA":         0x3d,
	"ETIME":           0x3e,
	"ENOSR":           0x3f,
	"ENONET":          0x40,
	"ENOPKG":          0x41,
	"EREMOTE":         0x42,
	"ENOLINK":         0x43,
	"EADV":            0x44,
	"ESRMNT":          0x45,
	"ECOMM":           0x46,
	"EPROTO":          0x47,
	"EMULTIHOP":       0x48,
	"EDOTDOT":         0x49,
	"EBADMSG":         0x4a,
	"EOVERFLOW":       0x4b,
	"ENOTUNIQ":        0x4c,
	"EBADFD":          0x4d,
	"EREMCHG":         0x4e,
	"ELIBACC":         0x4f,
	"ELIBBAD":         0x50,
	"ELIBSCN":         0x51,
	"ELIBMAX":         0x52,
	"ELIBEXEC":        0x53,
	"EILSEQ":          0x54,
	"ERESTART":        0x55,
	"ESTRPIPE":        0x56,
	"EUSERS":          0x57,
	"ENOTSOCK":        0x58,
	"EDESTADDRREQ":    0x59,
	"EMSGSIZE":        0x5a,
	"EPROTOTYPE":      0x5b,
	"ENOPROTOOPT":     0x5c,
	"EPROTONOSUPPORT": 0x5d,
	"ESOCKTNOSUPPORT": 0x5e,
	"EOPNOTSUPP":      0x5f,
	"EPFNOSUPPORT":    0x60,
	"EAFNOSUPPORT":    0x61,
	"EADDRINUSE":      0x62,
	"EADDRNOTAVAIL":   0x63,
	"ENETDOWN":        0x64,
	"ENETUNREACH":     0x65,
	"ENETRESET":       0x66,
	"ECONNABORTED":    0x67,
	"ECONNRESET":      0x68,
	"ENOBUFS":         0x69,
	"EISCONN":         0x6a,
	"ENOTCONN":        0x6b,
	"ESHUTDOWN":       0x6c,
	"ETOOMANYREFS":    0x6d,
	"ETIMEDOUT":       0x6e,
	"ECONNREFUSED":    0x6f,
	"EHOSTDOWN":       0x70,
	"EHOSTUNREACH":    0x71,
	"EALREADY":        0x72,
	"EINPROGRESS":     0x73,
	"ESTALE":          0x74,
	"EUCLEAN":         0x75,
	"ENOTNAM":         0x76,
	"ENAVAIL":         0x77,
	"EISNAM":          0x78,
	"EREMOTEIO":       0x79,
	"EDQUOT":          0x7a,
	"ENOMEDIUM":       0x7b,
	"EMEDIUMTYPE":     0x7c,
	"ECANCELED":       0x7d,
	"ENOKEY":          0x7e,
	"EKEYEXPIRED":     0x7f,
	"EKEYREVOKED":     0x80,
	"EKEYREJECTED":    0x81,
	"EOWNERDEAD":      0x82,
	"ENOTRECOVERABLE": 0x83,
	"ERFKILL":         0x84,
	"EHWPOISON":       0x85,
}

var AuditErrnoToName = map[int]string{
	1:   "EPERM",
	2:   "ENOENT",
	3:   "ESRCH",
	4:   "EINTR",
	5:   "EIO",
	6:   "ENXIO",
	7:   "E2BIG",
	8:   "ENOEXEC",
	9:   "EBADF",
	10:  "ECHILD",
	11:  "EAGAIN",
	12:  "ENOMEM",
	13:  "EACCES",
	14:  "EFAULT",
	15:  "ENOTBLK",
	16:  "EBUSY",
	17:  "EEXIST",
	18:  "EXDEV",
	19:  "ENODEV",
	20:  "ENOTDIR",
	21:  "EISDIR",
	22:  "EINVAL",
	23:  "ENFILE",
	24:  "EMFILE",
	25:  "ENOTTY",
	26:  "ETXTBSY",
	27:  "EFBIG",
	28:  "ENOSPC",
	29:  "ESPIPE",
	30:  "EROFS",
	31:  "EMLINK",
	32:  "EPIPE",
	33:  "EDOM",
	34:  "ERANGE",
	35:  "EDEADLK",
	36:  "ENAMETOOLONG",
	37:  "ENOLCK",
	38:  "ENOSYS",
	39:  "ENOTEMPTY",
	40:  "ELOOP",
	42:  "ENOMSG",
	43:  "EIDRM",
	44:  "ECHRNG",
	45:  "EL2NSYNC",
	46:  "EL3HLT",
	47:  "EL3RST",
	48:  "ELNRNG",
	49:  "EUNATCH",
	50:  "ENOCSI",
	51:  "EL2HLT",
	52:  "EBADE",
	53:  "EBADR",
	54:  "EXFULL",
	55:  "ENOANO",
	56:  "EBADRQC",
	57:  "EBADSLT",
	59:  "EBFONT",
	60:  "ENOSTR",
	61:  "ENODATA",
	62:  "ETIME",
	63:  "ENOSR",
	64:  "ENONET",
	65:  "ENOPKG",
	66:  "EREMOTE",
	67:  "ENOLINK",
	68:  "EADV",
	69:  "ESRMNT",
	70:  "ECOMM",
	71:  "EPROTO",
	72:  "EMULTIHOP",
	73:  "EDOTDOT",
	74:  "EBADMSG",
	75:  "EOVERFLOW",
	76:  "ENOTUNIQ",
	77:  "EBADFD",
	78:  "EREMCHG",
	79:  "ELIBACC",
	80:  "ELIBBAD",
	81:  "ELIBSCN",
	82:  "ELIBMAX",
	83:  "ELIBEXEC",
	84:  "EILSEQ",
	85:  "ERESTART",
	86:  "ESTRPIPE",
	87:  "EUSERS",
	88:  "ENOTSOCK",
	89:  "EDESTADDRREQ",
	90:  "EMSGSIZE",
	91:  "EPROTOTYPE",
	92:  "ENOPROTOOPT",
	93:  "EPROTONOSUPPORT",
	94:  "ESOCKTNOSUPPORT",
	95:  "EOPNOTSUPP",
	96:  "EPFNOSUPPORT",
	97:  "EAFNOSUPPORT",
	98:  "EADDRINUSE",
	99:  "EADDRNOTAVAIL",
	100: "ENETDOWN",
	101: "ENETUNREACH",
	102: "ENETRESET",
	103: "ECONNABORTED",
	104: "ECONNRESET",
	105: "ENOBUFS",
	106: "EISCONN",
	107: "ENOTCONN",
	108: "ESHUTDOWN",
	109: "ETOOMANYREFS",
	110: "ETIMEDOUT",
	111: "ECONNREFUSED",
	112: "EHOSTDOWN",
	113: "EHOSTUNREACH",
	114: "EALREADY",
	115: "EINPROGRESS",
	116: "ESTALE",
	117: "EUCLEAN",
	118: "ENOTNAM",
	119: "ENAVAIL",
	120: "EISNAM",
	121: "EREMOTEIO",
	122: "EDQUOT",
	123: "ENOMEDIUM",
	124: "EMEDIUMTYPE",
	125: "ECANCELED",
	126: "ENOKEY",
	127: "EKEYEXPIRED",
	128: "EKEYREVOKED",
	129: "EKEYREJECTED",
	130: "EOWNERDEAD",
	131: "ENOTRECOVERABLE",
	132: "ERFKILL",
	133: "EHWPOISON",
}
