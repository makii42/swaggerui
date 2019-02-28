# Keeping it simple. 

gen: 
	go-bindata \
		-modtime 1551345874 \
		-mode 420 \
		-pkg swaggerui \
		-o ./bindata.go \
		./assets
