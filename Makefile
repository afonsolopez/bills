all: clean install build compile run

test: clean build compile run

clean:
	rm -rf ./build ./reactjs/build
	echo "[✔️] Clean complete!"

install:
	cd ./reactjs && npm install
	echo "[✔️] Install complete!"	

build:
	cd ./reactjs && npm run build
	echo "[✔️] Build complete!"

compile:
	mkdir -p ./build
	go build -o build
	echo "[✔️] Compile complete!"


run:
	./build/bills
	echo "[✔️] App is running!"

react:
	cd ./reactjs && npm start