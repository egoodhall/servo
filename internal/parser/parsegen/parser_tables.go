// generated by Textmapper; DO NOT EDIT

package parsegen

import (
	"fmt"

	"github.com/egoodhall/servo/internal/parser/parsegen/token"
)

var tmNonterminals = [...]string{
	"Comment",
	"WhiteSpace",
	"WhiteSpace$1",
	"OptionName",
	"OptionString",
	"OptionBool",
	"OptionInt",
	"OptionFloat",
	"OptionValue",
	"Option",
	"TypeRef",
	"FieldName",
	"ScalarType",
	"MapKeyType",
	"MapValueType",
	"MapType",
	"ListElementType",
	"ListType",
	"FieldMod",
	"FieldType",
	"FieldDef",
	"FieldDefList",
	"MessageName",
	"FieldDefList_optlist",
	"Message",
	"UnionMemberName",
	"UnionMemberType",
	"UnionMember",
	"UnionMemberList",
	"UnionName",
	"Union",
	"UnionMemberList_optlist",
	"RpcName",
	"RpcRequest",
	"RpcResponse",
	"Method",
	"MethodList",
	"ServiceName",
	"MethodList_list",
	"Service",
	"EnumName",
	"EnumValue",
	"EnumField",
	"EnumValueList",
	"Enum",
	"EnumValueList_list",
	"AliasName",
	"AliasType",
	"Alias",
	"Definition",
	"DefinitionList",
	"DefinitionList_optlist",
	"ServoFile",
}

func symbolName(sym int32) string {
	if sym == noToken {
		return "<no-token>"
	}
	if sym < int32(token.NumTokens) {
		return token.Type(sym).String()
	}
	if i := int(sym) - int(token.NumTokens); i < len(tmNonterminals) {
		return tmNonterminals[i]
	}
	return fmt.Sprintf("nonterminal(%d)", sym)
}

var tmAction = []int32{
	-3, 6, 1, 0, 5, 191, -27, -73, -91, 4, 3, -1, -1, -1, -1, -1, -1, -1, 184,
	180, 181, 182, 183, 185, -109, 190, 165, -1, 45, -1, 66, -1, 158, -1, 7, -1,
	177, -1, 186, 187, 189, -1, -1, -133, -1, -147, -1, -1, -1, -1, -1, 166, -1,
	-1, -161, 176, -1, -1, 47, -1, -173, 72, -1, -187, -1, -1, -201, 160, -1, -1,
	8, 9, 10, 11, 12, 13, 14, 15, -1, 178, -1, -1, 168, -1, 170, 174, 175, -1,
	-1, -1, -1, 19, 51, -1, -213, 46, 47, -1, -1, -1, 52, 70, -1, -227, 71, 72,
	-1, 73, -1, -1, 157, 164, 159, -1, -1, 16, 179, 173, 167, -1, 172, 50, 41,
	42, -1, -1, 44, -1, 49, 69, 62, 63, -1, -1, 65, -1, 68, -1, -1, 163, -1, 162,
	171, 17, -1, 18, -1, 20, -241, -253, -265, -1, -1, 48, -1, 53, -1, -1, 67,
	-1, 74, -1, -1, 161, -277, 24, -1, -1, -1, 26, 27, 29, 31, 40, -1, -1, -1,
	-1, 61, -1, -1, -1, -1, -1, -1, -1, -1, -1, 25, 38, -1, 39, -1, 36, -1, 59,
	-1, 60, -1, 57, -1, -1, -1, -1, 155, -1, -1, -1, -1, -1, 22, -1, 37, 34, -1,
	35, 58, 55, -1, 56, -1, 135, -1, -1, -1, 75, -1, -1, 154, -1, 145, -1, -1,
	-1, -1, 115, -1, -1, 23, 33, 54, -1, -1, -1, 134, -1, 125, -1, -1, 153, -1,
	-1, -1, -1, -1, -1, 144, -1, 95, -1, -1, -1, -1, -1, 114, -1, 105, -1, -1,
	133, -1, -1, -1, -1, -1, -1, 124, 151, -1, 152, -1, 149, -1, -1, 143, -1, -1,
	-1, -1, -1, -1, 94, -1, 85, -1, -1, 113, -1, -1, -1, -1, -1, -1, 104, 131,
	-1, 132, -1, 129, -1, -1, 123, -1, -1, -1, 150, 147, -1, 148, 141, -1, 142,
	-1, 139, -1, -1, 93, -1, -1, -1, -1, -1, -1, 84, 111, -1, 112, -1, 109, -1,
	-1, 103, -1, -1, -1, 130, 127, -1, 128, 121, -1, 122, -1, 119, -1, 146, 140,
	137, -1, 138, 91, -1, 92, -1, 89, -1, -1, 83, -1, -1, -1, 110, 107, -1, 108,
	101, -1, 102, -1, 99, -1, 126, 120, 117, -1, 118, 136, 90, 87, -1, 88, 81,
	-1, 82, -1, 79, -1, 106, 100, 97, -1, 98, 116, 86, 80, 77, -1, 78, 96, 76,
	-1, -2,
}

