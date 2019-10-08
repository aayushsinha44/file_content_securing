rm -rf build/
mkdir src/
rm -rf src/file
mkdir src/file
rm -rf src/filestore
#mkdir src/filestore
# solc --abi --bin contracts/FileStore.sol -o build
solc --abi --bin contracts/File.sol -o build
# abigen --bin=build/FileStore.bin --abi=build/FileStore.abi --pkg=filestore --out=src/filestore/FileStore.go
abigen --bin=build/File.bin --abi=build/File.abi --pkg=file --out=src/file/file.go