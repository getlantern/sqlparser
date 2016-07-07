//line sql.y:6
package sqlparser

import __yyfmt__ "fmt"

//line sql.y:6
import "bytes"

func SetParseTree(yylex interface{}, stmt Statement) {
	yylex.(*Tokenizer).ParseTree = stmt
}

func SetAllowComments(yylex interface{}, allow bool) {
	yylex.(*Tokenizer).AllowComments = allow
}

func ForceEOF(yylex interface{}) {
	yylex.(*Tokenizer).ForceEOF = true
}

var (
	SHARE        = []byte("share")
	MODE         = []byte("mode")
	IF_BYTES     = []byte("if")
	VALUES_BYTES = []byte("values")
)

//line sql.y:31
type yySymType struct {
	yys         int
	empty       struct{}
	statement   Statement
	selStmt     SelectStatement
	byt         byte
	bytes       []byte
	bytes2      [][]byte
	str         string
	selectExprs SelectExprs
	selectExpr  SelectExpr
	columns     Columns
	colName     *ColName
	tableExprs  TableExprs
	tableExpr   TableExpr
	smTableExpr SimpleTableExpr
	tableName   *TableName
	indexHints  *IndexHints
	expr        Expr
	boolExpr    BoolExpr
	valExpr     ValExpr
	colTuple    ColTuple
	valExprs    ValExprs
	values      Values
	rowTuple    RowTuple
	subquery    *Subquery
	caseExpr    *CaseExpr
	whens       []*When
	when        *When
	orderBy     OrderBy
	order       *Order
	timerange   *TimeRange
	limit       *Limit
	insRows     InsertRows
	updateExprs UpdateExprs
	updateExpr  *UpdateExpr

	/*
	   for CreateTable
	*/
	createTableStmt   CreateTable
	columnDefinition  *ColumnDefinition
	columnDefinitions ColumnDefinitions
	columnAtts        ColumnAtts
}

const LEX_ERROR = 57346
const SELECT = 57347
const INSERT = 57348
const UPDATE = 57349
const DELETE = 57350
const FROM = 57351
const ASOF = 57352
const UNTIL = 57353
const WHERE = 57354
const GROUP = 57355
const HAVING = 57356
const ORDER = 57357
const BY = 57358
const LIMIT = 57359
const FOR = 57360
const ALL = 57361
const DISTINCT = 57362
const AS = 57363
const EXISTS = 57364
const IN = 57365
const IS = 57366
const LIKE = 57367
const BETWEEN = 57368
const NULL = 57369
const ASC = 57370
const DESC = 57371
const VALUES = 57372
const INTO = 57373
const DUPLICATE = 57374
const KEY = 57375
const DEFAULT = 57376
const SET = 57377
const LOCK = 57378
const ID = 57379
const STRING = 57380
const NUMBER = 57381
const VALUE_ARG = 57382
const LIST_ARG = 57383
const COMMENT = 57384
const LE = 57385
const GE = 57386
const NE = 57387
const NULL_SAFE_EQUAL = 57388
const PRIMARY = 57389
const UNIQUE = 57390
const UNION = 57391
const MINUS = 57392
const EXCEPT = 57393
const INTERSECT = 57394
const JOIN = 57395
const STRAIGHT_JOIN = 57396
const LEFT = 57397
const RIGHT = 57398
const INNER = 57399
const OUTER = 57400
const CROSS = 57401
const NATURAL = 57402
const USE = 57403
const FORCE = 57404
const ON = 57405
const OR = 57406
const AND = 57407
const NOT = 57408
const UNARY = 57409
const CASE = 57410
const WHEN = 57411
const THEN = 57412
const ELSE = 57413
const END = 57414
const CREATE = 57415
const ALTER = 57416
const DROP = 57417
const RENAME = 57418
const ANALYZE = 57419
const TABLE = 57420
const INDEX = 57421
const VIEW = 57422
const TO = 57423
const IGNORE = 57424
const IF = 57425
const USING = 57426
const SHOW = 57427
const DESCRIBE = 57428
const EXPLAIN = 57429
const BIT = 57430
const TINYINT = 57431
const SMALLINT = 57432
const MEDIUMINT = 57433
const INT = 57434
const INTEGER = 57435
const BIGINT = 57436
const REAL = 57437
const DOUBLE = 57438
const FLOAT = 57439
const UNSIGNED = 57440
const ZEROFILL = 57441
const DECIMAL = 57442
const NUMERIC = 57443
const DATE = 57444
const TIME = 57445
const TIMESTAMP = 57446
const DATETIME = 57447
const YEAR = 57448
const TEXT = 57449
const CHAR = 57450
const VARCHAR = 57451
const NULLX = 57452
const AUTO_INCREMENT = 57453
const BOOL = 57454
const APPROXNUM = 57455
const INTNUM = 57456

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LEX_ERROR",
	"SELECT",
	"INSERT",
	"UPDATE",
	"DELETE",
	"FROM",
	"ASOF",
	"UNTIL",
	"WHERE",
	"GROUP",
	"HAVING",
	"ORDER",
	"BY",
	"LIMIT",
	"FOR",
	"ALL",
	"DISTINCT",
	"AS",
	"EXISTS",
	"IN",
	"IS",
	"LIKE",
	"BETWEEN",
	"NULL",
	"ASC",
	"DESC",
	"VALUES",
	"INTO",
	"DUPLICATE",
	"KEY",
	"DEFAULT",
	"SET",
	"LOCK",
	"ID",
	"STRING",
	"NUMBER",
	"VALUE_ARG",
	"LIST_ARG",
	"COMMENT",
	"LE",
	"GE",
	"NE",
	"NULL_SAFE_EQUAL",
	"'('",
	"'='",
	"'<'",
	"'>'",
	"'~'",
	"PRIMARY",
	"UNIQUE",
	"UNION",
	"MINUS",
	"EXCEPT",
	"INTERSECT",
	"','",
	"JOIN",
	"STRAIGHT_JOIN",
	"LEFT",
	"RIGHT",
	"INNER",
	"OUTER",
	"CROSS",
	"NATURAL",
	"USE",
	"FORCE",
	"ON",
	"OR",
	"AND",
	"NOT",
	"'&'",
	"'|'",
	"'^'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'.'",
	"UNARY",
	"CASE",
	"WHEN",
	"THEN",
	"ELSE",
	"END",
	"CREATE",
	"ALTER",
	"DROP",
	"RENAME",
	"ANALYZE",
	"TABLE",
	"INDEX",
	"VIEW",
	"TO",
	"IGNORE",
	"IF",
	"USING",
	"SHOW",
	"DESCRIBE",
	"EXPLAIN",
	"BIT",
	"TINYINT",
	"SMALLINT",
	"MEDIUMINT",
	"INT",
	"INTEGER",
	"BIGINT",
	"REAL",
	"DOUBLE",
	"FLOAT",
	"UNSIGNED",
	"ZEROFILL",
	"DECIMAL",
	"NUMERIC",
	"DATE",
	"TIME",
	"TIMESTAMP",
	"DATETIME",
	"YEAR",
	"TEXT",
	"CHAR",
	"VARCHAR",
	"NULLX",
	"AUTO_INCREMENT",
	"BOOL",
	"APPROXNUM",
	"INTNUM",
	"')'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 257
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 658

