mfd-xml:
	@mfd-generator xml -c "postgres://postgres:postgres@localhost:5432/news?sslmode=disable" -m ./docs/model/news.mfd -n "news:news,categories,tags"

mfd-model:
	@mfd-generator model -m ./docs/model/news.mfd -p db -o ./pkg/db

mfd-repo:
	@mfd-generator repo -m ./docs/model/news.mfd -p db -o ./pkg/db