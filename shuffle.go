package shuffle 

import (
	"fmt"
	"math/rand"
	"time"
)

// 定义扑克牌结构体
type Card struct {
	Suit string // 花色
	Rank string // 点数
}

// 定义玩家手牌结构体
type PlayerHand struct {
	PlayerID int
	Cards    []Card
}

// 初始化一副标准扑克牌（52张）
func newDeck() []Card {
	suits := []string{"♠", "♥", "♦", "♣"}
	ranks := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	deck := make([]Card, 0, 52)

	for _, suit := range suits {
		for _, rank := range ranks {
			deck = append(deck, Card{Suit: suit, Rank: rank})
		}
	}
	return deck
}

// Fisher-Yates洗牌算法（保证完全随机）
func shuffleDeck(deck []Card) []Card {
	rand.Seed(time.Now().UnixNano())
	shuffled := make([]Card, len(deck))
	copy(shuffled, deck)

	for i := len(shuffled) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	}
	return shuffled
}

// 发牌函数
func dealCards(deck []Card, players, cardsPerPlayer int) ([]PlayerHand, error) {
	totalCards := players * cardsPerPlayer
	if totalCards > len(deck) {
		return nil, fmt.Errorf("需要%d张牌，但牌堆只有%d张", totalCards, len(deck))
	}

	hands := make([]PlayerHand, players)
	for i := range hands {
		hands[i].PlayerID = i + 1
	}

	// 轮询发牌（模拟真实发牌场景）
	for i := 0; i < cardsPerPlayer; i++ {
		for p := 0; p < players; p++ {
			hands[p].Cards = append(hands[p].Cards, deck[i*players+p])
		}
	}
	return hands, nil
}

// 分析牌型（示例功能）
func analyzeHand(hand []Card) string {
	maxRank := ""
	for _, card := range hand {
		if card.Rank > maxRank {
			maxRank = card.Rank
		}
	}
	return "高牌:" + maxRank
}

// 发牌结果
func printResults(hands []PlayerHand) {
	for _, hand := range hands {
		fmt.Printf("玩家 %d 的手牌: ", hand.PlayerID)
		for _, card := range hand.Cards {
			fmt.Printf("%s%s ", card.Suit, card.Rank)
		}
		fmt.Printf("| 牌型: %s\n", analyzeHand(hand.Cards))
	}
}
