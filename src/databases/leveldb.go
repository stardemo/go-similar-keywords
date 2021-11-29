package databases

import (
	"github.com/gogf/gf/frame/g"
	"github.com/stardemo/go-similar-keywords/src/utils"
	"github.com/syndtr/goleveldb/leveldb"
)

const (
	ConfigPrefixKeywordToIndexPath = "leveldb.KeywordToIndex"
	ConfigPrefixIndexToKeywordPath = "leveldb.IndexToKeyword"
)

type KeywordsDatabase struct {
	IndexToKeyword *leveldb.DB
	KeywordToIndex *leveldb.DB
}

var KeyWords *KeywordsDatabase

func init() {
	KeyWords = &KeywordsDatabase{
		IndexToKeyword: initIndexToKeywordDb(),
		KeywordToIndex: initKeywordToIndexDb(),
	}
}

func initIndexToKeywordDb() *leveldb.DB {
	dbIndexToKeyword, err := leveldb.OpenFile(utils.FindCfg(ConfigPrefixIndexToKeywordPath, "", false).String(), nil)
	if err != nil {
		g.Log().Panic(err)
	}
	return dbIndexToKeyword
}
func initKeywordToIndexDb() *leveldb.DB {
	dbKeywordToIndex, err := leveldb.OpenFile(utils.FindCfg(ConfigPrefixKeywordToIndexPath, "", false).String(), nil)
	if err != nil {
		g.Log().Panic(err)
	}
	return dbKeywordToIndex
}
