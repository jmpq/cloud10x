all:
	cd apiserver && go build -o ../cxserver
	cd adm && go build -o ../cxadm

clean:
	rm cxserver cxadm
