package card

import (
	"bank/pkg/bank/types"
	"fmt"
)


func ExempleTotal() {
	cards := []types.Card{
	{
		Balance: 10_000_00,
		Active: true,
 	},
	{
		Balance: 10_000_00,
		Active: true,
	},
	{
		Balance: 10_000_00,
		Active: true,
	},
  }
   fmt.Println(Total(cards))
   //Output: 2000000
}