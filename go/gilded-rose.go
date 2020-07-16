package main

type Item struct {
	name            string
	sellIn, quality int
}

//func IsSpecialItem(itemName string) bool {
//
//}

func UpdateQuality(items []*Item) {
	agedBrie := "Aged Brie"
	backstage := "Backstage passes to a TAFKAL80ETC concert"
	sulfuras := "Sulfuras, Hand of Ragnaros"

	for i := 0; i < len(items); i++ {

		if items[i].name != agedBrie && items[i].name != backstage {
			if items[i].quality > 0 && items[i].name != sulfuras {
				items[i].quality = items[i].quality - 1
			}
		} else {
			if items[i].quality < 50 {
				items[i].quality = items[i].quality + 1
				if items[i].name == backstage && items[i].sellIn < 11 && items[i].quality < 50 {
					items[i].quality = items[i].quality + 1
					if items[i].sellIn < 6 && items[i].quality < 50 {
						items[i].quality = items[i].quality + 1
					}
				}
			}
		}

		if items[i].name != sulfuras {
			items[i].sellIn = items[i].sellIn - 1
		}

		if items[i].sellIn < 0 {
			if items[i].name != agedBrie {
				if items[i].name != backstage && items[i].quality > 0 && items[i].name != sulfuras {
					items[i].quality = items[i].quality - 1
				} else {
					items[i].quality = items[i].quality - items[i].quality
				}
			} else {
				if items[i].quality < 50 {
					items[i].quality = items[i].quality + 1
				}
			}
		}
	}
}