var tmLalr = []int32{
	8, -1, 9, -1, 10, -1, 0, 191, 2, 191, 3, 191, 4, 191, 5, 191, 6, 191, 7, 191,
	11, 191, -1, -2, 8, -1, 9, -1, 10, -1, 0, 2, 2, 2, 3, 2, 4, 2, 5, 2, 6, 2, 7,
	2, 11, 2, 12, 2, 14, 2, 16, 2, 17, 2, 18, 2, 19, 2, 20, 2, 22, 2, 28, 2, 29,
	2, 30, 2, -1, -2, 2, -1, 3, -1, 4, -1, 5, -1, 6, -1, 7, -1, 11, -1, 0, 193,
	-1, -2, 2, -1, 3, -1, 4, -1, 5, -1, 6, -1, 7, -1, 11, -1, 0, 192, -1, -2, 8,
	-1, 9, -1, 10, -1, 0, 188, 2, 188, 3, 188, 4, 188, 5, 188, 6, 188, 7, 188,
	11, 188, -1, -2, 8, -1, 9, -1, 10, -1, 11, 47, 12, 47, 29, 47, -1, -2, 8, -1,
	9, -1, 10, -1, 11, 72, 12, 72, 29, 72, -1, -2, 8, -1, 9, -1, 10, -1, 20, 169,
	29, 169, -1, -2, 8, -1, 9, -1, 10, -1, 11, 47, 12, 47, 29, 47, -1, -2, 8, -1,
	9, -1, 10, -1, 11, 72, 12, 72, 29, 72, -1, -2, 8, -1, 9, -1, 10, -1, 17, 156,
	29, 156, -1, -2, 8, -1, 9, -1, 10, -1, 11, 43, 12, 43, 29, 43, -1, -2, 8, -1,
	9, -1, 10, -1, 11, 64, 12, 64, 29, 64, -1, -2, 13, -1, 8, 28, 9, 28, 10, 28,
	22, 28, -1, -2, 13, -1, 8, 30, 9, 30, 10, 30, 22, 30, -1, -2, 13, -1, 8, 32,
	9, 32, 10, 32, 22, 32, -1, -2, 15, 18, 16, 21, -1, -2,
}

var tmGoto = []int32{
	0, 2, 2, 6, 10, 14, 18, 22, 26, 220, 414, 608, 628, 750, 756, 764, 768, 810,
	826, 830, 846, 862, 864, 1070, 1072, 1074, 1076, 1078, 1080, 1096, 1134,
	1228, 1422, 1614, 1806, 1808, 1810, 1812, 1814, 1816, 1818, 1822, 1914, 1922,
	1930, 1932, 1934, 1942, 1944, 1952, 1958, 1966, 1974, 1982, 1984, 1992, 1996,
	2004, 2012, 2020, 2028, 2030, 2034, 2042, 2044, 2052, 2116, 2132, 2148, 2150,
	2158, 2162, 2164, 2180, 2196, 2212, 2216, 2224, 2226, 2228, 2232, 2236, 2240,
	2244, 2246,
}

