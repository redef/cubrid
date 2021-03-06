package gci

import "unsafe"

type AUTOCOMMIT_MODE int

const (
	AUTOCOMMIT_FALSE AUTOCOMMIT_MODE = iota
	AUTOCOMMIT_TRUE
)

const (
	TRAN_COMMIT   = 1
	TRAN_ROLLBACK = 2
)

const (
	GCI_CODE_SET        = 0x20
	GCI_CODE_MULTISET   = 0x40
	GCI_CODE_SEQUENCE   = 0x60
	GCI_CODE_COLLECTION = 0x60
)

type GCI_A_TYPE int

const (
	_                     = iota
	A_TYPE_STR GCI_A_TYPE = iota
	A_TYPE_INT
	A_TYPE_FLOAT
	A_TYPE_DOUBLE
	A_TYPE_BIT
	A_TYPE_DATE
	A_TYPE_SET
	A_TYPE_BIGINT
	A_TYPE_BLOB
	A_TYPE_CLOB
	A_TYPE_LAST
)

type GCI_ERROR_CODE int

const (
	GCI_ER_NO_ERROR        GCI_ERROR_CODE = 0
	GCI_ER_DBMS                           = -20001
	GCI_ER_CON_HANDLE                     = -20002
	GCI_ER_NO_MORE_MEMORY                 = -20003
	GCI_ER_COMMUNICATION                  = -20004
	GCI_ER_NO_MORE_DATA                   = -20005
	GCI_ER_TRAN_TYPE                      = -20006
	GCI_ER_STRING_PARAM                   = -20007
	GCI_ER_TYPE_CONVERSION                = -20008
	GCI_ER_BIND_INDEX                     = -20009
	GCI_ER_ATYPE                          = -20010
	GCI_ER_NOT_BIND                       = -20011
	GCI_ER_PARAM_NAME                     = -20012
	GCI_ER_COLUMN_INDEX                   = -20013
	GCI_ER_SCHEMA_TYPE                    = -20014
	GCI_ER_FILE                           = -20015
	GCI_ER_CONNECT                        = -20016

	GCI_ER_ALLOC_CON_HANDLE   = -20017
	GCI_ER_REQ_HANDLE         = -20018
	GCI_ER_INVALID_CURSOR_POS = -20019
	GCI_ER_OBJECT             = -20020
	GCI_ER_CAS                = -20021
	GCI_ER_HOSTNAME           = -20022
	GCI_ER_OID_CMD            = -20023

	GCI_ER_BIND_ARRAY_SIZE = -20024
	GCI_ER_ISOLATION_LEVEL = -20025

	GCI_ER_SET_INDEX     = -20026
	GCI_ER_DELETED_TUPLE = -20027

	GCI_ER_SAVEPOINT_CMD        = -20028
	GCI_ER_THREAD_RUNNING       = -20029
	GCI_ER_INVALID_URL          = -20030
	GCI_ER_INVALID_LOB_READ_POS = -20031
	GCI_ER_INVALID_LOB_HANDLE   = -20032

	GCI_ER_NO_PROPERY           = -20033
	GCI_ER_PROPERTY_TYPE        = -20034
	GCI_ER_INVALID_DATASOURCE   = -20035
	GCI_ER_DATASOURCE_TIMEOUT   = -20036
	GCI_ER_DATASOURCE_TIMEDWAIT = -20037

	GCI_ER_LOGIN_TIMEOUT = -20038
	GCI_ER_QUERY_TIMEOUT = -20039

	GCI_ER_RESULT_SET_CLOSED = -20040

	GCI_ER_INVALID_HOLDABILITY = -20041
	GCI_ER_NOT_UPDATABLE       = -20042

	GCI_ER_INVALID_ARGS    = -20043
	GCI_ER_USED_CONNECTION = -20044

	GCI_ER_NOT_IMPLEMENTED = -20099
	GCI_ER_END             = -20100
)

type GCI_CURSOR_POS int

const (
	GCI_CURSOR_FIRST GCI_CURSOR_POS = iota
	GCI_CURSOR_CURRENT
	GCI_CURSOR_LAST
)

type GCI_U_TYPE int

const (
	U_TYPE_FIRST   GCI_U_TYPE = 0
	U_TYPE_UNKNOWN GCI_U_TYPE = 0
	U_TYPE_NULL    GCI_U_TYPE = 0

	U_TYPE_CHAR      GCI_U_TYPE = 1
	U_TYPE_STRING    GCI_U_TYPE = 2
	U_TYPE_NCHAR     GCI_U_TYPE = 3
	U_TYPE_VARCHAR   GCI_U_TYPE = 4
	U_TYPE_BIT       GCI_U_TYPE = 5
	U_TYPE_VARBIT    GCI_U_TYPE = 6
	U_TYPE_NUMERIC   GCI_U_TYPE = 7
	U_TYPE_INT       GCI_U_TYPE = 8
	U_TYPE_SHORT     GCI_U_TYPE = 9
	U_TYPE_MONETARY  GCI_U_TYPE = 10
	U_TYPE_FLOAT     GCI_U_TYPE = 11
	U_TYPE_DOUBLE    GCI_U_TYPE = 12
	U_TYPE_DATE      GCI_U_TYPE = 13
	U_TYPE_TIME      GCI_U_TYPE = 14
	U_TYPE_TIMESTAMP GCI_U_TYPE = 15
	U_TYPE_SET       GCI_U_TYPE = 16
	U_TYPE_MULTISET  GCI_U_TYPE = 17
	U_TYPE_SEQUENCE  GCI_U_TYPE = 18
	U_TYPE_OBJECT    GCI_U_TYPE = 19
	U_TYPE_RESULTSET GCI_U_TYPE = 20
	U_TYPE_BIGINT    GCI_U_TYPE = 21
	U_TYPE_DATETIME  GCI_U_TYPE = 22
	U_TYPE_BLOB      GCI_U_TYPE = 23
	U_TYPE_CLOB      GCI_U_TYPE = 24
	U_TYPE_ENUM      GCI_U_TYPE = 25

	U_TYPE_LAST GCI_U_TYPE = U_TYPE_ENUM
)

