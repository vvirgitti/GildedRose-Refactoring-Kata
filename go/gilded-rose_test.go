package main

import "testing"

func Test_Foo(t *testing.T) {
	t.Run("normal items have sellin and quality decrease by 1", func(t *testing.T) {
		var items = []*Item{
			&Item{"+5 Dexterity Vest", 10, 20},
		}

		UpdateQuality(items)

		if items[0].sellIn != 9 {
			t.Errorf("Sellin: Expected %d but got %d ", 9, items[0].sellIn)
		}

		if items[0].quality != 19 {
			t.Errorf("Quality: Expected %d but got %d ", 19, items[0].sellIn)
		}
	})

	t.Run("even conjured items decrease in sellin and quality by 1 for now", func(t *testing.T) {
		var items = []*Item{
			&Item{"Conjured Mana Cake", 3, 6},
		}

		UpdateQuality(items)

		if items[0].sellIn != 2 {
			t.Errorf("Sellin: Expected %d but got %d ", 2, items[0].sellIn)
		}

		if items[0].quality != 5 {
			t.Errorf("Quality: Expected %d but got %d ", 5, items[0].sellIn)
		}
	})

	t.Run("sellin is passed, quality degrades faster", func(t *testing.T) {
		var items = []*Item{
			&Item{"+5 Dexterity Vest", -1, 4},
		}

		UpdateQuality(items)

		if items[0].sellIn != -2 {
			t.Errorf("Sellin: Expected %d but got %d ", -2, items[0].sellIn)
		}

		if items[0].quality != 2 {
			t.Errorf("Quality: Expected %d but got %d ", 2, items[0].sellIn)
		}
	})

	t.Run("sulfuras items are magical and never change", func(t *testing.T) {
		var items = []*Item{
			&Item{"Sulfuras, Hand of Ragnaros", 2, 80},
			&Item{"Sulfuras, Hand of Ragnaros", 2, 5},
		}

		UpdateQuality(items)

		t.Run("quality never goes over 80", func(t *testing.T) {
			if items[0].quality != 80 {
				t.Errorf("Quality: Expected %d but got %d ", 80, items[0].quality)
			}

			if items[1].quality != 5 {
				t.Errorf("Quality: Expected %d but got %d ", 5, items[1].quality)
			}
		})

		t.Run("never has to be sold", func(t *testing.T) {
			if items[0].sellIn != 2 {
				t.Errorf("Sellin: Expected %d but got %d ", 2, items[0].sellIn)
			}

			if items[1].sellIn != 2 {
				t.Errorf("Sellin: Expected %d but got %d ", 2, items[1].sellIn)
			}

		})

	})

	t.Run("a quality can never be negative", func(t *testing.T) {
		var items = []*Item{
			&Item{"Blah", -10, 0},
		}

		UpdateQuality(items)

		if items[0].quality != 0 {
			t.Errorf("Quality: Expected %d but got %d ", 0, items[0].sellIn)
		}
	})

	t.Run("Aged brie increases in quality, but never over 50", func(t *testing.T) {
		var items = []*Item{
			&Item{"Aged Brie", 2, 0},
			&Item{"Aged Brie", -1, 50},
		}

		UpdateQuality(items)

		if items[0].sellIn != 1 {
			t.Errorf("Sellin: Expected %d but got %d ", 1, items[0].sellIn)
		}

		if items[0].quality != 1 {
			t.Errorf("Quality: Expected %d but got %d ", 1, items[0].quality)
		}

		if items[1].sellIn != -2 {
			t.Errorf("Sellin: Expected %d but got %d ", -2, items[0].sellIn)
		}

		if items[1].quality != 50 {
			t.Errorf("Quality: Expected %d but got %d ", 50, items[0].quality)
		}
	})

	t.Run("Backstage passes", func(t *testing.T) {
		var items = []*Item{
			&Item{"Backstage passes to a TAFKAL80ETC concert", 10, 4},
			&Item{"Backstage passes to a TAFKAL80ETC concert", 5, 4},
			&Item{"Backstage passes to a TAFKAL80ETC concert", 0, 10},
			&Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		}

		UpdateQuality(items)

		t.Run("increase by 2, 10 days or less before the concert", func(t *testing.T) {
			if items[0].sellIn != 9 {
				t.Errorf("Sellin: Expected %d but got %d ", 9, items[0].sellIn)
			}

			if items[0].quality != 6 {
				t.Errorf("Quality: Expected %d but got %d ", 6, items[0].quality)
			}
		})

		t.Run("increase by 3, 5 days or less before the concert", func(t *testing.T) {
			if items[1].sellIn != 4 {
				t.Errorf("Sellin: Expected %d but got %d ", 4, items[1].sellIn)
			}

			if items[1].quality != 7 {
				t.Errorf("Quality: Expected %d but got %d ", 7, items[1].quality)
			}
		})

		t.Run("quality drops to 0 after the concert", func(t *testing.T) {
			if items[2].sellIn != -1 {
				t.Errorf("Sellin: Expected %d but got %d ", -1, items[2].sellIn)
			}

			if items[2].quality != 0 {
				t.Errorf("Quality: Expected %d but got %d ", 0, items[2].quality)
			}
		})

		t.Run("before 10 days, quality increases", func(t *testing.T) {
			if items[3].sellIn != 14 {
				t.Errorf("Sellin: Expected %d but got %d ", 14, items[3].sellIn)
			}

			if items[3].quality != 21 {
				t.Errorf("Quality: Expected %d but got %d ", 21, items[3].quality)
			}
		})
	})

	t.Run("new requirement for the conjured items", func(t *testing.T) {
		t.Skip()
		t.Run("conjured items degrade in quality twice as fast", func(t *testing.T) {
			// New requirement
			t.Skip()
			var items = []*Item{
				&Item{"Conjured Mana Cake", 3, 6},
			}

			UpdateQuality(items)

			if items[0].sellIn != 2 {
				t.Errorf("Sellin: Expected %d but got %d ", 2, items[0].sellIn)
			}

			if items[0].quality != 3 {
				t.Errorf("Quality: Expected %d but got %d ", 3, items[0].sellIn)
			}
		})

	})

}
