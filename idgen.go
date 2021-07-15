package crex

import (
	"github.com/qmarliu/crex/utils"
	"strconv"
)

var idGen *utils.IdGenerate

func SetIdGenerate(g *utils.IdGenerate) {
	idGen = g
}

func GenOrderId() string {
	id := idGen.Next()
	return strconv.Itoa(int(id))
}