const GCI_BIND_PTR int = 1

type GCI_ERROR struct {
	Code int
	Msg  string
}

type GCI_COL_INFO struct {
	u_type            GCI_U_TYPE
	is_non_null       string
	scale             int16
	precision         int
	col_name          string
	real_attr         string
	class_name        string
	default_value     string
	is_auto_increment bool
	is_unique_key     bool
	is_primary_key    bool
	is_foreign_key    bool
	is_reverse_index  bool
	is_reverse_unique bool
	is_shared         bool
}

type GCI_CUBRID_STMT int

const (
	CUBRID_STMT_ALTER_CLASS GCI_CUBRID_STMT = iota
	CUBRID_STMT_ALTER_SERIAL
	CUBRID_STMT_COMMIT_WORK
	CUBRID_STMT_REGISTER_DATABASE
	CUBRID_STMT_CREATE_CLASS
	CUBRID_STMT_CREATE_INDEX
	CUBRID_STMT_CREATE_TRIGGER
	CUBRID_STMT_CREATE_SERIAL
	CUBRID_STMT_DROP_DATABASE
	CUBRID_STMT_DROP_CLASS
	CUBRID_STMT_DROP_INDEX
	CUBRID_STMT_DROP_LABEL
	CUBRID_STMT_DROP_TRIGGER
	CUBRID_STMT_DROP_SERIAL
	CUBRID_STMT_EVALUATE
	CUBRID_STMT_RENAME_CLASS
	CUBRID_STMT_ROLLBACK_WORK
	CUBRID_STMT_GRANT
	CUBRID_STMT_REVOKE
	CUBRID_STMT_STATISTICS
	CUBRID_STMT_INSERT
	CUBRID_STMT_SELECT
	CUBRID_STMT_UPDATE
	CUBRID_STMT_DELETE
	CUBRID_STMT_CALL
	CUBRID_STMT_GET_ISO_LVL
	CUBRID_STMT_GET_TIMEOUT
	CUBRID_STMT_GET_OPT_LVL
	CUBRID_STMT_SET_OPT_LVL
	CUBRID_STMT_SCOPE
	CUBRID_STMT_GET_TRIGGER
	CUBRID_STMT_SET_TRIGGER
	CUBRID_STMT_SAVEPOINT
	CUBRID_STMT_PREPARE
	CUBRID_STMT_ATTACH
	CUBRID_STMT_USE
	CUBRID_STMT_REMOVE_TRIGGER
	CUBRID_STMT_RENAME_TRIGGER
	CUBRID_STMT_ON_LDB
	CUBRID_STMT_GET_LDB
	CUBRID_STMT_SET_LDB
	CUBRID_STMT_GET_STATS
	CUBRID_STMT_CREATE_USER
	CUBRID_STMT_DROP_USER
	CUBRID_STMT_ALTER_USER
)

type GCI_DATE struct {
	yr  int
	mon int
	day int
	hh  int
	mm  int
	ss  int
	ms  int
}

/*
 cubrid manual
 - 비트열은 0과 1로 이루어진 이진 값의 순열(sequence) 이다.
 - 2진수 형식으로 사용할 때에는 다음과 같이 문자 B뒤에 0과 1로 이루어진 문자열을 붙이거나,
   0b 뒤에 값을 붙여 표현한다.
   ex) B'1010'
       0b1010
 - 16진수 형식은 대문자 X 뒤에 0-9 그리고 A-F 문자로 이루어진 문자열을 붙이거나
   0x 뒤에 값을 붙여 표현한다.
   ex) X'a'
       0xA
*/
type GCI_BIT struct {
	size int
	buf  []byte
}

type GCI_SET unsafe.Pointer

type GCI_BLOB unsafe.Pointer

type GCI_CLOB unsafe.Pointer

func (date *GCI_DATE) Yr() uint {
	return uint(date.yr)
}

func (date *GCI_DATE) Mon() uint {
	return uint(date.mon)
}

func (date *GCI_DATE) Day() uint {
	return uint(date.day)
}

func (date *GCI_DATE) Hh() uint {
	return uint(date.hh)
}

func (date *GCI_DATE) Mm() uint {
	return uint(date.mm)
}

func (date *GCI_DATE) Ss() uint {
	return uint(date.ss)
}

func (date *GCI_DATE) Ms() uint {
	return uint(date.ms)
}

func (bit *GCI_BIT) Size() int {
	return int(bit.size)
}

func (bit *GCI_BIT) Buf() string {
	return string(bit.buf)
}
