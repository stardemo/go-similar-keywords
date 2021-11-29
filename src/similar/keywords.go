package similar

import (
	"github.com/gogf/gf/frame/g"
	"github.com/huichen/sego"
	"github.com/stardemo/go-similar-keywords/pkg/annoyindex"
	"github.com/stardemo/go-similar-keywords/src/databases"
	"github.com/stardemo/go-similar-keywords/src/utils"
	"math"
)

const (
	defaultNumReturnKeywords   = 10
	vecDim                     = 200
	kSearch                    = 10000
	ConfigPrefixAnnoyVecDim    = "annoy.vecDim"
	ConfigPrefixAnnoyIndexPath = "annoy.IndexPath"
)

var (
	segmenter  sego.Segmenter
	annoyIndex annoyindex.AnnoyIndex
)

type KeywordsResponse struct {
	Keywords []Keyword `json:"keywords"`
}

type Keyword struct {
	Word       string  `json:"word"`
	Similarity float32 `json:"similarity"`
}

func init() {
	annoyIndex = annoyindex.NewAnnoyIndexAngular(utils.FindCfg(ConfigPrefixAnnoyVecDim, "", false).Int())
	annoyIndex.Load(utils.FindCfg(ConfigPrefixAnnoyIndexPath, "", false).String())
}

func DoGetSimilarKeywords(req *ReqSimilarKeywords) (*KeywordsResponse, error) {
	if req.Num == 0 {
		g.Log().Warningf("Req num not set,set to default %d.", defaultNumReturnKeywords)
		req.Num = defaultNumReturnKeywords
	}
	g.Log().Debugf("DoGetSimilarKeywords keyword:%s,reqCount:%d", req.KeyWords, req.Num)
	data, err := getSimilarKeyword(req.KeyWords, req.Num)
	if err != nil {
		g.Log().Errorf("DoGetSimilarKeywords Error:%s", err.Error())
		return nil, err
	}
	return data, nil
}

func getSimilarKeywords(keywords []string) (*KeywordsResponse, error) {
	//TODO
	return nil, nil
}

func getSimilarKeyword(keyword string, num uint) (*KeywordsResponse, error) {
	wordVec := make([]float32, vecDim)
	_, err := databases.KeyWords.KeywordToIndex.Get([]byte(keyword), nil)
	if err != nil {
		// 如果是索引未找到则分词查询
		if err.Error() == "leveldb: not found" {
			segments := segmenter.Segment([]byte(keyword))
			return getSimilarKeywords(sego.SegmentsToSlice(segments, false))
		}
		return nil, err
	}
	// TODO
	// 多关键词情况下处理

	id, err := databases.KeyWords.KeywordToIndex.Get([]byte(keyword), nil)
	if err != nil {
		return nil, err
	}
	index := utils.Uint32frombytes(id)
	var wv []float32
	annoyIndex.GetItem(int(index), &wv)
	for i, v := range wv {
		wordVec[i] = wordVec[i] + v
	}

	var result []int
	annoyIndex.GetNnsByVector(wordVec, int(num), kSearch, &result)
	var sim KeywordsResponse
	for _, k := range result {
		keyword, err := databases.KeyWords.IndexToKeyword.Get(utils.Uint32bytes(uint32(k)), nil)
		if err != nil {
			return nil, err
		}
		similarityScore := getCosineSimilarityByVector(wordVec, k)
		sim.Keywords = append(sim.Keywords, Keyword{
			Word:       string(keyword),
			Similarity: similarityScore,
		})
	}
	return &sim, nil
}

func getCosineSimilarityByVector(vec []float32, j int) float32 {
	var vec2 []float32
	annoyIndex.GetItem(j, &vec2)

	var a, b, c float32
	for id, v := range vec {
		a = a + v*vec2[id]
		b = b + v*v
		c = c + vec2[id]*vec2[id]
	}

	return a / float32(math.Sqrt(float64(b))*math.Sqrt(float64(c)))
}
