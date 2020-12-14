package abtest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joselee214/j7f/components/http/server"
	"hash/crc32"
	"math"
	"math/rand"
	"time"
)

var pollOffset int = -1
var cw int = 0
var gcd int = 2

type AbController struct {
	server.Controller
}

type AbRequest struct {
	Uuid int64 `json:"uid" form:"uid"`
}

func Init(s *gin.Engine) {
	b := &AbController{}
	s.POST("/abtest/options", b.getOptions) //
}

func (c *AbController) getOptions(context *gin.Context) {
	var req AbRequest
	context.Bind(&req)

	// 获取策略

	options := Opts.AbTestOpt

	out := make(map[string]string)

	for _, singleOpt := range options {

		data := explain(singleOpt, context)
		out[singleOpt.Key] = data
	}

	c.Data = out
	c.ResponseSuccess(context)
}

func explain(abtest AbTest, context *gin.Context) string {

	mode := abtest.Mode

	distributions := abtest.Distribution
	length := len(distributions)

	if length == 0 {
		return ""
	}

	//explainResult := make(map[string] interface{})

	//explainResult["key"] = abtest.Key
	var explainResult string

	switch mode {
	case "random":
		rand.Seed(time.Now().Unix())
		randIndex := WeightedRandomIndex(distributions)
		//randIndex := rand.Intn(length - 1)
		explainResult = distributions[randIndex].Chromosome
		break
	case "polling":

		if pollOffset > length-1 {
			pollOffset = 0
		}

		for {
			pollOffset = (pollOffset + 1) % length
			if pollOffset == 0 {
				cw = cw - gcd
				if cw <= 0 {
					cw = getMaxWeight(distributions)
					if cw == 0 {
						return ""
					}
				}
			}
			// end
			if weight := distributions[pollOffset].Weight; weight >= cw {
				explainResult = distributions[pollOffset].Chromosome
				break
			}
		}
		//fmt.Println(pollOffset)
		//fmt.Println(cw)
		//fmt.Println(gcd)
		break
	case "remainder":
		disKey := abtest.Remainder.DividendKey
		disKeyV := context.PostForm(disKey)
		if disKeyV != "" {

			//hash := xxhash.New64()
			//_, _ = hash.Write([]byte(disKeyV))
			//hashRet := hash.Sum64()
			hashRet := crc32.ChecksumIEEE([]byte(disKeyV))
			hs := hashRet ^ 0x7FFFFFFF
			i := uint64(hs)
			l := uint64(length)
			l = uint64(math.Pow(2, float64(l)))
			index := i & (l - 1)
			explainResult = distributions[index].Chromosome

			fmt.Println(disKeyV)

		}
		break
	}
	return explainResult

}

func WeightedRandomIndex(l []Distribution) int {
	var sum float32 = 0.00
	for _, singleDis := range l {
		var weightFormat = float32(singleDis.Weight / 10)
		sum += weightFormat
	}
	r := rand.Float32() * sum

	var t float32 = 0.0
	for i, singleDis := range l {
		var weightFormat = float32(singleDis.Weight / 10)
		t += weightFormat
		if t > r {
			return i
		}
	}
	return len(l) - 1
}

func getMaxWeight(dis []Distribution) int {
	max := 0

	for _, v := range dis {

		if weight := v.Weight; weight >= max {
			max = weight
		}

	}
	return max
}