var yyAct = [...]int{

	95, 295, 159, 433, 92, 359, 93, 51, 63, 162,
	250, 197, 369, 246, 365, 103, 237, 177, 287, 86,
	81, 208, 161, 3, 261, 262, 263, 264, 265, 441,
	266, 267, 136, 135, 52, 53, 441, 444, 428, 66,
	82, 408, 71, 65, 64, 74, 364, 340, 342, 78,
	441, 54, 29, 30, 31, 32, 185, 429, 398, 87,
	130, 77, 69, 256, 405, 399, 229, 298, 124, 293,
	130, 120, 44, 387, 45, 386, 385, 341, 70, 121,
	128, 73, 123, 404, 406, 132, 130, 229, 47, 48,
	49, 350, 227, 163, 50, 344, 46, 164, 238, 238,
	285, 443, 272, 397, 148, 149, 150, 119, 442, 158,
	160, 134, 171, 66, 168, 117, 66, 65, 181, 180,
	65, 175, 440, 42, 113, 72, 135, 217, 228, 428,
	136, 135, 349, 87, 203, 181, 195, 382, 346, 299,
	207, 292, 282, 215, 216, 201, 219, 220, 221, 222,
	223, 224, 225, 226, 210, 205, 206, 400, 280, 230,
	179, 136, 135, 39, 288, 41, 204, 252, 231, 87,
	87, 202, 218, 115, 66, 66, 352, 127, 65, 244,
	334, 384, 242, 383, 288, 335, 253, 233, 235, 60,
	338, 191, 332, 245, 254, 241, 337, 333, 248, 146,
	147, 148, 149, 150, 336, 392, 14, 15, 16, 17,
	189, 271, 231, 192, 259, 201, 275, 276, 178, 273,
	143, 144, 145, 146, 147, 148, 149, 150, 210, 129,
	115, 229, 279, 429, 274, 393, 18, 87, 143, 144,
	145, 146, 147, 148, 149, 150, 354, 290, 116, 111,
	200, 211, 114, 284, 76, 281, 173, 209, 286, 294,
	199, 291, 258, 14, 115, 188, 190, 187, 174, 419,
	329, 418, 331, 417, 201, 328, 201, 165, 130, 375,
	348, 29, 30, 31, 32, 370, 366, 182, 351, 20,
	21, 23, 22, 24, 66, 200, 356, 169, 355, 357,
	360, 25, 26, 27, 79, 199, 167, 166, 110, 361,
	438, 314, 315, 316, 317, 318, 319, 320, 321, 322,
	323, 367, 368, 324, 325, 309, 310, 311, 312, 313,
	308, 306, 307, 411, 378, 371, 372, 373, 376, 374,
	377, 395, 396, 424, 425, 410, 409, 415, 330, 234,
	388, 98, 72, 67, 270, 389, 102, 251, 345, 108,
	343, 391, 327, 133, 446, 326, 85, 99, 100, 101,
	269, 261, 262, 263, 264, 265, 90, 266, 267, 72,
	106, 194, 447, 427, 193, 176, 143, 144, 145, 146,
	147, 148, 149, 150, 125, 421, 360, 122, 118, 422,
	61, 89, 416, 80, 75, 104, 105, 83, 112, 426,
	390, 353, 109, 14, 59, 423, 278, 231, 434, 434,
	434, 66, 435, 436, 432, 65, 430, 107, 450, 437,
	57, 98, 439, 431, 183, 126, 102, 296, 240, 108,
	55, 212, 448, 213, 214, 33, 85, 99, 100, 101,
	414, 451, 452, 297, 247, 413, 90, 380, 178, 232,
	106, 35, 36, 37, 38, 347, 14, 143, 144, 145,
	146, 147, 148, 149, 150, 381, 62, 449, 420, 14,
	34, 89, 403, 98, 402, 104, 105, 83, 102, 362,
	303, 108, 109, 305, 14, 304, 401, 407, 67, 99,
	100, 101, 363, 301, 302, 19, 98, 107, 90, 249,
	300, 102, 106, 184, 108, 40, 102, 255, 186, 108,
	43, 67, 99, 100, 101, 68, 67, 99, 100, 101,
	243, 90, 172, 89, 445, 106, 165, 104, 105, 394,
	106, 358, 412, 277, 109, 143, 144, 145, 146, 147,
	148, 149, 150, 379, 283, 170, 89, 236, 97, 107,
	104, 105, 94, 96, 289, 104, 105, 109, 102, 91,
	239, 108, 109, 137, 88, 257, 339, 198, 67, 99,
	100, 101, 107, 260, 196, 84, 268, 107, 165, 131,
	56, 28, 106, 58, 13, 12, 11, 10, 9, 8,
	138, 142, 140, 141, 7, 6, 5, 4, 2, 1,
	0, 0, 0, 0, 0, 0, 0, 104, 105, 0,
	154, 155, 156, 157, 109, 151, 152, 153, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 107,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 139,
	143, 144, 145, 146, 147, 148, 149, 150,
}
var yyPact = [...]int{

	201, -1000, -1000, 227, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	70, -23, 3, -5, 1, -1000, -1000, -1000, 474, 421,
	-1000, -1000, -1000, 410, -1000, 383, 363, 467, 316, -36,
	-16, 315, -1000, -12, 315, -1000, 367, -37, 315, -37,
	366, -1000, -1000, -1000, -1000, -1000, 409, -1000, 266, 363,
	373, 43, 363, 172, -1000, 200, -1000, 34, 361, 35,
	315, -1000, -1000, 360, -1000, -28, 357, 413, 108, 315,
	-1000, 220, -1000, -1000, 342, 30, 60, 577, -1000, 484,
	461, -1000, -1000, -1000, 541, 260, 259, -1000, 250, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 541,
	-1000, 221, 316, 348, 446, 316, 541, 315, 240, 412,
	-43, -1000, 176, -1000, 347, -1000, -1000, 344, -1000, 213,
	409, -1000, -1000, 315, 88, 484, 484, 541, 210, 418,
	541, 541, 100, 541, 541, 541, 541, 541, 541, 541,
	541, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 577,
	-38, -2, 29, 577, -1000, 489, 329, 409, -1000, 474,
	15, 165, 408, 316, 316, 206, -1000, 439, 484, -1000,
	165, -1000, 320, -1000, 98, 315, -1000, -33, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 204, 312, 333, 258,
	21, -1000, -1000, -1000, -1000, -1000, 55, 165, -1000, 489,
	-1000, -1000, 210, 541, 541, 165, 472, -1000, 389, 123,
	123, 123, 26, 26, -1000, -1000, -1000, -1000, -1000, 541,
	-1000, 165, -1000, 28, 409, 12, 14, -1000, 484, 95,
	230, 227, 115, 11, -1000, 439, 420, 437, 60, 9,
	-1000, 208, 328, -1000, -1000, 325, -1000, 446, 213, 310,
	213, -1000, -1000, 133, 121, 145, 137, 131, -20, -1000,
	323, -35, 321, 8, -1000, 165, 394, 541, -1000, 165,
	-1000, 2, -1000, 4, -1000, 541, 91, -1000, 379, 188,
	-1000, -1000, -1000, 316, 420, -1000, 541, 541, 320, -1000,
	-1000, -67, -1000, -1000, 239, -1000, 239, 239, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 238, 238, 238, 232, 232, -1000, -1000, 444, 312,
	464, 68, -1000, 124, -1000, 122, -1000, -1000, -1000, -1000,
	-18, -19, -21, -1000, -1000, -1000, -1000, 541, 165, -1000,
	-1000, 165, 541, 377, 230, -1000, -1000, 147, 177, -1000,
	313, -1000, 31, -73, -1000, -1000, 307, -1000, -1000, -1000,
	306, -1000, -1000, -1000, -1000, 294, -1000, -1000, -1000, 441,
	434, 309, 484, -1000, -1000, 226, 224, 222, 165, 165,
	471, -1000, 541, 541, -1000, -1000, -1000, 388, -1000, 305,
	-1000, -1000, -1000, -1000, 376, -1000, 350, -1000, -1000, -92,
	175, -1, 439, 484, 541, -1000, 60, 315, 315, 315,
	316, 165, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 271,
	420, 60, 173, -8, -1000, -22, -29, 172, -93, 346,
	-1000, 315, -1000, -1000, -1000, -1000, 470, 405, -1000, -1000,
	315, 315, -1000,
}
var yyPgo = [...]int{

	0, 609, 608, 22, 607, 606, 605, 604, 599, 598,
	597, 596, 595, 594, 445, 593, 591, 590, 20, 40,
	589, 586, 585, 584, 11, 583, 577, 189, 576, 3,
	17, 575, 19, 574, 573, 570, 569, 2, 21, 9,
	564, 6, 563, 15, 562, 4, 558, 557, 16, 555,
	554, 553, 542, 13, 541, 5, 539, 1, 534, 532,
	530, 18, 8, 44, 254, 525, 520, 518, 517, 515,
	513, 0, 7, 510, 10, 509, 505, 14, 504, 503,
	502, 497, 496, 495, 493, 12, 490, 489, 484, 482,
	480,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 3, 3, 4, 4, 5, 6, 7,
	81, 81, 73, 73, 73, 86, 86, 86, 86, 86,
	78, 78, 78, 79, 79, 83, 83, 83, 83, 83,
	83, 83, 84, 84, 84, 84, 84, 84, 84, 85,
	85, 77, 77, 80, 80, 87, 87, 87, 87, 87,
	87, 87, 82, 82, 88, 88, 89, 89, 74, 75,
	75, 76, 8, 8, 8, 9, 9, 9, 10, 11,
	11, 11, 12, 13, 13, 13, 90, 14, 15, 15,
	16, 16, 16, 16, 16, 17, 17, 18, 18, 19,
	19, 19, 22, 22, 20, 20, 20, 23, 23, 24,
	24, 24, 24, 21, 21, 21, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 26, 26, 26, 27, 27,
	28, 28, 28, 28, 29, 29, 30, 30, 32, 32,
	32, 32, 32, 33, 33, 33, 33, 33, 33, 33,
	33, 33, 33, 34, 34, 34, 34, 34, 34, 34,
	38, 38, 38, 43, 39, 39, 37, 37, 37, 37,
	37, 37, 37, 37, 37, 37, 37, 37, 37, 37,
	37, 37, 37, 42, 42, 44, 44, 44, 46, 49,
	49, 47, 47, 48, 50, 50, 45, 45, 36, 36,
	36, 36, 51, 51, 52, 52, 53, 53, 54, 54,
	55, 56, 56, 56, 31, 31, 31, 57, 57, 57,
	58, 58, 58, 59, 59, 60, 60, 61, 61, 35,
	35, 40, 40, 41, 41, 62, 62, 63, 64, 64,
	65, 65, 66, 66, 67, 67, 67, 67, 67, 68,
	68, 69, 69, 70, 70, 71, 72,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 13, 3, 7, 7, 8, 7, 3,
	0, 1, 3, 1, 1, 1, 1, 1, 1, 1,
	2, 2, 1, 2, 1, 1, 1, 1, 1, 1,
	1, 1, 2, 2, 2, 2, 2, 2, 2, 0,
	5, 0, 3, 0, 1, 0, 3, 2, 3, 3,
	2, 2, 1, 1, 2, 1, 1, 2, 3, 1,
	3, 7, 1, 8, 4, 6, 7, 4, 5, 4,
	5, 5, 3, 2, 2, 2, 0, 2, 0, 2,
	1, 2, 1, 1, 1, 0, 1, 1, 3, 1,
	2, 3, 1, 1, 0, 1, 2, 1, 3, 3,
	3, 3, 5, 0, 1, 2, 1, 1, 2, 3,
	2, 3, 2, 2, 2, 1, 3, 1, 1, 3,
	0, 5, 5, 5, 1, 3, 0, 2, 1, 3,
	3, 2, 3, 3, 3, 4, 3, 4, 5, 6,
	3, 4, 2, 1, 1, 1, 1, 1, 1, 1,
	3, 1, 1, 3, 1, 3, 1, 1, 1, 3,
	3, 3, 3, 3, 3, 3, 3, 2, 3, 4,
	5, 4, 1, 1, 1, 1, 1, 1, 5, 0,
	1, 1, 2, 4, 0, 2, 1, 3, 1, 1,
	1, 1, 0, 3, 0, 2, 0, 3, 1, 3,
	2, 0, 1, 1, 0, 2, 4, 0, 2, 4,
	0, 2, 4, 0, 3, 1, 3, 0, 5, 2,
	1, 1, 3, 3, 1, 1, 3, 3, 0, 2,
	0, 3, 0, 1, 1, 1, 1, 1, 1, 0,
	1, 0, 1, 0, 2, 1, 0,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -12, -13, 5, 6, 7, 8, 35, -76,
	88, 89, 91, 90, 92, 100, 101, 102, -16, 54,
	55, 56, 57, -14, -90, -14, -14, -14, -14, 93,
	-69, 95, 53, -66, 95, 97, 93, 93, 94, 95,
	93, -72, -72, -72, -3, 19, -17, 20, -15, 31,
	-27, 37, 9, -62, -63, -45, -71, 37, -65, 98,
	94, -71, 37, 93, -71, 37, -64, 98, -71, -64,
	37, -18, -19, 78, -22, 37, -32, -37, -33, 72,
	47, -36, -45, -41, -44, -71, -42, -46, 22, 38,
	39, 40, 27, -43, 76, 77, 51, 98, 30, 83,
	42, -27, 35, 81, -27, 58, 48, 81, 37, 72,
	-71, -72, 37, -72, 96, 37, 22, 69, -71, 9,
	58, -20, -71, 21, 81, 71, 70, -34, 23, 72,
	25, 26, 24, 73, 74, 75, 76, 77, 78, 79,
	80, 48, 49, 50, 43, 44, 45, 46, -32, -37,
	-32, -3, -39, -37, -37, 47, 47, 47, -43, 47,
	-49, -37, -59, 35, 47, -62, 37, -30, 12, -63,
	-37, -71, 47, 22, -70, 99, -67, 91, 89, 34,
	90, 15, 37, 37, 37, -72, -23, -24, -26, 47,
	37, -43, -19, -71, 78, -32, -32, -37, -38, 47,
	-43, 41, 23, 25, 26, -37, -37, 27, 72, -37,
	-37, -37, -37, -37, -37, -37, -37, 130, 130, 58,
	130, -37, 130, -18, 20, -18, -47, -48, 84, -35,
	30, -3, -62, -60, -45, -30, -53, 15, -32, -75,
	-74, 37, 69, -71, -72, -68, 96, -31, 58, 10,
	-25, 59, 60, 61, 62, 63, 65, 66, -21, 37,
	21, -24, 81, -39, -38, -37, -37, 71, 27, -37,
	130, -18, 130, -50, -48, 86, -32, -61, 69, -40,
	-41, -61, 130, 58, -53, -57, 17, 16, 58, 130,
	-73, -79, -78, -86, -83, -84, 123, 124, 122, 117,
	118, 119, 120, 121, 103, 104, 105, 106, 107, 108,
	109, 110, 111, 112, 115, 116, 37, 37, -30, -24,
	38, -24, 59, 64, 59, 64, 59, 59, 59, -28,
	67, 97, 68, 37, 130, 37, 130, 71, -37, 130,
	87, -37, 85, 32, 58, -45, -57, -37, -54, -55,
	-37, -74, -87, -80, 113, -77, 47, -77, -77, -85,
	47, -85, -85, -85, -77, 47, -85, -77, -72, -51,
	13, 11, 69, 59, 59, 94, 94, 94, -37, -37,
	33, -41, 58, 58, -56, 28, 29, 72, 27, 34,
	126, -82, -88, -89, 52, 33, 53, -81, 114, 39,
	39, 39, -52, 14, 16, 38, -32, 47, 47, 47,
	7, -37, -55, 27, 38, 39, 33, 33, 130, 58,
	-53, -32, -39, -29, -71, -29, -29, -62, 39, -57,
	130, 58, 130, 130, 130, -58, 18, 36, -71, 7,
	23, -71, -71,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 86, 86, 86, 86, 86, 72,
	251, 242, 0, 0, 0, 256, 256, 256, 0, 90,
	92, 93, 94, 95, 88, 0, 0, 0, 0, 240,
	0, 0, 252, 0, 0, 243, 0, 238, 0, 238,
	0, 83, 84, 85, 14, 91, 0, 96, 87, 0,
	0, 128, 0, 19, 235, 0, 196, 255, 0, 0,
	0, 256, 255, 0, 256, 0, 0, 0, 0, 0,
	82, 0, 97, 99, 104, 255, 102, 103, 138, 0,
	0, 166, 167, 168, 0, 196, 0, 182, 0, 198,
	199, 200, 201, 234, 185, 186, 187, 183, 184, 189,
	89, 223, 0, 0, 136, 0, 0, 0, 0, 0,
	253, 74, 0, 77, 0, 79, 239, 0, 256, 0,
	0, 100, 105, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 153, 154, 155, 156, 157, 158, 159, 141, 0,
	0, 0, 0, 164, 177, 0, 0, 0, 152, 0,
	0, 190, 0, 0, 0, 136, 129, 206, 0, 236,
	237, 197, 0, 241, 0, 0, 256, 249, 244, 245,
	246, 247, 248, 78, 80, 81, 214, 107, 113, 0,
	125, 127, 98, 106, 101, 139, 140, 143, 144, 0,
	161, 162, 0, 0, 0, 146, 0, 150, 0, 169,
	170, 171, 172, 173, 174, 175, 176, 142, 163, 0,
	233, 164, 178, 0, 0, 0, 194, 191, 0, 227,
	0, 230, 227, 0, 225, 206, 217, 0, 137, 0,
	69, 0, 0, 254, 75, 0, 250, 136, 0, 0,
	0, 116, 117, 0, 0, 0, 0, 0, 130, 114,
	0, 0, 0, 0, 145, 147, 0, 0, 151, 165,
	179, 0, 181, 0, 192, 0, 0, 15, 0, 229,
	231, 16, 224, 0, 217, 18, 0, 0, 0, 71,
	55, 53, 23, 24, 51, 34, 51, 51, 32, 25,
	26, 27, 28, 29, 35, 36, 37, 38, 39, 40,
	41, 49, 49, 49, 49, 49, 256, 76, 202, 108,
	215, 111, 118, 0, 120, 0, 122, 123, 124, 109,
	0, 0, 0, 115, 110, 126, 160, 0, 148, 180,
	188, 195, 0, 0, 0, 226, 17, 218, 207, 208,
	211, 70, 68, 20, 54, 33, 0, 30, 31, 42,
	0, 43, 44, 45, 46, 0, 47, 48, 73, 204,
	0, 0, 0, 119, 121, 0, 0, 0, 149, 193,
	0, 232, 0, 0, 210, 212, 213, 0, 57, 0,
	60, 61, 62, 63, 0, 65, 66, 22, 21, 0,
	0, 0, 206, 0, 0, 216, 112, 0, 0, 0,
	0, 219, 209, 56, 58, 59, 64, 67, 52, 0,
	217, 205, 203, 0, 134, 0, 0, 228, 0, 220,
	131, 0, 132, 133, 50, 13, 0, 0, 135, 221,
	0, 0, 222,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 80, 73, 3,
	47, 130, 78, 76, 58, 77, 81, 79, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	49, 48, 50, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 75, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 74, 3, 51,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 52, 53, 54, 55, 56,
	57, 59, 60, 61, 62, 63, 64, 65, 66, 67,
	68, 69, 70, 71, 72, 82, 83, 84, 85, 86,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100, 101, 102, 103, 104, 105, 106,
	107, 108, 109, 110, 111, 112, 113, 114, 115, 116,
	117, 118, 119, 120, 121, 122, 123, 124, 125, 126,
	127, 128, 129,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:183
		{
			SetParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:189
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 13:
		yyDollar = yyS[yypt-13 : yypt+1]
		//line sql.y:205
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, From: yyDollar[6].tableExprs, TimeRange: yyDollar[7].timerange, Where: NewWhere(AST_WHERE, yyDollar[8].boolExpr), GroupBy: GroupBy(yyDollar[9].valExprs), Having: NewWhere(AST_HAVING, yyDollar[10].boolExpr), OrderBy: yyDollar[11].orderBy, Limit: yyDollar[12].limit, Lock: yyDollar[13].str}
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:209
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 15:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:215
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows, OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 16:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:219
		{
			cols := make(Columns, 0, len(yyDollar[6].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[6].updateExprs))
			for _, col := range yyDollar[6].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 17:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:231
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(AST_WHERE, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 18:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:237
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(AST_WHERE, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:243
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 20:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:248
		{
			yyVAL.str = ""
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:252
		{
			yyVAL.str = AST_ZEROFILL
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:257
		{
			yyVAL.str = yyDollar[1].str
			if yyDollar[2].str != "" {
				yyVAL.str += " " + yyDollar[2].str
			}
			if yyDollar[3].str != "" {
				yyVAL.str += " " + yyDollar[3].str
			}
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:271
		{
			yyVAL.str = AST_DATE
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:275
		{
			yyVAL.str = AST_TIME
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:279
		{
			yyVAL.str = AST_TIMESTAMP
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:283
		{
			yyVAL.str = AST_DATETIME
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:287
		{
			yyVAL.str = AST_YEAR
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:293
		{
			if yyDollar[2].str == "" {
				yyVAL.str = AST_CHAR
			} else {
				yyVAL.str = AST_CHAR + yyDollar[2].str
			}
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:301
		{
			if yyDollar[2].str == "" {
				yyVAL.str = AST_VARCHAR
			} else {
				yyVAL.str = AST_VARCHAR + yyDollar[2].str
			}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:309
		{
			yyVAL.str = AST_TEXT
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:315
		{
			yyVAL.str = yyDollar[1].str + yyDollar[2].str
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:319
		{
			yyVAL.str = yyDollar[1].str
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:325
		{
			yyVAL.str = AST_BIT
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:329
		{
			yyVAL.str = AST_TINYINT
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:333
		{
			yyVAL.str = AST_SMALLINT
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:337
		{
			yyVAL.str = AST_MEDIUMINT
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:341
		{
			yyVAL.str = AST_INT
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:345
		{
			yyVAL.str = AST_INTEGER
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:349
		{
			yyVAL.str = AST_BIGINT
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:355
		{
			yyVAL.str = AST_REAL + yyDollar[2].str
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:359
		{
			yyVAL.str = AST_DOUBLE + yyDollar[2].str
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:363
		{
			yyVAL.str = AST_FLOAT + yyDollar[2].str
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:367
		{
			yyVAL.str = AST_DECIMAL + yyDollar[2].str
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:371
		{
			yyVAL.str = AST_DECIMAL + yyDollar[2].str
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:375
		{
			yyVAL.str = AST_NUMERIC + yyDollar[2].str
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:379
		{
			yyVAL.str = AST_NUMERIC + yyDollar[2].str
		}
	case 49:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:384
		{
			yyVAL.str = ""
		}
	case 50:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:388
		{
			yyVAL.str = "(" + string(yyDollar[2].bytes) + ", " + string(yyDollar[4].bytes) + ")"
		}
	case 51:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:393
		{
			yyVAL.str = ""
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:397
		{
			yyVAL.str = "(" + string(yyDollar[2].bytes) + ")"
		}
	case 53:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:402
		{
			yyVAL.str = ""
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:406
		{
			yyVAL.str = AST_UNSIGNED
		}
	case 55:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:411
		{
			yyVAL.columnAtts = ColumnAtts{}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:415
		{
			yyVAL.columnAtts = append(yyVAL.columnAtts, AST_NOT_NULL)
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:421
		{
			node := StrVal(yyDollar[3].bytes)
			yyVAL.columnAtts = append(yyVAL.columnAtts, "default "+String(node))
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:426
		{
			node := NumVal(yyDollar[3].bytes)
			yyVAL.columnAtts = append(yyVAL.columnAtts, "default "+String(node))
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:431
		{
			yyVAL.columnAtts = append(yyVAL.columnAtts, AST_AUTO_INCREMENT)
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:435
		{
			yyVAL.columnAtts = append(yyVAL.columnAtts, yyDollar[2].str)
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:441
		{
			yyVAL.str = AST_PRIMARY_KEY
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:445
		{
			yyVAL.str = AST_UNIQUE_KEY
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:459
		{
			yyVAL.columnDefinition = &ColumnDefinition{ColName: string(yyDollar[1].bytes), ColType: yyDollar[2].str, ColumnAtts: yyDollar[3].columnAtts}
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:465
		{
			yyVAL.columnDefinitions = ColumnDefinitions{yyDollar[1].columnDefinition}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:469
		{
			yyVAL.columnDefinitions = append(yyVAL.columnDefinitions, yyDollar[3].columnDefinition)
		}
	case 71:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:475
		{
			yyVAL.statement = &CreateTable{Name: yyDollar[4].bytes, ColumnDefinitions: yyDollar[6].columnDefinitions}
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:481
		{
			yyVAL.statement = yyDollar[1].statement
		}
	case 73:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:485
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].bytes, NewName: yyDollar[7].bytes}
		}
	case 74:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:490
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[3].bytes}
		}
	case 75:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:496
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[4].bytes, NewName: yyDollar[4].bytes}
		}
	case 76:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:500
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[4].bytes, NewName: yyDollar[7].bytes}
		}
	case 77:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:505
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 78:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:511
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].bytes, NewName: yyDollar[5].bytes}
		}
	case 79:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:517
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 80:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:521
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].bytes, NewName: yyDollar[5].bytes}
		}
	case 81:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:526
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:532
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:538
		{
			yyVAL.statement = &Other{}
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:542
		{
			yyVAL.statement = &Other{}
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:546
		{
			yyVAL.statement = &Other{}
		}
	case 86:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:551
		{
			SetAllowComments(yylex, true)
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:555
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 88:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:561
		{
			yyVAL.bytes2 = nil
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:565
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:571
		{
			yyVAL.str = AST_UNION
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:575
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 92:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:579
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:583
		{
			yyVAL.str = AST_EXCEPT
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:587
		{
			yyVAL.str = AST_INTERSECT
		}
	case 95:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:592
		{
			yyVAL.str = ""
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:596
		{
			yyVAL.str = AST_DISTINCT
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:602
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:606
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:612
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:616
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:620
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:626
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:630
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 104:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:635
		{
			yyVAL.bytes = nil
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:639
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 106:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:643
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:649
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:653
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:659
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:663
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:667
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 112:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:671
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 113:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:676
		{
			yyVAL.bytes = nil
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:680
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 115:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:684
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:690
		{
			yyVAL.str = AST_JOIN
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:694
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 118:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:698
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 119:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:702
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 120:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:706
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 121:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:710
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 122:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:714
		{
			yyVAL.str = AST_JOIN
		}
	case 123:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:718
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:722
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:728
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:732
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:736
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:742
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 129:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:746
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 130:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:751
		{
			yyVAL.indexHints = nil
		}
	case 131:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:755
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 132:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:759
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 133:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:763
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:769
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:773
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 136:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:778
		{
			yyVAL.boolExpr = nil
		}
	case 137:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:782
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:789
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 140:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:793
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 141:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:797
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:801
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:807
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:811
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].colTuple}
		}
	case 145:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:815
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].colTuple}
		}
	case 146:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:819
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 147:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:823
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 148:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:827
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 149:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:831
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:835
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 151:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:839
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 152:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:843
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:849
		{
			yyVAL.str = AST_EQ
		}
	case 154:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:853
		{
			yyVAL.str = AST_LT
		}
	case 155:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:857
		{
			yyVAL.str = AST_GT
		}
	case 156:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:861
		{
			yyVAL.str = AST_LE
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:865
		{
			yyVAL.str = AST_GE
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:869
		{
			yyVAL.str = AST_NE
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:873
		{
			yyVAL.str = AST_NSE
		}
	case 160:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:879
		{
			yyVAL.colTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:883
		{
			yyVAL.colTuple = yyDollar[1].subquery
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:887
		{
			yyVAL.colTuple = ListArg(yyDollar[1].bytes)
		}
	case 163:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:893
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 164:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:899
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 165:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:903
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:909
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 167:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:913
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 168:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:917
		{
			yyVAL.valExpr = yyDollar[1].rowTuple
		}
	case 169:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:921
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 170:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:925
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 171:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:929
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:933
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 173:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:937
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 174:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:941
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 175:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:945
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 176:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:949
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 177:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:953
		{
			if num, ok := yyDollar[2].valExpr.(NumVal); ok {
				switch yyDollar[1].byt {
				case '-':
					yyVAL.valExpr = append(NumVal("-"), num...)
				case '+':
					yyVAL.valExpr = num
				default:
					yyVAL.valExpr = &UnaryExpr{Operator: yyDollar[1].byt, Expr: yyDollar[2].valExpr}
				}
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: yyDollar[1].byt, Expr: yyDollar[2].valExpr}
			}
		}
	case 178:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:968
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 179:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:972
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 180:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:976
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 181:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:980
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 182:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:984
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 183:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:990
		{
			yyVAL.bytes = IF_BYTES
		}
	case 184:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:994
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 185:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1000
		{
			yyVAL.byt = AST_UPLUS
		}
	case 186:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1004
		{
			yyVAL.byt = AST_UMINUS
		}
	case 187:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1008
		{
			yyVAL.byt = AST_TILDA
		}
	case 188:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1014
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 189:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1019
		{
			yyVAL.valExpr = nil
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1023
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1029
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 192:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1033
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 193:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1039
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 194:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1044
		{
			yyVAL.valExpr = nil
		}
	case 195:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1048
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 196:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1054
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 197:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1058
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 198:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1064
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 199:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1068
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1072
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 201:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1076
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 202:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1081
		{
			yyVAL.valExprs = nil
		}
	case 203:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1085
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 204:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1090
		{
			yyVAL.boolExpr = nil
		}
	case 205:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1094
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 206:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1099
		{
			yyVAL.orderBy = nil
		}
	case 207:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1103
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 208:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1109
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 209:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1113
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 210:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1119
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 211:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1124
		{
			yyVAL.str = AST_ASC
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1128
		{
			yyVAL.str = AST_ASC
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1132
		{
			yyVAL.str = AST_DESC
		}
	case 214:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1137
		{
			yyVAL.timerange = nil
		}
	case 215:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1141
		{
			yyVAL.timerange = &TimeRange{From: string(yyDollar[2].bytes)}
		}
	case 216:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1145
		{
			yyVAL.timerange = &TimeRange{From: string(yyDollar[2].bytes), To: string(yyDollar[4].bytes)}
		}
	case 217:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1150
		{
			yyVAL.limit = nil
		}
	case 218:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1154
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 219:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1158
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 220:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1163
		{
			yyVAL.str = ""
		}
	case 221:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1167
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 222:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1171
		{
			if !bytes.Equal(yyDollar[3].bytes, SHARE) {
				yylex.Error("expecting share")
				return 1
			}
			if !bytes.Equal(yyDollar[4].bytes, MODE) {
				yylex.Error("expecting mode")
				return 1
			}
			yyVAL.str = AST_SHARE_MODE
		}
	case 223:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1184
		{
			yyVAL.columns = nil
		}
	case 224:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1188
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 225:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1194
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 226:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1198
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 227:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1203
		{
			yyVAL.updateExprs = nil
		}
	case 228:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1207
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 229:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1213
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 230:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1217
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 231:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1223
		{
			yyVAL.values = Values{yyDollar[1].rowTuple}
		}
	case 232:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1227
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].rowTuple)
		}
	case 233:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1233
		{
			yyVAL.rowTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 234:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1237
		{
			yyVAL.rowTuple = yyDollar[1].subquery
		}
	case 235:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1243
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 236:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1247
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 237:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1253
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 238:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1258
		{
			yyVAL.empty = struct{}{}
		}
	case 239:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1260
		{
			yyVAL.empty = struct{}{}
		}
	case 240:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1263
		{
			yyVAL.empty = struct{}{}
		}
	case 241:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1265
		{
			yyVAL.empty = struct{}{}
		}
	case 242:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1268
		{
			yyVAL.empty = struct{}{}
		}
	case 243:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1270
		{
			yyVAL.empty = struct{}{}
		}
	case 244:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1274
		{
			yyVAL.empty = struct{}{}
		}
	case 245:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1276
		{
			yyVAL.empty = struct{}{}
		}
	case 246:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1278
		{
			yyVAL.empty = struct{}{}
		}
	case 247:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1280
		{
			yyVAL.empty = struct{}{}
		}
	case 248:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1282
		{
			yyVAL.empty = struct{}{}
		}
	case 249:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1285
		{
			yyVAL.empty = struct{}{}
		}
	case 250:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1287
		{
			yyVAL.empty = struct{}{}
		}
	case 251:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1290
		{
			yyVAL.empty = struct{}{}
		}
	case 252:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1292
		{
			yyVAL.empty = struct{}{}
		}
	case 253:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1295
		{
			yyVAL.empty = struct{}{}
		}
	case 254:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1297
		{
			yyVAL.empty = struct{}{}
		}
	case 255:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1301
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 256:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1306
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