var tmFromTo = []int16{
	410, 411, 7, 11, 8, 11, 7, 12, 8, 12, 7, 13, 8, 13, 7, 14, 8, 14, 7, 15, 8,
	15, 7, 16, 8, 16, 0, 1, 6, 9, 24, 1, 27, 1, 29, 1, 31, 1, 33, 1, 41, 1, 43,
	1, 45, 1, 47, 1, 53, 1, 54, 1, 57, 1, 60, 1, 63, 1, 66, 1, 69, 1, 93, 1, 94,
	1, 102, 1, 103, 1, 108, 1, 124, 1, 132, 1, 137, 1, 151, 1, 152, 1, 156, 1,
	157, 1, 161, 1, 162, 1, 168, 1, 176, 1, 177, 1, 181, 1, 182, 1, 183, 1, 186,
	1, 192, 1, 198, 1, 201, 1, 203, 1, 206, 1, 207, 1, 208, 1, 220, 1, 223, 1,
	226, 1, 227, 1, 229, 1, 232, 1, 234, 1, 237, 1, 242, 1, 243, 1, 245, 1, 248,
	1, 252, 1, 254, 1, 255, 1, 257, 1, 260, 1, 262, 1, 263, 1, 265, 1, 268, 1,
	272, 1, 274, 1, 275, 1, 280, 1, 283, 1, 287, 1, 289, 1, 290, 1, 292, 1, 295,
	1, 299, 1, 301, 1, 302, 1, 307, 1, 310, 1, 314, 1, 322, 1, 325, 1, 329, 1,
	331, 1, 332, 1, 337, 1, 340, 1, 344, 1, 352, 1, 363, 1, 366, 1, 370, 1, 378,
	1, 394, 1, 0, 2, 6, 2, 24, 2, 27, 2, 29, 2, 31, 2, 33, 2, 41, 2, 43, 2, 45,
	2, 47, 2, 53, 2, 54, 2, 57, 2, 60, 2, 63, 2, 66, 2, 69, 2, 93, 2, 94, 2, 102,
	2, 103, 2, 108, 2, 124, 2, 132, 2, 137, 2, 151, 2, 152, 2, 156, 2, 157, 2,
	161, 2, 162, 2, 168, 2, 176, 2, 177, 2, 181, 2, 182, 2, 183, 2, 186, 2, 192,
	2, 198, 2, 201, 2, 203, 2, 206, 2, 207, 2, 208, 2, 220, 2, 223, 2, 226, 2,
	227, 2, 229, 2, 232, 2, 234, 2, 237, 2, 242, 2, 243, 2, 245, 2, 248, 2, 252,
	2, 254, 2, 255, 2, 257, 2, 260, 2, 262, 2, 263, 2, 265, 2, 268, 2, 272, 2,
	274, 2, 275, 2, 280, 2, 283, 2, 287, 2, 289, 2, 290, 2, 292, 2, 295, 2, 299,
	2, 301, 2, 302, 2, 307, 2, 310, 2, 314, 2, 322, 2, 325, 2, 329, 2, 331, 2,
	332, 2, 337, 2, 340, 2, 344, 2, 352, 2, 363, 2, 366, 2, 370, 2, 378, 2, 394,
	2, 0, 3, 6, 3, 24, 3, 27, 3, 29, 3, 31, 3, 33, 3, 41, 3, 43, 3, 45, 3, 47, 3,
	53, 3, 54, 3, 57, 3, 60, 3, 63, 3, 66, 3, 69, 3, 93, 3, 94, 3, 102, 3, 103,
	3, 108, 3, 124, 3, 132, 3, 137, 3, 151, 3, 152, 3, 156, 3, 157, 3, 161, 3,
	162, 3, 168, 3, 176, 3, 177, 3, 181, 3, 182, 3, 183, 3, 186, 3, 192, 3, 198,
	3, 201, 3, 203, 3, 206, 3, 207, 3, 208, 3, 220, 3, 223, 3, 226, 3, 227, 3,
	229, 3, 232, 3, 234, 3, 237, 3, 242, 3, 243, 3, 245, 3, 248, 3, 252, 3, 254,
	3, 255, 3, 257, 3, 260, 3, 262, 3, 263, 3, 265, 3, 268, 3, 272, 3, 274, 3,
	275, 3, 280, 3, 283, 3, 287, 3, 289, 3, 290, 3, 292, 3, 295, 3, 299, 3, 301,
	3, 302, 3, 307, 3, 310, 3, 314, 3, 322, 3, 325, 3, 329, 3, 331, 3, 332, 3,
	337, 3, 340, 3, 344, 3, 352, 3, 363, 3, 366, 3, 370, 3, 378, 3, 394, 3, 7,
	17, 8, 17, 59, 90, 62, 99, 89, 90, 97, 90, 98, 99, 106, 99, 127, 90, 135, 99,
	11, 26, 12, 28, 13, 30, 14, 32, 15, 34, 16, 36, 59, 91, 62, 100, 64, 107, 89,
	91, 97, 91, 98, 100, 106, 100, 124, 143, 127, 91, 132, 143, 135, 100, 137,
	143, 144, 143, 146, 143, 152, 143, 154, 143, 157, 143, 159, 143, 162, 143,
	175, 143, 180, 143, 185, 143, 187, 143, 203, 143, 220, 143, 224, 143, 227,
	143, 229, 143, 234, 143, 241, 143, 243, 143, 245, 143, 251, 143, 253, 143,
	255, 143, 257, 143, 261, 143, 263, 143, 265, 143, 271, 143, 273, 143, 275,
	143, 286, 143, 288, 143, 290, 143, 292, 143, 298, 143, 300, 143, 302, 143,
	313, 143, 328, 143, 330, 143, 332, 143, 343, 143, 369, 143, 148, 169, 149,
	169, 150, 169, 124, 144, 146, 144, 152, 144, 175, 144, 167, 188, 211, 238,
	93, 124, 102, 132, 125, 152, 133, 157, 166, 187, 183, 203, 201, 220, 205,
	227, 206, 229, 208, 234, 222, 243, 223, 245, 231, 255, 232, 257, 236, 263,
	237, 265, 247, 275, 259, 290, 260, 292, 267, 302, 294, 332, 47, 64, 65, 64,
	68, 64, 69, 64, 109, 64, 113, 64, 114, 64, 140, 64, 108, 137, 138, 162, 161,
	183, 182, 201, 184, 206, 186, 208, 202, 223, 207, 232, 209, 237, 233, 260,
	41, 51, 52, 51, 56, 51, 57, 51, 81, 51, 87, 51, 88, 51, 119, 51, 37, 50, 17,
	38, 53, 82, 78, 115, 80, 116, 83, 118, 90, 122, 99, 130, 151, 173, 156, 178,
	168, 189, 174, 191, 176, 193, 177, 195, 179, 197, 181, 199, 183, 204, 190,
	212, 192, 213, 194, 215, 196, 216, 198, 217, 200, 219, 201, 221, 205, 228,
	206, 230, 208, 235, 214, 239, 218, 240, 222, 244, 223, 246, 226, 249, 231,
	256, 232, 258, 236, 264, 237, 266, 242, 269, 247, 276, 248, 277, 250, 279,
	252, 281, 254, 284, 259, 291, 260, 293, 262, 296, 267, 303, 268, 304, 270,
	306, 272, 308, 274, 311, 278, 315, 280, 316, 282, 318, 283, 319, 285, 321,
	287, 323, 289, 326, 294, 333, 295, 334, 297, 336, 299, 338, 301, 341, 305,
	345, 307, 346, 309, 348, 310, 349, 312, 351, 314, 353, 317, 355, 320, 356,
	322, 357, 324, 359, 325, 360, 327, 362, 329, 364, 331, 367, 335, 371, 337,
	372, 339, 374, 340, 375, 342, 377, 344, 379, 347, 381, 350, 382, 352, 383,
	354, 385, 358, 386, 361, 387, 363, 388, 365, 390, 366, 391, 368, 393, 370,
	395, 373, 397, 376, 398, 378, 399, 380, 401, 384, 402, 389, 403, 392, 404,
	394, 405, 396, 407, 400, 408, 406, 409, 35, 49, 49, 70, 49, 71, 49, 72, 49,
	73, 27, 41, 29, 43, 31, 45, 33, 47, 42, 57, 44, 60, 46, 63, 48, 69, 17, 39,
	56, 85, 59, 92, 62, 101, 68, 111, 81, 117, 88, 120, 89, 121, 90, 123, 97,
	128, 98, 129, 99, 131, 106, 136, 109, 139, 114, 141, 119, 142, 127, 153, 135,
	158, 140, 163, 50, 79, 124, 145, 132, 145, 137, 145, 144, 164, 146, 145, 152,
	145, 154, 145, 157, 145, 159, 145, 162, 145, 175, 145, 180, 145, 185, 145,
	187, 145, 203, 145, 220, 145, 224, 145, 227, 145, 229, 145, 234, 145, 241,
	145, 243, 145, 245, 145, 251, 145, 253, 145, 255, 145, 257, 145, 261, 145,
	263, 145, 265, 145, 271, 145, 273, 145, 275, 145, 286, 145, 288, 145, 290,
	145, 292, 145, 298, 145, 300, 145, 302, 145, 313, 145, 328, 145, 330, 145,
	332, 145, 343, 145, 369, 145, 0, 4, 6, 10, 24, 4, 27, 4, 29, 4, 31, 4, 33, 4,
	41, 4, 43, 4, 45, 4, 47, 4, 53, 4, 54, 4, 57, 4, 60, 4, 63, 4, 66, 4, 69, 4,
	93, 4, 94, 4, 102, 4, 103, 4, 108, 4, 124, 4, 132, 4, 137, 4, 151, 4, 152, 4,
	156, 4, 157, 4, 161, 4, 162, 4, 168, 4, 176, 4, 177, 4, 181, 4, 182, 4, 183,
	4, 186, 4, 192, 4, 198, 4, 201, 4, 203, 4, 206, 4, 207, 4, 208, 4, 220, 4,
	223, 4, 226, 4, 227, 4, 229, 4, 232, 4, 234, 4, 237, 4, 242, 4, 243, 4, 245,
	4, 248, 4, 252, 4, 254, 4, 255, 4, 257, 4, 260, 4, 262, 4, 263, 4, 265, 4,
	268, 4, 272, 4, 274, 4, 275, 4, 280, 4, 283, 4, 287, 4, 289, 4, 290, 4, 292,
	4, 295, 4, 299, 4, 301, 4, 302, 4, 307, 4, 310, 4, 314, 4, 322, 4, 325, 4,
	329, 4, 331, 4, 332, 4, 337, 4, 340, 4, 344, 4, 352, 4, 363, 4, 366, 4, 370,
	4, 378, 4, 394, 4, 0, 5, 24, 40, 27, 42, 29, 44, 31, 46, 33, 48, 41, 52, 43,
	58, 45, 61, 47, 65, 53, 83, 54, 84, 57, 87, 60, 96, 63, 105, 66, 110, 69,
	113, 93, 125, 94, 126, 102, 133, 103, 134, 108, 138, 124, 146, 132, 154, 137,
	159, 151, 174, 152, 175, 156, 179, 157, 180, 161, 184, 162, 185, 168, 190,
	176, 194, 177, 196, 181, 200, 182, 202, 183, 205, 186, 209, 192, 214, 198,
	218, 201, 222, 203, 224, 206, 231, 207, 233, 208, 236, 220, 241, 223, 247,
	226, 250, 227, 251, 229, 253, 232, 259, 234, 261, 237, 267, 242, 270, 243,
	271, 245, 273, 248, 278, 252, 282, 254, 285, 255, 286, 257, 288, 260, 294,
	262, 297, 263, 298, 265, 300, 268, 305, 272, 309, 274, 312, 275, 313, 280,
	317, 283, 320, 287, 324, 289, 327, 290, 328, 292, 330, 295, 335, 299, 339,
	301, 342, 302, 343, 307, 347, 310, 350, 314, 354, 322, 358, 325, 361, 329,
	365, 331, 368, 332, 369, 337, 373, 340, 376, 344, 380, 352, 384, 363, 389,
	366, 392, 370, 396, 378, 400, 394, 406, 0, 6, 24, 6, 27, 6, 29, 6, 31, 6, 33,
	6, 41, 6, 43, 6, 45, 6, 47, 6, 53, 6, 54, 6, 57, 6, 60, 6, 63, 6, 66, 6, 69,
	6, 93, 6, 94, 6, 102, 6, 103, 6, 108, 6, 124, 6, 132, 6, 137, 6, 151, 6, 152,
	6, 156, 6, 157, 6, 161, 6, 162, 6, 168, 6, 176, 6, 177, 6, 181, 6, 182, 6,
	183, 6, 186, 6, 192, 6, 198, 6, 201, 6, 203, 6, 206, 6, 207, 6, 208, 6, 220,
	6, 223, 6, 226, 6, 227, 6, 229, 6, 232, 6, 234, 6, 237, 6, 242, 6, 243, 6,
	245, 6, 248, 6, 252, 6, 254, 6, 255, 6, 257, 6, 260, 6, 262, 6, 263, 6, 265,
	6, 268, 6, 272, 6, 274, 6, 275, 6, 280, 6, 283, 6, 287, 6, 289, 6, 290, 6,
	292, 6, 295, 6, 299, 6, 301, 6, 302, 6, 307, 6, 310, 6, 314, 6, 322, 6, 325,
	6, 329, 6, 331, 6, 332, 6, 337, 6, 340, 6, 344, 6, 352, 6, 363, 6, 366, 6,
	370, 6, 378, 6, 394, 6, 15, 35, 49, 74, 49, 75, 49, 76, 49, 77, 49, 78, 7,
	18, 8, 18, 124, 147, 132, 155, 137, 160, 144, 165, 146, 147, 152, 147, 154,
	155, 157, 155, 159, 160, 162, 160, 175, 147, 180, 155, 185, 160, 187, 210,
	203, 225, 220, 225, 224, 225, 227, 225, 229, 225, 234, 225, 241, 225, 243,
	225, 245, 225, 251, 225, 253, 225, 255, 225, 257, 225, 261, 225, 263, 225,
	265, 225, 271, 225, 273, 225, 275, 225, 286, 225, 288, 225, 290, 225, 292,
	225, 298, 225, 300, 225, 302, 225, 313, 225, 328, 225, 330, 225, 332, 225,
	343, 225, 369, 225, 59, 93, 89, 93, 97, 93, 127, 93, 124, 148, 146, 148, 152,
	148, 175, 148, 144, 166, 187, 211, 124, 149, 146, 149, 152, 149, 175, 149,
	144, 167, 124, 150, 146, 150, 152, 150, 175, 150, 148, 170, 149, 171, 150,
	172, 124, 151, 146, 168, 152, 176, 175, 192, 59, 94, 89, 94, 97, 94, 127, 94,
	59, 95, 89, 95, 97, 95, 127, 95, 12, 29, 43, 59, 58, 89, 60, 97, 96, 127, 7,
	19, 8, 19, 62, 102, 98, 102, 106, 102, 135, 102, 132, 156, 154, 177, 157,
	181, 180, 198, 62, 103, 98, 103, 106, 103, 135, 103, 62, 104, 98, 104, 106,
	104, 135, 104, 13, 31, 7, 20, 8, 20, 45, 62, 61, 98, 63, 106, 105, 135, 64,
	108, 137, 161, 159, 182, 162, 186, 185, 207, 203, 226, 220, 242, 224, 248,
	227, 252, 229, 254, 234, 262, 241, 268, 243, 272, 245, 274, 251, 280, 253,
	283, 255, 287, 257, 289, 261, 295, 263, 299, 265, 301, 271, 307, 273, 310,
	275, 314, 286, 322, 288, 325, 290, 329, 292, 331, 298, 337, 300, 340, 302,
	344, 313, 352, 328, 363, 330, 366, 332, 370, 343, 378, 369, 394, 47, 66, 65,
	66, 68, 66, 69, 66, 109, 66, 113, 66, 114, 66, 140, 66, 47, 67, 65, 67, 68,
	112, 69, 67, 109, 112, 113, 67, 114, 112, 140, 112, 14, 33, 47, 68, 65, 109,
	69, 114, 113, 140, 7, 21, 8, 21, 11, 27, 41, 53, 52, 53, 56, 53, 57, 53, 81,
	53, 87, 53, 88, 53, 119, 53, 41, 54, 52, 54, 56, 54, 57, 54, 81, 54, 87, 54,
	88, 54, 119, 54, 41, 55, 52, 55, 56, 86, 57, 55, 81, 86, 87, 55, 88, 86, 119,
	86, 7, 22, 8, 22, 41, 56, 52, 81, 57, 88, 87, 119, 16, 37, 50, 80, 7, 23, 8,
	23, 7, 24, 8, 24, 7, 25, 8, 25, 0, 7, 5, 8, 0, 410,
}

