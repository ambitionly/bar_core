##这里是处理bar的核心API

##生成PB文件

cd proto
protoc --go_out=plugins=grpc:../data ./*
