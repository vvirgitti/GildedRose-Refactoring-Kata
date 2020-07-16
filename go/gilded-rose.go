package main

type Item struct {
	name            string
	sellIn, quality int
}

const agedBrie = "Aged Brie"
const backstage = "Backstage passes to a TAFKAL80ETC concert"
const sulfuras = "Sulfuras, Hand of Ragnaros"

func isAgedBrie(itemName string) bool {
	return itemName == agedBrie
}

func isBackstage(itemName string) bool {
	return itemName == backstage
}

func isSulfuras(itemName string) bool {
	return itemName == sulfuras
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		itemName := items[i].name

		if !isAgedBrie(itemName) && !isBackstage(itemName) {
			if items[i].quality > 0 && !isSulfuras(itemName) {
				items[i].quality = items[i].quality - 1
			}
		} else {
			if items[i].quality < 50 {
				items[i].quality = items[i].quality + 1
				if isBackstage(itemName) && items[i].sellIn < 11 && items[i].quality < 50 {
					items[i].quality = items[i].quality + 1
					if items[i].sellIn < 6 && items[i].quality < 50 {
						items[i].quality = items[i].quality + 1
					}
				}
			}
		}

		if !isSulfuras(itemName) {
			items[i].sellIn = items[i].sellIn - 1
		}

		if items[i].sellIn < 0 {
			if !isAgedBrie(itemName) {
				if !isBackstage(itemName) && items[i].quality > 0 && !isSulfuras(itemName) {
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
