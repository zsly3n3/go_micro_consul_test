#!/usr/bin/bash
proto_path="./pb/"
des_path="./pb"

filelist=`ls ${proto_path}`
for file in ${filelist}; do
	suffix=${file##*.}
	if [ ${suffix} == "proto" ]; then
		protoc --proto_path=${proto_path} --micro_out=${proto_path} --go_out=${des_path} ${proto_path}${file}
	fi
done