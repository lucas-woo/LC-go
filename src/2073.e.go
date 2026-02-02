package main

func timeRequiredToBuy(tickets []int, k int) int {
	turns := 0;
	for tickets[k] != 0 {
		for i := 0; i < len(tickets); i++ {
			if tickets[i] == 0 {
				continue;
			}
			if i == k && tickets[i] == 1 {
				return turns + 1;
			}
			turns++;
			tickets[i]--;
		}
	}
	return turns
}