package service

import (
	"context"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/define/cachekey"
	"fgzs-single/internal/errorx"
	"fgzs-single/pkg/util/jsonutil"
	"github.com/importcjj/sensitive"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
	"sync"
)

var _ iSensitiveWordService = (*SensitiveWordService)(nil)

type SensitiveWordService struct {
	ctx   context.Context
	db    *gorm.DB
	redis *redis.Redis
}
type iSensitiveWordService interface {
	Check(word string) (*SensitiveWordCheck, error)
}

func NewSensitiveWordService(ctx context.Context, DB *gorm.DB, rd *redis.Redis) *SensitiveWordService {
	return &SensitiveWordService{ctx: ctx, db: DB, redis: rd}
}

type SensitiveCheck struct {
	len    int
	filter *sensitive.Filter
}

var lock sync.Mutex
var sc = new(SensitiveCheck)

type SensitiveWordCheck struct {
	Result  bool   `json:"Result,omitempty"`  //是否有敏感词
	Replace string `json:"Replace,omitempty"` //替换后词语
	Filter  string `json:"Filter,omitempty"`  //移除后词语
}

func (s *SensitiveWordService) Words() ([]string, error) {
	cacheKey := cachekey.SensitiveWord.BuildCacheKey()
	words := make([]string, 0)
	err := cacheKey.AutoCache(s.redis, &words, func() (string, error) {
		ws := make([]string, 0)
		sensitiveWordDao := dao.Use(s.db).SensitiveWord
		err := sensitiveWordDao.WithContext(s.ctx).Pluck(sensitiveWordDao.Text, &ws)
		if err != nil {
			return "", errorx.DataSqlErr.WithDetail(err)
		}
		toString, err := jsonutil.EncodeToString(ws)
		if err != nil {
			return "", err
		}
		return toString, nil
	})
	if err != nil {
		return nil, err
	}
	return words, nil
}

func (s *SensitiveWordService) Check(word string) (*SensitiveWordCheck, error) {
	words, err := s.Words()
	if err != nil {
		return nil, err
	}
	if sc.len != len(words) {
		lock.Lock()
		defer lock.Unlock()
		sc = &SensitiveCheck{
			len:    len(words),
			filter: sensitive.New(),
		}
		sc.filter.AddWord(words...)
		sc.filter.UpdateNoisePattern(`x`)
	}
	validate, _ := sc.filter.Validate(word)
	if validate {
		return &SensitiveWordCheck{
			Result:  false,
			Replace: "",
			Filter:  "",
		}, nil
	}
	replace := sc.filter.Replace(word, '*')
	filterStr := sc.filter.Filter(word)
	return &SensitiveWordCheck{
		Result:  true,
		Replace: replace,
		Filter:  filterStr,
	}, nil
}
