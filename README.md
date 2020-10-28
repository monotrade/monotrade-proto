https://colobu.com/2015/01/07/Protobuf-language-guide/





sudo apt-get install autoconf automake libtool curl make g++ unzip
git clone https://gitee.com/meiqicheng1216/protobuf.git
cd protobuf
git submodule update --init --recursive
./autogen.sh
./configure
make
make check
sudo make install
sudo ldconfig # refresh shared library cache.