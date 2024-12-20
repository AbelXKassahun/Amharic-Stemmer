package miscellaneous

import (
	"fmt"
	"github.com/AbelXKassahun/Amharic-Stemmer/stemmer"
	"strings"
)

//	var amh_prefix_inAmh := []string{
//		"ስለማይ",
//		"የማት",
//		"ዕንደ",
//		"የተ",
//		"በማ",
//		"በተ",
//		"ዐል",
//		"ስለ",
//		"መስ",
//		"ዕየ",
//		"ዐስ",
//		"ዐት",
//		"ዐን",
//		"ዐይ",
//		"ያል",
//		"ሳት",
//		"ሳን",
//		"ሳይ",
//		"ሳል",
//		"ያስ",
//		"የ",
//		"ለ",
//		"ከ",
//		"እን",
//		"ዕን",
//		"ዐል",
//		"ይ",
//		"ት",
//		"አ",
//		"እ",
//	}
func GetPrefixListFromAmhList() {
	amh_prefix_inCV := "ስልኧምኣይ|ይኧምኣት|ዕንድኧ|ይኧትኧ|ብኧምኣ|ብኧትኧ|ዕኧል|ስልኧ|ምኧስ|ዕይኧ|ዕኧስ|ዕኧት|ዕኧን|ዕኧይ|ይኣል|ስኣት|ስኣን|ስኣይ|ስኣል|ይኣስ|ይኧ|ልኧ|ክኧ|እን|ዕን|ዐል|ይ|ት|አ|እ"

	for _, val := range strings.Split(amh_prefix_inCV, "|") {
		pref, _ := stemmer.ToEng(val)
		pref_inAmh, _ := stemmer.ToAmh(pref)
		fmt.Printf("%v, %v, %v\n", stemmer.EngToArray(pref), val, pref_inAmh)
	}
}

////
////func GetEngSuffixListFromAmhCVFormList() {
////	ar := "ኢዕኧልኧሽ|ኣውኢው|ኣችኧውኣል|ኧችኣት|ኧችኣችህኡ|ኧችኣችኧው|ኣልኧህኡ|" +
////		"ኣውኦች|ኣልኧህ|ኣልኧሽ|ኣልችህኡ|ኣልኣልኧች|ብኣችኧውስ|ብኣችኧው|ኣችኧውን|ኣልኧች|" +
////		"ኣልኧን|ኣልኣችህኡ|ኣችህኡን|ኣችህኡ|ኣችህኡት|ውኦችንንኣ|ውኦችን|ኣችኧው|ውኦችኡን|" +
////		"ውኦችኡ|ኧውንኣ|ኦችኡን|ኦውኦች|ኧኝኣንኧትም|ኧኝኣንኣ|ኧኝኣንኧት|ኧኝኣን|ኧኝኣውም|" +
////		"ኧኝኣው|ኝኣውኣ|ብኧትን|ኣችህኡም|ኦውኣ|ኧችው|ኧችኡ|ኤችኡ|ንኧው|ንኧት|ኣልኡ|ኣችን|" +
////		"ክኡም|ክኡት|ክኧው|ኧችን|ኧችም|ኧችህ|ኧችሽ|ኧችን|ኧችው|ይኡሽን|ይኡሽ|ኧውኢ" +
////		"|ኦችንንኣ|ኣውኢ|ብኧት|ኦች|ኦችኡ|ውኦን|ኧኝኣ|ኝኣውን|ኝኣው|ኦችን|ኣል|ኧም|ሽው|" +
////		"ክም|ኧው|ትም|ውኦ|ውም|ውን|ንም|ሽን|ኣች|ኡት|ኢት|ክኡ|ኤ|ህ|ሽ|ኡ|ሽ|ክ|ኧ|ኧች|" +
////		"ኡን|ን|ም|ንኣ|ው"
////	ar2 := []string{
////		"ኢዐለሽ",
////		"ኣዊው",
////		"ኣቸዋል",
////		"አቻት",
////		"አቻችሁ",
////		"አቻቸው",
////		"ኣለሁ",
////		"ኣዎች",
////		"ኣለህ",
////		"ኣለሽ",
////		"ኣልችሁ",
////		"ኣላለች",
////		"ባቸውስ",
////		"ባቸው",
////		"ኣቸውን",
////		"ኣለች",
////		"ኣለን",
////		"ኣላችሁ",
////		"ኣችሁን",
////		"ኣችሁ",
////		"ኣችሁት",
////		"ዎችንና",
////		"ዋችን",
////		"ኣቸው",
////		"ዎቹን",
////		"ዎቹ",
////		"አውና",
////		"ኦቹን",
////		"ኦዎች",
////		"አኛነትም",
////		"አኛና",
////		"አኛነት",
////		"አኛን",
////		"አኛውም",
////		"አኛው",
////		"ኛዋ",
////		"በትን",
////		"ኣችሁም",
////		"ኦዋ",
////		"አችው",
////		"አቹ",
////		"ኤቹ",
////		"ነው",
////		"ነት",
////		"ኣሉ",
////		"ኣችን",
////		"ኩም",
////		"ኩት",
////		"ከው",
////		"አችን",
////		"አችም",
////		"አችህ",
////		"አችሽ",
////		"አችን",
////		"አችው",
////		"ዩሽን",
////		"ዩሽ",
////		"አዊ",
////		"ኦችንና",
////		"ኣዊ",
////		"በት",
////		"ኦች",
////		"ኦቹ",
////		"ዎን",
////		"አኛ",
////		"ኛውን",
////		"ኛው",
////		"ኦችን",
////		"ኣል",
////		"አም",
////		"ሽው",
////		"ክም",
////		"አው",
////		"ትም",
////		"ዎ",
////		"ውም",
////		"ውን",
////		"ንም",
////		"ሽን",
////		"ኣች",
////		"ኡት",
////		"ኢት",
////		"ኩ",
////		"ኤ",
////		"ህ",
////		"ሽ",
////		"ኡ",
////		"ሽ",
////		"ክ",
////		"አ",
////		"አች",
////		"ኡን",
////		"ን",
////		"ም",
////		"ና",
////		"ው",
////	}
////	suf := strings.Split(ar, "|")
////	fmt.Printf("%v, %v\n", len(suf), len(ar2))
////	for i, val := range ar2 {
////		pref, _ := stemmer.ToEng(val)
////		fmt.Printf("%v, , %v, %v\n", strings.Join(pref, "|"), suf[i], val)
////	}
////}