var tmRuleLen = []int8{
	1, 1, 1, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 1, 1, 1, 1, 1, 5, 1, 3,
	1, 2, 1, 2, 1, 2, 1, 7, 6, 6, 5, 6, 5, 5, 4, 2, 2, 1, 2, 1, 2, 0, 7, 6, 6, 5,
	1, 1, 7, 6, 6, 5, 6, 5, 5, 4, 2, 2, 1, 2, 1, 7, 6, 6, 5, 2, 0, 1, 1, 1, 14,
	13, 13, 12, 13, 12, 12, 11, 10, 9, 13, 12, 12, 11, 12, 11, 11, 10, 9, 8, 13,
	12, 12, 11, 12, 11, 11, 10, 9, 8, 12, 11, 11, 10, 11, 10, 10, 9, 8, 7, 13,
	12, 12, 11, 12, 11, 11, 10, 9, 8, 12, 11, 11, 10, 11, 10, 10, 9, 8, 7, 12,
	11, 11, 10, 11, 10, 10, 9, 8, 7, 11, 10, 10, 9, 10, 9, 9, 8, 7, 6, 1, 2, 1,
	2, 1, 7, 6, 6, 5, 1, 1, 3, 2, 1, 2, 7, 6, 6, 5, 2, 1, 1, 1, 5, 1, 1, 1, 1, 1,
	1, 2, 2, 1, 2, 2, 0, 2, 1,
}

