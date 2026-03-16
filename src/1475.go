package main

func finalPrices(prices []int) []int {
	var monotonicStack []int = make([]int, len(prices));
	monotonicStack[0] = 0
	var p int;
	
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[monotonicStack[p]] {
			p++;
			monotonicStack[p] = i
			continue;
		}
		for j := p; j >= 0; j-- {
			if prices[i] <= prices[monotonicStack[p]] {
				prices[monotonicStack[p]] -= prices[i];
				p--;
			}
			if p == -1 || prices[i] > prices[monotonicStack[p]]{
				p++;
				monotonicStack[p] = i;
				break;
			}
		}
	}
	return prices
}
