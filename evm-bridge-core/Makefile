build:
ifeq ($(OS),Windows_NT)
	go build -o build/swap-backend.exe main.go
else
	go build -o build/swap-backend main.go
endif

install:
ifeq ($(OS),Windows_NT)
	go install main.go
else
	go install main.go
endif

build-abi:
ifeq ($(OS),Windows_NT)
	abigen --abi=abi/ERC721SwapAgent.json --type=ERC721SwapAgent --pkg=abi --out=abi/ERC721SwapAgent.go
	abigen --abi=abi/ERC721Token.json --type=ERC721Token --pkg=abi --out=abi/ERC721Token.go
	abigen --abi=abi/ERC1155SwapAgent.json --type=ERC1155SwapAgent --pkg=abi --out=abi/ERC1155SwapAgent.go
	abigen --abi=abi/ERC1155Token.json --type=ERC1155Token --pkg=abi --out=abi/ERC1155Token.go
else
	abigen --abi=abi/ERC721SwapAgent.json --type=ERC721SwapAgent --pkg=abi --out=abi/ERC721SwapAgent.go
	abigen --abi=abi/ERC721Token.json --type=ERC721Token --pkg=abi --out=abi/ERC721Token.go
	abigen --abi=abi/ERC1155SwapAgent.json --type=ERC1155SwapAgent --pkg=abi --out=abi/ERC1155SwapAgent.go
	abigen --abi=abi/ERC1155Token.json --type=ERC1155Token --pkg=abi --out=abi/ERC1155Token.go
endif

.PHONY: build install
