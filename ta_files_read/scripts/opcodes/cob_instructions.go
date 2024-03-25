package opcodes

// All the instruction opcodes for the COB script machine.
const (
	// no args
	CI_SLEEP               = 0x10013000
	CI_ALLOC_LOCAL_VAR     = 0x10022000
	CI_ADD                 = 0x10031000
	CI_SUB                 = 0x10032000
	CI_MUL                 = 0x10033000
	CI_DIV                 = 0x10034000
	CI_BITWISE_OR          = 0x10036000
	CI_RAND                = 0x10041000
	CI_GET_VALUE           = 0x10042000
	CI_GET_VALUE_WITH_ARGS = 0x10043000
	CI_CMP_LESS            = 0x10051000
	CI_CMP_LEQ             = 0x10052000
	CI_CMP_GREATER         = 0x10053000
	CI_CMP_GEQ             = 0x10054000
	CI_CMP_EQ              = 0x10055000
	CI_CMP_NEQ             = 0x10056000
	CI_LOGICAL_AND         = 0x10057000
	CI_LOGICAL_OR          = 0x10058000
	CI_LOGICAL_XOR         = 0x10059000
	CI_NEG                 = 0x1005A000
	CI_SETSIGMASK          = 0x10068000
	CI_SIGNAL              = 0x10067000
	CI_RETURN              = 0x10065000
	CI_SET_VALUE           = 0x10082000
	CI_ATTACH_UNIT         = 0x10083000
	CI_DROP_UNIT           = 0x10084000

	// 1 arg
	CI_SHOW_OBJECT         = 0x10005000
	CI_HIDE_OBJECT         = 0x10006000
	CI_CACHE               = 0x10007000
	CI_DONTCACHE           = 0x10008000
	CI_SHADE               = 0x1000D000
	CI_DONTSHADE           = 0x1000E000
	CI_EMIT_SFX_FROM_PIECE = 0x1000F000
	CI_PUSH_CONST          = 0x10021001
	CI_PUSH_LOCAL_VAR      = 0x10021002
	CI_PUSH_STATIC_VAR     = 0x10021004
	CI_POP_LOCAL_VAR       = 0x10023002
	CI_POP_STATIC_VAR      = 0x10023004
	CI_JMP                 = 0x10064000
	CI_JMP_IF_FALSE        = 0x10066000
	CI_EXPLODE_PIECE       = 0x10071000

	// 2 args
	CI_MOVE_OBJECT      = 0x10001000
	CI_ROTATE_OBJECT    = 0x10002000
	CI_SPIN_OBJECT      = 0x10003000
	CI_STOP_SPIN_OBJECT = 0x10004000
	CI_MOVE_NOW         = 0x1000B000
	CI_TURN_NOW         = 0x1000C000
	CI_WAIT_FOR_TURN    = 0x10011000
	CI_WAIT_FOR_MOVE    = 0x10012000
	CI_START_SCRIPT     = 0x10061000
	CI_CALL_SCRIPT      = 0x10062000
)
