package project

import "news/pkg/db"

type News struct {
	db.News
	Tags []db.Tag
}

type NewsList []News

func (nl NewsList) TagIDs() map[int64]int64 {
	var out = make(map[int64]int64)
	for _, v := range nl {
		for _, s := range v.TagIDs {
			out[s] = s
		}
	}

	return out
}

func (nl NewsList) SetTags(tags map[int64]db.Tag) {
	for i, v := range nl {
		v.setTags(tags)
		nl[i] = v
	}
}

func newNews(in *db.News) *News {
	if in == nil {
		return nil
	}

	return &News{
		News: *in,
	}
}

func newNewsList(nl []db.News) NewsList {
	var nnl NewsList
	for _, v := range nl {
		nnl = append(nnl, *newNews(&v))
	}
	return nnl
}

func (n *News) setTags(tagsMap map[int64]db.Tag) {
	var out []db.Tag

	for _, v := range n.TagIDs {
		if value, is := tagsMap[v]; is != false {
			out = append(out, value)
		}
	}

	n.Tags = out
}
