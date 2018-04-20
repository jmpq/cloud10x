all:
	cd apiserver && go build -o ../cxserver
	cd cli && go build -o ../cxi

clean:
	rm cxserver cxi
