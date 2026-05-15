if which go > /dev/null; then
  export GOPATH=$(go env GOPATH)
  export PATH="$PATH:$GOPATH/bin"


  export GO111MODULE=on
  export GOPRIVATE="$GOPRIVATE,github.com/cao7113/golangprivatemod"
  # export GOPROXY=https://goproxy.io
  export GOPROXY=https://goproxy.cn
fi
