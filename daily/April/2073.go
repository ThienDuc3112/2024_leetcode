package main

func timeRequiredToBuy(tickets []int, k int) int {
    res := 0
    for i, time := range tickets {
        res += min(tickets[k], time)
        if i > k && tickets[k] < time {
            res--
        }
    }
    return res
}
