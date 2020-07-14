// Code generated by "stringer -type=Item -output=token_string.go"; DO NOT EDIT.

package tl

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ItemIllegal-0]
	_ = x[ItemWhitespace-1]
	_ = x[ItemEOF-2]
	_ = x[ItemUnderscore-3]
	_ = x[ItemColon-4]
	_ = x[ItemSemicolon-5]
	_ = x[ItemOpenPar-6]
	_ = x[ItemClosePar-7]
	_ = x[ItemOpenBracket-8]
	_ = x[ItemCloseBracket-9]
	_ = x[ItemOpenBrace-10]
	_ = x[ItemCloseBrace-11]
	_ = x[ItemLeftAngle-12]
	_ = x[ItemRightAngle-13]
	_ = x[ItemTripleMinus-14]
	_ = x[ItemEquals-15]
	_ = x[ItemHash-16]
	_ = x[ItemExclMark-17]
	_ = x[ItemQuestionMark-18]
	_ = x[ItemPercent-19]
	_ = x[ItemPlus-20]
	_ = x[ItemComma-21]
	_ = x[ItemDot-22]
	_ = x[ItemAsterisk-23]
	_ = x[ItemNatConst-24]
	_ = x[ItemLowerIdent-25]
	_ = x[ItemUpperIdent-26]
	_ = x[ItemFinal-27]
	_ = x[ItemNew-28]
	_ = x[ItemEmpty-29]
}

const _Item_name = "ItemIllegalItemWhitespaceItemEOFItemUnderscoreItemColonItemSemicolonItemOpenParItemCloseParItemOpenBracketItemCloseBracketItemOpenBraceItemCloseBraceItemLeftAngleItemRightAngleItemTripleMinusItemEqualsItemHashItemExclMarkItemQuestionMarkItemPercentItemPlusItemCommaItemDotItemAsteriskItemNatConstItemLowerIdentItemUpperIdentItemFinalItemNewItemEmpty"

var _Item_index = [...]uint16{0, 11, 25, 32, 46, 55, 68, 79, 91, 106, 122, 135, 149, 162, 176, 191, 201, 209, 221, 237, 248, 256, 265, 272, 284, 296, 310, 324, 333, 340, 349}

func (i Item) String() string {
	if i < 0 || i >= Item(len(_Item_index)-1) {
		return "Item(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Item_name[_Item_index[i]:_Item_index[i+1]]
}