var tmRuleSymbol = []int32{
	31, 31, 32, 33, 33, 33, 33, 34, 35, 36, 37, 38, 39, 39, 39, 39, 40, 41, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 50, 50, 50, 50, 50, 51, 51, 51, 51, 51,
	51, 51, 51, 51, 51, 52, 52, 53, 54, 54, 55, 55, 55, 55, 56, 57, 58, 58, 58,
	58, 58, 58, 58, 58, 58, 58, 59, 59, 60, 61, 61, 61, 61, 62, 62, 63, 64, 65,
	66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66,
	66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66,
	66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66,
	66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66,
	66, 66, 66, 66, 67, 67, 68, 69, 69, 70, 70, 70, 70, 71, 72, 73, 73, 74, 74,
	75, 75, 75, 75, 76, 76, 77, 78, 79, 80, 80, 80, 80, 80, 80, 80, 80, 81, 81,
	82, 82, 83, 83,
}

var tmRuleType = [...]NodeType{
	0,            // Comment : BlockComment
	0,            // Comment : EolComment
	0,            // WhiteSpace : WhiteSpace$1
	0,            // WhiteSpace$1 : WhiteSpace$1 Comment
	0,            // WhiteSpace$1 : WhiteSpace$1 WS
	0,            // WhiteSpace$1 : Comment
	0,            // WhiteSpace$1 : WS
	OptionName,   // OptionName : Name
	OptionString, // OptionString : StringLiteral
	OptionBool,   // OptionBool : BoolLiteral
	OptionInt,    // OptionInt : IntLiteral
	OptionFloat,  // OptionFloat : FloatLiteral
	0,            // OptionValue : OptionString
	0,            // OptionValue : OptionBool
	0,            // OptionValue : OptionInt
	0,            // OptionValue : OptionFloat
	0,            // Option : 'option' OptionName '=' OptionValue ';'
	0,            // TypeRef : Name
	0,            // TypeRef : Primitive
	FieldName,    // FieldName : Name
	ScalarType,   // ScalarType : TypeRef
	MapKeyType,   // MapKeyType : Primitive
	MapValueType, // MapValueType : TypeRef
	0,            // MapType : '[' MapKeyType ':' MapValueType ']'
	ListElement,  // ListElementType : TypeRef
	0,            // ListType : '[' ListElementType ']'
	FieldMod,     // FieldMod : Modifier
	0,            // FieldType : ScalarType FieldMod
	0,            // FieldType : ScalarType
	0,            // FieldType : MapType FieldMod
	0,            // FieldType : MapType
	0,            // FieldType : ListType FieldMod
	0,            // FieldType : ListType
	0,            // FieldDef : FieldName WhiteSpace ':' WhiteSpace FieldType WhiteSpace ';'
	0,            // FieldDef : FieldName WhiteSpace ':' WhiteSpace FieldType ';'
	0,            // FieldDef : FieldName WhiteSpace ':' FieldType WhiteSpace ';'
	0,            // FieldDef : FieldName WhiteSpace ':' FieldType ';'
	0,            // FieldDef : FieldName ':' WhiteSpace FieldType WhiteSpace ';'
	0,            // FieldDef : FieldName ':' WhiteSpace FieldType ';'
	0,            // FieldDef : FieldName ':' FieldType WhiteSpace ';'
	0,            // FieldDef : FieldName ':' FieldType ';'
	0,            // FieldDef : error ';'
	0,            // FieldDef : error '}'
	0,            // FieldDefList : FieldDef
	0,            // FieldDefList : FieldDef WhiteSpace
	MessageName,  // MessageName : Name
	0,            // FieldDefList_optlist : FieldDefList_optlist FieldDefList
	0,            // FieldDefList_optlist :
	0,            // Message : 'message' MessageName WhiteSpace '{' WhiteSpace FieldDefList_optlist '}'
	0,            // Message : 'message' MessageName WhiteSpace '{' FieldDefList_optlist '}'
	0,            // Message : 'message' MessageName '{' WhiteSpace FieldDefList_optlist '}'
	0,            // Message : 'message' MessageName '{' FieldDefList_optlist '}'
	FieldName,    // UnionMemberName : Name
	ScalarType,   // UnionMemberType : TypeRef
	0,            // UnionMember : UnionMemberName WhiteSpace ':' WhiteSpace UnionMemberType WhiteSpace ';'
	0,            // UnionMember : UnionMemberName WhiteSpace ':' WhiteSpace UnionMemberType ';'
	0,            // UnionMember : UnionMemberName WhiteSpace ':' UnionMemberType WhiteSpace ';'
	0,            // UnionMember : UnionMemberName WhiteSpace ':' UnionMemberType ';'
	0,            // UnionMember : UnionMemberName ':' WhiteSpace UnionMemberType WhiteSpace ';'
	0,            // UnionMember : UnionMemberName ':' WhiteSpace UnionMemberType ';'
	0,            // UnionMember : UnionMemberName ':' UnionMemberType WhiteSpace ';'
	0,            // UnionMember : UnionMemberName ':' UnionMemberType ';'
	0,            // UnionMember : error ';'
	0,            // UnionMember : error '}'
	0,            // UnionMemberList : UnionMember
	0,            // UnionMemberList : UnionMember WhiteSpace
	UnionName,    // UnionName : Name
	0,            // Union : 'union' UnionName WhiteSpace '{' WhiteSpace UnionMemberList_optlist '}'
	0,            // Union : 'union' UnionName WhiteSpace '{' UnionMemberList_optlist '}'
	0,            // Union : 'union' UnionName '{' WhiteSpace UnionMemberList_optlist '}'
	0,            // Union : 'union' UnionName '{' UnionMemberList_optlist '}'
	0,            // UnionMemberList_optlist : UnionMemberList_optlist UnionMemberList
	0,            // UnionMemberList_optlist :
	RpcName,      // RpcName : Name
	RpcRequest,   // RpcRequest : TypeRef
	RpcResponse,  // RpcResponse : TypeRef
	0,            // Method : 'rpc' RpcName WhiteSpace '(' WhiteSpace RpcRequest WhiteSpace ')' WhiteSpace ':' WhiteSpace RpcResponse WhiteSpace ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' WhiteSpace RpcRequest WhiteSpace ')' WhiteSpace ':' WhiteSpace RpcResponse ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' WhiteSpace RpcRequest WhiteSpace ')' WhiteSpace ':' RpcResponse WhiteSpace ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' WhiteSpace RpcRequest WhiteSpace ')' WhiteSpace ':' RpcResponse ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' WhiteSpace RpcRequest WhiteSpace ')' ':' WhiteSpace RpcResponse WhiteSpace ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' WhiteSpace RpcRequest WhiteSpace ')' ':' WhiteSpace RpcResponse ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' WhiteSpace RpcRequest WhiteSpace ')' ':' RpcResponse WhiteSpace ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' WhiteSpace RpcRequest WhiteSpace ')' ':' RpcResponse ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' WhiteSpace RpcRequest WhiteSpace ')' WhiteSpace ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' WhiteSpace RpcRequest WhiteSpace ')' ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' WhiteSpace RpcRequest ')' WhiteSpace ':' WhiteSpace RpcResponse WhiteSpace ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' WhiteSpace RpcRequest ')' WhiteSpace ':' WhiteSpace RpcResponse ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' WhiteSpace RpcRequest ')' WhiteSpace ':' RpcResponse WhiteSpace ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' WhiteSpace RpcRequest ')' WhiteSpace ':' RpcResponse ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' WhiteSpace RpcRequest ')' ':' WhiteSpace RpcResponse WhiteSpace ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' WhiteSpace RpcRequest ')' ':' WhiteSpace RpcResponse ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' WhiteSpace RpcRequest ')' ':' RpcResponse WhiteSpace ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' WhiteSpace RpcRequest ')' ':' RpcResponse ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' WhiteSpace RpcRequest ')' WhiteSpace ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' WhiteSpace RpcRequest ')' ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' RpcRequest WhiteSpace ')' WhiteSpace ':' WhiteSpace RpcResponse WhiteSpace ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' RpcRequest WhiteSpace ')' WhiteSpace ':' WhiteSpace RpcResponse ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' RpcRequest WhiteSpace ')' WhiteSpace ':' RpcResponse WhiteSpace ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' RpcRequest WhiteSpace ')' WhiteSpace ':' RpcResponse ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' RpcRequest WhiteSpace ')' ':' WhiteSpace RpcResponse WhiteSpace ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' RpcRequest WhiteSpace ')' ':' WhiteSpace RpcResponse ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' RpcRequest WhiteSpace ')' ':' RpcResponse WhiteSpace ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' RpcRequest WhiteSpace ')' ':' RpcResponse ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' RpcRequest WhiteSpace ')' WhiteSpace ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' RpcRequest WhiteSpace ')' ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' RpcRequest ')' WhiteSpace ':' WhiteSpace RpcResponse WhiteSpace ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' RpcRequest ')' WhiteSpace ':' WhiteSpace RpcResponse ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' RpcRequest ')' WhiteSpace ':' RpcResponse WhiteSpace ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' RpcRequest ')' WhiteSpace ':' RpcResponse ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' RpcRequest ')' ':' WhiteSpace RpcResponse WhiteSpace ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' RpcRequest ')' ':' WhiteSpace RpcResponse ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' RpcRequest ')' ':' RpcResponse WhiteSpace ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' RpcRequest ')' ':' RpcResponse ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' RpcRequest ')' WhiteSpace ';'
	0,            // Method : 'rpc' RpcName WhiteSpace '(' RpcRequest ')' ';'
	0,            // Method : 'rpc' RpcName '(' WhiteSpace RpcRequest WhiteSpace ')' WhiteSpace ':' WhiteSpace RpcResponse WhiteSpace ';'
	0,            // Method : 'rpc' RpcName '(' WhiteSpace RpcRequest WhiteSpace ')' WhiteSpace ':' WhiteSpace RpcResponse ';'
	0,            // Method : 'rpc' RpcName '(' WhiteSpace RpcRequest WhiteSpace ')' WhiteSpace ':' RpcResponse WhiteSpace ';'
	0,            // Method : 'rpc' RpcName '(' WhiteSpace RpcRequest WhiteSpace ')' WhiteSpace ':' RpcResponse ';'
	0,            // Method : 'rpc' RpcName '(' WhiteSpace RpcRequest WhiteSpace ')' ':' WhiteSpace RpcResponse WhiteSpace ';'
	0,            // Method : 'rpc' RpcName '(' WhiteSpace RpcRequest WhiteSpace ')' ':' WhiteSpace RpcResponse ';'
	0,            // Method : 'rpc' RpcName '(' WhiteSpace RpcRequest WhiteSpace ')' ':' RpcResponse WhiteSpace ';'
	0,            // Method : 'rpc' RpcName '(' WhiteSpace RpcRequest WhiteSpace ')' ':' RpcResponse ';'
	0,            // Method : 'rpc' RpcName '(' WhiteSpace RpcRequest WhiteSpace ')' WhiteSpace ';'
	0,            // Method : 'rpc' RpcName '(' WhiteSpace RpcRequest WhiteSpace ')' ';'
	0,            // Method : 'rpc' RpcName '(' WhiteSpace RpcRequest ')' WhiteSpace ':' WhiteSpace RpcResponse WhiteSpace ';'
	0,            // Method : 'rpc' RpcName '(' WhiteSpace RpcRequest ')' WhiteSpace ':' WhiteSpace RpcResponse ';'
	0,            // Method : 'rpc' RpcName '(' WhiteSpace RpcRequest ')' WhiteSpace ':' RpcResponse WhiteSpace ';'
	0,            // Method : 'rpc' RpcName '(' WhiteSpace RpcRequest ')' WhiteSpace ':' RpcResponse ';'
	0,            // Method : 'rpc' RpcName '(' WhiteSpace RpcRequest ')' ':' WhiteSpace RpcResponse WhiteSpace ';'
	0,            // Method : 'rpc' RpcName '(' WhiteSpace RpcRequest ')' ':' WhiteSpace RpcResponse ';'
	0,            // Method : 'rpc' RpcName '(' WhiteSpace RpcRequest ')' ':' RpcResponse WhiteSpace ';'
	0,            // Method : 'rpc' RpcName '(' WhiteSpace RpcRequest ')' ':' RpcResponse ';'
	0,            // Method : 'rpc' RpcName '(' WhiteSpace RpcRequest ')' WhiteSpace ';'
	0,            // Method : 'rpc' RpcName '(' WhiteSpace RpcRequest ')' ';'
	0,            // Method : 'rpc' RpcName '(' RpcRequest WhiteSpace ')' WhiteSpace ':' WhiteSpace RpcResponse WhiteSpace ';'
	0,            // Method : 'rpc' RpcName '(' RpcRequest WhiteSpace ')' WhiteSpace ':' WhiteSpace RpcResponse ';'
	0,            // Method : 'rpc' RpcName '(' RpcRequest WhiteSpace ')' WhiteSpace ':' RpcResponse WhiteSpace ';'
	0,            // Method : 'rpc' RpcName '(' RpcRequest WhiteSpace ')' WhiteSpace ':' RpcResponse ';'
	0,            // Method : 'rpc' RpcName '(' RpcRequest WhiteSpace ')' ':' WhiteSpace RpcResponse WhiteSpace ';'
	0,            // Method : 'rpc' RpcName '(' RpcRequest WhiteSpace ')' ':' WhiteSpace RpcResponse ';'
	0,            // Method : 'rpc' RpcName '(' RpcRequest WhiteSpace ')' ':' RpcResponse WhiteSpace ';'
	0,            // Method : 'rpc' RpcName '(' RpcRequest WhiteSpace ')' ':' RpcResponse ';'
	0,            // Method : 'rpc' RpcName '(' RpcRequest WhiteSpace ')' WhiteSpace ';'
	0,            // Method : 'rpc' RpcName '(' RpcRequest WhiteSpace ')' ';'
	0,            // Method : 'rpc' RpcName '(' RpcRequest ')' WhiteSpace ':' WhiteSpace RpcResponse WhiteSpace ';'
	0,            // Method : 'rpc' RpcName '(' RpcRequest ')' WhiteSpace ':' WhiteSpace RpcResponse ';'
	0,            // Method : 'rpc' RpcName '(' RpcRequest ')' WhiteSpace ':' RpcResponse WhiteSpace ';'
	0,            // Method : 'rpc' RpcName '(' RpcRequest ')' WhiteSpace ':' RpcResponse ';'
	0,            // Method : 'rpc' RpcName '(' RpcRequest ')' ':' WhiteSpace RpcResponse WhiteSpace ';'
	0,            // Method : 'rpc' RpcName '(' RpcRequest ')' ':' WhiteSpace RpcResponse ';'
	0,            // Method : 'rpc' RpcName '(' RpcRequest ')' ':' RpcResponse WhiteSpace ';'
	0,            // Method : 'rpc' RpcName '(' RpcRequest ')' ':' RpcResponse ';'
	0,            // Method : 'rpc' RpcName '(' RpcRequest ')' WhiteSpace ';'
	0,            // Method : 'rpc' RpcName '(' RpcRequest ')' ';'
	0,            // MethodList : Method
	0,            // MethodList : Method WhiteSpace
	ServiceName,  // ServiceName : Name
	0,            // MethodList_list : MethodList_list MethodList
	0,            // MethodList_list : MethodList
	0,            // Service : 'service' ServiceName WhiteSpace '{' WhiteSpace MethodList_list '}'
	0,            // Service : 'service' ServiceName WhiteSpace '{' MethodList_list '}'
	0,            // Service : 'service' ServiceName '{' WhiteSpace MethodList_list '}'
	0,            // Service : 'service' ServiceName '{' MethodList_list '}'
	EnumName,     // EnumName : Name
	EnumValue,    // EnumValue : Value
	0,            // EnumField : EnumValue WhiteSpace ';'
	0,            // EnumField : EnumValue ';'
	0,            // EnumValueList : EnumField
	0,            // EnumValueList : EnumField WhiteSpace
	0,            // Enum : 'enum' EnumName WhiteSpace '{' WhiteSpace EnumValueList_list '}'
	0,            // Enum : 'enum' EnumName WhiteSpace '{' EnumValueList_list '}'
	0,            // Enum : 'enum' EnumName '{' WhiteSpace EnumValueList_list '}'
	0,            // Enum : 'enum' EnumName '{' EnumValueList_list '}'
	0,            // EnumValueList_list : EnumValueList_list EnumValueList
	0,            // EnumValueList_list : EnumValueList
	AliasName,    // AliasName : Name
	AliasType,    // AliasType : Primitive
	0,            // Alias : 'alias' AliasName '->' AliasType ';'
	0,            // Definition : Message
	0,            // Definition : Union
	0,            // Definition : Service
	0,            // Definition : Enum
	0,            // Definition : Option
	0,            // Definition : Alias
	0,            // Definition : error ';'
	0,            // Definition : error '}'
	0,            // DefinitionList : Definition
	0,            // DefinitionList : Definition WhiteSpace
	0,            // DefinitionList_optlist : DefinitionList_optlist DefinitionList
	0,            // DefinitionList_optlist :
	0,            // ServoFile : WhiteSpace DefinitionList_optlist
	0,            // ServoFile : DefinitionList_optlist
}

// set(follow ERROR) = SEMICOLON, RBRACE
var afterErr = []token.Type{
	22, 29,
}
